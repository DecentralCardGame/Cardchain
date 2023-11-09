#!/usr/bin/env bash

Cardchaind config output "JSON"

echo "validator operators:" 
Cardchaind q staking validators | jq -r '.validators[] | .operator_address'

echo "delegators:"
Cardchaind q staking validators | jq -r '.validators[] | .operator_address' | while read object; do
    Cardchaind query staking delegations-to "$object" | jq -r '.delegation_responses[].delegation.delegator_address'
done

