#!/bin/bash

curl -X GET http://localhost:1317/auth/accounts/$(cscli keys show alice -a) | jq '.'
