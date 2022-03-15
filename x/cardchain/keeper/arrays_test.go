package keeper_test

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/stretchr/testify/require"
)

func TestArrays(t *testing.T) {
	// UintItemInArr
	uintList := []uint64{1, 5, 8, 9, 4}
	require.EqualValues(t, true, keeper.UintItemInArr(5, uintList))
	require.EqualValues(t, false, keeper.UintItemInArr(10, uintList))

	// stringList
	stringList := []string{"Hello", "world", "I", "hope", "this", "works"}
	require.EqualValues(t, true, keeper.StringItemInArr("hope", stringList))
	require.EqualValues(t, false, keeper.StringItemInArr("not", stringList))

	// IndexOfId
	require.EqualValues(t, 2, keeper.IndexOfId(8, len(uintList), func(i int) uint64 {return uintList[i]}))
	require.EqualValues(t, -1, keeper.IndexOfId(10, len(uintList), func(i int) uint64 {return uintList[i]}))
	// require.EqualValues(t, -1, keeper.IndexOfId(8, nil))

	// UintPopItemFromArr
	arr, err := keeper.UintPopItemFromArr(9, uintList)
	require.EqualValues(t, []uint64{1, 5, 8, 4}, arr)
	require.EqualValues(t, nil, err)
}
