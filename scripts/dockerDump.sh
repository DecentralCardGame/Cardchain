#!/bin/bash

docker-compose exec blockchain pkill -f csd; csd export --for-zero-height > genesis.json
