#!/bin/bash

curl -X GET http://127.0.0.1:1317/txs/$1 | jq .
