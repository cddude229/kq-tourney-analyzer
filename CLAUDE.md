# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`kq-tourney-analyzer` analyzes tournament statistics for the Killer Queen arcade game. It processes exports from the KQ Hivemind platform and computes per-player stats (kills, deaths, warrior uptime, berry counts, snail distance, etc.).

## Commands

All Go commands run from the `analyzer/` directory:

```bash
cd analyzer && go build -v ./...   # build
cd analyzer && go test ./...       # run all tests
cd analyzer && go test -run TestGame1652756Accuracy ./...  # run single test
```

Python scripts run from the repo root:

```bash
python scripts/tourneys.py         # fetch tournament metadata from Hivemind API
```

## Architecture

The pipeline has four stages:

**1. Parsing (`hivemind/`)** — Opens ZIP exports from Hivemind (containing `gameevent.csv`, `game.csv`, `usergame.csv`), parses rows into `HivemindEvent` structs, and groups them by `TourneyMatch`.

**2. State machine (`models/`)** — Each `TourneyMatch` is replayed through a state machine. Events are dispatched to handlers (`handle_fighting.go`, `handle_berries.go`, `handle_snail.go`, `handle_gates.go`, `handle_lifecycle.go`) that update per-player state and accumulate stats. The state machine is the authoritative source of all computed statistics.

**3. Aggregation (`aggregation/`)** — Extracts `PlayerAndStats` from each state machine, applies name mappings from a CSV file (maps cab position → real player name), then merges duplicate player entries across games using `(Name, Team)` as the merge key.

**4. Output (`main.go`)** — Sorts merged stats and renders an ASCII table via `go-pretty`.

## Key data model notes

- **`PlayerId`** is a 1–10 enum: positions 1–5 are Gold (Queen, Stripes, Abs, Skulls, Chex), 6–10 are Blue.
- Kill/death counts are stored as 2D arrays indexed by `[killer_warrior_state][victim_warrior_state]`, so mil-kills (warrior vs. drone/queen) require summing the appropriate slices.
- Uptime is accumulated in nanoseconds using event timestamps; percentages are computed at output time.
- "Mil" refers to warrior (military) stats throughout the codebase.

## Input data

Tournament ZIP exports go in `tourney_data/`. Use `scripts/tourneys.py` to fetch the list of tournaments from `https://kqhivemind.com/api/`. Player name mapping CSVs are loaded at runtime and map `(team, cab_position)` to a real player name.