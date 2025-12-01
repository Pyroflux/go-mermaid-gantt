# Rendering Contract — Alpine 中文渲染

## Invocation (library)

- **Function**: `Render(ctx, Input)`  
- **Required fields**:
  - `Input.Source` (Mermaid Gantt，含中文)
  - `Input.OutputPath` 或 `Input.Writer`
- **Font selection (ordered, no silent fallback)**:
  1) `Input.FontPath`（显式路径，失效即报错）  
  2) 环境变量 `GGM_FONT_PATH`（失效即报错）  
  3) 常见可用字体路径（找到则使用并返回 warning；未找到则报错）  
- **Behavior**:
  - 选定字体不可读/损坏 → 返回错误（包含路径/原因），不生成输出
  - 字体加载成功 → 渲染 PNG，保证中文可读（无方块/乱码）
  - 若使用环境变量或自动发现的字体，`RenderResult.Warnings` 记录实际路径来源

## Error Contracts

- **Font not found**: error message MUST include missing path.
- **Font unreadable/no permission**: error message MUST include path and permission hint.
- **Font load failure (damaged/unsupported)**: error message MUST include font path and load failure reason.
- **No usable font**: explicit failure; do not produce PNG bytes (包含尝试的路径列表或来源说明)。

## Determinism

- 同样的 Source/Theme/Font 组合在 Alpine 与 glibc 环境应产出一致的 PNG（尺寸与内容一致性用于测试比对）。

## Quick Validation Example

- Input: 中文标题与任务的示例 `.gantt`
- Font: 提供的可读中文字体（挂载到 `/tmp/font.ttf`）
- Expected: `Render` 返回 PNG bytes 非空且中文可读；如路径为 `/notfound.ttf` 则报错并指明路径。
