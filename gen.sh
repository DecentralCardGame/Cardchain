#!/usr/bin/sh

#ignite scaffold message UserCreate newUser:string alias:string
#ignite scaffold message CardSchemeBuy bid:coin --response cardId:uint
#ignite scaffold message CardSaveContent cardId:uint content:string notes:string artist:string balanceAnchor:bool --response airdropClaimed:bool
#ignite scaffold message CardVote vote:SingleVote  --response airdropClaimed:bool
#ignite scaffold message CardTransfer cardId:uint receiver:string
#ignite scaffold message CardDonate cardId:uint amount:coin
#ignite scaffold message CardArtworkAdd cardId:uint image:string fullArt:bool
#ignite scaffold message CardArtistChange cardId:uint artist:string
#ignite scaffold message CouncilRegister
#ignite scaffold message CouncilDeregister
#ignite scaffold message MatchReport matchId:uint playedCardsA:uints \
#    playedCardsB:uints outcome:int
#ignite scaffold message CouncilCreate cardId:uint
#ignite scaffold message MatchReporterAppoint reporter:string
#ignite scaffold message SetCreate name artist storyWriter contributors:strings
#ignite scaffold message SetCardAdd setId:uint cardId:uint
#ignite scaffold message SetCardRemove setId:uint cardId:uint
#ignite scaffold message SetContributorAdd setId:uint user
#ignite scaffold message SetContributorRemove setId:uint user
#ignite scaffold message SetFinalize setId:uint
#ignite scaffold message SetArtworkAdd setId:uint image
#ignite scaffold message SetStoryAdd setId:uint story
#ignite scaffold message BoosterPackBuy setId:uint --response airdropClaimed:bool
#ignite scaffold message SellOfferCreate cardId:uint price:coin
#ignite scaffold message SellOfferBuy sellOfferId:uint
#ignite scaffold message SellOfferRemove sellOfferId:uint
#ignite scaffold message CardRaritySet cardId:uint setId:uint rarity:int
#ignite scaffold message CouncilResponseCommit councilId:uint reponse \
#    suggestion
#ignite scaffold message CouncilResponseReveal councilId:uint reponse:int secret
#ignite scaffold message CouncilRestart councilId:uint
#ignite scaffold message MatchConfirm matchId:uint outcome:int votedCards:array.int
#ignite scaffold message ProfileCardSet cardId:uint
#ignite scaffold message ProfileWebsiteSet website
#ignite scaffold message ProfileBioSet bio
#ignite scaffold message BoosterPackOpen boosterPackId:uint --response cardIds:uints
#ignite scaffold message BoosterPackTransfer boosterPackId:uint receiver
#ignite scaffold message SetStoryWriterSet setId:uint storyWriter
#ignite scaffold message SetArtistSet setId:uint artist
#ignite scaffold message CardVoteMulti votes:array.int  --response airdropClaimed:bool
#ignite scaffold message MatchOpen playerA:string playerB:string playerADeck:uints playerBDeck:uints --response matchId:uint
#ignite scaffold message SetNameSet setId:uint name
#ignite scaffold message ProfileAliasSet alias
#ignite scaffold message EarlyAccessInvite user
#ignite scaffold message EarlyAccessDisinvite user
#ignite scaffold message ZealyConnect zealyId
#ignite scaffold message EncounterCreate name drawlist:uints parameters image
#ignite scaffold message EncounterDo encounterId:uint user
#ignite scaffold message EncounterClose encounterId:uint user won:bool
#ignite scaffold message CardBan cardId:uint --signer authority -y
#ignite scaffold message CardBan cardId:uint --signer authority -y
