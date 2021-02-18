package types

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

var (
	// Keys for store prefixes
	CnsKey = []byte{0x01} // prefix for each key
)

//TODO(sahith) Is using owner and addr for key an overkill??
func GetStoreKey(name, addr string) []byte {
	return append(CnsKey, append([]byte(name), []byte(addr)...)...)
}
