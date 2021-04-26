
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

  //console.log(content)

  if (content.Action) {
    content.Action.Class = {
      "Nature": content.Action.CostType.Lumber,
      "Mysticism": content.Action.CostType.Mana,
      "Technology": content.Action.CostType.Iron || content.Action.CostType.Energy,
      "Culture": content.Action.CostType.Food
    }
    if (! (content.Action.Class.Nature || content.Action.Class.Mysticism || content.Action.Class.Technology || content.Action.Class.Culture) ) {
      content.Action.Class.Culture =  true
    }
    content.Action.Effects = []
    content.Action.Keywords = []
    content.Action.RulesTexts = ""

    delete content.Action.CostType

    //console.log(content)
  }
  else if (content.Place) {

    content.Place.Class = {
      "Nature": content.Place.CostType.Lumber,
      "Mysticism": content.Place.CostType.Mana,
      "Technology": content.Place.CostType.Iron || content.Place.CostType.Energy,
      "Culture": content.Place.CostType.Food
    }
    if (! (content.Place.Class.Nature || content.Place.Class.Mysticism || content.Place.Class.Technology || content.Place.Class.Culture) ) {
      content.Place.Class.Culture =  true
    }
    content.Place.Abilities = []
    content.Place.Keywords = []
    content.Place.RulesTexts = ""

    delete content.Place.CostType

  }
  else if (content.Headquarter) {
    content.Headquarter.Delay = 0
    content.Headquarter.Class = {
      "Nature": content.Headquarter.CostType.Lumber,
      "Mysticism": content.Headquarter.CostType.Mana,
      "Technology": content.Headquarter.CostType.Iron || content.Headquarter.CostType.Energy,
      "Culture": content.Headquarter.CostType.Food
    }
    if (! (content.Headquarter.Class.Nature || content.Headquarter.Class.Mysticism || content.Headquarter.Class.Technology || content.Headquarter.Class.Culture) ) {
      content.Headquarter.Class.Culture =  true
    }
    content.Headquarter.Abilities = []
    content.Headquarter.Keywords = []
    content.Headquarter.RulesTexts = ""

    delete content.Headquarter.StartingHandSize
    delete content.Headquarter.Growth
    delete content.Headquarter.Wisdom
    delete content.Headquarter.CostType
  }
  else if (content.Entity) {
    content.Entity.Class = {
      "Nature": content.Entity.CostType.Lumber,
      "Mysticism": content.Entity.CostType.Mana,
      "Technology": content.Entity.CostType.Iron || content.Entity.CostType.Energy,
      "Culture": content.Entity.CostType.Food
    }
    if (! (content.Entity.Class.Nature || content.Entity.Class.Mysticism || content.Entity.Class.Technology || content.Entity.Class.Culture) ) {
      content.Entity.Class.Culture =  true
    }
    content.Entity.Abilities = []
    content.Entity.Keywords = []
    content.Entity.RulesTexts = ""

    delete content.Entity.CostType


    //console.log(content)
  }

  x.Content = btoa(JSON.stringify(content))
  return x
  //console.log(content)
}, genesisnew.app_state.cardservice.card_records)

/*
console.log('users written in the file:', genesisnew.app_state.cardservice.users)
//console.log('addresses:', genesisnew.app_state.cardservice.addresses)
console.log('cards:', R.map(x => {
    return atob(x.Content ? x.Content : '')
  }, genesisnew.app_state.cardservice.card_records))
*/
//ids = R.map(x => x.Owner, genesisnew.app_state.cardservice.card_records)
//console.log(ids)

fs.writeFileSync('./genesis.merged.json', JSON.stringify(genesisnew) , 'utf-8');
