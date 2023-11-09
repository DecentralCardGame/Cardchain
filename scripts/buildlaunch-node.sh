#!/usr/bin/env bash
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

cd "$parent_path"

docker-compose build --no-cache
docker-compose up -d
sleep 1
bash register_faucet.sh