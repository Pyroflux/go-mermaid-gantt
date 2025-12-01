package render_test

import (
	"os"
	"path/filepath"
	"testing"
)

// 确保提供字体时，中文渲染可产生可读 PNG（若字体缺失则跳过）。
func TestRender_WithFont_ChineseReadable(t *testing.T) {
	fontPath := firstExisting(
		filepath.Join("..", "..", "testdata", "fonts", "NotoSansSC-Regular.otf"),
		filepath.Join("..", "..", "testdata", "fonts", "NotoSansSC-Regular.ttf"),
		filepath.Join("..", "..", "testdata", "fonts", "font.ttf"),
	)
	if fontPath == "" {
		t.Skip("no test font available")
	}
	srcPath := filepath.Join("..", "..", "testdata", "alpine_font", "chinese_sample.gantt")

	data, err := os.ReadFile(srcPath)
	if err != nil {
		t.Fatalf("read sample: %v", err)
	}

	res, _ := renderWithFont(t, string(data), fontPath)
	if len(res.Bytes) == 0 {
		t.Fatalf("expected png bytes")
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
