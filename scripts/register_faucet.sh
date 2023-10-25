#!/usr/bin/env bash
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

cd "$parent_path"

mapfile -t rpcs < <(
	jq -r '.rpcs[]' peer_nodes.json
)

for i in "${!rpcs[@]}"; do
 	if curl --output /dev/null --silent --head --fail --connect-timeout 5 ${rpcs[$i]}; then
   		echo "URL exists: ${rpcs[$i]}"
 		RPC=${rpcs[$i]}

 		
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

Cardchaind tx cardchain createuser $(cat ../backup/faucetaddress.txt) faucet --from jannik --gas auto --node tcp://$RPC:26657