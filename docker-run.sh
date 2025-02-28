#!/bin/bash

#set -eo pipefail

echo -e "\033[0;32mfasten your seatbelts\033[0m"
FAUCET_SECRET_KEY="0x6F1f5bd93f3D59d6eed1d5ec40E29C1821029759"
USE_SNAP=false

if [ -z "$FAUCET_SECRET_KEY" ]
then
	echo -e "\033[0;31mNO SECRET KEY FOR FAUCET CONFIGURED! \033[0m"
	exit 1
fi

#we can derive the peer id from the rpc
mapfile -t peers < <(
	jq -r '.peers[]' peer_nodes.json
)

for item in "${peers[@]}"; do
    echo "peers: ${item}"
done

mapfile -t rpcs < <(
	jq -r '.rpcs[]' peer_nodes.json
)

echo "peers" $peers

for i in "${!rpcs[@]}"; do
 	if curl --output /dev/null --silent --head --fail --connect-timeout 5 ${rpcs[$i]}; then
   		echo "URL exists: ${rpcs[$i]}"
 		PEER_ID=$(curl -s ${rpcs[$i]}"/status" | jq -r .result.node_info.id)

 		PEERS=$PEER_ID"@"${peers[$i]}
 		break
 	else
   		echo "not reachable $i"
 	fi
done
if [ -z "$PEERS" ]
then
	echo -e "\033[0;31mNo PEERS available\033[0m"
	exit
fi

SEEDS=""
echo "peers is:" $PEERS
sed -i.bak -e "s/^seeds *=.*/seeds = \"$SEEDS\"/; s/^persistent_peers *=.*/persistent_peers = \"$PEERS\"/" $HOME/.cardchaind/config/config.toml

if [ -z $USE_SNAP ] 
then

	mapfile -t snap_rpcs < <(
	  jq -r '.snap_rpcs[]' peer_nodes.json
	)

	for i in "${snap_rpcs[@]}"; do
	 	if curl --output /dev/null --silent --head --fail --connect-timeout 5 $i; then
	   		echo "URL exists: $i"
	 		SNAP_RPC=$i
	 		break
	 	else
	   		echo "not reachable $i"
	 	fi
	done
	if [ -z "$SNAP_RPC" ]
	then
		echo -e "\033[0;31mNo SNAP_RPC available\033[0m"
		exit
	fi

	LATEST_HEIGHT=$(curl -s $SNAP_RPC/block | jq -r .result.block.header.height)
	echo $LATEST_HEIGHT
	BLOCK_HEIGHT=$((LATEST_HEIGHT)); \
	TRUST_HASH=$(curl -s "$SNAP_RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)
	echo -e "\033[0;36mlatest height: $LATEST_HEIGHT \nblock height: $BLOCK_HEIGHT \ntrust hash: $TRUST_HASH \033[0m"

	sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
	s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP_RPC,$SNAP_RPC\"| ; \
	s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
	s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"|" $HOME/.cardchaind/config/config.toml ; \

fi

# copy genesis.json to nginx server
cp $HOME/.cardchaind/config/genesis.json $HOME/files/genesis.json 

# config pruning
indexer="kv"
pruning="custom"
pruning_keep_recent="100"
pruning_keep_every="0"
pruning_interval="10"

sed -i -e "s/^indexer *=.*/indexer = \"$indexer\"/" $HOME/.cardchaind/config/config.toml
sed -i -e "s/^pruning *=.*/pruning = \"$pruning\"/" $HOME/.cardchaind/config/app.toml
sed -i -e "s/^pruning-keep-recent *=.*/pruning-keep-recent = \"$pruning_keep_recent\"/" $HOME/.cardchaind/config/app.toml
sed -i -e "s/^pruning-keep-every *=.*/pruning-keep-every = \"$pruning_keep_every\"/" $HOME/.cardchaind/config/app.toml
sed -i -e "s/^pruning-interval *=.*/pruning-interval = \"$pruning_interval\"/" $HOME/.cardchaind/config/app.toml

echo -e "\033[0;32mstarting faucet \033[0m"
sed -i -e "s/^SECRET_KEY *=.*/SECRET_KEY = \"$FAUCET_SECRET_KEY\"/" go-faucet-master/.env
cd go-faucet-master
./go-faucet &
echo -e "\033[0;31mfaucet adress: \033[0;36m $(cardchaind keys show alice --address) \033[0;31m must be registered!\033[0m use scripts/register_faucet.sh for that"
echo $(cardchaind keys show alice --address) > /backup/faucetaddress.txt

echo -e "\033[0;32mstarting Blockchain\033[0m"
cardchaind start

# backup area (this will be executed if the cardchaind process is killed)
now=$(date +"%d.%m.%Y")
cardchaind export > /backup/genesis$now.json
echo "BACKUP should be in /backup/genesis$now - don't forget to use migrate_with_data.py script in case you need it"
echo "fail? is backup folder owned by root? (no idea how this happens though)"
