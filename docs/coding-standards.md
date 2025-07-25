# 编码规范

## 1. 命名规范

### 1.1 包命名
- 包名应简洁、小写，不使用下划线
- 包名应与其所在目录名一致
- 避免使用复数形式

```go
// 好的示例
package handlers
package models
package utils

// 避免的示例
package Handlers  // 不要使用大写
package util      // 应该是 utils
package my_handlers // 不要使用下划线
```

### 1.2 接口命名
- 单个方法的接口以 "er" 后缀命名
- 多个方法的接口使用描述性名称

```go
// 单个方法接口
type Reader interface {
    Read(p []byte) (n int, err error)
}

// 多个方法接口
type UserService interface {
    Create(user *User) error
    Get(id uint) (*User, error)
    Update(user *User) error
    Delete(id uint) error
}
```

### 1.3 结构体命名
- 使用驼峰命名法
- 名称应清晰表达其用途

```go
type User struct {
    ID       uint
    Username string
    Email    string
}

type GameSession struct {
    ID        uint
    PlayerID  uint
    StartTime time.Time
    EndTime   *time.Time
}
```

### 1.4 函数和方法命名
- 使用驼峰命名法
- 函数名应清晰表达其功能
- Getter 方法省略 Get 前缀

```go
// 好的示例
func (u *User) IsValid() bool { ... }
func (gs *GameSession) Duration() time.Duration { ... }
func NewUserService() UserService { ... }

// 避免的示例
func (u *User) GetIsValid() bool { ... }  // 不需要 Get 前缀
```

### 1.5 变量命名
- 使用驼峰命名法
- 局部变量可使用简短命名
- 导出变量使用描述性命名

```go
// 循环变量可以简短
for i, user := range users { ... }

// 函数参数可以简短
func CreateUser(u *User) error { ... }

// 导出变量应具有描述性
var (
    MaxRetryAttempts = 3
    DefaultTimeout   = 30 * time.Second
)
```

## 2. 代码格式化

### 2.1 使用 gofmt
所有代码必须使用 `gofmt` 格式化：

```bash
gofmt -w .
```

### 2.2 行长度
- 每行代码不应超过 120 个字符
- 注释行不应超过 80 个字符

### 2.3 缩进
- 使用制表符进行缩进
- 一个制表符等于 8 个空格

### 2.4 空行
- 不同的代码块之间使用空行分隔
- 函数之间使用空行分隔
- 逻辑相关的代码块不需要空行分隔

```go
func Example() {
    // 变量声明
    var a int
    var b string

    // 逻辑处理
    if a > 0 {
        b = "positive"
    } else {
        b = "non-positive"
    }

    // 返回结果
    return b
}
```

## 3. 注释规范

### 3.1 包注释
每个包都应该有包注释，位于 package 语句之前：

```go
// Package handlers 提供 HTTP 请求处理功能
package handlers
```

### 3.2 函数注释
导出函数应有详细注释，说明功能、参数和返回值：

```go
// CreateUser 创建一个新的用户
// 参数:
//   - username: 用户名
//   - email: 邮箱地址
// 返回值:
//   - *User: 创建的用户对象
//   - error: 错误信息
func CreateUser(username, email string) (*User, error) {
    // 实现
}
```

### 3.3 结构体注释
导出结构体应有注释说明其用途：

```go
// User 表示系统中的用户
type User struct {
    // ID 是用户的唯一标识符
    ID uint

    // Username 是用户的登录名
    Username string

    // Email 是用户的邮箱地址
    Email string
}
```

### 3.4 内联注释
- 内联注释应简洁明了
- 注释应解释"为什么"而不是"是什么"
- 复杂逻辑应有详细注释

```go
// 使用互斥锁保护并发访问
mu.Lock()
defer mu.Unlock()

// 重试机制，最多尝试3次
for i := 0; i < 3; i++ {
    if err := doSomething(); err != nil {
        time.Sleep(time.Second)
        continue
    }
    break
}
```

## 4. 错误处理

### 4.1 错误处理原则
- 错误应被显式处理
- 不要忽略错误返回值
- 使用自定义错误类型提供更多信息

```go
// 好的示例
if err := doSomething(); err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// 避免的示例
doSomething() // 忽略错误
```

### 4.2 自定义错误类型
定义具有上下文信息的错误类型：

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}
```

### 4.3 错误日志记录
- 记录错误日志便于调试
- 包含足够的上下文信息
- 避免记录敏感信息

```go
if err := db.Create(user); err != nil {
    log.Printf("Failed to create user %s: %v", user.Username, err)
    return fmt.Errorf("failed to create user: %w", err)
}
```

## 5. 测试规范

### 5.1 测试文件命名
测试文件应与被测试文件同名，但以 `_test.go` 结尾：

```
user.go
user_test.go
```

### 5.2 测试函数命名
测试函数应以 `Test` 开头，后跟被测试函数名：

```go
func TestCreateUser(t *testing.T) { ... }
func TestUser_IsValid(t *testing.T) { ... }
```

### 5.3 使用测试库
推荐使用 `testify` 库简化测试代码：

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
    user, err := CreateUser("testuser", "test@example.com")

    require.NoError(t, err)
    assert.Equal(t, "testuser", user.Username)
    assert.Equal(t, "test@example.com", user.Email)
}
```

### 5.4 表驱动测试
对于多个测试用例，使用表驱动测试：

```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        email   string
        isValid bool
    }{
        {"valid@example.com", true},
        {"invalid.email", false},
        {"", false},
    }

    for _, tt := range tests {
        t.Run(tt.email, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if tt.isValid {
                assert.NoError(t, err)
            } else {
                assert.Error(t, err)
            }
        })
    }
}
```

## 6. 性能优化

### 6.1 避免不必要的内存分配
- 重用对象而不是频繁创建新对象
- 使用对象池管理频繁创建的对象

```go
// 使用 strings.Builder 避免字符串拼接的内存分配
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")
result := builder.String()
```

### 6.2 合理使用 Goroutines
- 不要为每个请求都创建新的 Goroutine
- 使用 worker pool 模式处理并发任务
- 注意 Goroutine 泄漏问题

```go
// 使用 worker pool 处理并发任务
func workerPool(jobs <-chan Job, results chan<- Result, numWorkers int) {
    for i := 0; i < numWorkers; i++ {
        go func() {
            for job := range jobs {
                result := processJob(job)
                results <- result
            }
        }()
    }
}
```

### 6.3 数据库查询优化
- 使用预编译语句
- 合理使用索引
- 避免 N+1 查询问题

```go
// 使用预编译语句
stmt, err := db.Prepare("SELECT id, username FROM users WHERE id = ?")
if err != nil {
    return nil, err
}
defer stmt.Close()

// 批量查询避免 N+1 问题
func GetUsersWithPosts(userIDs []uint) ([]User, error) {
    // 一次性查询所有用户及其文章
    // 而不是为每个用户单独查询文章
}
```

## 7. 安全规范

### 7.1 输入验证
- 所有外部输入都应验证
- 使用正则表达式验证格式
- 限制输入长度

```go
func ValidateUsername(username string) error {
    if len(username) < 3 || len(username) > 20 {
        return errors.New("username length must be between 3 and 20")
    }

    matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
    if !matched {
        return errors.New("username can only contain letters, numbers, and underscores")
    }

    return nil
}
```

### 7.2 敏感信息处理
- 不要在日志中记录敏感信息
- 使用环境变量存储密钥
- 加密存储敏感数据

```go
// 不要在日志中记录密码
log.Printf("User login: %s", username) // 好的
// log.Printf("User login: %s, password: %s", username, password) // 避免的

// 使用环境变量存储密钥
jwtSecret := os.Getenv("JWT_SECRET")
```

### 7.3 SQL 注入防护
- 使用参数化查询
- 不要拼接 SQL 字符串

```go
// 好的示例
row := db.QueryRow("SELECT id, username FROM users WHERE id = ?", userID)

// 避免的示例
query := fmt.Sprintf("SELECT id, username FROM users WHERE id = %d", userID)
row := db.QueryRow(query)
```

## 8. 日志规范

### 8.1 日志级别
- Debug: 调试信息，仅在开发环境使用
- Info: 一般信息，记录重要业务流程
- Warn: 警告信息，可预期的错误
- Error: 错误信息，需要关注的问题

### 8.2 结构化日志
使用结构化日志便于分析：

```go
log.Printf(`{"level":"info","event":"user_login","user_id":%d,"timestamp":"%s"}`,
    userID, time.Now().Format(time.RFC3339))
```

### 8.3 日志上下文
在日志中包含足够的上下文信息：

```go
// 包含请求ID便于追踪
log.Printf("RequestID: %s, User login: %s", requestID, username)
