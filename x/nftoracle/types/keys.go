package types

const (
	// ModuleName defines the module name
	ModuleName = "nftoracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nftoracle"

	KeyPermissionOracle = "oracle"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MintRequestKey      = "MintRequest-value-"
	MintRequestCountKey = "MintRequest-count-"
)

const (
	ActionRequestKey      = "ActionRequest-value-"
	ActionRequestCountKey = "ActionRequest-count-"
)
