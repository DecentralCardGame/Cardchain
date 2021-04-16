
// file system module to perform file operations
const fs = require('fs');
const R = require('ramda');
const atob = require('atob')
const btoa = require('btoa')

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
*/
//console.log('merged:', genesisnew.app_state)
console.log('users before sanitation:', genesisnew.app_state.cardservice.users)

genesisnew.app_state.cardservice.users = R.map(function(x) { x.OwnedCards = x.OwnedCards ? R.uniq(x.OwnedCards) : []; return x }, genesisnew.app_state.cardservice.users)
genesisnew.app_state.cardservice.users = R.map(function(x) { x.OwnedCardSchemes = x.OwnedCardSchemes ? R.uniq(x.OwnedCardSchemes) : []; return x }, genesisnew.app_state.cardservice.users)
genesisnew.app_state.cardservice.users = R.map(function(x) { x.OwnedCardSchemes = x.OwnedCardSchemes && x.OwnedCards ? R.without(x.OwnedCards, x.OwnedCardSchemes) : x.OwnedCardSchemes; return x }, genesisnew.app_state.cardservice.users)

// card model merger
genesisnew.app_state.cardservice.card_records = R.map(x => {
  //console.log('decoded:', atob(x.Content != null ? x.Content : btoa('{}')))

  let content = JSON.parse(atob(x.Content != null ? x.Content : btoa('{}')))

  if (content.Action) {
    console.log('Action')
    costType = content.Action.CostType
    if (!costType.Energy && !costType.Food && !costType.Lumber && !costType.Mana && !costType.Iron) {
      console.log('fail!')

      costType.Food = true
      console.log(content)
    }
  }
  else if (content.Place) {
    console.log('Place')
    costType = content.Place.CostType
    if (!costType.Energy && !costType.Food && !costType.Lumber && !costType.Mana && !costType.Iron) {
      console.log('fail!')

      costType.Food = true
      console.log(content)
    }
  }
  else if (content.Headquarter) {
    console.log('Headquarter')
    costType = content.Headquarter.CostType
    if (!costType.Energy && !costType.Food && !costType.Lumber && !costType.Mana && !costType.Iron) {
      console.log('fail!')

      costType.Food = true
      console.log(content)
    }
  }
  else if (content.Entity) {
    console.log('Entity')
    costType = content.Entity.CostType
    if (!costType.Energy && !costType.Food && !costType.Lumber && !costType.Mana && !costType.Iron) {
      console.log('fail!')

      costType.Food = true
      console.log(content)
    }
  }

  x.Content = btoa(JSON.stringify(content))
  return x
  //console.log(content)
}, genesisnew.app_state.cardservice.card_records)


console.log('users written in the file:', genesisnew.app_state.cardservice.users)
//console.log('addresses:', genesisnew.app_state.cardservice.addresses)
console.log('cards:', R.map(x => {
    return atob(x.Content ? x.Content : '')
  }, genesisnew.app_state.cardservice.card_records))

//ids = R.map(x => x.Owner, genesisnew.app_state.cardservice.card_records)
//console.log(ids)

fs.writeFileSync('./genesis.merged.json', JSON.stringify(genesisnew) , 'utf-8');
