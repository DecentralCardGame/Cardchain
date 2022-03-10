#!/usr/bin/env python3
"""
Imports a old genesis' cards into cardchain
get_card_content has to be build first
"""

import json
import sys
import os
import ctypes
import requests

lib = ctypes.cdll.LoadLibrary('./get_card_content.so')
lib.get_card_content.restype = ctypes.c_char_p
lib.make_add_artwork_request.restype = ctypes.c_char_p
lib.make_save_card_content_request.restype = ctypes.c_char_p

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

        content = lib.get_card_content(card["Content"].encode('utf-8'))
        notes = card['Notes'].encode("utf-8")
        id = i+1
        lib.make_save_card_content_request("Cooler Typ".encode("utf-8"), i+1, content, notes, artist.encode("utf-8"))

        artwork = lib.get_card_content(card["Image"].encode("utf-8")).decode("utf-8")
        full_art = card["FullArt"]
        lib.make_add_artwork_request("Cooler Artist".encode("utf-8"), i+1, artwork.encode("utf-8"), full_art)


if __name__ == "__main__":
    with open(sys.argv[1], "r") as file:
        card_records = json.load(file)["app_state"]["cardservice"]["card_records"]

    main(card_records)
