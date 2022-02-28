package keeper

import (
  "crypto/sha256"
  "encoding/hex"

  "github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)

// GetResponseHash generates a hash form a response and a secret
func GetResponseHash(response types.Response, secret string) string {
	clearResponse := response.String() + secret
	hashResponse := sha256.Sum256([]byte(clearResponse))
	hashStringResponse := hex.EncodeToString(hashResponse[:])
  return hashStringResponse
}
