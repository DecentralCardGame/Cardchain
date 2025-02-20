#!/usr/bin/env bash

cardchaind config output "JSON"

echo "validator operators:" 
cardchaind q staking validators | jq -r '.validators[] | .operator_address'

echo "delegators:"
cardchaind q staking validators | jq -r '.validators[] | .operator_address' | while read object; do
    cardchaind query staking delegations-to "$object" | jq -r '.delegation_responses[].delegation.delegator_address'
done

