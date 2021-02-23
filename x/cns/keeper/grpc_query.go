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

//TODO(sahith) Add query for all chain infos
