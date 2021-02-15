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

	_, err := k.GetChainInfo(store, cInfo.ChainName)
	if err != nil {
		return err
	}

	err = cInfo.ValidateBasic()
	if err != nil {
		return err
	}

	//TODO(sahith): don't overwrite - compare and retain old values
	err = k.SetChainInfo(store, cInfo)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetChainInfo(store sdk.KVStore, cName string) (types.ChainInfo, error) {
	//TODO(sahith): update to use store key
	bytes := store.Get(types.KeyPrefix(cName))

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

	ownerAddr, _ := sdk.AccAddressFromBech32(info.Owner)
	store.Set(types.GetStorKey(info.ChainName, ownerAddr), bytes)
	return nil
}
