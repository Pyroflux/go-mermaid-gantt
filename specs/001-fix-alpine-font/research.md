# Phase 0 Research — 修复 Alpine 中文渲染乱码

## 决策与结论

- **Decision**: 字体加载顺序为：Input.FontPath（显式参数）优先，其次环境变量 `GGM_FONT_PATH`，最后内置字体回退；若选中的路径不可用则错误并中止，不生成乱码文件。  
  **Rationale**: 遵循用户显式意图，避免静默降级导致乱码；与现有配置方式兼容。  
  **Alternatives considered**: 自动扫描系统字体路径（/usr/share/fonts 等）但在 Alpine 极简镜像中可靠性低且增加不确定性；未采纳。

- **Decision**: 在 Alpine/musl 环境下渲染前检查字体文件可读性与必需 glyph 覆盖（通过 freetype 加载失败即报错）。  
  **Rationale**: 早期失败比生成乱码更可诊断，符合宪章的可观测性原则。  
  **Alternatives considered**: 静默尝试多个字体并继续渲染，但会产生不可控输出；未采纳。

- **Decision**: 新增最小复现输入与 PNG 基线断言：中文标题与任务，输出写入 `os.TempDir()`；测试比对 PNG 头/尺寸/非空，并在字体缺失时断言返回错误。  
  **Rationale**: 防止回归并验证字体路径/环境优先级；保持纯 Go 测试。  
  **Alternatives considered**: 仅依赖手工查看图片；不符合自动回归要求，未采纳。

- **Decision**: 提供 Alpine 容器 quickstart（复制/挂载字体、验证命令、常见错误提示）。  
  **Rationale**: 降低首次集成门槛，减少支持请求。  
  **Alternatives considered**: 仅在 README 增补一句字体说明，缺乏可操作步骤；未采纳。

## 需要澄清/假设

- 无需额外澄清：目标是修复提供字体仍乱码的场景，保持纯 Go，不新增外部依赖。
