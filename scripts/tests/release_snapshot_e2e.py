#!/usr/bin/env python3
"""Bootstrap the public snapshot with a fresh validator; no changes to live homes."""
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
import urllib.parse
import urllib.request

SCRIPT_ROOT = Path(__file__).resolve().parents[1]
parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('--binary', type=Path, required=True)
parser.add_argument('--genesis', type=Path, required=True)
parser.add_argument('--work-dir', type=Path)
args = parser.parse_args()
ROOT = (args.work_dir or Path(tempfile.mkdtemp(prefix='cardchain-snapshot-'))).resolve()
ROOT.mkdir(parents=True, exist_ok=True)
BIN = args.binary.resolve()
SNAPSHOT = args.genesis.resolve()
HOME = ROOT / ('release-preflight-' + datetime.datetime.now().strftime('%Y%m%d-%H%M%S'))
CHAIN = 'release-preflight-1'
RPC = 'http://127.0.0.1:27657'
API = 'http://127.0.0.1:1417'
env = {**os.environ, 'GOMAXPROCS': '2', 'GOMEMLIMIT': '6GiB'}

def cli(*args):
    result = subprocess.run([str(BIN), *args, '--home', str(HOME)], capture_output=True, text=True, env=env, timeout=180)
    if result.returncode: raise RuntimeError(result.stderr)
    return result.stdout

def get(path):
    with urllib.request.urlopen(path, timeout=20) as r: return json.load(r)

HOME.mkdir()
(HOME / 'config').mkdir()
(HOME / 'config/client.toml').write_text('keyring-backend = "test"\n')
cli('init', 'release-preflight', '--chain-id', CHAIN)
prepared = HOME / 'prepared-genesis.json'
subprocess.run(['python3', str(SCRIPT_ROOT/'prepare-deployment-genesis.py'), str(SNAPSHOT), str(prepared),
                '--expected-cards', '552', '--chain-id', CHAIN, '--validator-key', str(HOME/'config/priv_validator_key.json')], check=True)
prepared.replace(HOME / 'config/genesis.json')
cli('genesis', 'validate-genesis')
config = HOME / 'config/config.toml'
config.write_text(config.read_text().replace('pex = true', 'pex = false').replace('pprof_laddr = "localhost:6060"', 'pprof_laddr = ""'))
log = (HOME / 'node.log').open('w')
node = subprocess.Popen([str(BIN), 'start', '--home', str(HOME), '--minimum-gas-prices', '0ubpf',
    '--rpc.laddr', 'tcp://127.0.0.1:27657', '--p2p.laddr', 'tcp://127.0.0.1:27656',
    '--grpc.address', '127.0.0.1:9190', '--api.enable', '--api.address', 'tcp://127.0.0.1:1417',
    '--grpc-web.enable=false', '--log_level', 'warn'], stdout=log, stderr=log, env=env)

def wait(fn, timeout=360):
    deadline = time.monotonic() + timeout
    last = None
    while time.monotonic() < deadline:
        if node.poll() is not None: raise RuntimeError(f'Node exited: {HOME}/node.log')
        try:
            result = fn()
            if result: return result
        except (urllib.error.URLError, TimeoutError, ConnectionError) as e: last = e
        time.sleep(2)
    raise TimeoutError(str(last))

try:
    prefix = API + '/DecentralCardGame/cardchain/cardchain/'
    ids = list(map(int, wait(lambda: get(prefix+'cards?sortBy=Id')['cardIds'])))
    source = json.loads(SNAPSHOT.read_text())
    cards = source['app_state']['cardchain']['cardRecords']
    expected = {i for i,c in enumerate(cards) if c['content']}
    assert set(ids) == expected and len(ids) == 552
    for offset in range(0,len(ids),40):
        batch = ids[offset:offset+40]
        contents = get(prefix+'card_contents?'+urllib.parse.urlencode([('cardIds',i) for i in batch]))['cardContents']
        assert len(contents) == len(batch)
        for i,entry in zip(batch,contents): assert json.loads(entry['content']) == json.loads(base64.b64decode(cards[i]['content'])), i
    versions = get(API+'/cosmos/upgrade/v1beta1/module_versions?module_name=cardchain')['module_versions']
    assert any(v['name']=='cardchain' and int(v['version'])==2 for v in versions), versions
    owner = 'cc14km80077s0hch3sh38wh2hfk7kxfau4456r3ej'
    assert not get(prefix+'owners/'+owner+'/run_decks')['decks']
    assert not get(prefix+'owners/'+owner+'/all_stars')['all_stars']
    first = int(get(RPC+'/status')['result']['sync_info']['latest_block_height'])
    assert first >= int(source['initial_height'])
    last = wait(lambda: (height if (height := int(get(RPC+'/status')['result']['sync_info']['latest_block_height'])) > first else None), 60)
    print(f'PASS: fresh-validator deployment, 552 exact cards, empty new stores, module version 2, blocks {first}->{last}.')
    print('Evidence:', HOME)
finally:
    node.terminate()
    node.wait(timeout=60)
    log.close()
