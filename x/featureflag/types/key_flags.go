package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FlagsKeyPrefix is the prefix to retrieve all Flags
	FlagsKeyPrefix = "Flags/value/"
)

// FlagsKey returns the store key to retrieve a Flags from the index fields
func FlagsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
