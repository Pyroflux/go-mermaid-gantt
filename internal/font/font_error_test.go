package font

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSelectFontPath_NotFound(t *testing.T) {
	t.Setenv("GGM_FONT_PATH", "")
	_, err := SelectFontPath("/path/does/not/exist.ttf")
	if err == nil {
		t.Fatalf("expected error for missing font")
	}
	if !strings.Contains(err.Error(), "/path/does/not/exist.ttf") {
		t.Fatalf("error should mention missing path, got: %v", err)
	}
}

func TestSelectFontPath_Damaged(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "broken.ttf")
	if writeErr := os.WriteFile(tmp, []byte("not-a-font"), 0o600); writeErr != nil {
		t.Fatalf("write tmp: %v", writeErr)
	}
	t.Setenv("GGM_FONT_PATH", "")
	_, err := SelectFontPath(tmp)
	if err == nil {
		t.Fatalf("expected error for damaged font")
	}
	if !strings.Contains(err.Error(), tmp) {
		t.Fatalf("error should mention path: %v", err)
	}
}
