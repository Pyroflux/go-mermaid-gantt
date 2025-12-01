# Implementation Plan: 修复 Alpine 中文渲染乱码

**Branch**: `001-fix-alpine-font` | **Date**: 2025-12-01 | **Spec**: specs/001-fix-alpine-font/spec.md
**Input**: Feature specification from `/specs/001-fix-alpine-font/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

在 Alpine 等最小镜像中渲染包含中文的甘特图不应出现乱码或方块；当字体路径无效或缺字时应给出可诊断的错误，并提供容器内快速验证的指引。方案保持纯 Go，依赖用户指定字体或明确回退，不新增外部依赖，并通过示例/夹具和 PNG 断言验证。

## Technical Context

<!--
  ACTION REQUIRED: Replace the content in this section with the technical details
  for the project. The structure here is presented in advisory capacity to guide
  the iteration process.
-->

**Language/Version**: Go 1.24  
**Primary Dependencies**: stdlib；github.com/golang/freetype；golang.org/x/image  
**Storage**: N/A  
**Testing**: go test ./...；表驱动解析测试 + `testdata` 夹具；渲染 PNG 字节/尺寸差异断言  
**Target Platform**: Alpine Linux 容器（musl）及通用 Linux  
**Project Type**: 单一 Go 库  
**Performance Goals**: 保持现有渲染性能与确定性；字体加载不显著增加启动/渲染时间（<5% 回归）  
**Constraints**: 纯 Go，不引入外部 CLI/Node；字体路径/环境必须显式可诊断；输出可复现  
**Scale/Scope**: 库级特性改进，无新增服务或二进制

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- 保持纯 Go；不得引入外部 CLI/Node 依赖，如需新增需在计划中说明理由与替代方案，并获得宪章豁免。
- 计划必须列出测试策略：`go test ./...` 全量、解析改动的表驱动 + `testdata/*.gantt` 夹具、渲染改动的 PNG 字节/尺寸断言，含清理临时文件。
- 明确时间/时区/工作日处理与字体策略，不依赖本地时区；说明 `FontPath`/`GGM_FONT_PATH` 的配置或回退。
- 标记潜在可见行为变更及对应的 SemVer 影响和验证证据（README/示例更新）。

当前方案符合门槛：保持现有 Go 依赖；新增 testdata 与渲染断言；明确 `FontPath`/`GGM_FONT_PATH` 优先级与回退；任何默认字体或错误信息的可见变更将记录并视为可观察变更。

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)

```text
.
├── renderer.go / theme.go / types.go
├── internal/
│   ├── parser/        # Mermaid DSL 解析与调度
│   ├── render/        # 绘制与布局
│   └── font/          # 字体加载与回退
├── examples/          # 示例（含 basic）
├── testdata/          # 解析/渲染夹具（mermaid_full/*.gantt）
└── specs/001-fix-alpine-font/  # 本特性文档
```

**Structure Decision**: 单一 Go 库，改动集中在 `internal/font`、`internal/render`、`renderer.go` 及示例/夹具。

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |
