package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SellOrderKeyPrefix is the prefix to retrieve all SellOrder
	SellOrderKeyPrefix = "SellOrder/value/"
)

// SellOrderKey returns the store key to retrieve a SellOrder from the index fields
func SellOrderKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
