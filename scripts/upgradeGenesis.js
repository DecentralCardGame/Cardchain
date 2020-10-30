
// file system module to perform file operations
const fs = require('fs');
const R = require('ramda');

var genesisold = require('./genesis.old.json');
var genesisnew = require('./genesis.new.json');


console.log('old:', genesisold.app_state)
console.log('new:', genesisnew.app_state)

// TODO double ids for owned cards
// TODO BUY CARD SCHEMES OVERWRITE REAL CARDS
genesisnew.app_state.auth.accounts = genesisold.app_state.auth.accounts
genesisnew.app_state.cardservice.users = genesisold.app_state.cardservice.users
genesisnew.app_state.cardservice.addresses = genesisold.app_state.cardservice.addresses
genesisnew.app_state.cardservice.card_records = genesisold.app_state.cardservice.card_records
/*
genesisnew.app_state.cardservice.users.push(JSON.parse(JSON.stringify(genesisnew.app_state.cardservice.users[0])))
genesisnew.app_state.cardservice.users.push(JSON.parse(JSON.stringify(genesisnew.app_state.cardservice.users[0])))
genesisnew.app_state.cardservice.users[0].OwnedCards = ['1', '2', '3']
genesisnew.app_state.cardservice.users[1].OwnedCards = ['4', '5']
genesisnew.app_state.cardservice.users[2].OwnedCards = R.map(JSON.stringify, R.range(6, 23))
genesisnew.app_state.cardservice.addresses.push(genesisnew.app_state.cardservice.card_records[3].Owner)
genesisnew.app_state.cardservice.addresses.push(genesisnew.app_state.cardservice.card_records[5].Owner)
*/
console.log('merged:', genesisnew.app_state)
//console.log('cards:', genesisnew.app_state.cardservice.card_records)
console.log('users:', genesisnew.app_state.cardservice.users)
console.log('addresses:', genesisnew.app_state.cardservice.addresses)

//ids = R.map(x => x.Owner, genesisnew.app_state.cardservice.card_records)
//console.log(ids)

fs.writeFileSync('./genesis.merged.json', JSON.stringify(genesisnew) , 'utf-8');
