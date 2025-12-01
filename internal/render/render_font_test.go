package render_test

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	gantt "github.com/pyroflux/go-mermaid-gantt"
)

// renderWithFont 渲染到临时文件并验证 PNG 可读性，返回结果与文件路径（已注册清理）。
func renderWithFont(t *testing.T, src, fontPath string) (gantt.RenderResult, string) {
	t.Helper()
	requireReadableFont(t, fontPath)

	out := filepath.Join(os.TempDir(), fmt.Sprintf("%s.png", sanitizeName(t.Name())))
	res, err := gantt.Render(t.Context(), gantt.Input{
		Source:     src,
		OutputPath: out,
		FontPath:   fontPath,
	})
	if res.OutputPath != "" {
		t.Cleanup(func() { _ = os.Remove(res.OutputPath) })
	}
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if res.OutputPath == "" {
		t.Fatalf("expected output path")
	}
	ensurePNGReadable(t, res.OutputPath)
	return res, res.OutputPath
}

func requireReadableFont(t *testing.T, path string) {
	t.Helper()
	if path == "" {
		t.Skip("font path not provided")
	}
	if _, err := os.Stat(path); err != nil {
		t.Skipf("font not available: %v", err)
	}
}

func ensurePNGReadable(t *testing.T, path string) {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open output: %v", err)
	}
	defer f.Close()
	if _, err := png.Decode(f); err != nil {
		t.Fatalf("decode png: %v", err)
	}
}

func sanitizeName(name string) string {
	return strings.ReplaceAll(name, "/", "_")
}
