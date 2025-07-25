# 贡献指南

感谢您对本项目的关注！我们欢迎任何形式的贡献，包括但不限于提交 Issue、提出功能建议、修复 Bug、添加新功能等。

## 开发流程

1. Fork 本仓库
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 代码规范

### Go 代码规范

- 遵循 [Effective Go](https://golang.org/doc/effective_go.html) 指南
- 使用 `gofmt` 格式化代码
- 使用 `golint` 检查代码风格
- 使用 `go vet` 检查代码错误

### 提交信息规范

我们使用 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/) 规范：

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

常见的提交类型：
- `feat`: 新功能
- `fix`: 修复 Bug
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

## 代码审查

所有 Pull Request 都需要经过代码审查：

1. 至少需要 2 人审查通过
2. 审查人员会检查代码质量、功能实现、测试覆盖等
3. 审查通过后才能合并到主分支

## 测试

- 所有新功能都需要包含相应的单元测试
- 确保所有测试通过后再提交 Pull Request
- 测试覆盖率应保持在 80% 以上

## 文档

- 新功能需要提供相应的文档说明
- API 变更需要更新 API 文档
- 重要的设计决策需要记录在 ADR (Architecture Decision Records) 中

## 问题报告

如果您发现了 Bug 或有功能建议，请通过 GitHub Issues 提交：

1. 搜索是否已存在相关 Issue
2. 如果不存在，创建新的 Issue
3. 按照模板填写相关信息

## 行为准则

请遵守我们的行为准则，共同维护一个开放、友好的社区环境。
