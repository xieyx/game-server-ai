# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- 项目初始化
- README.md 文档
- 用户管理功能实现
  - 用户模型定义
  - 数据库连接和迁移
  - 用户服务层实现
  - HTTP处理层实现
  - 主程序入口和路由配置
- 测试框架搭建
  - 用户服务测试用例
- 开发环境配置
  - .env.example环境变量模板
  - docker-compose.yml容器编排配置
- JWT认证功能
  - JWT工具包实现
  - 登录接口返回token

### Changed
- 完善README.md文档，添加详细使用说明
- 扩展GitHub Actions工作流，增加代码质量检查和安全扫描
- 修复用户服务中检查用户名和邮箱是否已存在的逻辑错误
- 实现GetUser方法，从数据库获取真实用户信息
- 修正测试代码中的问题

### Deprecated

### Removed

### Fixed
- 修复用户服务中检查用户名和邮箱是否已存在的逻辑错误
- 修复GetUser方法中使用硬编码测试数据的问题
- 修复测试代码中的MockDB结构和方法调用问题

### Security
- 用户密码加密存储
- 添加JWT token认证机制
