#!/usr/bin/env python3

import sys
import json

args = sys.argv

assert len(args) == 3, f"Error: Syntax: {args[0]} [old_genesis] [new_genesis]"

del_cards = [370, 346, 258]

file_path_old = args[1]
file_path_new = args[2]

with open(file_path_old, "r") as file:
    old_dict = json.load(file)

with open(file_path_new, "r") as file:
    new_dict = json.load(file)

new_dict["app_state"]["cardchain"] = old_dict["app_state"]["cardchain"].copy()
new_dict["app_state"]["cardchain"]["addresses"] = []
new_dict["app_state"]["cardchain"]["users"] = []
for card in del_cards:
    new_dict["app_state"]["cardchain"]["cardRecords"][card] = {}

for idx, addr in enumerate(old_dict["app_state"]["cardchain"]["addresses"]):
    new_dict["app_state"]["cardchain"]["addresses"].append(addr)
    new_dict["app_state"]["cardchain"]["users"].append(old_dict["app_state"]["cardchain"]["users"][idx])
    for i in old_dict["app_state"]["auth"]["accounts"]:
        if i.get("address") == addr :
            new_dict["app_state"]["auth"]["accounts"].append(i)
            break
    for i in old_dict["app_state"]["bank"]["balances"]:
        if i["address"] == addr :
            for idx, coin in enumerate(i["coins"]):
                if coin["denom"] == "ubpf":
                    i["coins"][idx]["amount"] = "5000000"
            new_dict["app_state"]["bank"]["balances"].append(i)
            break

with open(file_path_new, "w") as file:
    json.dump(new_dict, file, indent=2)
