#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify set id via argument"
  exit 2
fi

Cardchaind tx cardchain finalize-set $1 --from jannik

Cardchaind tx cardchain submit-set-proposal $1 10000000ubpf --from jannik

Cardchaind q gov proposals