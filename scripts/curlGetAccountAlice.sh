curl -XGET http://127.0.0.1:1317/auth/accounts/$(cscli keys show alice --address) | jq .
