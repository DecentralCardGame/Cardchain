#!/usr/bin/env python3

import sys
import json
import csv
import os

args = sys.argv

assert len(args) == 3, f"Error: Syntax: {args[0]} [old_genesis] [new_genesis]"

__location__ = os.path.realpath(os.path.join(os.getcwd(), os.path.dirname(__file__)))

gameserver_addr = ["cc1z94z55n2rr4rmjf4ea0m7ykgh9a8urwzrlsxt4", "cc1ch66e3f0szxy8q976rsq5y07esmgdqzj70dfpu"]
alpha_creator = "cc14km80077s0hch3sh38wh2hfk7kxfau4456r3ej"
del_cards = []  # [370, 346, 258]

file_path_old = args[1]
file_path_new = args[2]


# early access, boosterpacks, airdrop from zealy

airdrop_accs = []
boosterpack_accs = []
early_access_addr = ["cc14km80077s0hch3sh38wh2hfk7kxfau4456r3ej", "cc1tmhtms6ahkrxltx3hkmmf2dteqj4pv0thwhdxa"]
with open(os.path.join(__location__, "./zealy.tsv"), "r", encoding="utf8") as zealy_file:
    tsv_reader = csv.DictReader(zealy_file, delimiter="\t")
    for entry in tsv_reader:
        airdrop_accs.append((entry["CCAddress"], entry["Airdrop"]))
        boosterpack_accs.append([entry["CCAddress"], entry["BoosterPacks"]])
        if entry["EarlyAccessToGame"] == "TRUE":
            early_access_addr.append(entry["CCAddress"])

genesisAccs = []
# here we load the balances of addresses that start with balances on CC
with open(os.path.join(__location__, "./genesis_balances.tsv"), "r", encoding="utf8") as genesis_file:
    tsv_reader = csv.DictReader(genesis_file, delimiter="\t")
    for entry in tsv_reader:
        genesisAccs.append((entry["Address"], entry["Balance"]))
        # print(f"{genesisAddresses} has {genesisBalances}")

rarities = []
# here we load the table with card rarities
with open(os.path.join(__location__, "./card_rarities.tsv"), "r", encoding="utf8") as rarity_file:
    tsv_reader = csv.DictReader(rarity_file, delimiter="\t")
    for entry in tsv_reader:
        rarities.append((entry["CardId"], entry["Rarity"]))

starters = []
# here we load the table with starter card
with open(os.path.join(__location__, "./card_starters.tsv"), "r", encoding="utf8") as starter_file:
    tsv_reader = csv.DictReader(starter_file, delimiter="\t")
    for entry in tsv_reader:
        starters.append(entry["CardId"])

# this loads the old genesis file
with open(file_path_old, "r") as file:
    old_dict = json.load(file)
    # old_dict = json.loads(file.read().replace("collection", "set").replace("Collection", "Set"))

# this loads the new genesis file
with open(file_path_new, "r") as file:
    new_dict = json.load(file)

# delete all sets
# old_dict["app_state"]["cardchain"]["sets"] = []


params = new_dict["app_state"]["cardchain"]["params"]
new_dict["app_state"]["featureflag"] = old_dict["app_state"].get("featureflag", new_dict["app_state"]["featureflag"])
new_dict["app_state"]["cardchain"] = old_dict["app_state"]["cardchain"].copy()
new_dict["app_state"]["cardchain"]["addresses"] = []
new_dict["app_state"]["cardchain"]["users"] = []

if "RunningAverages" in new_dict["app_state"]["cardchain"]:
    new_dict["app_state"]["cardchain"]["runningAverages"] = new_dict["app_state"]["cardchain"]["RunningAverages"].copy()
    del new_dict["app_state"]["cardchain"]["RunningAverages"]
if "Servers" in new_dict["app_state"]["cardchain"]:
    new_dict["app_state"]["cardchain"]["servers"] = new_dict["app_state"]["cardchain"]["Servers"].copy()
    del new_dict["app_state"]["cardchain"]["Servers"]

for card in del_cards:
    new_dict["app_state"]["cardchain"]["cardRecords"][card] = {}

if new_dict["app_state"]["cardchain"]["cardRecords"]:
    # write rarities into cards
    for card in rarities:
        if card[1] == "C":
            new_dict["app_state"]["cardchain"]["cardRecords"][int(card[0])]["rarity"] = "common"
        if card[1] == "U":
            new_dict["app_state"]["cardchain"]["cardRecords"][int(card[0])]["rarity"] = "uncommon"
        if card[1] == "R":
            new_dict["app_state"]["cardchain"]["cardRecords"][int(card[0])]["rarity"] = "rare"

    # set starter cards
    for card in starters:
        new_dict["app_state"]["cardchain"]["cardRecords"][int(card)]["starterCard"] = True

for param in params:
    if param in old_dict["app_state"]["cardchain"]["params"]:
        params[param] = old_dict["app_state"]["cardchain"]["params"][param]
new_dict["app_state"]["cardchain"]["params"] = params

# set balanceAnchor, for alpha creator true, for all others false, deactivated because already fine for most cards
# for key in new_dict["app_state"]["cardchain"]["cardRecords"]:
#     if key["owner"] == alpha_creator:
#         key["balanceAnchor"] = True
#     else:
#         key["balanceAnchor"] = False

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
    # limit balances to 5k for all old accounts (genesis accs + alice and bob will have more)
    for i in old_dict["app_state"]["bank"]["balances"]:
        if i["address"] == addr:
            for idx, coin in enumerate(i["coins"]):
                # adjust BPFs
                if coin["denom"] == "ubpf":
                    bpf = 0
                    # use real bpf value for genesisAddresses
                    for acc in genesisAccs:
                        if acc[0] == addr:
                            bpf += int(acc[1]) * 1000000
                    # zealy airdrop
                    for acc in airdrop_accs:
                        if acc[0] == addr:
                            bpf += int(acc[1]) * 1000000

                    # use flat value for all others (TODO ON LAUNCH THIS SHOULD BE 0)
                    if bpf < 5000000:
                        bpf = 5000000

                    i["coins"][idx]["amount"] = str(bpf)

                    # give bpf to alpha creator (jannik)
                    if addr == alpha_creator:
                        i["coins"][idx]["amount"] = "1000000000"
                # adjust Credits
                if coin["denom"] == "ucredits":
                    if addr == alpha_creator:
                        i["coins"][idx]["amount"] = "100000000000"
            new_dict["app_state"]["bank"]["balances"].append(i)
            break

# Remove deprecated voteRights from users and more shenanigans like booster packs and early access
for addr, user in zip(new_dict["app_state"]["cardchain"]["addresses"], new_dict["app_state"]["cardchain"]["users"]):

    if "voteRights" in user:
        del user["voteRights"]

    for entry in boosterpack_accs:
        if entry[0] == addr:
            num_packs = int(entry[1])
            for x in range(num_packs):
                user["boosterPacks"].append({'dropRatiosPerPack': ['150', '50', '1'], 'raritiesPerPack': ['4', '2', '1'], 'setId': '1', 'timeStamp': '0'})

    user["earlyAccess"] = user.get(
        "earlyAccess",
        {"active": False, "invitedUser": "", "invitedByUser": ""}
    )  # add earlyAccess
    if addr in early_access_addr:
        user["earlyAccess"]["active"] = True


coinMap = {}
for account in new_dict["app_state"]["bank"]["balances"]:
    for coin in account["coins"]:
        coinMap[coin["denom"]] = coinMap.get(coin["denom"], 0) + int(coin["amount"])

new_dict["app_state"]["bank"]["supply"] = [{"denom": denom, "amount": str(amount)} for denom, amount in coinMap.items()]

with open(file_path_new, "w") as file:
    json.dump(new_dict, file, indent=2)
