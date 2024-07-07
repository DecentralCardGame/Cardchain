const fs = require('node:fs');
const R = require('ramda');


const testFolder = './snapshots/';
let files = fs.readdirSync(testFolder);

let data = [];
files.forEach(file => {
	let content = fs.readFileSync('snapshots/'+file, "utf8");
	let contentArray = R.map(R.split('\t'), R.split('\n', content))
	data.push(contentArray)
});

//console.log(R.mergeAll(data))

let saveableData = []
R.forEachObjIndexed(entry => {
		console.log(entry)
		saveableData.push(entry)
	},
	R.mergeAll(data)
)
saveableData = R.join("\n", R.map(R.join("	"), saveableData))

console.log(saveableData)
fs.writeFileSync("genesis_balances.tsv", JSON.stringify(saveableData)); 

