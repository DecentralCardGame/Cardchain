#!/usr/bin/env python3

import sys
import json
import csv

args = sys.argv

assert len(args) == 3, f"Error: Syntax: {args[0]} [old_genesis] [new_genesis]"

gameserver_addr = ["cc1z94z55n2rr4rmjf4ea0m7ykgh9a8urwzrlsxt4", "cc1ch66e3f0szxy8q976rsq5y07esmgdqzj70dfpu"]
alpha_creator = "cc14km80077s0hch3sh38wh2hfk7kxfau4456r3ej"
del_cards = [] #[370, 346, 258]

file_path_old = args[1]
file_path_new = args[2]

with open("./genesis_balances.tsv", "r", encoding="utf8") as genesis_file:
    tsv_reader = csv.DictReader(genesis_file, delimiter="\t")
    for entry in tsv_reader:
        addr = entry["Address"]
        balance = entry["Balance"]
        print(f"{addr} has {balance}")

with open(file_path_old, "r") as file:
	old_dict = json.load(file)

with open(file_path_new, "r") as file:
	new_dict = json.load(file)

new_dict["app_state"]["cardchain"] = old_dict["app_state"]["cardchain"].copy()
new_dict["app_state"]["cardchain"]["addresses"] = []
new_dict["app_state"]["cardchain"]["users"] = []
for card in del_cards:
	new_dict["app_state"]["cardchain"]["cardRecords"][card] = {}

params = new_dict["app_state"]["cardchain"]["params"]

for param in params:
	if param in old_dict["app_state"]["cardchain"]["params"]:
		params[param] = old_dict["app_state"]["cardchain"]["params"][param]
new_dict["app_state"]["cardchain"]["params"] = params

# set balanceAnchor, for alpha creator true, for all others false
for key in new_dict["app_state"]["cardchain"]["cardRecords"]:
	if key["owner"] == alpha_creator:
		key["balanceAnchor"] = True
	else:
		key["balanceAnchor"] = False


for idx, addr in enumerate(old_dict["app_state"]["cardchain"]["addresses"]):
	# set reportmatches to true for gameserver addresses
	if addr in gameserver_addr:
		old_dict["app_state"]["cardchain"]["users"][idx]["ReportMatches"] = True

	new_dict["app_state"]["cardchain"]["addresses"].append(addr)
	new_dict["app_state"]["cardchain"]["users"].append(old_dict["app_state"]["cardchain"]["users"][idx])
	for i in old_dict["app_state"]["auth"]["accounts"]:
		if i.get("address") == addr:
			new_dict["app_state"]["auth"]["accounts"].append(i)
			break
	# limit balances to 5k for all old accounts (only alice and bob will have more)
	for i in old_dict["app_state"]["bank"]["balances"]:
		if i["address"] == addr:
			for idx, coin in enumerate(i["coins"]):
				if coin["denom"] == "ubpf":
					i["coins"][idx]["amount"] = "5000000"
			new_dict["app_state"]["bank"]["balances"].append(i)
			break


with open(file_path_new, "w") as file:
	json.dump(new_dict, file, indent=2)
