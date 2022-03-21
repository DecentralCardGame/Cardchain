package keeper

import (
	"golang.org/x/exp/slices"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)

func PopItemFromArr[E comparable](item E, arr []E) ([]E, error) {
	idx := slices.Index(arr, item)
	if idx == -1 {
		return arr, types.ErrCardNotThere
	}
	arr = slices.Delete(arr, idx, idx+1)
	return arr, nil
}
