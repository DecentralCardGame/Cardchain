package keeper

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/cardobject"
	"github.com/DecentralCardGame/cardobject/keywords"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper Yeah the keeper
type Keeper struct {
	cdc                     codec.BinaryCodec // The wire codec for binary encoding/decoding.
	UsersStoreKey           sdk.StoreKey
	CardsStoreKey           sdk.StoreKey
	MatchesStoreKey         sdk.StoreKey
	CollectionsStoreKey     sdk.StoreKey
	InternalStoreKey        sdk.StoreKey
	SellOffersStoreKey      sdk.StoreKey
	PoolsStoreKey           sdk.StoreKey
	CouncilsStoreKey        sdk.StoreKey
	RunningAveragesStoreKey sdk.StoreKey
	paramstore              paramtypes.Subspace

	PoolKeys           []string
	RunningAverageKeys []string

	BankKeeper types.BankKeeper
}

// NewKeeper Constructor for Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	usersStoreKey,
	cardsStoreKey sdk.StoreKey,
	matchesStorekey sdk.StoreKey,
	collectionsStoreKey sdk.StoreKey,
	sellOffersStoreKey sdk.StoreKey,
	poolsStoreKey sdk.StoreKey,
	councilsStoreKey sdk.StoreKey,
	runningAveragesStoreKey sdk.StoreKey,
	internalStoreKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                     cdc,
		UsersStoreKey:           usersStoreKey,
		CardsStoreKey:           cardsStoreKey,
		MatchesStoreKey:         matchesStorekey,
		CollectionsStoreKey:     collectionsStoreKey,
		SellOffersStoreKey:      sellOffersStoreKey,
		PoolsStoreKey:           poolsStoreKey,
		CouncilsStoreKey:        councilsStoreKey,
		RunningAveragesStoreKey: runningAveragesStoreKey,
		InternalStoreKey:        internalStoreKey,
		paramstore:              ps,
		PoolKeys:                []string{PublicPoolKey, WinnersPoolKey, BalancersPoolKey},
		RunningAverageKeys:      []string{Games24ValueKey, Votes24ValueKey},
		BankKeeper:              bankKeeper,
	}
}

// Logger Tendermint logger for logging in the cosmos log
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TransferSchemeToCard Makes a users cardscheme a card
func (k Keeper) TransferSchemeToCard(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)

	var gottenUser types.User
	k.cdc.MustUnmarshal(bz, &gottenUser)

	idPosition := IndexOfId(cardId, gottenUser.OwnedCardSchemes)

	if idPosition >= 0 {
		gottenUser.OwnedPrototypes = append(gottenUser.OwnedPrototypes, cardId)
		gottenUser.OwnedCardSchemes = append(gottenUser.OwnedCardSchemes[:idPosition], gottenUser.OwnedCardSchemes[idPosition+1:]...)

		store.Set(address, k.cdc.MustMarshal(&gottenUser))
	}
}

type candidate struct {
	id    uint64
	votes int64
}

// SetLastCardScheme Sets the current id of the last bought card scheme
func (k Keeper) SetLastVotingResults(ctx sdk.Context, results types.VotingResults) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("lastVotingResults"), k.cdc.MustMarshal(&results))
}

// GetLastVotingResults Returns the current price of the card scheme auction
func (k Keeper) GetLastVotingResults(ctx sdk.Context) (results types.VotingResults) {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("lastVotingResults"))
	k.cdc.MustUnmarshal(bz, &results)
	return
}

// NerfBuffCards Nerfes or buffs certain cards
func (k Keeper) NerfBuffCards(ctx sdk.Context, cardIds []uint64, buff bool) {
	store := ctx.KVStore(k.CardsStoreKey)

	for _, val := range cardIds {
		bz := store.Get(sdk.Uint64ToBigEndian(val))
		var buffCard types.Card
		k.cdc.MustUnmarshal(bz, &buffCard)

		cardobj, err := keywords.Unmarshal(buffCard.Content)
		if err != nil {
			fmt.Println("error on card content:", err, "with card", buffCard.Content)
		}

		buffnerfCost := func(cost *cardobject.CastingCost) {
			update := *cost
			if buff {
				update -= 1
			} else {
				update += 1
			}
			// only apply the buffed/nerfed value if the new value validates without error
			if update.ValidateType(nil) == nil {
				*cost = update
			}
		}

		if cardobj.Action != nil {
			buffnerfCost(&cardobj.Action.CastingCost)
		}
		if cardobj.Entity != nil {
			buffnerfCost(&cardobj.Entity.CastingCost)
		}
		if cardobj.Place != nil {
			buffnerfCost(&cardobj.Place.CastingCost)
		}
		if cardobj.Headquarter != nil {
			updateHealth := cardobj.Headquarter.Health
			updateDelay := cardobj.Headquarter.Delay
			if buff {
				updateDelay -= 1
				updateHealth += 1
			} else {
				updateDelay += 1
				updateHealth -= 1
			}
			if updateDelay.ValidateType(nil) == nil {
				cardobj.Headquarter.Delay = updateDelay
			}
			if updateHealth.ValidateType(nil) == nil {
				cardobj.Headquarter.Health = updateHealth
			}
		}

		cardJSON, _ := json.Marshal(cardobj)
		buffCard.Content = cardJSON

		if buff {
			buffCard.Nerflevel -= 1
		} else {
			buffCard.Nerflevel += 1
		}

		store.Set(sdk.Uint64ToBigEndian(val), k.cdc.MustMarshal(&buffCard))
	}
}

// UpdateBanStatus Bans cards
func (k Keeper) UpdateBanStatus(ctx sdk.Context, newBannedIds []uint64) {
	cardsStore := ctx.KVStore(k.CardsStoreKey)
	usersStore := ctx.KVStore(k.UsersStoreKey)

	// go through all cards and find already marked cards
	iterator := k.GetCardsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var gottenCard types.Card
		k.cdc.MustUnmarshal(iterator.Value(), &gottenCard)

		if gottenCard.Status == types.Status_bannedVerySoon {
			address, _ := sdk.AccAddressFromBech32(gottenCard.Owner)

			// remove the card from the Cards store
			var emptyCard types.Card
			cardsStore.Set(iterator.Key(), k.cdc.MustMarshal(&emptyCard))

			// remove the card from the ownedCards of the owner
			bz2 := usersStore.Get(address)
			var gottenUser types.User
			k.cdc.MustUnmarshal(bz2, &gottenUser)

			idPosition := IndexOfId(binary.BigEndian.Uint64(iterator.Key()), gottenUser.OwnedCardSchemes)
			if idPosition >= 0 {
				gottenUser.OwnedPrototypes = append(gottenUser.OwnedCardSchemes[:idPosition], gottenUser.OwnedCardSchemes[idPosition+1:]...)
				usersStore.Set(address, k.cdc.MustMarshal(&gottenUser))
			} else {
				fmt.Println("trying to delete card id:", binary.BigEndian.Uint64(iterator.Key()), " of owner", address, " but does not exist")
			}
		} else if gottenCard.Status == types.Status_bannedSoon {
			gottenCard.Status = types.Status_bannedVerySoon
			cardsStore.Set(iterator.Key(), k.cdc.MustMarshal(&gottenCard))
		}
	}
}

// GetOPandUPCards Gets OP and UP cards
func (k Keeper) GetOPandUPCards(ctx sdk.Context) (buffbois []uint64, nerfbois []uint64, fairbois []uint64, banbois []uint64) {
	var OPcandidates []candidate
	var UPcandidates []candidate
	var IAcandidates []candidate

	//var votingResults VotingResults
	votingResults := types.NewVotingResults()

	var µUP float64 = 0
	var µOP float64 = 0

	iterator := k.GetCardsIterator(ctx)

	// go through all cards and collect candidates
	for ; iterator.Valid(); iterator.Next() {
		var gottenCard types.Card
		k.cdc.MustUnmarshal(iterator.Value(), &gottenCard)

		id := binary.BigEndian.Uint64(iterator.Key())
		nettoOP := int64(gottenCard.OverpoweredVotes - gottenCard.FairEnoughVotes - gottenCard.UnderpoweredVotes)
		nettoUP := int64(gottenCard.UnderpoweredVotes - gottenCard.FairEnoughVotes - gottenCard.OverpoweredVotes)
		nettoIA := int64(gottenCard.InappropriateVotes - gottenCard.FairEnoughVotes - gottenCard.OverpoweredVotes - gottenCard.UnderpoweredVotes)

		votingResults.TotalFairEnoughVotes += gottenCard.FairEnoughVotes
		votingResults.TotalOverpoweredVotes += gottenCard.OverpoweredVotes
		votingResults.TotalUnderpoweredVotes += gottenCard.UnderpoweredVotes
		votingResults.TotalInappropriateVotes += gottenCard.InappropriateVotes
		votingResults.TotalVotes += gottenCard.FairEnoughVotes + gottenCard.OverpoweredVotes + gottenCard.UnderpoweredVotes + gottenCard.InappropriateVotes

		//fmt.Println("id:",id," - op:",nettoOP," / up:", nettoUP);
		//fmt.Println(gottenCard)

		// all candidates are added to the results log
		if nettoIA > 0 || nettoOP > 0 || nettoUP > 0 {
			votingResults.CardResults = append(votingResults.CardResults, &types.VotingResult{
				CardId:             id,
				FairEnoughVotes:    gottenCard.FairEnoughVotes,
				OverpoweredVotes:   gottenCard.OverpoweredVotes,
				UnderpoweredVotes:  gottenCard.UnderpoweredVotes,
				InappropriateVotes: gottenCard.InappropriateVotes,
				Result:             "fair_enough",
			})

			// sort candidates into the specific arrays
			if nettoIA > 0 {
				IAcandidates = append(IAcandidates, candidate{id: id, votes: nettoIA})
			} else if nettoOP > 0 {
				µOP += float64(nettoOP)
				OPcandidates = append(OPcandidates, candidate{id: id, votes: nettoOP})
			} else if nettoUP > 0 {
				µUP += float64(nettoUP)
				UPcandidates = append(UPcandidates, candidate{id: id, votes: nettoUP})
			}
		}
	}

	// go through all OP candidates and calculate the cutoff value and collect all above this value
	if len(OPcandidates) > 0 {
		// µ is the average, so it must be divided by n, but we can do this only after all cards are counted
		µOP /= float64(len(OPcandidates))

		sort.Slice(OPcandidates, func(i, j int) bool {
			return OPcandidates[i].votes < OPcandidates[j].votes
		})

		var giniOPsum float64
		for i := 1; i <= len(OPcandidates); i++ {
			giniOPsum += float64(OPcandidates[i-1].votes) * float64(2*i-len(OPcandidates)-1)
		}

		giniOP := giniOPsum / float64(len(OPcandidates)*len(OPcandidates)) / µOP
		cutvalue := giniOP * float64(OPcandidates[len(OPcandidates)-1].votes)

		//fmt.Println("µ/gini/cut: ",µOP, giniOP,cutvalue)

		for i := 0; i < len(OPcandidates); i++ {
			if float64(OPcandidates[i].votes) > cutvalue {
				nerfbois = append(nerfbois, OPcandidates[i].id)
			} else {
				fairbois = append(fairbois, OPcandidates[i].id)
			}
		}
	}
	// go through all UP candidates and calculate the cutoff value and collect all above this value
	if len(UPcandidates) > 0 {
		µUP /= float64(len(UPcandidates))

		sort.Slice(UPcandidates, func(i, j int) bool {
			return UPcandidates[i].votes < UPcandidates[j].votes
		})

		var giniUPsum float64
		for i := 1; i <= len(UPcandidates); i++ {
			giniUPsum += float64(UPcandidates[i-1].votes) * float64(2*i-len(UPcandidates)-1)
		}

		giniUP := giniUPsum / float64(len(UPcandidates)*len(UPcandidates)) / µUP
		cutvalue := giniUP * float64(UPcandidates[len(UPcandidates)-1].votes)

		//fmt.Println("µ/gini/cut: ",µUP, giniUP,cutvalue)

		for i := 0; i < len(UPcandidates); i++ {
			if float64(UPcandidates[i].votes) > cutvalue {
				buffbois = append(buffbois, UPcandidates[i].id)
			} else {
				fairbois = append(fairbois, UPcandidates[i].id)
			}
		}
	}
	// go through all IA candidates and collect them (there is no cutoff here)
	if len(IAcandidates) > 0 {
		for i := 0; i < len(IAcandidates); i++ {
			banbois = append(banbois, IAcandidates[i].id)
		}
	}

	// add the result to the voting log
	for i := 0; i < len(votingResults.CardResults); i++ {
		for j := 0; j < len(buffbois); j++ {
			if votingResults.CardResults[i].CardId == buffbois[j] {
				votingResults.CardResults[i].Result = "buff"
			}
		}
		for j := 0; j < len(nerfbois); j++ {
			if votingResults.CardResults[i].CardId == nerfbois[j] {
				votingResults.CardResults[i].Result = "nerf"
			}
		}
		for j := 0; j < len(banbois); j++ {
			if votingResults.CardResults[i].CardId == banbois[j] {
				votingResults.CardResults[i].Result = "ban"
			}
		}
	}

	// and save the log
	k.SetLastVotingResults(ctx, votingResults)

	return
}
