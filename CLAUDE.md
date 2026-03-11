# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Pure Go Mermaid-style Gantt chart renderer. Parses Mermaid Gantt DSL and outputs PNG directly — no Node.js or mermaid-cli required. Supports themes, Chinese fonts, configurable time axes, weekend/timezone handling.

## Build & Test Commands

```bash
go test ./...                                    # Run full test suite (required before commit)
go test ./internal/parser -run TestTimeFormat     # Run specific parser tests
go vet ./...                                     # Static analysis
gofmt -w .                                       # Format code
go mod tidy                                      # After dependency changes
go run ./examples/basic                          # Render sample PNG to temp path
```

## Architecture

```
User Code → Render(ctx, Input)
  ├─ parser.Parse(source)        → Model (AST, unresolved times)
  ├─ parser.ResolveSchedule(m)   → Model (resolved Start/End via topological sort)
  ├─ font.SelectFontPath()       → Font path (explicit > GGM_FONT_PATH env > system candidates)
  ├─ render.RenderModel(m, opts) → []byte (PNG)
  └─ Write to OutputPath / Writer
```

### Package Layout

- **Root package** (`types.go`, `renderer.go`, `theme.go`): Public API — `Render()`, `Input`, `Theme`, `RenderResult`
- **`internal/parser`**: Mermaid Gantt DSL parsing (`gantt_parser.go`), AST types (`ast.go`), dependency resolution & time scheduling (`scheduler.go`)
- **`internal/render`**: PNG drawing (`render.go` ~32KB, core drawing logic), calendar-aware duration (`calendar.go`), time axis utils (`timeutil.go`)
- **`internal/font`**: Font loading with fallback chain (`font.go`), hex color parsing

### Key Design Decisions

- **Pure Go only**: No external CLI dependencies. Only `github.com/golang/freetype` and `golang.org/x/image` as runtime deps.
- **Reproducible output**: Same input + theme + font = identical PNG bytes. No local-time assumptions.
- **Font errors are loud**: Missing explicit/env font → hard error, not silent fallback. Auto-discovery failure lists all attempted paths.
- **Mermaid syntax compatibility**: Parser aligns with Mermaid.js Gantt spec — dayjs `dateFormat`, strftime `axisFormat`, `excludes`/`includes`, task statuses (`crit`, `done`, `active`, `milestone`, `vert`).

## Testing Conventions

- Table-driven tests preferred for parser/render inputs
- Fixtures live in `testdata/mermaid_full/*.gantt` — use fixtures over embedded strings for syntax coverage
- Rendering tests write to `os.TempDir()` and assert PNG byte characteristics (size/colors), not manual viewing
- Bug fixes: add reproduction test case first, then fix

## Project Constraints (from constitution)

- `Render`/`Input`/`Theme` public API must remain backward-compatible
- Visible behavior changes (including theme color values) require version bump per SemVer
- New DSL directives must be documented in README/examples with fallback or clear error
- Time/timezone handling must be explicit and cross-platform reproducible
- Do not commit generated PNGs; example outputs go to temp directories
- Go 1.24+ required
