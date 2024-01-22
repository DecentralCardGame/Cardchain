#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify set id via argument"
  exit 2
fi

cardchaind tx cardchain finalize-set $1 --from jannik --gas auto

cardchaind tx gov submit-legacy-proposal set-proposal 1 10000000ubpf --from jannik --gas auto

cardchaind q gov proposals