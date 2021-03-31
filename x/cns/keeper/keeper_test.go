package keeper_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/cns/app"
	"github.com/tendermint/cns/x/cns/types"
	ibctesting "github.com/tendermint/cns/x/ibc/testing"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

type TestSuite struct {
	suite.Suite
	app         *app.App
	ctx         sdk.Context
	queryClient types.QueryClient
	cInfo       types.ChainInfo

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
	chainC *ibctesting.TestChain
}

func (suite *TestSuite) SetupTest() {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.CNSkeeper)
	queryClient := types.NewQueryClient(queryHelper)
	addr := "cosmos1wrx0x9m9ykdhw9sg04v7uljme53wuj03cfqce6"
	ownerAddr, err := sdk.AccAddressFromBech32(addr)
	suite.Require().NoError(err)
	suite.Require().NoError(app.BankKeeper.SetBalances(ctx, ownerAddr,
		sdk.NewCoins(sdk.NewInt64Coin("stake", 100000000))))
	suite.cInfo = types.ChainInfo{
		ChainName:          "test",
		Expiration:         0,
		Owner:              addr,
		CanonicalIbcClient: "test",
		Seed:               []string{"xyz@abc:1"},
		SourceCodeUrl:      "github.com/tendermint/cns",
		Version: &types.VersionInfo{
			Version:          1,
			SourceCommitHash: "d29ef87107420fe2e5443f7436da2ff69ce3a5f8",
			GenesisHash:      "d29ef87107420fe2e5443f7436da2ff69ce3a5f8",
		},
	}
	// register test info
	err = app.CNSkeeper.Register(ctx, suite.cInfo)
	suite.Require().NoError(err)

	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainC = suite.coordinator.GetChain(ibctesting.GetChainID(2))
	suite.ctx, suite.app, suite.queryClient = ctx, app, queryClient
}
func (suite *TestSuite) TestKeeper() {
	k := suite.app.CNSkeeper
	addr := "cosmos1wrx0x9m9ykdhw9sg04v7uljme53wuj03cfqce6"
	testCases := []struct {
		msg      string
		cInfo    types.ChainInfo
		register bool
		expPass  bool
	}{
		{
			msg: "fetch valid chain info",
			cInfo: types.ChainInfo{
				ChainName: "test",
				Owner:     addr,
			},
			expPass: true,
		},
		{
			msg: "Fetch random chain name",
			cInfo: types.ChainInfo{
				ChainName: "random",
			},
			expPass: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				info, err := k.GetChainInfo(suite.ctx, tc.cInfo.ChainName, tc.cInfo.Owner)
				suite.Require().NoError(err)
				suite.Require().Equal(suite.cInfo, info)
			} else {
				info, err := k.GetChainInfo(suite.ctx, tc.cInfo.ChainName, tc.cInfo.Owner)
				suite.Require().Error(err)
				suite.Require().Empty(info)
			}
		})
	}
	// test IterateAllInfos
	var infos []types.ChainInfo
	k.IterateAllInfos(suite.ctx, func(info types.ChainInfo) bool {
		infos = append(infos, info)
		return false
	})
	suite.Require().Equal(1, len(infos))
}
func (suite *TestSuite) TestUpdate() {
	k := suite.app.CNSkeeper
	addr := "cosmos1wrx0x9m9ykdhw9sg04v7uljme53wuj03cfqce6"
	updatedInfo := types.ChainInfo{
		ChainName:          "test",
		Expiration:         0,
		Owner:              addr,
		CanonicalIbcClient: "test2",
		Seed:               []string{"abc@xyz:1"},
		SourceCodeUrl:      "github.com/tendermint/starport",
		Version: &types.VersionInfo{
			Version:          2,
			SourceCommitHash: "d29ef87107420fe2e5443f7436da2rg49ce3a5f8",
			GenesisHash:      "d29ef87107420fe2e5443f7436da2rg49ce3a5f8",
		},
	}

	testCases := []struct {
		msg      string
		newInfo  types.ChainInfo
		cInfo    types.ChainInfo
		register bool
		expPass  bool
	}{
		{
			msg: "update chain info",
			cInfo: types.ChainInfo{
				ChainName: "test",
				Owner:     addr,
			},
			expPass: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				err := k.Update(suite.ctx, updatedInfo)
				suite.Require().NoError(err)
				info, err := k.GetChainInfo(suite.ctx, tc.cInfo.ChainName, tc.cInfo.Owner)
				suite.Require().Equal(updatedInfo, info)
			} else {
				info, err := k.GetChainInfo(suite.ctx, tc.cInfo.ChainName, tc.cInfo.Owner)
				suite.Require().Error(err)
				suite.Require().Empty(info)
			}
		})
	}
}
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
