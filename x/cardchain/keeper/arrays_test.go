package keeper_test

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/stretchr/testify/require"
)

func TestArrays(t *testing.T) {
	// UintItemInList
	uintList := []uint64{1, 5, 8, 9, 4}
	require.EqualValues(t, true, keeper.UintItemInList(5, uintList))
	require.EqualValues(t, false, keeper.UintItemInList(10, uintList))

	// stringList
	stringList := []string{"Hello", "world", "I", "hope", "this", "works"}
	require.EqualValues(t, true, keeper.StringItemInList("hope", stringList))
	require.EqualValues(t, false, keeper.StringItemInList("not", stringList))

	// IndexOfId
	require.EqualValues(t, 2, keeper.IndexOfId(8, uintList))
	require.EqualValues(t, -1, keeper.IndexOfId(10, uintList))
	require.EqualValues(t, -1, keeper.IndexOfId(8, nil))

	// UintPopElementFromArr
	arr, err := keeper.UintPopElementFromArr(9, uintList)
	require.EqualValues(t, []uint64{1, 5, 8, 4}, arr)
	require.EqualValues(t, nil, err)
}
