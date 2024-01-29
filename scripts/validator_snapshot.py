import requests
import bech32
import csv
import hashlib
import base64
import bech32 
from datetime import datetime
import sys
import os

api = "http://202.61.225.157:1317/"
pagination_limit = 150
uptime_percentage = 0.8
signing_window = 30000
USAGE = f"Usage: python {sys.argv[0]} [airdrop_amount]"
subfolder_name = "genesis_balances"

def get_validator_info(api_url, params, signing_info, airdrop_amount):
    try:
        response = requests.get(api_url, params=params)

        # Check if the request was successful (status code 200)
        if response.status_code == 200:
            data = response.json()

            # Extract ccvaloper and pubkey and convert them to address and ccvalcons
            validators_info = []
            for validator in data.get("validators", []):
                valoperaddress = validator.get("operator_address")
                address = convert_valoper_address(valoperaddress)
                consensus_pubkey = validator.get("consensus_pubkey", {}).get("key")
                valconsaddress = convert_pubkey_valcons(consensus_pubkey)

                # filter validators by uptime by matching signing_info to validator_info
                for validator in signing_info:
                    if valconsaddress == validator[0] and ((1 - int(validator[1])/signing_window)> uptime_percentage):
                        validators_info.append([address, airdrop_amount])
                
                
            return validators_info
        else:
            print(f"Error: {response.status_code} - {response.text}")
            return None

    except requests.RequestException as e:
        print(f"An error occurred: {e}")
        return None

def get_signing_info(api_url, params):
    try:
        response = requests.get(api_url, params=params)

        # Check if the request was successful (status code 200)
        if response.status_code == 200:
            data = response.json()

            # Extracting ccvalcons and missed blocks, this is needed to evaluate uptime of validators
            signing_info = []
            for validator in data.get("info", []):
                address = validator.get("address")
                missed_blocks_counter = validator.get("missed_blocks_counter")
                signing_info.append([address, missed_blocks_counter])

            return signing_info
        else:
            print(f"Error: {response.status_code} - {response.text}")
            return None

    except requests.RequestException as e:
        print(f"An error occurred while getting the validator signing info: {e}")
        return None

# Convert ccvaloper to cc address
def convert_valoper_address(address):
    _, data = bech32.bech32_decode(address)
    return bech32.bech32_encode("cc", data)

# Convert pubkey to ccvalcons
def convert_pubkey_valcons(pubkey):
    data = hashlib.sha256(base64.b64decode(pubkey)).digest()[:20]
    return bech32.bech32_encode("ccvalcons", bech32.convertbits(data, 8, 5))

def main():
    args = sys.argv[1:]
    if not len(args) == 1 or not args[0].isdigit() :
        raise SystemExit(USAGE)
    airdrop_amount = args[0]
    api_url_sign = api + "cosmos/slashing/v1beta1/signing_infos"
    params_sign = {
        "pagination.limit": pagination_limit
    }

    api_url_val = api + "cosmos/staking/v1beta1/validators"
    params_val = {
        "status": "BOND_STATUS_BONDED",
        "pagination.limit": pagination_limit
    }
    
    signing_info = get_signing_info(api_url_sign, params_sign)
    
    validator_info = get_validator_info(api_url_val, params_val, signing_info, airdrop_amount)

    if validator_info is not None:
        scripts_dir = os.path.dirname(os.path.abspath(__file__))
        subfolder_path = os.path.join(scripts_dir, subfolder_name)
        os.makedirs(subfolder_path, exist_ok=True)

        # Append the current date to the file name
        current_date = datetime.now().strftime("%Y_%m_%d")
        file_name = os.path.join(subfolder_path, f"validator_airdrop_{current_date}.tsv")

        # Write the converted addresses to a TSV file
        with open(file_name, mode="w", newline="") as tsv_file:
            tsv_writer = csv.writer(tsv_file, delimiter='\t')
            tsv_writer.writerow(["Address", "Balance"])

            for validator in validator_info:
                tsv_writer.writerow(validator)

        print(f"Snapshot written to {file_name}")

if __name__ == "__main__":
    main()