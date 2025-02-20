#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify proposal id via argument"
  exit 2
fi

echo "\033[0;36mvoting! (I hope jannik has staked ubpf)\033[0m"

cardchaind tx gov vote $1 yes --from jannik