package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "cns"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	// Keys for store prefixes
	CnsKey = []byte{0x01} // prefix for each key
)

func GetStorKey(name string, addr sdk.AccAddress) []byte {
	return append(CnsKey, append(addr.Bytes(), []byte(name)...)...)
}
