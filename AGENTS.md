# Repository Guidelines

## Project Structure & Module Organization
- Root package (`renderer.go`, `theme.go`, `types.go`) exposes the public `Render` API and theme helpers.
- `internal/parser` contains the Mermaid-style Gantt grammar, AST, and scheduling/time utilities.
- `internal/render` draws timelines, tasks, and axes; `internal/font` handles font selection and fallbacks.
- Tests sit next to code; fixture-driven cases live in `testdata/mermaid_full`. Examples are under `examples/` (see `examples/basic/main.go` and `examples/full_mermaid.gantt`).

## Build, Test, and Development Commands
- `go test ./...` – run the full suite (requires Go 1.24+).
- `go test ./internal/parser -run TestTimeFormat` – target parser time/format cases when iterating.
- `go vet ./...` – static checks; run before PRs.
- `gofmt -w .` – enforce standard Go formatting; `go mod tidy` after dependency changes.
- `go run ./examples/basic` – render a sample PNG locally (writes to a temp path).

## Coding Style & Naming Conventions
- Use gofmt defaults (tabs, Canonical Go); keep files ASCII unless rendering requires otherwise.
- Exported identifiers use CamelCase and have brief doc comments; package-private helpers stay lowercase and unstuttered.
- Prefer small, focused functions; keep renderer/parser surfaces minimal and consistent with `Input`/`Theme` types.
- Error messages should be concise and user-oriented (what failed and why).

## Testing Guidelines
- Follow Go test naming: `*_test.go`, `TestXxx`, and subtests for variants; favor table-driven cases for parser/render inputs.
- Use `testdata/mermaid_full/*.gantt` fixtures when adding syntax coverage instead of embedding long strings.
- For rendering tests, write outputs to `os.TempDir()` and clean up (see `render_test.go`); assert PNG bytes (size/colors) rather than relying on manual viewing.
- If adding new fixtures, keep names descriptive (e.g., `timezone_offsets.gantt`) and document intent in comments near the test.

## Commit & Pull Request Guidelines
- Commit messages in this repo are short and imperative (e.g., `readme`, `change`); keep them concise but clear.
- PRs should state intent, summarize key changes, and include test results (`go test ./...`). Link issues when relevant.
- Provide screenshots or output paths for visual/theme changes; avoid committing generated PNGs.
- Keep `go.mod`/`go.sum` updates in the same PR as the code that needs them.

## Security & Configuration Tips
- Rendering is pure Go; no external CLI deps. Fonts come from `FontPath` or `GGM_FONT_PATH`—note non-default fonts in PR descriptions.
- Validate time zone-sensitive changes with explicit test cases; avoid hard-coded local time assumptions.
- Do not store secrets or private assets in fixtures; prefer local overrides or environment variables.

## Active Technologies
- Go 1.24 + stdlib；github.com/golang/freetype；golang.org/x/image (001-fix-alpine-font)

## Recent Changes
- 001-fix-alpine-font: Added Go 1.24 + stdlib；github.com/golang/freetype；golang.org/x/image
