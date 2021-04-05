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

	// create header from chainA and update clientAB on chain B
	header1, err := suite.chainB.ConstructUpdateTMClientHeader(suite.chainA, clientAB)
	suite.Require().NoError(err)
	suite.chainB.App.IBCKeeper.ClientKeeper.UpdateClient(suite.chainB.GetContext(), clientAB, header1)
	suite.chainB.NextBlock()

	clientstate, proofBz := suite.chainB.QueryClientStateProof(clientAB)
	suite.T().Log(clientstate)
	suite.Require().NoError(err)

	// TODO: DEBUG - without this header update, the consensus state doesn't exist.
	header, err := suite.chainC.ConstructUpdateTMClientHeader(suite.chainA, clientAC)
	suite.Require().NoError(err)
	suite.T().Log(header)

	err = suite.chainC.App.IBCKeeper.ClientKeeper.UpdateClient(suite.chainC.GetContext(), clientAC, header1)
	suite.Require().NoError(err)

	clientstate1, _ := suite.chainB.QueryClientStateProof(clientAC)

	// fetch consensus state - is same on both clients
	cons1, _ := suite.chainB.App.IBCKeeper.ClientKeeper.GetClientConsensusState(suite.chainB.GetContext(), clientAB, header1.GetHeight())
	cons2, _ := suite.chainC.App.IBCKeeper.ClientKeeper.GetClientConsensusState(suite.chainC.GetContext(), clientAC, header1.GetHeight())
	suite.T().Log(clientstate)
	suite.T().Log(clientstate1)
	suite.T().Log(cons1)
	suite.T().Log(cons2)

	suite.chainA.NextBlock()
	suite.chainB.NextBlock()
	suite.chainC.NextBlock()
	err = suite.chainC.App.CNSkeeper.ProveClientState(suite.chainC.GetContext(), clientAC, header1.GetHeight(), clientstate, proofBz, header1)
	suite.T().Log(err)
	suite.Require().NoError(err)
}
