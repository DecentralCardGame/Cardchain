#!/usr/bin/env python3
"""Prepare the exported local snapshot for this release; never edits the input.
Optional rekeying bootstraps a separate single-validator network using a fresh key.
No private key material is copied into the output.
"""
import argparse
import base64
import datetime
import hashlib
import json
from pathlib import Path


def consensus_address(public_key, hrp):
    raw = hashlib.sha256(base64.b64decode(public_key, validate=True)).digest()[:20]
    value = int.from_bytes(raw, 'big')
    data = [(value >> shift) & 31 for shift in range(155, -1, -5)]
    words = [ord(c) >> 5 for c in hrp] + [0] + [ord(c) & 31 for c in hrp] + data + [0] * 6
    checksum = 1
    generators = [0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3]
    for word in words:
        top = checksum >> 25
        checksum = ((checksum & 0x1ffffff) << 5) ^ word
        for i, generator in enumerate(generators):
            if (top >> i) & 1:
                checksum ^= generator
    checksum ^= 1
    alphabet = 'qpzry9x8gf2tvdw0s3jn54khce6mua7l'
    return raw.hex().upper(), hrp + '1' + ''.join(alphabet[v] for v in data + [(checksum >> shift) & 31 for shift in range(25, -1, -5)])


def replace_address(value, old, new):
    if isinstance(value, dict): return {k: replace_address(v, old, new) for k, v in value.items()}
    if isinstance(value, list): return [replace_address(v, old, new) for v in value]
    return new if value == old else value


def prepare(genesis, expected_cards, chain_id=None, public_key=None):
    state = genesis['app_state']
    cards = state['cardchain']['cardRecords']
    actual = sum(bool(c['content']) for c in cards)
    if actual != expected_cards:
        raise ValueError(f'Expected {expected_cards} active cards, found {actual}')
    # Preserve any existing decks/stars when this tool is reused for later exports.
    if state['cardchain'].get('run_decks') is None:
        state['cardchain']['run_decks'] = {'decks': [], 'all_stars': []}
    ibc = state['ibc']
    clients, connections, channels = ibc['client_genesis'], ibc['connection_genesis'], ibc['channel_genesis']
    if any(c['client_id'] != '09-localhost' for c in clients['clients']):
        raise ValueError('External IBC clients require a separate migration')
    if any(c['id'] != 'connection-localhost' for c in connections['connections']):
        raise ValueError('External IBC connections require a separate migration')
    if clients['clients_consensus'] or clients['clients_metadata'] or connections['client_connection_paths']:
        raise ValueError('Unexpected IBC client metadata or connection paths')
    for field in ('channels', 'acknowledgements', 'commitments', 'receipts', 'send_sequences', 'recv_sequences', 'ack_sequences'):
        if channels[field]: raise ValueError(f'IBC {field} require a separate migration')
    # SDK exports connection-localhost but its genesis validator rejects that ID.
    clients['clients'] = []
    clients['create_localhost'] = True
    connections['connections'] = []
    if bool(chain_id) != bool(public_key):
        raise ValueError('Provide both --chain-id and --validator-key, or neither')
    if public_key:
        if chain_id == genesis['chain_id']:
            raise ValueError('Rekeying requires a distinct chain ID (new network)')
        if public_key['type'] != 'tendermint/PubKeyEd25519':
            raise ValueError('Only Ed25519 validator keys are supported')
        if len(base64.b64decode(public_key['value'], validate=True)) != 32:
            raise ValueError('Invalid Ed25519 public key length')
        validators = genesis['consensus']['validators']
        staking = state['staking']['validators']
        signing_infos = state['slashing']['signing_infos']
        if len(validators) != 1 or len(staking) != 1 or len(signing_infos) != 1:
            raise ValueError('Rekeying supports only this single-validator snapshot')
        old_address = signing_infos[0]['address']
        hrp = old_address.rsplit('1', 1)[0]
        address_hex, address = consensus_address(public_key['value'], hrp)
        validators[0].update(address=address_hex, pub_key=public_key)
        staking[0]['consensus_pubkey']['key'] = public_key['value']
        for module in ('slashing', 'distribution'):
            state[module] = replace_address(state[module], old_address, address)
        genesis['chain_id'] = chain_id
    genesis['app_hash'] = ''
    genesis['genesis_time'] = datetime.datetime.now(datetime.timezone.utc).isoformat().replace('+00:00', 'Z')
    return genesis


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('source', type=Path)
    parser.add_argument('output', type=Path, help='Must not already exist')
    parser.add_argument('--expected-cards', type=int, required=True)
    parser.add_argument('--chain-id', help='Distinct chain ID for a new network')
    parser.add_argument('--validator-key', type=Path, help='Fresh local priv_validator_key.json; only pub_key is read')
    args = parser.parse_args()
    if args.output.exists(): parser.error('Refusing to overwrite output')
    if args.source.resolve() == args.output.resolve(): parser.error('Source and output must differ')
    public_key = json.loads(args.validator_key.read_text())['pub_key'] if args.validator_key else None
    genesis = prepare(json.loads(args.source.read_text()), args.expected_cards, args.chain_id, public_key)
    args.output.parent.mkdir(parents=True, exist_ok=True)
    with args.output.open('x') as output: json.dump(genesis, output, separators=(',', ':'))
    print(f"Prepared {args.output}: chain {genesis['chain_id']}, initial height {genesis['initial_height']}, {args.expected_cards} cards")


if __name__ == '__main__': main()
