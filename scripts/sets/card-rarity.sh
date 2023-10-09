#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify set id via argument"
  exit 2
fi

Cardchaind tx cardchain set-card-rarity 12 $1 common --from jannik
Cardchaind tx cardchain set-card-rarity 13 $1 common --from jannik
Cardchaind tx cardchain set-card-rarity 14 $1 common --from jannik
Cardchaind tx cardchain set-card-rarity 15 $1 uncommon --from jannik
Cardchaind tx cardchain set-card-rarity 16 $1 rare --from jannik