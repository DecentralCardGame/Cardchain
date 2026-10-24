#!/usr/bin/env python3
"""Signed transaction integration test on an isolated temporary local chain.
Never uses or modifies the running development node or any real player keys.
"""
import argparse
import base64
import datetime
import json
import os
from pathlib import Path
import subprocess
import time
import tempfile
import urllib.error
import urllib.request

parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('--binary', type=Path, required=True)
parser.add_argument('--work-dir', type=Path)
args = parser.parse_args()
ROOT = (args.work_dir or Path(tempfile.mkdtemp(prefix='cardchain-e2e-'))).resolve()
ROOT.mkdir(parents=True, exist_ok=True)
BIN = args.binary.resolve()
HOME = ROOT / ('run-deck-e2e-' + datetime.datetime.now().strftime('%Y%m%d-%H%M%S'))
HOME.mkdir()
(HOME / 'config').mkdir()
(HOME / 'config/client.toml').write_text('keyring-backend = "test"\n')
CHAIN = 'run-deck-integration-1'
RPC = 'http://127.0.0.1:27657'
API = 'http://127.0.0.1:1417/DecentralCardGame/cardchain/cardchain/'
env = {**os.environ, 'GOMAXPROCS': '2', 'GOMEMLIMIT': '2GiB'}

def cli(*args):
    result = subprocess.run([str(BIN), *args, '--home', str(HOME)], capture_output=True, text=True, env=env, timeout=90)
    if result.returncode:
        raise RuntimeError(f'CLI {args[:3]} failed: {result.stderr}')
    return result.stdout

def get(url):
    with urllib.request.urlopen(url, timeout=10) as response:
        return json.load(response)

def wait_for(fn, seconds=120):
    deadline = time.monotonic() + seconds
    last = None
    while time.monotonic() < deadline:
        try:
            value = fn()
            if value: return value
        except (urllib.error.URLError, TimeoutError, ConnectionError) as e:
            last = e
        time.sleep(1)
    raise TimeoutError(f'Condition not ready: {last}')

cli('init', 'run-deck-integration', '--chain-id', CHAIN)
addresses = {}
for key in ('validator', 'reporter', 'owner'):
    addresses[key] = json.loads(cli('keys', 'add', key, '--keyring-backend', 'test', '--output', 'json'))['address']
    cli('genesis', 'add-genesis-account', addresses[key], '1000000000stake,1000000000ucredits')
cli('genesis', 'gentx', 'validator', '10000000stake', '--chain-id', CHAIN, '--keyring-backend', 'test')
cli('genesis', 'collect-gentxs')
p = HOME / 'config/genesis.json'
g = json.loads(p.read_text())
s = g['app_state']['cardchain']
s['pools'] = [{'denom': 'ucredits', 'amount': '0'} for _ in range(3)]
s['cardAuctionPrice'] = {'denom': 'ucredits', 'amount': '1000000'}
s['addresses'] = [addresses['reporter'], addresses['owner']]
s['users'] = [{'alias': 'reporter', 'reportMatches': True}, {'alias': 'owner'}]
content = {'Action': {'CardName': 'Test Base', 'CastingCost': 1, 'Class': {'Mysticism': True},
                      'Effects': [{'Sacrifice': {}}], 'FlavourText': '', 'Tags': []}}
s['cardRecords'] = [{'owner': addresses['owner'], 'status': 'permanent',
                     'content': base64.b64encode(json.dumps(content).encode()).decode(),
                     'votePool': {'denom': 'ucredits', 'amount': '0'}}]
p.write_text(json.dumps(g))
p = HOME / 'config/config.toml'
p.write_text(p.read_text().replace('pex = true', 'pex = false').replace('timeout_commit = "5s"', 'timeout_commit = "1s"').replace('pprof_laddr = "localhost:6060"', 'pprof_laddr = ""'))
log = (HOME / 'node.log').open('w')
process = subprocess.Popen([str(BIN), 'start', '--home', str(HOME), '--minimum-gas-prices', '0stake',
                           '--rpc.laddr', 'tcp://127.0.0.1:27657', '--p2p.laddr', 'tcp://127.0.0.1:27656',
                           '--grpc.address', '127.0.0.1:9190', '--api.enable', '--api.address', 'tcp://127.0.0.1:1417',
                           '--grpc-web.enable=false', '--log_level', 'warn'], stdout=log, stderr=log, env=env)

def tx(sender, command, flags, should_pass=True):
    result = json.loads(cli('tx', 'cardchain', 'run-decks', command, *flags,
        '--from', sender, '--keyring-backend', 'test', '--chain-id', CHAIN,
        '--node', RPC, '--gas', '1000000', '--fees', '0stake', '--yes', '--output', 'json'))
    if int(result.get('code', 0)) != 0:
        assert not should_pass, result
        return result
    committed = wait_for(lambda: get(RPC + '/tx?hash=0x' + result['txhash']))
    code = int(committed['result']['tx_result']['code'])
    assert (code == 0) == should_pass, committed
    return committed

try:
    wait_for(lambda: int(get(RPC + '/status')['result']['sync_info']['latest_block_height']) > 0)
    wait_for(lambda: get(API + 'owners/' + addresses['owner'] + '/run_decks'))
    recipe = {'Generated': {'Effects': [{'Sacrifice': {'Target': 'THIS'}}], 'EffectParameters': [{'EffectIndex': 0, 'Amount': 2}]}}
    occurrence = {'instance_id': '42', 'base_card_id': '0', 'forge_recipes': [json.dumps(recipe), json.dumps(recipe)]}
    flags = ['--owner', addresses['owner'], '--run-ref', 'run-1', '--recipe-format-version', '1', '--cards', json.dumps(occurrence)]
    tx('owner', 'submit-run-deck', flags, False)
    tx('reporter', 'submit-run-deck', flags)
    tx('reporter', 'submit-run-deck', flags, False)
    deck = get(API + 'run_decks/0')['deck']
    assert len(deck['cards'][0]['forge_recipes']) == 2
    tx('reporter', 'distill-run-deck', ['--deck-id', '0', '--instance-id', '42'], False)
    tx('owner', 'distill-run-deck', ['--deck-id', '0', '--instance-id', '42'])
    tx('owner', 'distill-run-deck', ['--deck-id', '0', '--instance-id', '42'], False)
    star = get(API + 'all_stars/0')['all_star']
    assert star['card'] == deck['cards'][0]
    assert get(API + 'run_decks/0')['deck']['consumed'] is True
    print('PASS: signed reporter submit, unauthorized/replay rejection, owner-only one-time distillation, exact recipe copy and REST queries.')
    print('Evidence:', HOME)
finally:
    process.terminate()
    process.wait(timeout=30)
    log.close()
