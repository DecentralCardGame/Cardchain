#!/usr/bin/env python3
"""
Imports a old genesis' cards into cardchain
get_card_content has to be build first
"""

import json
import sys
import os
import ctypes
from requester import requester

artist = "Cooler Artist".encode("utf-8")
creator = "Cooler Typ".encode("utf-8")

noah_jannik = {
     "cosmos1aka9p2tc2td923044ve0508xnn8zuaftc2knxd": "cc13jx3az0ajw938t4k0kw2gtyw0vnjs54spedvky".encode("utf-8"), # noah
     "cosmos15ymvugyn9r0h5e44627aclzp9se3dstnwm9syg": "cc14km80077s0hch3sh38wh2hfk7kxfau4456r3ej".encode("utf-8")  # jannik
}

def register_cards(card_records, address_records):
    bid = "80000000000000000000000000000000ucredits".encode("utf-8")
    cards = [c for c in card_records if c["Content"] not in ["", "e30=", None] and c["Image"] not in ["", None]]
    for i, card in enumerate(cards):
        # artist = address_records[card["Owner"]].strip().encode("utf-8")
        requester.make_buy_card_scheme_request(creator, bid)

        content = card["Content"].encode('utf-8')
        notes = card['Notes'].encode("utf-8")
        id = i+8
        requester.make_save_card_content_request(creator, id, content, notes, artist)

        artwork = card["Image"].encode("utf-8")
        full_art = card["FullArt"]
        requester.make_add_artwork_request(artist, id, artwork, full_art)

        transfer_user = noah_jannik.get(card["Owner"])

        if transfer_user is not None:
            requester.make_transfer_card_request(creator, id, transfer_user)

        # requester.make_transfer_card_request(creator, i+1, artist)


def register_users(user_records, address_records):
    user_aliases = [""]
    for user in user_records:
        alias = user["Alias"].strip()
        if alias in user_aliases:
            continue
        os.system(f"Cardchaind keys add '{alias}'")
        print("")
        requester.make_create_user_request("alice".encode("utf-8"), alias.encode("utf-8"))
        user_aliases.append(alias)


if __name__ == "__main__":
    with open(sys.argv[1], "r") as file:
        records = json.load(file)["app_state"]["cardservice"]

    card_records = records["card_records"]
    user_records = records["users"]
    address_records = records["addresses"]
    user_dict = {address_records[i]: user["Alias"] for i, user in enumerate(user_records)}

    #register_users(user_records, address_records)
    register_cards(card_records, user_dict)
