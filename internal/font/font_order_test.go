package font

import (
	"os"
	"path/filepath"
	"testing"
)

// 显式 FontPath 必须优先于环境变量。
func TestSelectFontPath_FontPathOverridesEnv(t *testing.T) {
	fontPath := firstExisting(
		filepath.Join("..", "..", "testdata", "fonts", "NotoSansSC-Regular.otf"),
		filepath.Join("..", "..", "testdata", "fonts", "NotoSansSC-Regular.ttf"),
		filepath.Join("..", "..", "testdata", "fonts", "font.ttf"),
	)
	if fontPath == "" {
		t.Skip("font not available")
	}

	t.Setenv("GGM_FONT_PATH", "/nonexistent/env-font.ttf")
	selected, err := SelectFontPath(fontPath)
	if err != nil {
		t.Fatalf("select font: %v", err)
	}
	if selected != fontPath {
		t.Fatalf("expected %s, got %s", fontPath, selected)
	}
}

func firstExisting(paths ...string) string {
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}
