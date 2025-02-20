#!/usr/bin/env bash
if [ $# -lt 1 ]; then
  echo 1>&2 "Please specify set id via argument"
  exit 2
fi

cardchaind tx cardchain add-story-to-set $1 "Once upon a time a great set was created. It was the set of all sets that do not contain themselves. End of story." --from jannik