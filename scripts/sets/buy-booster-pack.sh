#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify set id via argument"
  exit 2
fi

cardchaind tx cardchain buy-booster-pack 1 --from jannik --gas auto