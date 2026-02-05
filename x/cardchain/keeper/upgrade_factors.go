package keeper

import (
	"sort"

	"github.com/DecentralCardGame/cardchain/x/cardchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const NERF_FACTOR float32 = 0.9

type normalizedUpgradeFactors struct {
	UpgradeFactor   *types.UpgradeFactor
	NormalizedPicks float32
}

func (k Keeper) UpdateUpgradeCostFactors(ctx sdk.Context) {
	factors := k.UpgradeFactorK.GetAll(ctx)

	var upgradeCandidates []normalizedUpgradeFactors
	for _, factor := range factors {
		if factor.Occurrence >= 3 {
			upgradeCandidates = append(
				upgradeCandidates,
				normalizedUpgradeFactors{
					factor,
					factor.Picks / float32(factor.Occurrence),
				},
			)
		}
	}

	sort.Slice(upgradeCandidates, func(i, j int) bool {
		return upgradeCandidates[i].NormalizedPicks < upgradeCandidates[j].NormalizedPicks
	})

	var totalSum float32
	var weightedSum float32
	n := float32(len(upgradeCandidates))

	if n == 0 {
		k.logger.Warn("no upgrade factors available")
		return
	}

	for idx, candidate := range upgradeCandidates {
		picks := candidate.NormalizedPicks
		totalSum += picks
		i := float32(idx + 1)
		weightedSum += (2*i - n - 1) * picks
	}

	if totalSum == 0 {
		return
	}

	giniCoefficient := weightedSum / (n * totalSum)

	if giniCoefficient <= 0.1 {
		return
	}

	for idx, candidate := range upgradeCandidates {
		buff(candidate, idx <= len(upgradeCandidates)/2)
	}

	for _, factor := range factors {
		factor.Picks = 0
		factor.Occurrence = 0

		k.UpgradeFactorK.Set(ctx, factor.Name, factor)
	}
}

func buff(candidate normalizedUpgradeFactors, buff bool) {
	if candidate.UpgradeFactor.Cost > 0 {
		if buff {
			candidate.UpgradeFactor.Cost *= NERF_FACTOR
		} else {
			candidate.UpgradeFactor.Cost /= NERF_FACTOR
		}
	} else {
		if buff {
			candidate.UpgradeFactor.Cost /= NERF_FACTOR
		} else {
			candidate.UpgradeFactor.Cost *= NERF_FACTOR
		}
	}
}
