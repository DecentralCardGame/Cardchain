
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

R.map(function(x) { console.log(x.Owner) }, genesisnew.app_state.cardservice.card_records)


// card model merger
let id = 0

genesisnew.app_state.cardservice.card_records = R.map(x => {
  //console.log('decoded:', atob(x.Content != null ? x.Content : btoa('{}')))
  //console.log(id)
  id++

  let content = JSON.parse(atob(x.Content != null ? x.Content : btoa('{}')))

  let illegalWords = R.map(s => s.charAt(0).toUpperCase() + s.slice(1), [
    "arm",
    ])

  let illegalTags = R.map(s => s.toUpperCase(), [
    "dwarf",
    "engineer",
    "equipment",
    "farm",
    "knight",
    "farm",
    "magic",
    "militant",
    "primitive",
    "range",
    "spiritual"
  ])

  let oldWords = R.map(s => s.charAt(0).toUpperCase() + s.slice(1), [
    "armor",
    ])

  let filterFunction = entry => {
    console.log(entry)
    return R.map(ability => {
      return JSON.parse(R.replace('EQUIPMENT', 'WEAPON', JSON.stringify(ability)))
    }, entry)
  }

  let illegalArm = entry => {
    let isIllegal = false
    entry = R.map(ability => {
      //let checkstring = JSON.stringify(ability)
      let includesKeyword = x => R.any(R.identity, R.map(R.includes(R.__, JSON.stringify(x)), illegalWords))

      if (includesKeyword(ability)) {
        isIllegal = true
        //console.log("Updated:", util.inspect(entry, {showHidden: false, depth: null}))
        return
      }
      return ability
    }, entry)

    return isIllegal
  }

  let oldTags = entry => {
    return R.without(illegalTags, entry)
  }

  let illegalArmor = entry => {
    if (!entry) return false
    let isIllegal = false
    entry = R.map(ability => {

      let includesKeyword = x => R.any(R.identity, R.map(R.includes(R.__, JSON.stringify(x)), oldWords))

      if (includesKeyword(ability)) {
        isIllegal = true
        console.log("Updated:", util.inspect(entry, {showHidden: false, depth: null}))
        return
      }
      return ability
    }, entry)

    return isIllegal
  }

  if (content.Action) {

  }
  else if (content.Place) {

  }
  else if (content.Headquarter) {

  }
  else if (content.Entity) {

  }

  x.Content = btoa(JSON.stringify(content))
  x.FullArt = "true"
  //console.log(x.Owner)
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
