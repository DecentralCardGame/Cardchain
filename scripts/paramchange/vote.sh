#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify proposal id via argument"
  exit 2
fi

Cardchaind tx gov vote $1 yes --from jannik