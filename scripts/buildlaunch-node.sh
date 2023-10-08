#!/usr/bin/env bash
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

cd "$parent_path"

docker-compose build
docker-compose up -d
sleep 1
bash register_faucet.sh