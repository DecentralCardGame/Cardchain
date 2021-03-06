
// file system module to perform file operations
const fs = require('fs');
const R = require('ramda');
const util = require('util')
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
let id = 0

genesisnew.app_state.cardservice.card_records = R.map(x => {
  //console.log('decoded:', atob(x.Content != null ? x.Content : btoa('{}')))
  console.log(id)
  id++

  let content = JSON.parse(atob(x.Content != null ? x.Content : btoa('{}')))

  let filterWords = R.map(s => s.charAt(0).toUpperCase() + s.slice(1), [
    "insight",
    "produce",
    ])

  let fixClassIds = R.map(x => x+1,[10, 17, 68, 83, 112, 127, 134, 146, 170, 172, 183, 194, 198, 202, 211, 225, 227, 233, 258, 260])

  let fixAdditionalCostsIds = R.map(x => x+1,[260, 259, 258, 233, 232, 231, 230, 229, 228, 227, 226, 225, 224, 223, 222, 221, 220, 219, 113, 112, 100, 91, 50])

  let filterFunction = entry => {
    entry = R.map(ability => {
      //let checkstring = JSON.stringify(ability)
      includesKeyword = x => R.any(R.identity, R.map(R.includes(R.__, JSON.stringify(x)), filterWords))

      if (includesKeyword(ability)) {
        console.log("Updating:", util.inspect(entry, {showHidden: false, depth: null}))

        let recursiveFilter = (x, name) => {
          if (includesKeyword(name) && R.any(R.identity, R.map(y => y === "ManaAmount", R.keys(x)))) {
            x.ManaAmount = {
              SimpleIntValue: x.ManaAmount
            }
          }
          else if (includesKeyword(name) && R.any(R.identity, R.map(y => y === "WisdomAmount", R.keys(x)))) {
            x.WisdomAmount = {
              SimpleIntValue: x.WisdomAmount
            }
          }
          else {
            R.map(y => recursiveFilter(x[y], y), R.keys(x))
          }
        }

        recursiveFilter(ability, "root")

        console.log("Updated:", util.inspect(entry, {showHidden: false, depth: null}))
      }

      return ability
    }, entry)

    return entry
  }

  if (content.Action) {
    //content.Action.Effects = filterFunction(content.Action.Effects)
    //content.Action.Keywords = filterFunction(content.Action.Keywords)
    //content.Action.RulesTexts = filterFunction(content.Action.RulesTexts)

  }
  else if (content.Place && content.Place.Abilities) {
    //content.Place.Abilities = filterFunction(content.Place.Abilities)
    //content.Place.Keywords = filterFunction(content.Place.Keywords)
    //content.Place.RulesTexts = filterFunction(content.Place.RulesTexts)

  }
  else if (content.Headquarter) {
    //content.Headquarter.Abilities = filterFunction(content.Headquarter.Abilities)
    //content.Headquarter.Keywords = filterFunction(content.Headquarter.Keywords)
    //content.Headquarter.RulesTexts = filterFunction(content.Headquarter.RulesTexts)

  }
  else if (content.Entity) {
    //content.Entity.Abilities = filterFunction(content.Entity.Abilities)
    //content.Entity.Keywords = filterFunction(content.Entity.Keywords)
    //content.Entity.RulesTexts = filterFunction(content.Entity.RulesTexts)
    
  }

  x.Content = btoa(JSON.stringify(content))
  return x
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
