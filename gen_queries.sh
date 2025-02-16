set -o errexit
set -o pipefail

#ignite scaffold query card cardId:uint --response card:CardWithImage
# ---ignite scaffold query card_content cardId:uint --response card:Card
#ignite scaffold query user address:string --response user:User
#ignite scaffold query cards owner --response cardIds:uints
#ignite scaffold query match matchId --response match:Match
#ignite scaffold query set setId --response match:SetWithArtwork
#ignite scaffold query sellOffer sellOfferId:uint --response sellOffer:SellOffer
#ignite scaffold query council councilId:uint --response council:Council
#ignite scaffold query server serverId:uint --response server:Server
#ignite scaffold query encounter encounterId:uint --response encounter:Encounter
#ignite scaffold query encounters --response encounters
#ignite scaffold query encounterWithImage --response encounter:EncounterWithImage
#ignite scaffold query encountersWithImage --response encounters
#ignite scaffold query cardchainInfo --response cardAuctionPrice:coin activeSets:uints cardsNumber:uint matchesNumber:uint sellOffersNumber:uint councilsNumber:uint lastCardModified:uint
#ignite scaffold query cardContent cardId:uint --response encounters
#ignite scaffold query setRarityDistribution setId --response current:uints wanted:uints
#ignite scaffold query accountFromZealy zealyId:string --response address:string
#ignite scaffold query votingResults --response lastVotingResults:VotingResults
#ignite scaffold query matches timestampDown:uint timestampUp:uint containsUsers:strings reporter outcome cardsPlayed:uints  --response matches
ignite scaffold query sets status contributors:strings containsCards:uints owner --response sets
