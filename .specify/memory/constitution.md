<!--
Sync Impact Report
- Version: N/A -> 1.0.0
- Modified Principles: 模板占位 -> 纯Go渲染与稳定API；模板占位 -> Mermaid DSL一致性与时区确定性；模板占位 -> 测试优先与基准用例；模板占位 -> 可观测性与调试友好；模板占位 -> 版本化与发布证据
- Added Sections: 附加约束与运行环境; 开发流程与质量门禁
- Removed Sections: 无
- Templates: ✅ .specify/templates/plan-template.md; ✅ .specify/templates/tasks-template.md; ⚠ pending: 无
- TODOs: 无
-->

# go-mermaid-gantt 项目宪章

## Core Principles

### I. 纯Go渲染与稳定API
保持纯Go实现，不新增外部CLI/Node依赖；`Render/Input/Theme` 等公开接口需保持向后兼容，
破坏性变更仅随 MAJOR 版本发布；同输入、同主题、同字体时输出 PNG 必须可复现。理由：减少
供应链风险，确保库能嵌入任意 Go 项目并在 CI 中确定性运行。

### II. Mermaid DSL一致性与时区确定性
解析与排期需保持 Mermaid Gantt 语法一致；新增指令或字段必须在 README/示例中记录，并提供
回退或明确错误信息。时间、时区、工作日/节假日处理必须显式（不可依赖本地时区），跨平台运
行结果保持一致。理由：保证输入 DSL 的可移植性与用户预期一致。

### III. 测试优先与基准用例
`go test ./...` 为提交前必经门槛；解析/排期改动需增加表驱动单测与 `.gantt` 夹具（放在
`testdata/`），渲染改动需断言 PNG 字节特征或尺寸而非人工目视。修复缺陷时先添加复现用例。
理由：快速回归、防止时间/区域性回归。

### IV. 可观测性与调试友好
错误需包含原因与输入片段定位；缺失字体或路径异常必须显式失败或说明回退逻辑，避免静默降
级。输出路径/Writer 行为要有可追踪的日志或测试覆盖。理由：便于集成方排查渲染问题。

### V. 版本化与发布证据
遵循 SemVer：解析/渲染可见行为变更按 Major/Minor/Patch 区分；主题或默认样式的色值变化
视为可见变更，需记录在变更说明（README 示例或发行说明）。升级 Go 版本或依赖需写明动机与
影响。理由：给集成方可预测的升级路径。

## 附加约束与运行环境
Go 版本固定为 1.24+（见 `go.mod`）；仅允许 `github.com/golang/freetype` 与
`golang.org/x/image` 为运行时依赖。字体可通过 `FontPath` 或 `GGM_FONT_PATH` 配置，缺省时
遵循内置回退，不可在代码中硬编码私有路径。不提交生成的 PNG；示例输出应写入临时目录或被
.gitignore。

## 开发流程与质量门禁
新特性需同时更新 README 示例或新增 `examples/` 以覆盖用法；任何影响 DSL/时间轴/主题的改动
必须在 `testdata/` 添加覆盖并在 PR 描述中附测试命令与结果。代码评审需检查：是否保持纯 Go、
是否有再现性测试、是否记录行为变更。发布前确认 `go fmt`、`go vet ./...`、`go test ./...`
通过。

## Governance
本宪章优先于团队习惯。任何修订必须通过 PR 记录变更、版本号调整与理据，并更新相关模板/指
南。版本号按 SemVer 维护，`Last Amended` 必须在修订日更新。评审需验证提交是否遵守核心原则
与附加约束，并在计划/任务文档中标注宪章检查项已通过。每个发布周期至少复核一次宪章适用性。

**Version**: 1.0.0 | **Ratified**: 2025-12-01 | **Last Amended**: 2025-12-01
