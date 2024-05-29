package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BuyOrderKeyPrefix is the prefix to retrieve all BuyOrder
	BuyOrderKeyPrefix = "BuyOrder/value/"
)

// BuyOrderKey returns the store key to retrieve a BuyOrder from the index fields
func BuyOrderKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
