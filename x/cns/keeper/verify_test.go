package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	commitmenttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/23-commitment/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	ibctesting "github.com/tendermint/cns/x/ibc/testing"
)

func (suite *TestSuite) Test2chains() {
	var (
		channelA, channelB ibctesting.TestChannel
	)

	_, clientAB, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
	channelA, channelB = suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

	coinFromBToA := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))
	transferMsg := types.NewMsgTransfer(channelB.PortID, channelB.ID, coinFromBToA, suite.chainB.SenderAccount.GetAddress(), suite.chainA.SenderAccount.GetAddress().String(), clienttypes.NewHeight(0, 110), 0)
	err := suite.coordinator.SendMsg(suite.chainB, suite.chainA, channelA.ClientID, transferMsg)
	suite.Require().NoError(err) // message committed

	// query client state of clientAB on chainB
	ctx := sdk.WrapSDKContext(suite.chainA.GetContext())
	cliRes, err := suite.chainB.App.IBCKeeper.ClientKeeper.ClientState(ctx, &clienttypes.QueryClientStateRequest{ClientId: clientAB})
	clientState, err := clienttypes.UnpackClientState(cliRes.ClientState)
	suite.T().Log(clientState)
	suite.T().Log(cliRes.ProofHeight)
	suite.Require().NoError(err)

	_, proofBz := suite.chainB.QueryClientStateProof(clientAB)

	// query client state of clientAB on chain A
	ctx = sdk.WrapSDKContext(suite.chainA.GetContext())
	clientStateA, ok := suite.chainA.App.IBCKeeper.ClientKeeper.GetClientState(suite.chainA.GetContext(), clientAB)
	suite.Require().Equal(true, ok)

	k := suite.chainA.App.CNSkeeper
	prefix := k.ConnKeeper.GetCommitmentPrefix()
	prefixBz := commitmenttypes.NewMerklePrefix(prefix.Bytes())

	// verify the client states match and proof is verified
	err = clientStateA.VerifyClientState(k.ClientKeeper.ClientStore(suite.chainA.GetContext(), clientAB), k.Cdc, clientState.GetLatestHeight(),
		&prefixBz, clientAB, proofBz, clientState)
	suite.Require().NoError(err)
}
