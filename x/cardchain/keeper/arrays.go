package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"golang.org/x/exp/slices"
)

func PopItemFromArr[E comparable](item E, arr []E) ([]E, error) {
	idx := slices.Index(arr, item)
	if idx == -1 {
		return arr, types.ErrCardNotThere
	}
	arr = slices.Delete(arr, idx, idx+1)
	return arr, nil
}
