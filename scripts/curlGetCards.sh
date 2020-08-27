#!/bin/bash
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo -e ${GREEN}all cards:${NC}
curl -X GET http://127.0.0.1:1317/cardservice/cardList | jq .

echo -e ${GREEN}owned by alice:${NC}
curl -X GET 'http://127.0.0.1:1317/cardservice/cardList?status=&owner='$(cscli keys show alice --address)'&' | jq .

echo -e ${GREEN}prototype, owned by alice, nameContains Archer:${NC}
curl -X GET 'http://127.0.0.1:1317/cardservice/cardList?status=prototype&owner='$(cscli keys show alice --address)'&nameContains=Archer' | jq .

echo -e ${GREEN}nameContains fuxer:${NC}
curl -X GET 'http://127.0.0.1:1317/cardservice/cardList?&nameContains=fuxer' | jq .
