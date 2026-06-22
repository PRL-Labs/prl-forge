# PRL Forge Development Guide

## Статус

Дата: 22.06.2026

---

# Архитектура

```
Miner
    │
    ▼
Stratum Adapter
    │
    ▼
PRL Forge Core
    │
    ▼
pearl-node RPC
    │
    ▼
submitblock()
```

---

# Core (ЗАМРАЗЕН)

internal/pool/
- [x] engine.go
- [x] builder.go
- [x] job.go
- [x] job_manager.go
- [x] pool.go

internal/updater/
- [x] updater.go

internal/pearl/
- [x] rpc.go
- [x] client.go
- [x] getblocktemplate()

---

# Stratum

internal/stratum/

- [x] session.go
- [x] server.go
- [x] dispatcher.go
- [x] authorize.go
- [x] subscribe.go
- [x] notify.go
- [ ] submit.go (финализиране)
- [ ] share validation

---

# Adapter

internal/adapter/stratum/

- [x] adapter.go
- [x] srbminer.go
- [x] bzminer.go
- [x] generic.go

---

# RPC

- [x] getblocktemplate()
- [x] BuildJob()
- [x] Broadcast()

---

# Build

- [x] go build ./...
- [x] go run ./cmd/forge

---

# Следващи Commit-и

## Commit 003
- [ ] cmd/stratum-test

## Commit 004
- [ ] Notify Adapter Integration

## Commit 005
- [ ] Submit Adapter Integration

## Commit 006
- [ ] Share Validation

## Commit 007
- [ ] submitblock()

---

# Финални тестове

- [ ] Local Stratum Test Client
- [ ] HiveOS Test
- [ ] Accepted Share
- [ ] Block Found
- [ ] submitblock()
- [ ] Block Accepted

---

# Правила

- Не променяме архитектурата.
- Един commit = една задача.
- Само цели файлове.
- Build след всеки commit.
- Runtime тест след всеки commit.

---

# Забранено

- Втори Engine
- Втори JobManager
- Bitcoin логика в Core
- Patch-ове
- Частични файлове

---

# Цел

```
Connected
↓

Subscribe
↓

Authorize
↓

Notify
↓

Submit
↓

Accepted Share
↓

Block Found
↓

submitblock()

↓

Block Accepted
```