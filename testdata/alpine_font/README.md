# Alpine 中文渲染夹具

此目录存放与 Alpine/musl 环境相关的中文渲染输入与说明。

## 使用方式

1. 准备字体：将开源中文字体放到 `testdata/fonts/`（示例名称 `font.ttf` 或 `NotoSansSC-Regular.otf`），在容器内挂载到 `/tmp/font.ttf`。
2. 渲染示例：
```bash
go run ./examples/basic -font /tmp/font.ttf -output /tmp/gantt.png
```
或通过环境变量：
```bash
GGM_FONT_PATH=/tmp/font.ttf go run ./examples/basic -output /tmp/gantt.png
```
3. 预期：输出 PNG 中中文可读无乱码；若字体路径无效或缺字，错误信息包含路径与原因。
