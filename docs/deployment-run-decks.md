# Deploy the reviewed card snapshot with Run Decks / All-Stars

This release combines the keyword-foundation Cardobject with the completed Run Deck /
All-Star chain core. Frame purchases and content-author permissions are unchanged.
No gameclient integration or All-Star frontend UI is included.

## Source and dependencies

- PR base: Cardchain `development`.
- Go toolchain: **1.24.5** (see `go.mod`).
- Cardobject is pinned to **`v0.5.4-0.20261024120417-dad764e41bae`**, merged master commit
  `dad764e41bae3da9e929a34f32193c7f6f84e9b4` (Cardobject PR #107).
- No sibling checkout, local `replace` directive, private module or vendored snapshot is needed.
- Generated protobuf and gateway code is committed. Regeneration additionally requires
  `protoc`; use `bash scripts/generate-run-deck-proto.sh`.

```sh
go test -mod=readonly ./...
python3 -B -m unittest discover -s scripts/tests -p test_deployment_genesis.py
go build -mod=readonly -trimpath -o build/cardchaind ./cmd/cardchaind
```

## Genesis artifact

The separate `genesis.json` / `genesis.json.gz` artifact is derived from a fresh export
of the reviewed local chain, not from the original, uncurated testnet genesis.
Its fingerprint and inventory are recorded in `release-snapshot.json` alongside this document.
That manifest intentionally retains the original snapshot's build provenance (Cardobject
`c8c0da79c0db`); updating the dependency does not regenerate or relabel the existing artifact.
The file is hundreds of MB because historical image data is retained; it is deliberately
not committed to Git. Obtain the artifact accompanying the PR/release, then verify:

```sh
sha256sum -c SHA256SUMS
# When only the compressed artifact was downloaded:
# gzip -dc genesis.json.gz > genesis.json
# sha256sum -c genesis.sha256
```

The snapshot contains **552 nonempty cards**, **562 empty slots**, and **1114 total ID slots**.
All surviving IDs and card contents are retained. Historic records and image data remain.
Run Deck and All-Star stores are explicitly initialized empty. Existing authorized match
reporters remain the authorization boundary for submitting completed runs.

Preparation changes only the genesis header (`app_hash` and fresh genesis time), adds the
new empty stores when absent, and recreates the SDK's built-in localhost IBC client and
connection. The SDK exports `connection-localhost` but rejects it during genesis validation.
No external IBC state exists in this snapshot; the preparation tool refuses such state.
Bank balances, user addresses, existing permissions and validator operator ownership are
not reset. This is a stateful snapshot, **not a fresh production economic configuration**.

## New external single-validator network (fresh private key)

The provided genesis references the local validator's **public** key. It does not contain
its private key, so simply starting that file on a new validator cannot produce blocks.
Use the preparation tool to replace the consensus public key with your own freshly generated
key and select a distinct chain ID. All validators of the same network must use the same
resulting genesis; do not independently rekey every node.

```sh
BIN="$PWD/build/cardchaind"
NODE_HOME="$HOME/.cardchain-allstars"
CHAIN_ID="cardchain-allstars-1"
test ! -e "$NODE_HOME" || { echo 'Use a fresh node home'; exit 1; }
"$BIN" init validator --home "$NODE_HOME" --chain-id "$CHAIN_ID"
python3 scripts/prepare-deployment-genesis.py genesis.json prepared-genesis.json \
  --expected-cards 552 --chain-id "$CHAIN_ID" \
  --validator-key "$NODE_HOME/config/priv_validator_key.json"
install -m 0644 prepared-genesis.json "$NODE_HOME/config/genesis.json"
"$BIN" genesis validate-genesis --home "$NODE_HOME"
# Choose your fee/security/network policy before exposing services publicly.
"$BIN" start --home "$NODE_HOME" --minimum-gas-prices 0.01ubpf \
  --rpc.laddr tcp://127.0.0.1:26657 --api.enable \
  --api.address tcp://127.0.0.1:1317
```

The rekeying path supports this single-validator export only. It updates the consensus
validator key/address, staking consensus key and consensus-address references in slashing/
distribution. It does **not** transfer the inherited staking operator account, delegation
ownership, existing user accounts or governance voting power to a new operator. Provision
those permissions deliberately before treating a fork as a production network.

Private validator keys, keyrings, node databases and signing state must never be committed
or bundled with the public genesis. Back them up separately using the operator's secure process.

## Starting the exact snapshot with the existing local validator

For the local continuation test, retain the existing validator key and last-sign state in
an isolated new home, use the supplied genesis unchanged, and stop the previous process
first. The exported initial height must exceed its last signed height. Do not run old and
new homes simultaneously. Keep the previous home and binary intact for rollback.

This is a deliberate local export/import continuation, **not** an in-place upgrade protocol
for a live public network. A coordinated production upgrade still needs its upgrade handler
and module migrations. Fresh genesis initialization invokes the module version map with
Cardchain consensus version **2**; the old state is not opened directly by this procedure.

## Smoke checks

```sh
curl -fsS 'http://127.0.0.1:1317/DecentralCardGame/cardchain/cardchain/cards?sortBy=Id'
curl -fsS 'http://127.0.0.1:1317/cosmos/upgrade/v1beta1/module_versions?module_name=cardchain'
# Substitute an existing cc address; both lists start empty in this snapshot.
cardchaind query cardchain run-decks run-decks-by-owner --owner cc1...
cardchaind query cardchain run-decks all-stars --owner cc1...
```

Verify continued block production, the exact 552 active IDs and unchanged contents, empty
new stores, and module consensus version 2. Core tests cover signer extraction, reporter
permission, replay prevention, one-time owner-only distillation, exact ordered recipe copying,
pagination, strict bounded recipe validation, and genesis round-tripping. Signed integration
tests run on a separate disposable chain, not by granting reporters on the exported
reviewed snapshot:

```sh
# Run sequentially; these tests use localhost ports 27656, 27657, 1417 and 9190.
python3 -B scripts/tests/release_snapshot_e2e.py \
  --binary build/cardchaind --genesis genesis.json
python3 -B scripts/tests/run_decks_e2e.py --binary build/cardchaind
```

Both create independent homes, shut their test node down, and print the evidence directory.
The signed test uses fresh test-only accounts and the noninteractive test keyring backend.
It checks unauthorized submission, authorized submission, replay rejection, foreign-owner
rejection, successful owner distillation, double-consumption rejection and exact recipe
copying through the real transaction pipeline. The snapshot test checks rekeying, all 552
retained card contents, empty new stores and module version 2.

See `run-decks-all-stars.md` for the full API contract and trust boundary.
