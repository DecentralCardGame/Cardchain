package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	errorsmod "cosmossdk.io/errors"
)

const ArtworkMaxSize = 500000

func (i Image) GetHash() string {
	sum := md5.Sum(i.Image)
	return hex.EncodeToString(sum[:])
}

type HasImage interface {
	GetImage() []byte
}

func ValidateImage(msg HasImage) error {
	imageLen := len(msg.GetImage())

	if imageLen > ArtworkMaxSize {
		return errorsmod.Wrap(ErrImageSizeExceeded, fmt.Sprint(imageLen))
	}
	return nil
}
