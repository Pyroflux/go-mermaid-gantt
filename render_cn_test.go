package go_mermaid_gantt

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRender_ChineseText(t *testing.T) {
	src := `
gantt
    title 【All】初始号安全优化1期
    dateFormat YYYY-MM-DD
    excludes weekends
    section 研发
    【前端】初始号安全优化1期 PC :T8李政恩, 2026-03-13, 2026-03-19
    【移动端】功能开发 :T8廖检成, 2026-03-12, 2026-03-16
    【触屏端】初始号安全优化 :T8何俊文, 2026-03-11, 2026-03-16
    【后端】非搜索相关 :T8翁晓彤, 2026-03-11, 2026-03-16
    【后端】 搜索相关 :T8王建民, 2026-03-11, 2026-03-16
    section 测试
    【测试】功能验证 :T8曾令涛, 2026-03-17, 2026-03-18

`
	out := filepath.Join(os.TempDir(), "gantt_cn.png")
	res, err := Render(t.Context(), Input{
		Source:     src,
		OutputPath: out,
	})
	if err != nil {
		t.Fatalf("render chinese failed: %v", err)
	}
	if res.OutputPath == "" {
		t.Fatalf("expected output path")
	}
	if info, err := os.Stat(out); err != nil || info.Size() == 0 {
		t.Fatalf("output file missing or empty: %v", err)
	}
	t.Logf("rendered chinese case to %s", out)
}
