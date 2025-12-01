# Feature Specification: 修复 Alpine 中文渲染乱码

**Feature Branch**: `001-fix-alpine-font`  
**Created**: 2025-12-01  
**Status**: Draft  
**Input**: User description: "现在有个问题 中文在alpine镜像中渲染乱码 即使给定了字体文件路径也依然乱码"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Alpine 渲染中文可读 (Priority: P1)

作为在 Alpine 容器内使用库的开发者，我提供自带的中文字体文件，希望甘特图输出中的中文不再变成方块或乱码。

**Why this priority**: 直接影响用户可读性，是阻塞性的生产问题。

**Independent Test**: 在 Alpine 基础镜像中挂载指定字体路径，渲染包含中文标题和任务的示例，验证 PNG 中中文完全可读。

**Acceptance Scenarios**:

1. **Given** Alpine 镜像内可读的中文字体文件路径，**When** 渲染包含多行中文任务，**Then** 输出 PNG 中所有中文字符可读、无乱码/缺字。
2. **Given** 提供的字体覆盖常用中文字符，**When** 重复渲染同一输入，**Then** 输出保持一致，不因环境差异出现乱码。

---

### User Story 2 - 字体问题可诊断 (Priority: P2)

作为维护者，我需要在字体路径不可读或字体缺字时获得明确的错误或提示，便于快速修复而不是静默回退。

**Why this priority**: 缺乏诊断会导致重复踩坑、无法自助解决。

**Independent Test**: 提供不存在的字体路径或缺字字体，渲染任务并捕获返回的错误/提示，确认信息包含路径和原因。

**Acceptance Scenarios**:

1. **Given** 字体路径不存在或无权限，**When** 发起渲染，**Then** 返回信息包含问题路径和失败原因，且不会输出乱码文件。
2. **Given** 字体存在但缺少所需中文 glyph，**When** 渲染含该字符的输入，**Then** 能得到包含缺字说明的提示或明确的失败，而非静默生成乱码。

---

### User Story 3 - 容器集成指引 (Priority: P3)

作为初次集成的使用者，我希望有简明步骤在 Alpine/最小镜像里准备字体并验证输出，以避免踩坑。

**Why this priority**: 降低支持成本，提升首开体验。

**Independent Test**: 按文档在全新 Alpine 基础镜像中配置字体并运行示例，首次尝试即可得到可读的中文输出。

**Acceptance Scenarios**:

1. **Given** 新的 Alpine 容器和指引步骤，**When** 挂载/复制字体并运行示例命令，**Then** 一次即可生成包含中文的可读 PNG。
2. **Given** 指引中的检查命令，**When** 用户执行，**Then** 可确认字体文件存在且可读，减少渲染前的故障排查时间。

### Edge Cases

- 字体路径存在但无读取权限时的行为（应给出清晰错误而非乱码输出）。
- 字体文件损坏或缺少常用中文 glyph 时的提示与处理。
- 环境变量与输入字段同时提供字体时的优先级和提示一致性。
- 在无任何可用中文字体时的失败模式（不可接受生成乱码文件）。

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: 系统必须在 Alpine 等最小镜像中使用提供的字体文件正确渲染中文内容，输出无乱码或缺字。
- **FR-002**: 当字体路径不存在、不可读或文件损坏时，必须返回包含路径和原因的可操作错误信息，不得静默降级生成乱码。
- **FR-003**: 字体选择顺序需可说明（例如输入参数/环境变量/内置回退的优先级），并在输出或日志中记录实际使用的字体来源。
- **FR-004**: 提供最小可复现示例（含输入与期望输出描述），用于在 Alpine 环境验证中文渲染成功。
- **FR-005**: 补充容器集成指引，列出字体放置位置、权限要求和验证步骤，确保新用户一次通过。

### Key Entities *(include if feature involves data)*

- **渲染请求**: 包含原始 Mermaid Gantt 文本、输出目标、字体配置（路径或环境变量）。
- **字体来源**: 用户提供的字体文件或内置回退，需记录来源、可读性及是否覆盖中文 glyph。

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 在 Alpine 环境下使用指定字体渲染官方示例时，100% 中文字符可读且与基线输出一致。
- **SC-002**: 提供错误的字体路径时，错误信息包含路径和原因，用户可在一次重试内修复（无需查看源码）。
- **SC-003**: 按指引在全新 Alpine 环境配置字体并运行示例，首次成功率 ≥ 90%，操作步骤不超过 5 分钟。
- **SC-004**: 针对字体问题的支持请求/issue 数量在版本发布后一个周期内下降 ≥ 50%。
