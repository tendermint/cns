package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cns/x/cns/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryChainInfo(c context.Context, req *types.QueryChainInfoRequest) (*types.QueryChainInfoResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.ChainName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty chain name")
	}
	ctx := sdk.UnwrapSDKContext(c)

	info := k.GetChainInfoFromName(ctx, req.ChainName)

	return &types.QueryChainInfoResponse{Info: info}, nil
}

func (k Keeper) QueryBalances(c context.Context, req *types.QueryBalancesRequest) (*types.QueryBalancesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.Addr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty addr")
	}

	addr, err := sdk.AccAddressFromBech32(req.Addr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address")
	}

	ctx := sdk.UnwrapSDKContext(c)
	balances, err := k.FetchUpdatedBalances(ctx, addr)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBalancesResponse{Balances: balances}, nil
}

//TODO(sahith) Add query for all chain infos
