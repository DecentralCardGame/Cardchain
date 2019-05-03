
curl -XPUT \
  http://127.0.0.1:1317/cardservice/create_user \
  -d '{
  "base_req":{
      "from":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh",
      "password":"asdfasdf",
      "chain_id":"testCardchain",
      "sequence":"0",
      "account_number":"0",
      "gas": "auto",
      "gas_adjustment": "1.5"
    },
    "new_user":"cosmos1a2e728jkct4h45hhcpgz2l8ak9jhw5fkvyyy6k",
    "creator":"cosmos189ujy6ep9pflja9j00tt4pk3q0kqenayf3l7hh"
  }'
