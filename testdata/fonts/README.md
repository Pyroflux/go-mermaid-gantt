# 字体测试资源

此目录用于放置可读中文字体文件，供渲染与回归测试使用。请使用开源可再分发的字体（推荐：
Noto Sans SC Regular 或等价字体）。为了避免仓库膨胀，未随仓库提交字体二进制文件；请在本地或
CI 拉取后将字体放置于此目录，并在下方记录来源与许可证。

建议文件名：
- `NotoSansSC-Regular.otf` 或 `NotoSansSC-Regular.ttf`

仓库不包含字体二进制文件（如 `font.ttf`），请本地或 CI 准备后放置于此目录，并记录来源与许可证。

需要记录的元数据：
- 字体名称与版本
- 下载来源 URL
- 许可证（例如 SIL Open Font License 1.1）

测试依赖：
- 相关测试会引用 `testdata/fonts/NotoSansSC-Regular.otf`（或 .ttf），请确保路径与大小写一致。
