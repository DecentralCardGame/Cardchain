#!/bin/bash

curl -X GET http://127.0.0.1:1317/cardservice/cardList/prototype | jq .

curl -X GET http://127.0.0.1:1317/cardservice/cardList/scheme | jq .

curl -X GET http://127.0.0.1:1317/cardservice/cardList | jq .
