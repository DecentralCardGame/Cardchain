package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)


func UintItemInList(item uint64, list []uint64) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func StringItemInList(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func IndexOfId(cardID uint64, cards []uint64) int {
	if cards == nil {
		return -1
	}
	for i, b := range cards {
		if b == cardID {
			return i
		}
	}
	return -1
}

func UintPopElementFromArr(element uint64, arr []uint64) ([]uint64, error) {
	for idx, val := range arr {
		if element == val {
			return append(arr[:idx], arr[idx+1:]...), nil
		}
	}
	return []uint64{}, types.ErrCardNotThere
}
