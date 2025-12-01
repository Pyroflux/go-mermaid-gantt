---

description: "Task list template for feature implementation"
---

# Tasks: [FEATURE NAME]

**Input**: Design documents from `/specs/[###-feature-name]/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: 宪章要求涉及解析、排期或渲染的改动必须有可复现测试；仅当特性纯文档/样式且不触及逻辑时，可注明无需新增测试。示例包含测试任务，需根据特性具体化。

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

<!-- 
  ============================================================================
  IMPORTANT: The tasks below are SAMPLE TASKS for illustration purposes only.
  
  The /speckit.tasks command MUST replace these with actual tasks based on:
  - User stories from spec.md (with their priorities P1, P2, P3...)
  - Feature requirements from plan.md
  - Entities from data-model.md
  - Endpoints from contracts/
  
  Tasks MUST be organized by user story so each story can be:
  - Implemented independently
  - Tested independently
  - Delivered as an MVP increment
  
  DO NOT keep these sample tasks in the generated tasks.md file.
  ============================================================================
-->

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [ ] T001 根据计划调整目录（如新增 `internal/[feature]` 包）并保持 gofmt 通过
- [ ] T002 若新增依赖，更新 `go.mod`/`go.sum` 并记录用途
- [ ] T003 [P] 配置/验证 gofmt、`go vet ./...` 基线

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

Examples of foundational tasks (adjust based on your project):

- [ ] T004 建立解析/渲染所需的公共类型或常量（如时间格式、颜色定义）
- [ ] T005 [P] 补充基础错误处理与日志/调试输出钩子
- [ ] T006 [P] 准备必要的字体/资源查找策略或测试夹具
- [ ] T007 确认时间/时区/工作日配置的入口与默认值

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - [Title] (Priority: P1) 🎯 MVP

**Goal**: [Brief description of what this story delivers]

**Independent Test**: [How to verify this story works on its own]

### Tests for User Story 1 (OPTIONAL - only if tests requested) ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T010 [P] [US1] 表驱动解析测试，新增 `.gantt` 夹具到 `testdata/`
- [ ] T011 [P] [US1] 渲染测试断言 PNG 尺寸/特征，输出到临时目录并清理

### Implementation for User Story 1

- [ ] T012 [P] [US1] 在 `internal/parser` 添加/调整 AST 与调度逻辑
- [ ] T013 [P] [US1] 在 `internal/render` 实现绘制/布局改动
- [ ] T014 [US1] 更新公开入口（如 `renderer.go`）或新选项与默认值
- [ ] T015 [US1] 增加必要的校验与错误信息
- [ ] T016 [US1] 在 README/`examples/` 补充该故事用例

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - [Title] (Priority: P2)

**Goal**: [Brief description of what this story delivers]

**Independent Test**: [How to verify this story works on its own]

### Tests for User Story 2 (OPTIONAL - only if tests requested) ⚠️

- [ ] T018 [P] [US2] 针对新指令/配置的解析用例（表驱动 + 夹具）
- [ ] T019 [P] [US2] 渲染特性回归测试（PNG 字节/颜色差异）

### Implementation for User Story 2

- [ ] T020 [P] [US2] 扩展解析/调度器以支持新配置
- [ ] T021 [US2] 在渲染层实现对应表现或主题
- [ ] T022 [US2] 调整公共类型或默认值并记录兼容性
- [ ] T023 [US2] 如需，与 US1 组件集成并补充文档

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - [Title] (Priority: P3)

**Goal**: [Brief description of what this story delivers]

**Independent Test**: [How to verify this story works on its own]

### Tests for User Story 3 (OPTIONAL - only if tests requested) ⚠️

- [ ] T024 [P] [US3] 解析/渲染回归测试（包含边界时间/时区/工作日场景）
- [ ] T025 [P] [US3] 性能/尺寸基准或可视差异测试（视需求）

### Implementation for User Story 3

- [ ] T026 [P] [US3] 丰富/重构解析或渲染逻辑，保持接口稳定
- [ ] T027 [US3] 增加可配置项或主题，含回退逻辑
- [ ] T028 [US3] 补充示例、README 更新与 SemVer 影响说明

**Checkpoint**: All user stories should now be independently functional

---

[Add more user story phases as needed, following the same pattern]

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [ ] TXXX [P] 文档更新（README、examples/、testdata/ 说明）
- [ ] TXXX 代码清理与复用抽象
- [ ] TXXX 性能或内存优化，记录基准数据
- [ ] TXXX [P] 追加单测/基准测试，覆盖关键路径
- [ ] TXXX 安全与回退策略审查（字体/路径/时区）
- [ ] TXXX 验证 quickstart/示例可运行（输出到临时目录）

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2 → P3)
- **Polish (Final Phase)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - May integrate with US1 but should be independently testable
- **User Story 3 (P3)**: Can start after Foundational (Phase 2) - May integrate with US1/US2 but should be independently testable

### Within Each User Story

- Tests (if included) MUST be written and FAIL before implementation
- Models before services
- Services before endpoints
- Core implementation before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- All tests for a user story marked [P] can run in parallel
- Models within a story marked [P] can run in parallel
- Different user stories can be worked on in parallel by different team members

---

## Parallel Example: User Story 1

```bash
# Launch all tests for User Story 1 together (if tests requested):
Task: "Contract test for [endpoint] in tests/contract/test_[name].py"
Task: "Integration test for [user journey] in tests/integration/test_[name].py"

# Launch all models for User Story 1 together:
Task: "Create [Entity1] model in src/models/[entity1].py"
Task: "Create [Entity2] model in src/models/[entity2].py"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test independently → Deploy/Demo
4. Add User Story 3 → Test independently → Deploy/Demo
5. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together
2. Once Foundational is done:
   - Developer A: User Story 1
   - Developer B: User Story 2
   - Developer C: User Story 3
3. Stories complete and integrate independently

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Verify tests fail before implementing
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts, cross-story dependencies that break independence
