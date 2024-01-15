#!/bin/bash

echo "please give write permission to backup folder:"
sudo chmod o+w backup

PID=$(docker-compose exec blockchain pidof cardchaind)
echo PID:$PID

docker-compose exec blockchain pkill cardchaind

while $(docker-compose exec blockchain pkill -0 cardchaind); do
	sleep 1
done 

echo "writing backup"

# no longer necessary
#sh ./scripts/backupGenesis.sh
