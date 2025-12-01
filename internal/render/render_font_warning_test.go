package render_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	gantt "github.com/pyroflux/go-mermaid-gantt"
)

// 当未提供可用字体时应返回可诊断错误，而不是生成乱码输出。
func TestRender_NoUsableFont_ReturnsError(t *testing.T) {
	t.Setenv("GGM_FONT_PATH", "")
	srcPath := filepath.Join("..", "..", "testdata", "alpine_font", "chinese_sample.gantt")
	data, err := os.ReadFile(srcPath)
	if err != nil {
		t.Fatalf("read sample: %v", err)
	}

	_, err = gantt.Render(t.Context(), gantt.Input{
		Source:     string(data),
		OutputPath: filepath.Join(t.TempDir(), "out.png"),
		FontPath:   "/path/not/exist.ttf",
	})
	if err == nil {
		t.Fatalf("expected error when font missing")
	}
	if !strings.Contains(err.Error(), "/path/not/exist.ttf") {
		t.Fatalf("error should mention missing font path, got: %v", err)
	}
}
