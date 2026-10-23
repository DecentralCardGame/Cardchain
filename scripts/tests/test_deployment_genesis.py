import base64
import copy
import importlib.util
import json
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

SCRIPT = Path(__file__).resolve().parents[1] / 'prepare-deployment-genesis.py'
spec = importlib.util.spec_from_file_location('deployment', SCRIPT)
deployment = importlib.util.module_from_spec(spec)
spec.loader.exec_module(deployment)


def fixture():
    return {
        'chain_id': 'old-2', 'initial_height': '40', 'app_hash': 'old',
        'consensus': {'validators': [{}]},
        'app_state': {
            'cardchain': {'cardRecords': [{'content': 'unchanged'}, {'content': ''}]},
            'bank': {'balances': ['unchanged']},
            'staking': {'validators': [{'consensus_pubkey': {'key': 'old'}}]},
            'slashing': {'signing_infos': [{'address': 'ccvalcons1old'}]},
            'distribution': {'record': 'ccvalcons1old'},
            'ibc': {
                'client_genesis': {'clients': [{'client_id': '09-localhost'}], 'clients_consensus': [], 'clients_metadata': [], 'create_localhost': False},
                'connection_genesis': {'connections': [{'id': 'connection-localhost'}], 'client_connection_paths': []},
                'channel_genesis': {field: [] for field in ('channels', 'acknowledgements', 'commitments', 'receipts', 'send_sequences', 'recv_sequences', 'ack_sequences')},
            },
        },
    }


class DeploymentGenesisTests(unittest.TestCase):
    def test_preserves_cards_balances_and_existing_decks(self):
        source = fixture()
        existing = {'decks': [{'id': '0'}], 'all_stars': []}
        source['app_state']['cardchain']['run_decks'] = copy.deepcopy(existing)
        result = deployment.prepare(source, 1)
        self.assertEqual(result['app_state']['cardchain']['run_decks'], existing)
        self.assertEqual(result['app_state']['bank'], {'balances': ['unchanged']})
        self.assertEqual(result['app_state']['cardchain']['cardRecords'][0]['content'], 'unchanged')
        self.assertEqual(result['initial_height'], '40')
        self.assertEqual(result['chain_id'], 'old-2')
        self.assertEqual(result['app_hash'], '')

    def test_new_stores_and_safety_checks(self):
        result = deployment.prepare(fixture(), 1)
        self.assertEqual(result['app_state']['cardchain']['run_decks'], {'decks': [], 'all_stars': []})
        with self.assertRaises(ValueError): deployment.prepare(fixture(), 2)
        with self.assertRaises(ValueError): deployment.prepare(fixture(), 1, 'new-1')
        for branch, field, value in (
            ('client_genesis', 'clients', [{'client_id': 'external'}]),
            ('connection_genesis', 'connections', [{'id': 'connection-0'}]),
            ('channel_genesis', 'commitments', ['pending']),
        ):
            source = fixture()
            source['app_state']['ibc'][branch][field] = value
            with self.assertRaises(ValueError): deployment.prepare(source, 1)

    def test_cli_rekey_never_exports_private_key_or_overwrites(self):
        public = {'type': 'tendermint/PubKeyEd25519', 'value': base64.b64encode(bytes(32)).decode()}
        with self.assertRaises(ValueError): deployment.prepare(fixture(), 1, 'old-2', public)
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            source, key, output = root/'source.json', root/'key.json', root/'output.json'
            original = json.dumps(fixture())
            source.write_text(original)
            key.write_text(json.dumps({'pub_key': public, 'priv_key': {'value': 'DO_NOT_EXPORT_THIS_SECRET'}}))
            command = [sys.executable, str(SCRIPT), str(source), str(output), '--expected-cards', '1', '--chain-id', 'new-1', '--validator-key', str(key)]
            subprocess.run(command, check=True, capture_output=True)
            result = json.loads(output.read_text())
            self.assertNotIn('DO_NOT_EXPORT_THIS_SECRET', output.read_text())
            self.assertEqual(source.read_text(), original)
            self.assertEqual(result['chain_id'], 'new-1')
            self.assertEqual(result['consensus']['validators'][0]['pub_key'], public)
            self.assertEqual(result['app_state']['staking']['validators'][0]['consensus_pubkey']['key'], public['value'])
            self.assertEqual(result['app_state']['slashing']['signing_infos'][0]['address'], result['app_state']['distribution']['record'])
            self.assertNotEqual(subprocess.run(command, capture_output=True).returncode, 0)


if __name__ == '__main__': unittest.main()
