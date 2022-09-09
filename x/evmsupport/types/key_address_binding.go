package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AddressBindingKeyPrefix is the prefix to retrieve all AddressBinding
	AddressBindingKeyPrefix = "AddressBinding/value/"
)

// AddressBindingKey returns the store key to retrieve a AddressBinding from the index fields
func EthAddressBindingKey(
	ethAddress string,
	nativeAddress string,
) []byte {
	var key []byte

	ethAddressBytes := []byte(ethAddress)
	key = append(key, ethAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// AddressBindingKey returns the store key to retrieve a AddressBinding from the index fields
func AddressBindingKey(
	ethAddress string,
) []byte {
	var key []byte

	ethAddressBytes := []byte(ethAddress)
	key = append(key, ethAddressBytes...)
	key = append(key, []byte("/")...)

	// nativeAddressBytes := []byte(nativeAddress)
	// key = append(key, nativeAddressBytes...)
	// key = append(key, []byte("/")...)

	return key
}
