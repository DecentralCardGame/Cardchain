# Crowd Control Cardchain

This blockchain is based on Cosmos SDK. Current version is 0.38.	
The current features are:
- Buy card scheme - ongoing auction to sell rights to create new cards  
- Save content of a card - write content to a card scheme  
- Transfer card - change ownership of a card  
- Vote on a card - give your opinion if a card is overpowered, underpowered or fair enough  
- Donate to a card - donate credits on a card to incentivize voting  

All features are shown as examples in `/scripts`, REST and CLI examples are given. The same applies for query commands.

A testnet is live and our [website](https://www.crowdcontrol.network) connects to it.  
If you want to join the testnet or get in contact (dev with us): Join us at https://discord.gg/yPA3aKe 

## Building and running

Prerequisites: Go 1.13, some scripts require jq or JS

```bash
make intall  
sh scripts/setup.sh  
csd start
  
sh scripts/startRESTserverLocal.sh
```

The last command is only necessary if you want to run the curl scripts.

The blockchain can also be started in a Docker-container:

```bash
docker-compose build  
docker-compose up
```  

## Upgrading with data
```bash
df -h (no space on device = next step will fail)
sh docker-stop-and-export.sh  
mv blockchain-data/genesis.json scripts/genesis.old.json  
git pull  
docker-compose build  
docker-compose up -d  
sh docker-stop-and-export.sh  
mv blockchain-data/genesis.json scripts/genesis.new.json  
node scripts/upgradeGenesis.js  
cp genesis.merged.json blockchain-data/genesis.json  
docker-compose up -d  
```
 
