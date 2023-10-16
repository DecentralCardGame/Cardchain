#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify set id via argument"
  exit 2
fi

Cardchaind tx cardchain add-card-to-set $1 12  --from jannik
Cardchaind tx cardchain add-card-to-set $1 13  --from jannik
Cardchaind tx cardchain add-card-to-set $1 14  --from jannik
Cardchaind tx cardchain add-card-to-set $1 15  --from jannik
Cardchaind tx cardchain add-card-to-set $1 16  --from jannik
