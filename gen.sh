#!/usr/bin/sh

#ignite scaffold message UserCreate newUser:string alias:string
#ignite scaffold message CardSchemeBuy bid:coin --response cardId:uint
#ignite scaffold message CardSaveContent cardId:uint content:string notes:string artist:string balanceAnchor:bool --response airdropClaimed:bool
ignite scaffold message CardVote vote:SingleVote  --response airdropClaimed:bool
ignite scaffold message CardTransfer cardId:uint receiver:string
ignite scaffold message CardDonate cardId:uint amount:coin
ignite scaffold message CardArtworkAdd cardId:uint image:string fullArt:bool
ignite scaffold message CardArtistChange cardId:uint artist:string
ignite scaffold message CouncilRegister
ignite scaffold message CouncilDeregister
ignite scaffold message MatchReport matchId:uint playedCardsA:uints \
    playedCardsB:uints outcome:Outcome
