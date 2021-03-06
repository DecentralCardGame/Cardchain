package keeper

import (
	"fmt"
	"sort"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/DecentralCardGame/Cardchain/x/cardservice/internal/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper types.BankKeeper

	CardsStoreKey    sdk.StoreKey
	UsersStoreKey    sdk.StoreKey
	InternalStoreKey sdk.StoreKey

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the cardservice Keeper
func NewKeeper(coinKeeper types.BankKeeper, cardsStoreKey sdk.StoreKey, usersStoreKey sdk.StoreKey, internalStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: 			coinKeeper,
		CardsStoreKey:    cardsStoreKey,
		UsersStoreKey:    usersStoreKey,
		InternalStoreKey: internalStoreKey,
		cdc:        			cdc,
	}
}

// GetLastCardScheme - get the current id of the last bought card scheme
func (k Keeper) GetLastCardSchemeId(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("lastCardScheme"))
	return binary.BigEndian.Uint64(bz)
}

// SetLastCardScheme - sets the current id of the last bought card scheme
func (k Keeper) SetLastCardSchemeId(ctx sdk.Context, lastId uint64) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("lastCardScheme"), sdk.Uint64ToBigEndian(lastId))
}

// returns the current price of the card scheme auction
func (k Keeper) GetCardAuctionPrice(ctx sdk.Context) sdk.Coin {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("currentCardSchemeAuctionPrice"))
	var price sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &price)
	return price
}

// SetLastCardScheme - sets the current id of the last bought card scheme
func (k Keeper) SetLastVotingResults(ctx sdk.Context, results types.VotingResults) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("lastVotingResults"), k.cdc.MustMarshalBinaryBare(results))
}

// returns the current price of the card scheme auction
func (k Keeper) GetLastVotingResults(ctx sdk.Context) types.VotingResults {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("lastVotingResults"))
	var results types.VotingResults
	k.cdc.MustUnmarshalBinaryBare(bz, &results)
	return results
}

// sets the current price of the card scheme auction
func (k Keeper) SetCardAuctionPrice(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("currentCardSchemeAuctionPrice"), k.cdc.MustMarshalBinaryBare(price))
}

// gets the credits in the public pool
func (k Keeper) GetPublicPoolCredits(ctx sdk.Context) sdk.Coin {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	return amount
}

// sets the credits in the public pool
func (k Keeper) SetPublicPoolCredits(ctx sdk.Context, price sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(price))
}

// subtracts credits from the public pool
func (k Keeper) SubtractPublicPoolCredits(ctx sdk.Context, delta sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	newAmount := amount.Sub(delta)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(newAmount))
}

// adds or subtracts credits from the public pool
func (k Keeper) AddPublicPoolCredits(ctx sdk.Context, delta sdk.Coin) {
	store := ctx.KVStore(k.InternalStoreKey)
	bz := store.Get([]byte("publicPoolCredits"))
	var amount sdk.Coin
	k.cdc.MustUnmarshalBinaryBare(bz, &amount)
	newAmount := amount.Add(delta)
	store.Set([]byte("publicPoolCredits"), k.cdc.MustMarshalBinaryBare(newAmount))
}

func (k Keeper) GetCard(ctx sdk.Context, cardId uint64) types.Card{
	store := ctx.KVStore(k.CardsStoreKey)
	bz := store.Get(sdk.Uint64ToBigEndian(cardId))

	var gottenCard types.Card
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenCard)
	return gottenCard
}

func (k Keeper) SetCard(ctx sdk.Context, cardId uint64, newCard types.Card) {
	store := ctx.KVStore(k.CardsStoreKey)
	store.Set(sdk.Uint64ToBigEndian(cardId), k.cdc.MustMarshalBinaryBare(newCard))
}

func (k Keeper) GetUser(ctx sdk.Context, address sdk.AccAddress) types.User {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)

	var gottenUser types.User
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenUser)
	return gottenUser
}

func (k Keeper) SetUser(ctx sdk.Context, address sdk.AccAddress, userData types.User) {
	store := ctx.KVStore(k.UsersStoreKey)
	store.Set(address, k.cdc.MustMarshalBinaryBare(userData))
}


func (k Keeper) AddVoteRightsToAllUsers(ctx sdk.Context, expireBlock int64) {
	votingRights := k.GetVoteRightToAllCards(ctx, expireBlock)

	userStore := ctx.KVStore(k.UsersStoreKey)

	userIterator := sdk.KVStorePrefixIterator(userStore, nil)

	for ; userIterator.Valid(); userIterator.Next() {
		var user types.User
		k.cdc.MustUnmarshalBinaryBare(userIterator.Value(), &user)
		user.VoteRights = votingRights
		userStore.Set(userIterator.Key(), k.cdc.MustMarshalBinaryBare(user))
	}

	userIterator.Close()
}

func (k Keeper) GetVoteRightToAllCards(ctx sdk.Context, expireBlock int64) []types.VoteRight {
	cardStore := ctx.KVStore(k.CardsStoreKey)
	cardIterator := sdk.KVStorePrefixIterator(cardStore, nil)

	votingRights := []types.VoteRight{}

	for ; cardIterator.Valid(); cardIterator.Next() {
		// here only give right if card is not a scheme or banished
		var gottenCard types.Card
		k.cdc.MustUnmarshalBinaryBare(cardIterator.Value(), &gottenCard)

		if gottenCard.Status == "permanent" || gottenCard.Status == "trial" || gottenCard.Status == "prototype" {
			right := types.NewVoteRight(binary.BigEndian.Uint64(cardIterator.Key()), expireBlock)
			votingRights = append(votingRights, right)
		}
	}
	cardIterator.Close()

	return votingRights
}

func (k Keeper) RemoveVoteRight(ctx sdk.Context, userAddress sdk.AccAddress, rightsIndex int) {
	user := k.GetUser(ctx, userAddress)
	user.VoteRights[rightsIndex] = user.VoteRights[len(user.VoteRights)-1]
	//user.VoteRights[len(user.VoteRights)-1] = null
	user.VoteRights = user.VoteRights[:len(user.VoteRights)-1]
	userStore := ctx.KVStore(k.UsersStoreKey)
	userStore.Set(userAddress, k.cdc.MustMarshalBinaryBare(user))
}

func (k Keeper) SetUserName(ctx sdk.Context, address sdk.AccAddress, name string) {
	store := ctx.KVStore(k.UsersStoreKey)
	gottenUser := k.GetUser(ctx, address)
	gottenUser.Alias = name
	store.Set(address, k.cdc.MustMarshalBinaryBare(gottenUser))
}

func (k Keeper) InitUser(ctx sdk.Context, address sdk.AccAddress, alias string) {
	store := ctx.KVStore(k.UsersStoreKey)

	if alias == "" {
		alias = "newbie"
	}
	newUser := types.NewUser()
	newUser.Alias = alias
	k.CoinKeeper.AddCoins(ctx, address, sdk.Coins{sdk.NewInt64Coin("credits", 10000)})
	const votingRightsExpirationTime = 86000
	newUser.VoteRights = k.GetVoteRightToAllCards(ctx, ctx.BlockHeight()+votingRightsExpirationTime)		// TODO this might be a good thing to remove later, so that sybil voting is not possible

	store.Set(address, k.cdc.MustMarshalBinaryBare(newUser))
}

func (k Keeper) CreateUser(ctx sdk.Context, newUser sdk.AccAddress, alias string) types.User {
	// check if user already exists
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(newUser)

	if bz == nil {
		k.InitUser(ctx, newUser, alias)
	}
	return k.GetUser(ctx, newUser)
}

func (k Keeper) AddOwnedCardScheme(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)

	var gottenUser types.User
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenUser)

	gottenUser.OwnedCardSchemes = append(gottenUser.OwnedCardSchemes, cardId)

	store.Set(address, k.cdc.MustMarshalBinaryBare(gottenUser))
}

func (k Keeper) TransferSchemeToCard(ctx sdk.Context, cardId uint64, address sdk.AccAddress) {
	store := ctx.KVStore(k.UsersStoreKey)
	bz := store.Get(address)

	var gottenUser types.User
	k.cdc.MustUnmarshalBinaryBare(bz, &gottenUser)

	idPosition := indexOfId(cardId, gottenUser.OwnedCardSchemes)

	if idPosition >= 0 {
			gottenUser.OwnedCards = append(gottenUser.OwnedCards, cardId)
			gottenUser.OwnedCardSchemes = append(gottenUser.OwnedCardSchemes[:idPosition], gottenUser.OwnedCardSchemes[idPosition+1:]...)

			store.Set(address, k.cdc.MustMarshalBinaryBare(gottenUser))
	}
}

func indexOfId(cardID uint64, cards []uint64) int {
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

func (k Keeper) GetVoteRights(ctx sdk.Context, voter sdk.AccAddress) []types.VoteRight {
	user := k.GetUser(ctx, voter)
	return user.VoteRights
}

func (k Keeper) SetVoteRights(ctx sdk.Context, voteRights []types.VoteRight, voter sdk.AccAddress) {
	store := ctx.KVStore(k.UsersStoreKey)
	store.Set(voter, k.cdc.MustMarshalBinaryBare(voteRights))
}

type candidate struct {
    id uint64
    votes int64
}

func (k Keeper) GetOPandUPCards(ctx sdk.Context) ([]uint64, []uint64, []uint64, []uint64) {
	var OPcandidates []candidate
	var UPcandidates []candidate
	var IAcandidates []candidate

	var nerfbois []uint64
	var buffbois []uint64
	var fairbois []uint64
	var banbois []uint64

	//var votingResults VotingResults
	votingResults := types.NewVotingResults()

	var µUP float64 = 0
	var µOP float64 = 0

	iterator := k.GetCardsIterator(ctx)

	// go through all cards and collect candidates
	for ; iterator.Valid(); iterator.Next() {
		var gottenCard types.Card
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenCard)

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
		if (nettoIA > 0 || nettoOP > 0 || nettoUP > 0) {
			votingResults.CardResults = append(votingResults.CardResults, types.VotingResult{
				CardId: 						id,
				FairEnoughVotes:		gottenCard.FairEnoughVotes,
				OverpoweredVotes:		gottenCard.OverpoweredVotes,
				UnderpoweredVotes:	gottenCard.UnderpoweredVotes,
				InappropriateVotes:	gottenCard.InappropriateVotes,
				Result:							"fair enough",
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
	if(len(OPcandidates) > 0) {
			// µ is the average, so it must be divided by n, but we can do this only after all cards are counted
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

			//fmt.Println("µ/gini/cut: ",µOP, giniOP,cutvalue)

			for i := 0; i < len(OPcandidates); i++ {
				if float64(OPcandidates[i].votes) > cutvalue {
					nerfbois = append(nerfbois, OPcandidates[i].id)
				}	else {
					fairbois = append(fairbois, OPcandidates[i].id)
				}
			}
	}
	// go through all UP candidates and calculate the cutoff value and collect all above this value
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
	if(len(IAcandidates) > 0) {
		for i := 0; i < len(IAcandidates); i++ {
			banbois = append(banbois, IAcandidates[i].id)
		}
	}

	// add the result to the voting log
	for i := 1; i < len(votingResults.CardResults); i++ {
		for j := 1; j < len(buffbois); j++ {
			if (votingResults.CardResults[i].CardId == buffbois[j] ) {
				votingResults.CardResults[i].Result = "buff"
			}
		}
		for j := 1; j < len(nerfbois); j++ {
			if (votingResults.CardResults[i].CardId == nerfbois[j] ) {
				votingResults.CardResults[i].Result = "nerf"
			}
		}
		for j := 1; j < len(banbois); j++ {
			if (votingResults.CardResults[i].CardId == banbois[j] ) {
				votingResults.CardResults[i].Result = "ban"
			}
		}
	}

	// and save the log
	k.SetLastVotingResults(ctx, votingResults)

	return buffbois, nerfbois, fairbois, banbois
}

func (k Keeper) NerfBuffCards(ctx sdk.Context, cardIds []uint64, buff bool) {
	store := ctx.KVStore(k.CardsStoreKey)

	for _, val := range cardIds {
		bz := store.Get(sdk.Uint64ToBigEndian(val))
		var buffCard types.Card
		k.cdc.MustUnmarshalBinaryBare(bz, &buffCard)

		if buff {
			buffCard.Nerflevel -= 1
		} else {
			buffCard.Nerflevel += 1
		}

		store.Set(sdk.Uint64ToBigEndian(val), k.cdc.MustMarshalBinaryBare(buffCard))
	}
}

func (k Keeper) RemoveCards(ctx sdk.Context, cardIds []uint64) {
	cardsStore := ctx.KVStore(k.CardsStoreKey)
	usersStore := ctx.KVStore(k.UsersStoreKey)

	// go through all cardIds that should be removed
	for _, id := range cardIds {
		var removeCard types.Card
		bz := cardsStore.Get(sdk.Uint64ToBigEndian(id))
		k.cdc.MustUnmarshalBinaryBare(bz, &removeCard)
		address := removeCard.Owner

		// remove the card from the Cards store
		var emptyCard types.Card
		cardsStore.Set(sdk.Uint64ToBigEndian(id), k.cdc.MustMarshalBinaryBare(emptyCard))

		// remove the card from the ownedCards of the owner
		bz2 := usersStore.Get(address)
		var gottenUser types.User
		k.cdc.MustUnmarshalBinaryBare(bz2, &gottenUser)

		idPosition := indexOfId(id, gottenUser.OwnedCardSchemes)
		if idPosition >= 0 {
				gottenUser.OwnedCards = append(gottenUser.OwnedCardSchemes[:idPosition], gottenUser.OwnedCardSchemes[idPosition+1:]...)
				usersStore.Set(address, k.cdc.MustMarshalBinaryBare(gottenUser))
		} else {
			fmt.Println("trying to delete card id:",id," of owner",address," but does not exist");
		}
	}
}

func (k Keeper) ResetAllVotes(ctx sdk.Context) {
	store := ctx.KVStore(k.CardsStoreKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)

	for ; iterator.Valid(); iterator.Next() {
		var resetCard types.Card
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &resetCard)

		resetCard.FairEnoughVotes = 0
		resetCard.OverpoweredVotes = 0
		resetCard.UnderpoweredVotes = 0

		store.Set(iterator.Key(), k.cdc.MustMarshalBinaryBare(resetCard))
	}
}

func (k Keeper) GetCardsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.CardsStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetAllCards(ctx sdk.Context) []types.Card {
	var allCards []types.Card
	iterator := k.GetCardsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenCard types.Card
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenCard)

		allCards = append(allCards, gottenCard)
	}
	return allCards
}

func (k Keeper) GetUsersIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.UsersStoreKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetAllUsers(ctx sdk.Context) ([]types.User, []sdk.AccAddress) {
	var allUsers []types.User
	var allAddresses []sdk.AccAddress

	iterator := k.GetUsersIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		var gottenUser types.User
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &gottenUser)

		allUsers = append(allUsers, gottenUser)
		allAddresses = append(allAddresses, iterator.Key())
	}
	return allUsers, allAddresses
}
