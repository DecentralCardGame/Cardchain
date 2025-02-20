#!/usr/bin/env bash
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

cd "$parent_path"

cardchaind tx gov submit-legacy-proposal param-change paramchange.json --from jannik

cardchaind q gov proposals