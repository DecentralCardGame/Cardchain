#!/bin/bash

#set -eo pipefail

echo -e "\033[0;32mfasten your seatbelts\033[0m"
FAUCET_SECRET_KEY="0x6F1f5bd93f3D59d6eed1d5ec40E29C1821029759"
CHAIN_ID=Cardchain

if [ -z "$FAUCET_SECRET_KEY" ]
then
	echo -e "\033[0;31mNO SECRET KEY FOR FAUCET CONFIGURED! \033[0m"
	exit 1
fi

#we can derive the peer id from the address and should do so!
NODE_ADDR="lxgr.xyz"
PEER_ID=$(curl -s "http://"$NODE_ADDR":26657/status" | jq -r .result.node_info.id)
SEEDS=""
PEERS=$PEER_ID"@"$NODE_ADDR":26656"
echo "peers is:" $PEERS
sed -i.bak -e "s/^seeds *=.*/seeds = \"$SEEDS\"/; s/^persistent_peers *=.*/persistent_peers = \"$PEERS\"/" $HOME/.Cardchain/config/config.toml

SNAP_RPCs=("http://crowd.rpc.t.stavr.tech:21207"
"https://cardchain-testnet.nodejumper.io:443"
"https://cardchain-rpc.acloud.pp.ua:443")

# for i in "${SNAP_RPCs[@]}"; do
# 	if curl --output /dev/null --silent --head --fail --connect-timeout 5 $i; then
#   		echo "URL exists: $i"
# 		SNAP_RPC=$i
# 		break
# 	else
#   		echo "not reachable $i"
# 	fi
# done

# if [ -z "$SNAP_RPC" ]
# then
# 	echo -e "\033[0;31mNo SNAP_RPC available\033[0m"
# fi

SNAP_RPC="http://lxgr.xyz:26657"
LATEST_HEIGHT=$(curl -s $SNAP_RPC/block | jq -r .result.block.header.height)
echo $LATEST_HEIGHT
BLOCK_HEIGHT=$((LATEST_HEIGHT)); \
TRUST_HASH=$(curl -s "$SNAP_RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)
echo -e "\033[0;36mlatest height: $LATEST_HEIGHT \nblock height: $BLOCK_HEIGHT \ntrust hash: $TRUST_HASH \033[0m"

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP_RPC,$SNAP_RPC\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"|" $HOME/.Cardchain/config/config.toml; \

# config pruning
indexer="kv"
pruning="custom"
pruning_keep_recent="100"
pruning_keep_every="0"
pruning_interval="10"

sed -i -e "s/^indexer *=.*/indexer = \"$indexer\"/" $HOME/.Cardchain/config/config.toml
sed -i -e "s/^pruning *=.*/pruning = \"$pruning\"/" $HOME/.Cardchain/config/app.toml
sed -i -e "s/^pruning-keep-recent *=.*/pruning-keep-recent = \"$pruning_keep_recent\"/" $HOME/.Cardchain/config/app.toml
sed -i -e "s/^pruning-keep-every *=.*/pruning-keep-every = \"$pruning_keep_every\"/" $HOME/.Cardchain/config/app.toml
sed -i -e "s/^pruning-interval *=.*/pruning-interval = \"$pruning_interval\"/" $HOME/.Cardchain/config/app.toml

echo -e "\033[0;32mstarting faucet \033[0m"
sed -i -e "s/^SECRET_KEY *=.*/SECRET_KEY = \"$FAUCET_SECRET_KEY\"/" go-faucet-master/.env
cd go-faucet-master
./go-faucet &
echo -e "\033[0;31mfaucet adress: \033[0;36m $(Cardchaind keys show alice --address) \033[0;31m must be registered!\033[0m"

echo -e "\033[0;32mstarting Blockchain\033[0m"
Cardchaind start

# backup area (this will be executed if the Cardchaind process is killed)
now=$(date +"%d.%m.%Y")
Cardchaind export > /backup/genesis$now.json
echo "BACKUP should be in /backup/genesis$now"
