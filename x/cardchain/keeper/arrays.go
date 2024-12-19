package keeper

import (
	"fmt"
	"slices"
)

func PopItemFromArr[E comparable](item E, arr []E) ([]E, error) {
	idx := slices.Index(arr, item)
	if idx == -1 {
		return arr, fmt.Errorf("item not in array")
	}
	arr = slices.Delete(arr, idx, idx+1)
	return arr, nil
}
