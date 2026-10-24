# Challenge run decks and All-Star cards

Status: Chain core implemented (stores, signed transactions, queries, genesis, CLI and tests). All-Star frontend UI and the frame/author-policy transition are explicitly outside this change. See `deployment-run-decks.md` for the reviewed 552-card snapshot and bootstrap instructions.

## Confirmed requirements

- A completed challenge run can be retained as an immutable deck owned by a player.
- Store source card IDs, not frozen evaluated card content. Current base-card balance changes apply.
- Each card occurrence has its own identity and ordered Forge recipes. Repeated recipes and cards are significant.
- Use the existing gameclient recipe format, not offer/action messages or evaluated Enchantments.
- Only appointed gameclient reporters can submit completed runs.
- The owner can consume the whole deck once and retain exactly one chosen occurrence as an All-Star card.
- All-Star creation copies the base ID and complete ordered recipes without nerfs, conversion or rerolls.
- Player Card Frames will be replaced by this progression. Content authoring with Card Frames must remain available to qriouzz (`cc14km80077s0hch3sh38wh2hfk7kxfau4456r3ej`).
- Gameclient code and gameplay remain outside this implementation scope.

## Existing boundaries

### Authorization

`MsgMatchReporterAppoint` checks module authority and sets `User.ReportMatches` via `SetMatchReporter`.
`MatchOpen`, `MatchReport` and `UpgradePicksReport` already enforce that flag.
Reuse this existing privilege for run submission; do not use existence in `Servers` as authorization.
`GetServerId` creates server statistics records on demand and is not a permission check.
The current appointment path only enables reporters; revocation needs an explicit authority-controlled path before production deployment.

Signatures must cover the submitter through the SDK message signer annotation. Merely supplying
an authorized address in a payload is not sufficient. The deck owner need not be the reporter.
The owner signs consumption; the selected occurrence must belong to that owner's unconsumed deck.

### Gameclient source inspected read-only

Source: gameclient `main`, commit `b21ddc48c535b8b25b2140915f50bec9572c7ac3`.

- `Assets/CrowdControl/Scripts/Models/Run/RunCard.cs`: `baseId`, `id`, ordered `upgrades`.
- `Models/Run/ForgeRecipe.cs`: exactly one of `FixedPackageId` or `Generated`.
- `GeneratedForgeRecipe`: `Ability` or `Effects`, optional `AdditionalCost`, and `EffectParameters`.
- `ForgeEffectParameters`: `EffectIndex`, optional `Amount`, optional ordered `Zones`.
- `KeywordInterpreter/Limited/ForgeRecipeCodec.cs`: strict JSON, explicit adapters and reconstruction.
- `Models/Run/ForgeActionCodec.cs`: actions identify offers/targets; they are not durable upgrade payloads.
- `Models/Run/Run.cs`: stable `forgeIdentity` and ordered `decklist`; `Revision` is a persistence revision, not another completed run.
- `docs/forge-recipe-foundation.md`: base changes and current Forge valuations are deliberately followed.

The inspected codec still supports adapters for legacy Sacrifice/Void/Discard and Choose zone arrays.
Do not validate these payloads by unmarshalling them into the newly changed Cardobject structs:
that would discard or reject valid persisted recipe data. A versioned recipe contract must be distinct
from the editable base-card contract. No Unity/gameplay validation was performed here.

## Proposed state

Names are provisional until protobuf integration.

### SavedRunDeck

- chain-assigned `id`
- `owner` (validated player address)
- `reporter` (authenticated signer)
- `run_ref` (stable originating run identity, not report revision)
- `recipe_format_version`
- optional source catalog identity for provenance, NOT a pinned balance snapshot
- ordered `cards: RunDeckCard[]`
- `created_height`
- `consumed` and optional consumption provenance (`consumed_height`, `all_star_id` with explicit presence)

### RunDeckCard

- `instance_id`: occurrence ID, unique within this deck; zero may be a valid upstream ID
- `base_card_id`
- ordered `forge_recipes`

### Forge recipe storage boundary

Use a typed versioned envelope preserving the existing codec JSON for each recipe. Preserve
ordered arrays and exact gameplay values; do not cache evaluated mana costs, display names,
price coefficients or offer IDs. Specify a canonical JSON representation before hashing it:
reject duplicate members, unknown envelope members, ambiguous alternatives and unsupported versions.

Do not reorder/deduplicate recipes. Generated recipes need strict checks for alternatives,
required parameters, adapter indices and bounded depth/size. Validation must not silently replace
unknown typed effects. Fixed package IDs must be nonempty stable identifiers, never array indices.

The current Forge definition catalog is not stored in Cardchain. Therefore the chain cannot claim
to verify package existence, current valuations, or that a run was legitimately won merely by
accepting recipe JSON. The appointed reporter attests to the completed run and resolved recipes.
A future on-chain catalog/validator would be a separate feature. Retired package definitions must
remain resolvable in the serving catalog for retained decks and All-Stars.

### AllStarCard

- chain-assigned `id`
- `owner`
- `source_deck_id` and `source_instance_id`
- `base_card_id`, copied exactly
- `recipe_format_version` and ordered `forge_recipes`, copied exactly
- `created_height`

Keep this namespace separate from `CardK` (global authored content). No new global base-card
ID, editable frame, or inherited authoring permission is created by distillation.
No transfers or trading are implied by this first version.

## Transactions

### SubmitRunDeck

Signer: appointed reporter. Input: owner, run reference, ordered occurrences, versioned recipes.

Validate before writing:
- signer is a registered user with `ReportMatches=true`;
- owner is a registered player;
- nonempty bounded run reference;
- unique `(reporter, run_ref)` completion key; retries must not create duplicate decks;
- occurrence IDs unique; base cards exist and are not empty/removed slots;
- valid bounded versioned recipe payloads;
- limits for cards/deck, recipes/card, recipe bytes, aggregate bytes and nesting.

A reporter-scoped completion key prevents replays from that reporter. It cannot prove that two
trusted reporters did not attest to the same real-world run under different identities. Do not
claim global run authenticity without a shared authoritative run identity/issuer contract.
There is no ownership requirement on the referenced global base cards: players own the resulting
deck, not the content author's source cards. Allow unfinished/lost runs only if that is part of the
reporter's completion contract; the chain must not infer completion from a client-supplied flag.

### DistillRunDeck

Signer: deck owner. Input: deck ID, instance ID (not just base ID).

Validate owner, unconsumed state, and membership before allocating an All-Star ID. Copy the
occurrence and mark the deck consumed atomically in the SDK transaction. Failures leave no
partial mint or consumption. Reject second consumption even if the requested occurrence differs.
Preserve the consumed deck as historical provenance. No edit, refill or second extraction endpoint.

## Queries and persistence

- deck by ID; decks by owner, optional consumed filter, pagination;
- All-Star by ID; All-Stars by owner, pagination;
- maintain owner indexes and completion uniqueness index rather than unbounded full-store scans;
- include all records and consumed state in genesis/export;
- reconstruct secondary indexes with duplicate/reference checks on import;
- do not renumber IDs or resurrect consumed decks on restart/import.

Use the module's existing protobuf/codec conventions and store abstractions, with new isolated
prefixes and explicit query pagination. No generated protobuf files should be handwritten.

## Content authoring / frame replacement

Introduce a governance-configured content-author allowlist, initially containing qriouzz's address,
not a special case for a mutable account alias. Preserve his existing frames and content.

Retiring player frame purchases requires more than changing the frontend: assess
`CardSchemeBuy`, scheme transfers, existing unspent schemes, and content-save/artist authorization.
Do not disable transfers/saving globally and accidentally break qriouzz or commissioned artwork.
Treat this as a staged policy change after the deck/All-Star paths exist; preserve existing frame
records until an explicit migration handles them. Do not switch off the current economy in this
research/design step.

## Required tests before deployment

- unauthorized reporter, auto-created server statistics record, forged signer, unknown owner;
- repeated completion submission, stable run reference across revisions, size/depth limits;
- duplicate occurrence IDs, repeated valid base IDs, distinct upgrades on identical base IDs;
- recipe order and byte/gameplay-value preservation, fixed/generated alternatives and adapter validation;
- foreign-owner consumption, nonexistent occurrence, double consumption, failure atomicity;
- exact All-Star copy and immutable source deck;
- balance changes do not rewrite stored IDs/recipes;
- pagination, owner isolation, genesis/export round-trip, uniqueness reconstruction;
- future frame-policy changes must preserve qriouzz's authoring path and test transfers; this policy is deferred.

The existing test helper's outdated `NewKeeper` constructor setup has been repaired
(BankKeeper and FeatureFlagKeeper are now required).

## Delivery sequence

1. Freeze the versioned recipe boundary and limits using fixtures from the inspected codec.
2. Implement types, stores, genesis/export, message/query services and focused tests.
3. Add signer-aware reporter submission and owner-only atomic distillation.
4. Export the curated state and verify the complete binary against that snapshot in a separate home.
5. Supply the snapshot, checksums and deployment instructions, preserving the prior home for rollback.

Frontend deck/All-Star UI is outside this work. The configured frame-authoring policy
will be decided separately; existing frame/authoring behavior remains unchanged.

## Implemented API and limits

The core is in `proto/cardchain/cardchain/run_deck.proto`, `keeper/run_deck.go`,
`types/run_deck.go`, and `types/forge_recipe.go`. Generated Go/gateway sources are
regenerated with `scripts/generate-run-deck-proto.sh` (Go and protoc required).

- Msg service: `cardchain.cardchain.RunDeckMsg`.
- Messages: `MsgSubmitRunDeck` (signer reporter), `MsgDistillRunDeck` (signer owner).
- Query service: `cardchain.cardchain.RunDeckQuery`.
- REST: `/DecentralCardGame/cardchain/cardchain/run_decks/{id}` and `all_stars/{id}`.
- Owner lists: `/DecentralCardGame/cardchain/cardchain/owners/{owner}/run_decks` and `all_stars`.
- Owner lists use inclusive `cursor` and `limit` (default 20, maximum 100).
  Follow `next_cursor` only when `has_more=true`. Consumed decks remain in the list.
- CLI: `cardchaind tx cardchain run-decks ...` and `cardchaind query cardchain run-decks ...`.
- IDs start at zero. A zero All-Star ID is valid; use `consumed`, not `all_star_id != 0`.
- Format version 1 accepts the inspected durable recipe codec. It preserves the old
  Sacrifice/Discount/Choose adapters independently of current base-card schemas.
- Limits: 128 cards/deck; 64 recipe occurrences/card; 16 KiB/recipe; 256 KiB aggregate
  raw AND canonical recipe data/deck; 64 effects/recipe; 32 JSON nesting levels.
- Duplicate JSON members (including case variants), unknown members/alternatives,
  invalid enums/amounts/adapters, and unsupported versions are rejected.
- JSON object keys are canonicalized; ordered arrays and repeated occurrences are preserved.
- No Forge package catalog or gameplay evaluator runs on-chain. Reporter attestation
  remains the trust boundary for package existence, completion and balance interpretation.
- Reporter revocation must still be provided through an authority-controlled policy;
  this feature checks the current `ReportMatches` value on every submission.

Examples (use test keys and the intended chain ID/node):

```sh
cardchaind tx cardchain run-decks submit-run-deck \
  --from reporter --owner cc1... --run-ref RUN_UUID --recipe-format-version 1 \
  --cards '{"instance_id":"42","base_card_id":"168","forge_recipes":["{\"FixedPackageId\":\"stable.package.id\"}"]}'
cardchaind tx cardchain run-decks distill-run-deck \
  --from owner --deck-id 0 --instance-id 42
cardchaind query cardchain run-decks run-deck --id 0
cardchaind query cardchain run-decks run-decks-by-owner --owner cc1... --limit 20
```

## Upgrade/deployment boundary

Module consensus version is now 2; the registered 1-to-2 migration is additive and
initializes no existing values (new prefixes are initially empty). A real chain upgrade
must invoke the module migrations through its upgrade handler. Do not reset genesis or
simply treat a binary swap as a coordinated production upgrade. A reviewed export can
instead bootstrap a fresh application store with the new module version, as documented
in `deployment-run-decks.md`.

The old test-helper constructor mismatch was repaired using a real, isolated featureflag
keeper. Its bank dependency remains nil because that store-only fixture does not execute
bank transfers. Core tests cover real SDK signer extraction and gRPC gateway dispatch in
addition to keeper operations and strict recipe validation.

```sh
go test ./x/cardchain/types ./x/cardchain/keeper ./x/cardchain/module
go build ./cmd/cardchaind
```

The gameclient, frame purchases and existing authoring permissions are not changed by
this implementation. Reviewed card removals are supplied in the separate genesis artifact.
Portable isolated integration tests live in `scripts/tests/run_decks_e2e.py` and
`scripts/tests/release_snapshot_e2e.py`; see the deployment guide for invocation.
