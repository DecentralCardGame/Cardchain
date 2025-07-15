package types

import (
	"crypto/md5"
	"encoding/hex"
)

func (i Image) GetHash() string {
	sum := md5.Sum(i.Image)
	return hex.EncodeToString(sum[:])
}
