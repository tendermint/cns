package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	cnstypes "github.com/tendermint/cns/x/cns/types"
	ibctesting "github.com/tendermint/cns/x/ibc/testing"
)

func (suite *TestSuite) TestSendTokens() {
	var (
		channelA, channelB ibctesting.TestChannel
	)

	clientBA, clientAB, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
	channelA, channelB = suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
	_ = suite.chainB.SenderAccount.GetAddress().String()

	_, clientAC, _, _ := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainC, exported.Tendermint)
	_, clientBC, _, _ := suite.coordinator.SetupClientConnections(suite.chainB, suite.chainC, exported.Tendermint)

	coinFromBToA := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))
	transferMsg := types.NewMsgTransfer(channelB.PortID, channelB.ID, coinFromBToA, suite.chainB.SenderAccount.GetAddress(), suite.chainA.SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 110), 0)
	err := suite.coordinator.SendMsg(suite.chainB, suite.chainA, channelA.ClientID, transferMsg)
	suite.Require().NoError(err) // message committed

	// relay send packet
	fungibleTokenPacket := types.NewFungibleTokenPacketData(coinFromBToA.Denom, coinFromBToA.Amount.Uint64(), suite.chainB.SenderAccount.GetAddress().String(), suite.chainA.SenderAccount.GetAddress().String())
	packet := channeltypes.NewPacket(fungibleTokenPacket.GetBytes(), 1, channelB.PortID, channelB.ID, channelA.PortID, channelA.ID, clienttypes.NewHeight(0, 110), 0)
	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
	err = suite.coordinator.RelayPacket(suite.chainB, suite.chainA, clientAB, clientBA, packet, ack.GetBytes())
	suite.Require().NoError(err) // relay committed

	// register chains with diff client IDs
	err = suite.chainC.App.CNSkeeper.Register(suite.chainC.GetContext(), cnstypes.ChainInfo{
		ChainName:          suite.chainA.ChainID,
		Owner:              suite.chainA.SenderAccount.GetAddress().String(),
		CanonicalIbcClient: clientAC,
	})
	suite.Require().NoError(err)
	err = suite.chainC.App.CNSkeeper.Register(suite.chainC.GetContext(), cnstypes.ChainInfo{
		ChainName:          suite.chainB.ChainID,
		Owner:              suite.chainB.SenderAccount.GetAddress().String(),
		CanonicalIbcClient: clientBC,
	})
	suite.Require().NoError(err)

	//suite.T().Log(suite.chainA.GetContext().BlockHeight())
	//header2, err := suite.chainA.ConstructUpdateTMClientHeader(suite.chainC, clientAC)
	//suite.Require().NoError(err)

	ctx := sdk.WrapSDKContext(suite.chainB.GetContext())
	queryRes, err := suite.chainB.App.IBCKeeper.ClientKeeper.ClientState(ctx, &clienttypes.QueryClientStateRequest{ClientId: clientAB})
	suite.Require().NoError(err)
	header1, err := suite.chainB.ConstructUpdateTMClientHeader(suite.chainA, clientAB)
	suite.Require().NoError(err)
	suite.chainB.App.IBCKeeper.ClientKeeper.UpdateClient(suite.chainB.GetContext(), clientAB, header1)
	suite.chainB.NextBlock()
	_, proofBz := suite.chainB.QueryClientStateProof(clientAB)
	suite.T().Log(string(proofBz))
	suite.Require().NoError(err)

	queryRes, err = suite.chainB.App.IBCKeeper.ClientKeeper.ClientState(ctx, &clienttypes.QueryClientStateRequest{ClientId: clientAB})
	suite.Require().NoError(err)
	newCS, err := clienttypes.UnpackClientState(queryRes.ClientState)
	suite.Require().NoError(err)

	ctx = sdk.WrapSDKContext(suite.chainC.GetContext())
	queryRes2, err := suite.chainC.App.IBCKeeper.ClientKeeper.ClientState(ctx, &clienttypes.QueryClientStateRequest{ClientId: clientAC})
	suite.Require().NoError(err)
	suite.T().Log(queryRes2)
	suite.T().Log(queryRes)

	header, err := suite.chainC.ConstructUpdateTMClientHeader(suite.chainA, clientAC)
	suite.Require().NoError(err)
	suite.T().Log(header)

	suite.chainC.NextBlock()
	err = suite.chainC.App.CNSkeeper.ProveClientState(suite.chainC.GetContext(), clientAC, header1.GetHeight(), newCS, proofBz, header1)
	suite.T().Log(err)
}
