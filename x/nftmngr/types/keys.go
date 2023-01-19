package types

const (
	// ModuleName defines the module name
	ModuleName = "nftmngr"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nftmngr"
)

// KVStore keys
var (
	// BalancesPrefix is the prefix for the account balances store. We use a byte
	// (instead of `[]byte("balances")` to save some disk space).
	CollectionPrefix = []byte{0x1}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func CollectionkeyPrefix(nftSchemaCode []byte) []byte {
	return append(CollectionPrefix, nftSchemaCode...)
}

const (
	NFTFeeConfigKey = "NFTFeeConfig-value-"

	KeyPermissionNftFeeAdmin = "nft_fee_admin"
)

const (
	NFTFeeBalanceKey = "NFTFeeBalance-value-"
)

const (
	KeyMintPermissionOnlySystem = "system"
	KeyMintPermissionAll        = "all"
)
