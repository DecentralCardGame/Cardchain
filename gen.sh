#!/usr/bin/sh

#ignite scaffold message UserCreate newUser:string alias:string
#ignite scaffold message CardSchemeBuy bid:coin --response cardId:uint
#ignite scaffold message CardSaveContent cardId:uint content:string notes:string artist:string balanceAnchor:bool --response airdropClaimed:bool
ignite scaffold message CardVote vote:SingleVote  --response airdropClaimed:bool
