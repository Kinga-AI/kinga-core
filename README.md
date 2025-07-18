# kinga-core

**Shared utilities and contracts for the Kinga AI stack** — language registry,
policy & data schemas, authentication helpers, telemetry shims, and generated
SDKs (Go · Python · TypeScript) consumed by **Inzwa**, **Kinga**, and **Kweli**.

- ✅ **Language Pack v1.1:** en, fr, sw, ha, am, zu, sn, nd (pilot: Nigeria, Kenya, Zimbabwe)
- 📦 JSON Schema / Protobuf definitions: PromptLog, EvalRun, PackMeta, EvidenceSignature
- 🔐 Auth interfaces (token verification; pluggable IdPs)
- 📊 Observability helpers (OpenTelemetry)
- 🛠 Codegen targets for downstream repos

**License:** Apache-2.0 (code).  
**Contributing:** DCO required (`git commit -s`). See `CONTRIBUTING.md`.

> This repo is the *source of truth.* Product repos vendor or depend on tagged releases.

