# Tasks: 修复 Alpine 中文渲染乱码

**Input**: Design documents from `/specs/001-fix-alpine-font/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: 宪章要求涉及解析、排期或渲染的改动必须有可复现测试；本特性需新增字体相关渲染回归测试与错误路径测试。

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Go package根**: 源码在仓库根目录及 `internal/parser|render|font`。
- **测试**: 与代码同目录的 `*_test.go`；夹具放 `testdata/`（例如 `testdata/mermaid_full/*.gantt`）。
- **示例**: `examples/`（如 `examples/basic/main.go`）；输出写入临时目录或忽略文件。
- **命令**: 以 `go run ./...`、`go test ./...` 为主，必要时 `go vet ./...`。

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 在仓库根目录执行 `gofmt -w`、`go vet ./...` 确认基线无格式/静态检查问题
- [X] T002 创建 `testdata/fonts/README.md` 记录即将新增字体文件来源与许可
- [X] T003 [P] 创建 `testdata/alpine_font/` 目录用于中文甘特夹具与说明

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T004 将开源中文字体（如 NotoSansSC-Regular 或等价小体积字体）添加到 `testdata/fonts/` 并在 README 记录许可证
- [X] T005 [P] 添加中文甘特夹具 `testdata/alpine_font/chinese_sample.gantt`（含标题/任务/依赖）用于回归
- [X] T006 [P] 在 `internal/render/render_font_test.go` 添加渲染 PNG 解码与临时文件清理的测试工具函数
- [X] T007 [P] 在 `internal/font/font.go` 提取字体选择与校验的公共函数（显式路径 > 环境变量 > 回退），返回可诊断错误

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Alpine 渲染中文可读 (Priority: P1) 🎯 MVP

**Goal**: 提供字体后在 Alpine 渲染中文无乱码/方块

**Independent Test**: 使用 `testdata/fonts/` 字体和 `testdata/alpine_font/chinese_sample.gantt` 渲染到临时路径，PNG 可解码且中文可读、无空输出

### Tests for User Story 1 ⚠️

- [X] T008 [P] [US1] 在 `render_font_success_test.go` 添加渲染测试，使用 `FontPath` 指向 `testdata/fonts/*.ttf` 和中文夹具，断言 PNG 解码与尺寸非零
- [X] T009 [P] [US1] 在 `internal/font/font_order_test.go` 添加测试，验证 `Input.FontPath` 优先于 `GGM_FONT_PATH`，并记录实际使用的字体路径

### Implementation for User Story 1

- [X] T010 [P] [US1] 更新 `internal/font/font.go` 实现显式路径优先、失败即返回错误（不静默回退）
- [X] T011 [US1] 确保 `renderer.go`/`internal/render` 使用选定字体渲染并在成功路径返回稳定输出（无需触发默认回退）

**Checkpoint**: User Story 1 functional and testable independently

---

## Phase 4: User Story 2 - 字体问题可诊断 (Priority: P2)

**Goal**: 字体路径/缺字问题有明确错误或警告，避免静默生成乱码

**Independent Test**: 提供不存在/无权限/损坏字体时，渲染返回包含路径与原因的错误或警告，不产出乱码文件

### Tests for User Story 2 ⚠️

- [X] T012 [P] [US2] 在 `internal/font/font_error_test.go` 添加错误路径测试（不存在/无权限/损坏字体），断言错误包含路径与原因
- [X] T013 [P] [US2] 在 `render_font_warning_test.go` 添加无可用字体时的行为测试，验证返回的 `RenderResult.Warnings`/错误信息

### Implementation for User Story 2

- [X] T014 [US2] 改进 `internal/font/font.go` 与 `renderer.go` 的错误包装与警告收集，将路径/权限/加载原因写入返回值或 Warnings
- [X] T015 [US2] 更新 `specs/001-fix-alpine-font/contracts/rendering.md` 与 `README.md` 字体章节，描述错误/警告示例与可见行为

**Checkpoint**: User Stories 1 AND 2 functional and independently testable

---

## Phase 5: User Story 3 - 容器集成指引 (Priority: P3)

**Goal**: 提供 Alpine/最小镜像的字体配置与验证步骤，首次集成即可生成可读中文

**Independent Test**: 按指引在全新 Alpine 容器内复制/挂载字体并运行示例，首次即可生成可读 PNG

### Tests for User Story 3 ⚠️

- [X] T016 [P] [US3] 在 `examples/basic/main.go` 增加可选 font/output 标志或环境读取，验证可在容器内使用指定字体渲染

### Implementation for User Story 3

- [X] T017 [US3] 扩写 `specs/001-fix-alpine-font/quickstart.md`，补充 Alpine 挂载字体命令与验证步骤，并同步 README 字体说明
- [X] T018 [US3] 在 `testdata/alpine_font/README.md` 记录夹具用途、示例命令（含 `GGM_FONT_PATH`），便于复现实验

**Checkpoint**: All user stories should now be independently functional

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [X] T019 [P] 文档与示例收尾：检查 README/quickstart/contract 中的字体指引一致性
- [X] T020 统一执行 `go fmt`, `go vet ./...`, `go test ./...`，确保新夹具/字体文件被正确引用并无临时文件泄漏

---

## Dependencies & Execution Order

- Phase 1 → Phase 2 → User Stories (3→4→5) → Polish  
- User Story 1 (P1) 完成后可独立验证中文渲染；User Story 2 依赖 US1 的字体加载路径；User Story 3 依赖 US1 的成功渲染能力与 US2 的诊断结论。

### Within Each User Story

- 测试（若有）应先于实现并先行失败  
- 实现需遵守显式字体优先、失败可诊断、不静默回退的约束  
- 每个故事完成后可独立运行渲染用例验证

### Parallel Opportunities

- Setup 与 Foundational 中标记 [P] 的任务可并行（不同路径的目录/测试工具）  
- US1/US2 的测试任务可并行；US3 文档更新可在代码冻结后并行进行  
- 最终回归 (`go test ./...`) 必须串行在所有实现完成后执行

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. 完成 Phase 1 + Phase 2  
2. 完成 US1 测试与实现，验证中文渲染可读  
3. 暂停并评估是否满足紧急需求

### Incremental Delivery

1. 完成 US1（渲染可读）  
2. 完成 US2（诊断与错误信息）  
3. 完成 US3（容器指引与示例），更新文档  
4. 运行全量测试并收敛文档

### Parallel Team Strategy

- 一人推进 US1 渲染与字体加载逻辑；一人并行编写 US2 错误路径测试与消息；一人处理 US3 文档/示例。  
- 所有实现合并后统一回归测试与文档对齐。
