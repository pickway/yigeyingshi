# Implementation Plan: CineVerse Admin

## Phase 1: Security and API foundation

- [ ] Scaffold two isolated projects and environment examples.
- [ ] Write failing auth and CRUD integration tests.
- [ ] Implement config, models, token auth, dashboard and resource routes.
- [ ] Verify `go test ./...` and `go vet ./...`.

## Phase 2: Admin application

- [ ] Build login/session API boundary and router guard.
- [ ] Build editorial shell, dashboard and navigation.
- [ ] Build reusable table/editor pages for all managed resources.
- [ ] Add subscriber status management and confirmations.

## Phase 3: Verification

- [ ] Run frontend tests/build and dependency audit.
- [ ] Start admin server/UI with a temporary shared database copy.
- [ ] Browser-test login, CRUD, validation, responsive layout and console.
- [ ] Review diff and document setup without committing.

## Risks

- Shared SQLite concurrent writes: enable WAL and busy timeout.
- Default development credentials: require explicit secret in release mode.
- Generic CRUD becoming opaque: keep explicit resource routes and field schemas on both sides.
