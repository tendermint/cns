package keeper

import (
	"fmt"
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/tendermint/cns/x/cns/types"
)

type TestSuite struct {
	suite.Suite

	ctx         sdk.Context
	keeper      Keeper
	queryClient types.QueryClient
}

func (s *TestSuite) SetupTest() {
	keyIdentifier := sdk.NewKVStoreKey(types.StoreKey)
	memKeyIdentifier := sdk.NewKVStoreKey(types.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIdentifier, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(memKeyIdentifier, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "test"}, true, nil)

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	k := NewKeeper(
		marshaler,
		keyIdentifier,
		memKeyIdentifier,
	)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, interfaceRegistry)
	types.RegisterQueryServer(queryHelper, k)
	queryClient := types.NewQueryClient(queryHelper)

	s.ctx, s.keeper, s.queryClient = ctx, *k, queryClient
}

func (s *TestSuite) TestKeeper() {
	testInfo := types.ChainInfo{
		ChainName:          "test",
		Expiration:         0,
		Owner:              "addr",
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
	err := s.keeper.Register(s.ctx, testInfo)
	s.Require().NoError(err)

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
				Owner:     "addr",
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
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				info, err := s.keeper.GetChainInfo(s.ctx, tc.cInfo.ChainName, tc.cInfo.Owner)
				s.Require().NoError(err)
				s.Require().Equal(testInfo, info)
				s.T().Log(s.keeper.GetChainInfoFromName(s.ctx, tc.cInfo.ChainName))
			} else {
				info, err := s.keeper.GetChainInfo(s.ctx, tc.cInfo.ChainName, tc.cInfo.Owner)
				s.Require().Error(err)
				s.Require().Empty(info)
			}
		})
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
