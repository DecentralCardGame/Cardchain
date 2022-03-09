#!/usr/bin/env python3

import json
import sys
import os

def main(card_records):
    artist = os.popen('Cardchaind keys show "Cooler Artist" --address').read().strip()
    print(artist)
    for i, card in enumerate(card_records):
        if card["Content"] in ["", "e30="] or card["Image"] == "":
            continue
        ret = os.system('Cardchaind tx cardchain buy-card-scheme 8000000000ucredits --from "Cooler Typ"')
        print("")
        if ret != 0:
            raise KeyboardInterrupt

        content = os.popen(f'get_card_content {card["Content"]}').read()
        notes = card['Notes']
        id = i+1
        command = f"Cardchaind tx cardchain save-card-content \"{id}\" '{content}' \"{notes}\" {artist} --from 'Cooler Typ'"
        print(command)
        ret = os.system(command)
        if ret != 0:
            raise KeyboardInterrupt

        artwork = os.popen(f'get_card_content {card["Image"]}').read()
        full_art = card["FullArt"]
        command = f"Cardchaind tx cardchain add-artwork \"{id}\" '{artwork}' \"{full_art}\" --from 'Cooler Artist'"
        print(command)
        ret = os.system(command)
        if ret != 0:
            raise KeyboardInterrupt






with open(sys.argv[1], "r") as file:
    card_records = json.load(file)["app_state"]["cardservice"]["card_records"]

main(card_records)
