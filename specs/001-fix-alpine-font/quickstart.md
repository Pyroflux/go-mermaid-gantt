# Quickstart — Alpine 中文渲染修复验证

1) 准备字体  
- 将可读中文字体放置/挂载到容器：`/tmp/font.ttf`（可自定义路径）。  
- 确认可读：`ls -l /tmp/font.ttf && file /tmp/font.ttf`.

2) 在 Alpine 容器运行示例  
```bash
# 进入仓库根目录后，挂载字体并指定输出
go run ./examples/basic \
  -font /tmp/font.ttf \
  -output /tmp/gantt.png
```
（若不传 `-font`，可设置环境变量 `GGM_FONT_PATH=/tmp/font.ttf`；默认使用临时输出路径）。

3) 验证结果  
- 检查 `/tmp/gantt.png` 中文是否可读且无方块/缺字。  
- 若渲染失败，错误信息应包含字体路径及原因（不存在/不可读/损坏）；使用环境变量或自动发现字体时会在输出中给出警告。

4) 常见故障排查  
- 路径不存在或权限：`stat /tmp/font.ttf`。  
- 字体损坏：`ftvalid /tmp/font.ttf`（若可用）或尝试其他字体。  
- 未设置字体：确保传入 `-font` 或设置 `GGM_FONT_PATH`。

5) 清理  
- 删除输出文件：`rm -f /tmp/gantt.png`.
