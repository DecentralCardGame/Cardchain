package cardservice

import (
	"fmt"
	"sort"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type candidate struct {
    id uint64
    votes uint64
}

// handle donate to card message
func UpdateNerfLevels(ctx sdk.Context, keeper Keeper) sdk.Result {

	var OPcandidates []candidate
	var UPcandidates []candidate

	var nerfbois []uint64
	var buffbois []uint64
	var fairbois []uint64

	var µUP float64 = 0
	var µOP float64 = 0

	//analyse all cards
	iterator := keeper.GetCardsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {

		var gottenCard Card
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenCard)

		id := binary.BigEndian.Uint64(iterator.Key())
		nettoOP := gottenCard.OverpoweredVotes - gottenCard.FairEnoughVotes - gottenCard.UnderpoweredVotes
		nettoUP := gottenCard.UnderpoweredVotes - gottenCard.FairEnoughVotes - gottenCard.OverpoweredVotes

		fmt.Println(id);
		fmt.Println(gottenCard)

		if nettoOP > 0 {
			µOP += float64(nettoOP)
			OPcandidates = append(OPcandidates, candidate{id: id, votes: nettoOP})

		} else if nettoUP > 0 {
			µUP += float64(nettoUP)
			UPcandidates = append(UPcandidates, candidate{id: id, votes: nettoUP})

		}
	}

	// µ is actually the average, so it must be divided by n
	if(len(OPcandidates) > 0) {
			µOP /= float64(len(OPcandidates))

			sort.Slice(OPcandidates, func(i, j int) bool {
				return OPcandidates[i].votes < OPcandidates[j].votes
			})

			var giniOPsum float64
			for i := 1; i <= len(OPcandidates); i++ {
					giniOPsum += float64(OPcandidates[i-1].votes) * float64(2*i - len(OPcandidates) - 1)
			}

			giniOP := giniOPsum/float64(len(OPcandidates)*len(OPcandidates))/µOP
			cutvalue := giniOP*float64(OPcandidates[len(OPcandidates)-1].votes)

			fmt.Println(µOP, giniOP)
			fmt.Println(cutvalue)

			for i := 0; i < len(OPcandidates); i++ {
				if float64(OPcandidates[i].votes) > cutvalue {
					nerfbois = append(nerfbois, OPcandidates[i].id)
				}	else {
					fairbois = append(fairbois, OPcandidates[i].id)
				}
			}

	}
	if(len(UPcandidates) > 0) {
			µUP /= float64(len(UPcandidates))

			sort.Slice(UPcandidates, func(i, j int) bool {
				return UPcandidates[i].votes < UPcandidates[j].votes
			})

			var giniUPsum float64
			for i := 1; i <= len(UPcandidates); i++ {
					giniUPsum += float64(UPcandidates[i-1].votes) * float64(2*i - len(UPcandidates) - 1)
			}

			giniUP := giniUPsum/float64(len(UPcandidates)*len(UPcandidates))/µUP
			cutvalue := giniUP*float64(UPcandidates[len(UPcandidates)-1].votes)

			fmt.Println(µUP, giniUP)
			fmt.Println(cutvalue)

			for i := 0; i < len(UPcandidates); i++ {
				if float64(UPcandidates[i].votes) > cutvalue {
					buffbois = append(buffbois, UPcandidates[i].id)
				} else {
					fairbois = append(fairbois, UPcandidates[i].id)
				}
			}
	}

	fmt.Println("buff:")
	fmt.Println(buffbois)
	keeper.NerfBuffCards(ctx, buffbois, true)
	fmt.Println("nerf:")
	fmt.Println(nerfbois)
	keeper.NerfBuffCards(ctx, nerfbois, false)

	return sdk.Result{}
}
