# Implementation Plan: CineVerse 完善版

## Architecture Decisions

- Keep the current Vue/Gin/SQLite stack and `/api/v1` prefix.
- Additive API evolution: extend movie query parameters and add search/article detail endpoints without removing existing fields.
- Use reusable state and card components to reduce the four oversized views gradually.
- Keep generated imagery local to the repository; no runtime dependency on image hosts.

## Phase 1: Foundation and contracts

- [x] Document product scope and API behavior.
- [x] Add failing backend tests for search, validation, article detail and unified search.
- [x] Add frontend pure-helper tests and a test script.
- [x] Normalize API errors and validate public inputs.

Checkpoint: both test suites pass and existing API calls remain compatible.

## Phase 2: Complete content journeys

- [x] Implement keyword search and supported movie sorting end to end.
- [x] Implement article detail and global search end to end.
- [x] Add movie detail and article detail route pages.
- [x] Connect cards, navbar search and useful CTA links.

Checkpoint: discovery → list/search → detail journeys work locally.

## Phase 3: Visual system and resilience

- [x] Refresh global tokens, typography, navigation, footer, and home page composition.
- [x] Add reusable loading, error, empty and modal/search-overlay states.
- [x] Polish movie, learning and AI pages for responsive layout and accessibility.
- [x] Add project-owned generated hero artwork and optimized local assets.

Checkpoint: desktop/mobile browser smoke test is clean.

## Phase 4: Quality and handoff

- [x] Run formatting, full tests, production build, dependency audit and `go vet`.
- [x] Review diff for correctness, security, architecture, performance and dead code.
- [x] Replace template README with complete development/deployment documentation.
- [x] Report all changes without committing or pushing.

## Risks and Mitigations

| Risk | Mitigation |
|---|---|
| Seed data already exists in local DB | Make schema additions optional/additive and avoid relying on reseeding existing rows. |
| Large existing single-file views | Extract only reusable states/cards needed for journeys; avoid wholesale framework rewrites. |
| Anonymous write endpoint abuse | Validate email, normalize it, and apply a lightweight in-memory rate limiter. |
| Generated asset size | Inspect dimensions and optimize before use. |
