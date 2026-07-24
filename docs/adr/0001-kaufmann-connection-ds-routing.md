# ADR 0001 — Route the Kaufmann connection source by CloudEvent `ds`

- **Status:** Accepted
- **Date:** 2026-07-23
- **Deciders:** DIMO platform / Kaufmann oracle work
- **Related:** `pkg/kaufmann`, `pkg/modules/register.go`; kaufmann-oracle#164 (oracle
  side), kaufmann-oracle#165; the DIS wiring change (below)

## Context

The **Kaufmann connection license**
(`0x8D8cDB2B26423c8fDbb27321aF20b4659Ce919fD`) started out carrying a single
device family — Ruptela **Smart5** trackers — and was therefore registered
straight to the **ruptela** module in `pkg/modules/register.go`.

The Kaufmann oracle now also ingests **Queclink GV58 / GV300W** devices that
relay **Kamaleon** telemetry through a flespi proxy. Two facts make these
incompatible with the ruptela decoder:

1. **Different payload shape.** Ruptela telemetry is a `data.signals` map of
   numeric Ruptela IO ids (e.g. `"104"`, `"105"`, `"106"` for VIN) with
   Ruptela-specific encodings and scale factors. Kamaleon frames are a
   completely different ASCII record format.
2. **The oracle already decodes Kamaleon.** Rather than teach a server-side
   decoder the Kamaleon record schema, the oracle pre-decodes each frame into
   **VSS-named signals** and emits them in the DIS **default module** payload
   shape (`data.signals: [{name, value, timestamp}]`).

So one connection now carries **two payload families**, and they must be decoded
by two different modules. They are distinguished by the CloudEvent **`ds`**
(data version): the oracle tags Ruptela events `r/v0/*` and Kamaleon events
`kam/v0/*`.

The CloudEvent **header** is identical for both families: producer = the
**synthetic device** NFT DID, subject = the **vehicle** NFT DID, both built from
the `deviceTokenId` / `vehicleTokenId` in the envelope. Only the *payload
decoding* differs.

## Decision

Introduce **`pkg/kaufmann.Module`**, registered for `KaufmannSource` in all four
registries in place of the ruptela module. It implements the CloudEvent, Signal,
Fingerprint, and Event module interfaces and **routes on the `ds` prefix**:

| `ds` prefix | Handling |
|---|---|
| `r/…`   | Delegated to the **ruptela** module (unchanged historical behavior). |
| `kam/…` | Signal/event/fingerprint decoding delegated to the **default** module (the payload is already VSS); CloudEvent header built locally. |

Header construction is shared: for the `kam/` family the module builds the
status CloudEvent header directly (producer = synthetic DID, subject = vehicle
DID) using its configured chain id and contract addresses, mirroring exactly what
the ruptela module does for `r/v0/s`. The `r/` family delegates header
construction to the ruptela module.

`DIS` configures the CloudEvent registration for `KaufmannSource` with a
contract-configured `kaufmann.Module` (with `AftermarketContractAddr` set to the
**synthetic device** contract, since Kaufmann devices are synthetic devices),
replacing the previous `ruptela.Module` override. The signal/event/fingerprint
registries use the zero-value module from `register.go` (those paths don't build
DIDs, so they need no contract config).

## Consequences

**Positive**

- One connection license and one mTLS identity serve both hardware families; no
  new on-chain connection, cert, or synthetic-device minting token id is needed.
- All Kamaleon record knowledge stays in the **oracle**; model-garage only
  routes. DIS is unchanged apart from a dependency bump and the one-line
  registration swap.
- Existing Ruptela behavior is untouched and covered by a delegation regression
  test, so the change is low-risk for the existing fleet.
- Adding a future family on this connection is a new `ds` prefix and one more
  branch — no new source/license.

**Negative / trade-offs**

- The Kaufmann source now has slightly more logic than a plain module alias, and
  a reader must know that `ds` is the routing key.
- Header construction for `kam/` is a small duplication of the ruptela header
  logic (kept intentionally, to avoid coupling the two decoders).
- The oracle and this module share an implicit **wire contract** (the `kam/`
  envelope shape). It is documented in the oracle repo
  (`docs/flespi-to-dis-plan.md`) and pinned by golden tests on both sides.

## Alternatives considered

- **A. Masquerade Kamaleon as Ruptela** — map Kamaleon fields onto Ruptela
  numeric IO ids/encodings in the oracle so the existing ruptela decoder handles
  them. *Rejected:* semantically wrong (per-IO encodings and scale factors differ,
  e.g. Ruptela IO 114 odometer = value×5), lossy (only overlapping signals), and
  it pollutes the `r/` schema.
- **B. A second connection license / source** — give Kamaleon its own source
  address, which falls through to the default module for free. *Rejected:*
  operationally heavy — a new on-chain license, a second mTLS cert, a second
  connection token id for synthetic-device minting, and a fleet split across two
  connections — for no functional gain over routing by `ds`.
- **C. Route by `ds` in a Kaufmann module (chosen)** — confines the change to
  model-garage routing (~a module) plus a DIS dependency bump, keeps Kamaleon
  decoding in the oracle, and leaves Ruptela untouched.

## References

- `pkg/kaufmann/module.go`, `pkg/kaufmann/module_test.go`
- `pkg/modules/register.go` (`KaufmannSource` registration)
- Oracle side: kaufmann-oracle `internal/kamaleon/cloudevent.go` (emits the
  `kam/v0/s` envelope) and `docs/flespi-to-dis-plan.md` (Option C plan + wire
  contract)
- DIS side: `internal/processors/cloudeventconvert/cloudeventconvert.go` overrides
  `KaufmannSource` with a contract-configured `kaufmann.Module`
