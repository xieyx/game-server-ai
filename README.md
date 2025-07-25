# Game Server AI

## 项目介绍

这是一个基于 Go 语言开发的游戏服务器 AI 系统，旨在提供智能化的游戏逻辑处理和玩家行为分析功能。

## 功能特性

- 智能游戏逻辑处理
- 玩家行为分析
- 实时数据处理
- 可扩展的插件系统

## 技术栈

- Go 1.19+
- Gin Web Framework
- Redis for caching
- PostgreSQL for data storage
- Docker for containerization

## 快速开始

### 环境要求

- Go 1.19+
- Docker
- PostgreSQL
- Redis

### 安装步骤

1. 克隆项目
   ```bash
   git clone git@github.com:xieyx/game-server-ai.git
   cd game-server-ai
   ```

2. 安装依赖
   ```bash
   go mod tidy
   ```

3. 配置环境变量
   ```bash
   cp .env.example .env
   # 编辑 .env 文件，配置数据库和Redis连接信息
   ```

4. 运行服务
   ```bash
   go run cmd/main.go
   ```

## 项目结构

```
.
├── cmd/                  # 主程序入口
├── internal/            # 内部包
│   ├── handlers/        # HTTP请求处理
│   ├── services/        # 业务逻辑
│   └── models/          # 数据模型
├── pkg/                 # 公共包
│   └── utils/           # 工具函数
├── tests/               # 测试文件
├── docs/                # 文档
├── .github/             # GitHub 配置
└── README.md            # 项目说明
```

## 开发指南

### 分支策略

- `main`: 稳定版本分支
- `develop`: 开发分支
- `feature/*`: 功能开发分支
- `hotfix/*`: 紧急修复分支

### 提交规范

- 使用 conventional commits 规范
- 每个提交应该包含相关的 issue 编号

### 代码审查

- 所有 PR 需要至少 2 人审查
- 审查通过后才能合并到主分支

## 贡献指南

欢迎提交 Issue 和 Pull Request 来改进项目。

### 如何贡献

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 许可证

本项目采用 MIT 许可证，详情请见 LICENSE 文件。
