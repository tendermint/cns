package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/cns/x/cns/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Register(ctx sdk.Context, cInfo types.ChainInfo) error {
	store := ctx.KVStore(k.storeKey)

	if err := cInfo.ValidateBasic(); err != nil {
		return err
	}

	err := k.SetChainInfo(store, cInfo)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) Update(ctx sdk.Context, cInfo types.ChainInfo) error {
	store := ctx.KVStore(k.storeKey)

	storeInfo, err := k.GetChainInfo(store, cInfo.ChainName, cInfo.Owner)
	if err != nil {
		return err
	}

	err = cInfo.ValidateBasic()
	if err != nil {
		return err
	}

	storeInfo.Update(cInfo)
	err = k.SetChainInfo(store, storeInfo)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetChainInfo(store sdk.KVStore, cName, owner string) (types.ChainInfo, error) {
	bytes := store.Get(types.GetStorKey(cName, owner))

	var cInfo types.ChainInfo
	err := k.cdc.UnmarshalBinaryBare(bytes, &cInfo)
	if err != nil {
		return types.ChainInfo{}, err
	}

	return cInfo, err
}

func (k Keeper) SetChainInfo(store sdk.KVStore, info types.ChainInfo) error {
	bytes, err := k.cdc.MarshalBinaryBare(&info)
	if err != nil {
		return err
	}

	store.Set(types.GetStorKey(info.ChainName, info.Owner), bytes)
	return nil
}
