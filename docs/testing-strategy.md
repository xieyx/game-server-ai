# 测试策略

## 1. 测试原则

### 1.1 测试金字塔
遵循测试金字塔原则，合理分配不同层次的测试：

```
        ┌─────────────┐
        │   E2E测试   │  <- 少量
        ├─────────────┤
        │  集成测试   │  <- 适量
        ├─────────────┤
        │  单元测试   │  <- 大量
        └─────────────┘
```

### 1.2 测试覆盖目标
- 单元测试覆盖率：80%以上
- 关键业务逻辑：100%覆盖
- API接口测试：100%覆盖

### 1.3 测试独立性
- 测试之间相互独立，不依赖其他测试的执行结果
- 每个测试都应能独立运行
- 测试数据应能自动清理

## 2. 测试类型

### 2.1 单元测试
单元测试是最基础的测试类型，用于验证代码的最小可测试单元。

#### 2.1.1 测试范围
- 函数级别的测试
- 方法级别的测试
- 结构体行为测试

#### 2.1.2 测试规范
```go
func TestUserService_CreateUser(t *testing.T) {
    // 准备测试数据
    userService := NewUserService(mockDB)

    // 执行被测试函数
    user, err := userService.CreateUser("testuser", "test@example.com")

    // 验证结果
    assert.NoError(t, err)
    assert.Equal(t, "testuser", user.Username)
    assert.Equal(t, "test@example.com", user.Email)
}
```

#### 2.1.3 Mock策略
- 使用接口抽象外部依赖
- 使用 testify/mock 或 mockery 生成 mock 对象
- 避免直接 mock 具体实现

### 2.2 集成测试
集成测试验证不同模块之间的交互是否正确。

#### 2.2.1 测试范围
- 数据库操作测试
- 外部API调用测试
- 缓存交互测试

#### 2.2.2 测试环境
- 使用测试容器（testcontainers）模拟外部服务
- 使用独立的测试数据库
- 确保测试环境与生产环境一致性

#### 2.2.3 测试示例
```go
func TestUserRepository_CreateUser(t *testing.T) {
    // 启动测试数据库容器
    container := startPostgresContainer()
    defer container.Terminate()

    // 初始化数据库连接
    db := connectToTestDB(container)
    repo := NewUserRepository(db)

    // 执行测试
    user := &User{Username: "testuser", Email: "test@example.com"}
    createdUser, err := repo.CreateUser(user)

    // 验证结果
    assert.NoError(t, err)
    assert.NotZero(t, createdUser.ID)
    assert.Equal(t, "testuser", createdUser.Username)
}
```

### 2.3 端到端测试
端到端测试验证整个系统是否按预期工作。

#### 2.3.1 测试范围
- 完整的业务流程测试
- API接口测试
- 用户界面测试（如果有）

#### 2.3.2 测试工具
- 使用 httptest 包进行HTTP测试
- 使用 testify/assert 进行断言
- 使用 testify/suite 组织测试套件

#### 2.3.3 测试示例
```go
func TestUserAPI(t *testing.T) {
    // 启动测试服务器
    server := startTestServer()
    defer server.Close()

    // 创建HTTP客户端
    client := &http.Client{}

    // 测试创建用户
    user := User{Username: "testuser", Email: "test@example.com"}
    jsonData, _ := json.Marshal(user)

    resp, err := client.Post(server.URL+"/users", "application/json", bytes.NewBuffer(jsonData))

    // 验证响应
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
```

## 3. 测试工具

### 3.1 单元测试工具
- `testing` - Go标准库测试包
- `github.com/stretchr/testify` - 断言和mock库
- `github.com/golang/mock` - mock生成工具

### 3.2 集成测试工具
- `github.com/testcontainers/testcontainers-go` - 测试容器
- `github.com/DATA-DOG/go-sqlmock` - SQL mock
- `github.com/bxcodec/faker` - 测试数据生成

### 3.3 性能测试工具
- `testing` - Go标准库基准测试
- `github.com/uber-go/atomic` - 原子操作性能测试

## 4. 测试执行策略

### 4.1 本地测试
开发者在本地开发环境中应执行：

```bash
# 运行所有测试
go test ./...

# 运行指定包的测试
go test ./internal/user/...

# 运行测试并显示覆盖率
go test -cover ./...

# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 4.2 CI/CD测试
在CI/CD流水线中执行：

1. 单元测试 - 快速反馈
2. 代码质量检查 - golint, go vet
3. 安全扫描 - govulncheck
4. 集成测试 - 完整的集成验证
5. 性能测试 - 关键路径性能验证

### 4.3 测试环境管理
- 开发环境：本地Docker环境
- CI环境：GitHub Actions容器环境
- 预发布环境：与生产环境一致的测试环境

## 5. 测试数据管理

### 5.1 测试数据生成
- 使用faker库生成随机测试数据
- 使用fixtures管理固定测试数据
- 避免在测试中使用真实用户数据

### 5.2 测试数据清理
- 每个测试执行前准备干净的测试数据
- 每个测试执行后清理测试数据
- 使用事务回滚机制简化数据清理

```go
func TestWithTransaction(t *testing.T) {
    // 开始事务
    tx := db.Begin()
    defer tx.Rollback()

    // 在事务中执行测试
    repo := NewUserRepository(tx)
    user, err := repo.CreateUser("testuser", "test@example.com")

    // 验证结果
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)

    // 事务自动回滚，无需手动清理
}
```

## 6. 测试报告和监控

### 6.1 测试报告
- 生成详细的测试执行报告
- 记录测试覆盖率统计
- 提供失败测试的详细信息

### 6.2 持续监控
- 监控测试通过率趋势
- 监控代码覆盖率变化
- 设置测试质量门禁

### 6.3 告警机制
- 测试失败时及时通知
- 覆盖率下降时发出警告
- 性能退化时触发告警

## 7. 测试最佳实践

### 7.1 测试命名
- 使用描述性的测试名称
- 遵循 `Test[被测单元]_[测试场景]` 的命名规范
- 避免使用模糊的测试名称

```go
// 好的示例
func TestUserService_CreateUser_WithValidData(t *testing.T) { ... }
func TestUserService_CreateUser_WithDuplicateUsername(t *testing.T) { ... }

// 避免的示例
func TestCreateUser1(t *testing.T) { ... }
func TestCreateUser2(t *testing.T) { ... }
```

### 7.2 测试组织
- 按功能模块组织测试文件
- 使用测试套件组织相关的测试
- 保持测试文件与被测试文件的一致性

### 7.3 测试维护
- 定期审查和更新测试用例
- 删除过时的测试
- 重构测试代码以提高可维护性

## 8. 性能测试

### 8.1 基准测试
使用Go内置的基准测试功能：

```go
func BenchmarkUserService_CreateUser(b *testing.B) {
    userService := NewUserService(mockDB)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        userService.CreateUser(fmt.Sprintf("user%d", i), fmt.Sprintf("user%d@example.com", i))
    }
}
```

### 8.2 性能监控
- 监控关键API的响应时间
- 监控数据库查询性能
- 监控内存和CPU使用情况

### 8.3 性能优化
- 识别性能瓶颈
- 优化慢查询
- 减少不必要的内存分配
