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

func (suite *TestSuite) TestFetchBalances() {
	//addr := "cosmos1wrx0x9m9ykdhw9sg04v7uljme53wuj03cfqce6"
	//ownerAddr, err := sdk.AccAddressFromBech32(addr)
	//suite.Require().NoError(err)
	//suite.Require().NoError(suite.app.BankKeeper.SetBalances(suite.ctx, ownerAddr,
	//	sdk.NewCoins(sdk.NewInt64Coin("stake", 100000000), sdk.NewInt64Coin("ibc/F48A74D4F2A8B3D0CA0572F682A334C4F9595827BFE030BFE112863A2BC928C0", 100))))
	//coins, err := suite.app.CNSkeeper.FetchUpdatedBalances(suite.ctx, addr)
	//suite.Require().NoError(err)
	//suite.T().Log(coins)
}

func (suite *TestSuite) TestSendTokens() {
	var (
		channelA, channelB ibctesting.TestChannel
	)

	clientA, clientB, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
	channelA, channelB = suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
	_ = suite.chainB.SenderAccount.GetAddress().String() // must be explicitly changed in malleate

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

	//err = suite.chainB.App.CNSkeeper.Register(suite.chainB.GetContext(), cnstypes.ChainInfo{
	//	ChainName: suite.chainA.ChainID,
	//	Owner: suite.chainA.SenderAccount.GetAddress().String(),
	//	CanonicalIbcClient: clientC,
	//})
	suite.Require().NoError(err)

	suite.T().Log(suite.chainA.App.BankKeeper.GetAllBalances(suite.chainA.GetContext(), suite.chainA.SenderAccount.GetAddress()))

	coins, err := suite.chainA.App.CNSkeeper.FetchUpdatedBalances(suite.chainA.GetContext(), suite.chainA.SenderAccount.GetAddress())
	suite.Require().NoError(err)
	suite.T().Log(coins)
	suite.Require().Len(coins, 2)
}
