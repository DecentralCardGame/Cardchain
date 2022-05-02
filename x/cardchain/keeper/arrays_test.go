package keeper_test

import (
	"testing"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/stretchr/testify/require"
)

func TestArrays(t *testing.T) {
	// UintItemInArr
	uintList := []uint64{1, 5, 8, 9, 4}

	// PopItemFromArr
	arr, err := keeper.PopItemFromArr(9, uintList)
	require.EqualValues(t, []uint64{1, 5, 8, 4}, arr)
	require.EqualValues(t, nil, err)
}
