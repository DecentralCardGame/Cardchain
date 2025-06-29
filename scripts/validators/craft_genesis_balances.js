const fs = require('node:fs');
const R = require('ramda');


const testFolder = './snapshots/';
let files = fs.readdirSync(testFolder);

let data = [];
files.forEach(file => {
	let content = fs.readFileSync('snapshots/'+file, "utf8");
	let contentArray = R.map(R.split('\t'), R.split('\n', content))
	data = R.append(contentArray, data)
});

let dataObj = {}
data.forEach(file => {
	let fileObj = {}
	file.forEach(entry => {
		if (dataObj[entry[0]]) dataObj[entry[0]] += parseInt(entry[1])
		else dataObj[entry[0]] = parseInt(entry[1])	
	})
})

saveableData = ["Address\tBalance"]
R.mapObjIndexed((val, key) => saveableData.push(key+"\t"+val), dataObj)
saveableData = saveableData.join("\n")

console.log(saveableData)
fs.writeFileSync("../merged_genesis_balances.tsv", saveableData); 

