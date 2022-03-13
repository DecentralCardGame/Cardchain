package keeper

import (
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
)


func UintItemInArr(item uint64, list []uint64) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func StringItemInArr(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func IndexOfId(id uint64, length int, f func(int) uint64) int {
	for i := 0; i < length; i++ {
		if f(i) == id {
			return i
		}
	}
	return -1
}

func UintPopItemFromArr(element uint64, arr []uint64) ([]uint64, error) {
	for idx, val := range arr {
		if element == val {
			return append(arr[:idx], arr[idx+1:]...), nil
		}
	}
	return []uint64{}, types.ErrCardNotThere
}
