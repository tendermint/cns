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

	clientA, clientB, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
	channelA, channelB = suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
	_ = suite.chainB.SenderAccount.GetAddress().String()

	coinFromBToA := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))
	transferMsg := types.NewMsgTransfer(channelB.PortID, channelB.ID, coinFromBToA, suite.chainB.SenderAccount.GetAddress(), suite.chainA.SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 110), 0)
	err := suite.coordinator.SendMsg(suite.chainB, suite.chainA, channelA.ClientID, transferMsg)
	suite.Require().NoError(err) // message committed

	// relay send packet
	fungibleTokenPacket := types.NewFungibleTokenPacketData(coinFromBToA.Denom, coinFromBToA.Amount.Uint64(), suite.chainB.SenderAccount.GetAddress().String(), suite.chainA.SenderAccount.GetAddress().String())
	packet := channeltypes.NewPacket(fungibleTokenPacket.GetBytes(), 1, channelB.PortID, channelB.ID, channelA.PortID, channelA.ID, clienttypes.NewHeight(0, 110), 0)
	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
	err = suite.coordinator.RelayPacket(suite.chainB, suite.chainA, clientB, clientA, packet, ack.GetBytes())
	suite.Require().NoError(err) // relay committed

	// register chains with diff client IDs
	_, clientD, _, _ := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
	err = suite.chainA.App.CNSkeeper.Register(suite.chainA.GetContext(), cnstypes.ChainInfo{
		ChainName:          suite.chainB.ChainID,
		Owner:              suite.chainB.SenderAccount.GetAddress().String(),
		CanonicalIbcClient: clientD,
	})
	suite.Require().NoError(err)

	suite.T().Log(suite.chainA.App.BankKeeper.GetAllBalances(suite.chainA.GetContext(), suite.chainA.SenderAccount.GetAddress()))
	coins, err := suite.chainA.App.CNSkeeper.FetchUpdatedBalances(suite.chainA.GetContext(), suite.chainA.SenderAccount.GetAddress())
	suite.Require().NoError(err)
	suite.T().Log(coins)
	suite.Require().Len(coins, 2)
}
