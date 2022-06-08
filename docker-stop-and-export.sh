#!/bin/bash

PID=$(docker-compose exec blockchain pidof Cardchaind)
echo PID:$PID

docker-compose exec blockchain pkill Cardchaind

while $(docker-compose exec blockchain pkill -0 Cardchaind); do
	sleep 1
done 

echo "writing backup"

sh ./scripts/backupGenesis.sh
