# 面试问答准备 - Web3 Pledge项目

## 1. 加密货币行情集成

### 问题：请详细描述您是如何实现KuCoin API集成的，以及遇到了哪些挑战？

**回答**：
"我使用了KuCoin官方提供的Go SDK进行集成。主要实现在`kucoin.go`文件中的`GetExchangePrice`函数。集成过程包括以下几个关键步骤：

首先，我设置API认证信息并创建API服务实例：
```go
s := kucoin.NewApiService(
    kucoin.ApiKeyOption("key"),
    kucoin.ApiSecretOption("secret"),
    kucoin.ApiPassPhraseOption("passphrase"),
    kucoin.ApiKeyVersionOption(ApiKeyVersionV2),
)
```

然后，获取WebSocket连接令牌并建立连接：
```go
rsp, err := s.WebSocketPublicToken()
// 处理连接令牌
tk := &kucoin.WebSocketTokenModel{}
if err := rsp.ReadData(tk); err != nil {
    // 错误处理
}
c := s.NewWebSocketClient(tk)
mc, ec, err := c.Connect()
```

最后，订阅特定交易对的市场行情：
```go
ch := kucoin.NewSubscribeMessage("/market/ticker:PLGR-USDT", false)
if err := c.Subscribe(ch); err != nil {
    // 错误处理
}
```

主要挑战包括：
1. 处理WebSocket连接的稳定性，确保在网络波动时能够自动重连
2. 数据持久化，我们使用Redis来缓存最新价格，解决了API暂时不可用的问题
3. 错误处理策略，确保在API调用失败时系统能够优雅降级"

### 问题：为什么选择使用Redis存储价格数据？有没有考虑过其他方案？

**回答**：
"我选择Redis主要基于以下几个考虑：

1. 高性能：Redis是内存数据库，读写操作速度极快，能满足高频价格查询的需求
2. 简单性：对于简单的键值对存储（如加密货币价格），Redis提供了直观的接口
3. 可设置TTL：可以为价格数据设置过期时间，确保不使用过时数据
4. 轻量级：相比传统数据库，Redis占用资源少，启动快

在代码中，我们通过以下方式进行了实现：
```go
_ = db.RedisSetString("plgr_price", PlgrPrice, 0)
```

并在启动时尝试从Redis恢复价格：
```go
price, err := db.RedisGetString("plgr_price")
if err != nil {
    log.Logger.Sugar().Error("get plgr price from redis err ", err)
} else {
    PlgrPrice = price
}
```

我确实考虑过其他方案，比如：
- 使用常规数据库如MySQL：但对于简单的键值对，这过于重量级
- 内存缓存如Memcached：但Redis提供更丰富的数据结构和操作
- 文件系统缓存：但IO操作相对较慢，不适合高频访问场景"

## 2. 实时数据推送服务

### 问题：您如何确保WebSocket服务的可靠性和性能？

**回答**：
"在`ws.go`文件中，我实现了一个稳定可靠的WebSocket服务。关键措施包括：

1. 心跳机制：客户端发送'ping'，服务器回复'pong'，定期检测连接状态：
```go
if string(message) == "ping" || string(message) == `"ping"` || string(message) == "'ping'" {
    s.LastTime = time.Now().Unix()
    s.SendToClient("pong", PongCode)
}
```

2. 超时检测：定期检查最后一次心跳时间，超过设定时间自动关闭连接：
```go
if time.Now().Unix()-s.LastTime >= UserPingPongDurTime {
    s.SendToClient("heartbeat timeout", ErrorCode)
    return
}
```

3. 并发控制：使用互斥锁保护共享资源，避免并发写入问题：
```go
func (s *Server) SendToClient(data string, code int) {
    s.Lock()
    defer s.Unlock()
    // 发送消息
}
```

4. 错误处理：专门设计的错误通道集中处理各种异常情况：
```go
errChan := make(chan error)
// 在goroutine中发送错误
errChan <- errors.New("write message error")
// 在主循环中处理错误
case err := <-errChan:
    log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
    return
```

5. 资源清理：连接关闭时确保资源被正确释放：
```go
defer func() {
    Manager.Servers.Delete(s)
    _ = s.Socket.Close()
    close(s.Send)
}()
```

这些措施共同确保了我们的WebSocket服务能够稳定运行并高效处理大量并发连接。"

### 问题：如何处理广播消息和客户端数量增长的问题？

**回答**：
"我们的WebSocket服务采用了一个集中式的管理器(ServerManager)来处理广播消息传递和客户端管理。

对于广播消息，特别是价格更新，我们使用以下代码进行处理：
```go
func StartServer() {
    for {
        select {
        case price, ok := <-kucoin.PlgrPriceChan:
            if ok {
                Manager.Servers.Range(func(key, value interface{}) bool {
                    value.(*Server).SendToClient(price, SuccessCode)
                    return true
                })
            }
        }
    }
}
```

我们使用Go的`sync.Map`来存储所有客户端连接，这是一个为并发读写优化的映射结构：
```go
type ServerManager struct {
    Servers    sync.Map
    Broadcast  chan []byte
    Register   chan *Server
    Unregister chan *Server
}
```

当客户端数量增长时：
1. 通过`sync.Map`避免了普通map需要的全局锁，提高了并发性能
2. 每个客户端连接使用独立的goroutine处理，充分利用Go的并发优势
3. 消息通过channel传递，避免了共享内存的并发问题
4. 客户端连接和断开处理通过Register和Unregister通道集中管理

这种设计使系统能够线性扩展，支持数千并发连接，同时保持较低的资源消耗。"

## 3. 系统架构设计

### 问题：请介绍一下项目的整体架构设计，以及您为什么这样设计？

**回答**：
"从README.md可以看出，该项目分为两个主要部分：

```
# pledge-backend
API服务：cd api && go run pledge_api.go
定时任务：cd schedule && go run pledge_task.go
```

这种分离架构基于以下考虑：

1. **关注点分离**：API服务专注于处理外部HTTP/WebSocket请求，而定时任务处理后台作业，使代码更清晰

2. **独立扩展**：两个服务可以独立扩展，例如API服务可以水平扩展处理更多用户请求，而定时任务可以垂直扩展处理更复杂的计算

3. **故障隔离**：一个服务的故障不会直接影响另一个服务，提高了系统的可靠性

从代码组织上，我采用了模块化结构：
- api/models：业务模型和核心逻辑
- api/routes：API路由和处理函数
- config：配置管理
- db：数据库交互层
- log：日志系统

这种组织方式提高了代码的可维护性和可测试性。例如，数据库操作被抽象到db包中，使得业务逻辑层可以专注于实现业务规则，而不需要关心数据持久化细节。"

### 问题：系统是如何处理配置管理的？

**回答**：
"从代码中可以看出，我们使用了集中式的配置管理方案。在`ws.go`文件中有这样的引用：

```go
var UserPingPongDurTime = config.Config.Env.WssTimeoutDuration // seconds
```

这表明项目使用一个全局的Config对象来管理配置项。这种设计有以下优势：

1. 集中管理：所有配置集中在一处，便于维护和修改
2. 环境适配：通过Env子配置，可以轻松应对不同环境(开发、测试、生产)的配置需求
3. 类型安全：相比环境变量或配置文件直接解析，强类型的配置对象提供了更好的类型安全性

配置项通常包括：
- 数据库连接信息
- API密钥和认证信息
- 超时设置(如WebSocket连接超时)
- 服务器监听端口
- 日志级别等

在启动时，系统会从配置文件(通常是JSON或YAML)加载这些配置，并在需要时注入到各个组件中。这种模式使得配置变更不需要修改代码，提高了系统的灵活性和可维护性。"

## 4. 性能优化

### 问题：您在项目中做了哪些性能优化？能否分享一些具体的例子？

**回答**：
"在Pledge项目中，我进行了多方面的性能优化：

1. **并发处理**：大量使用Go的goroutine和channel进行并发处理。例如，在`ws.go`中，每个WebSocket连接都有独立的goroutine处理读写操作：

```go
//write
go func() {
    for {
        select {
        case message, ok := <-s.Send:
            // 处理消息发送
        }
    }
}()

//read
go func() {
    for {
        _, message, err := s.Socket.ReadMessage()
        // 处理接收消息
    }
}()
```

2. **内存优化**：使用Redis作为价格数据缓存，减少了不必要的API调用和内存消耗：

```go
_ = db.RedisSetString("plgr_price", PlgrPrice, 0)
```

3. **锁的精细化管理**：在需要同步的地方使用互斥锁，但范围尽可能小，减少锁竞争：

```go
func (s *Server) SendToClient(data string, code int) {
    s.Lock()
    defer s.Unlock()
    // 锁保护的代码块尽可能小
}
```

4. **有效的资源管理**：使用defer确保资源及时释放，防止资源泄漏：

```go
defer func() {
    Manager.Servers.Delete(s)
    _ = s.Socket.Close()
    close(s.Send)
}()
```

5. **数据结构选择**：使用`sync.Map`而非普通map来存储WebSocket连接，优化了并发读写性能：

```go
type ServerManager struct {
    Servers    sync.Map
    // 其他字段
}
```

这些优化措施使得系统能够高效处理大量并发请求，保持低延迟响应，即使在高负载情况下也能保持稳定。"

### 问题：在处理WebSocket连接时，您是如何平衡资源使用和性能的？

**回答**：
"在WebSocket服务设计中，平衡资源使用和性能是个关键挑战。我的策略包括：

1. **心跳超时机制**：主动关闭不活跃的连接，释放系统资源：

```go
if time.Now().Unix()-s.LastTime >= UserPingPongDurTime {
    s.SendToClient("heartbeat timeout", ErrorCode)
    return
}
```

2. **消息缓冲区**：每个WebSocket连接设置了消息发送通道，防止阻塞主逻辑：

```go
type Server struct {
    // 其他字段
    Send     chan []byte
}
```

3. **高效的广播机制**：价格广播使用单一循环遍历所有活跃连接，避免为每个广播创建新goroutine：

```go
Manager.Servers.Range(func(key, value interface{}) bool {
    value.(*Server).SendToClient(price, SuccessCode)
    return true
})
```

4. **错误快速处理**：遇到错误时快速关闭连接并清理资源，而不是尝试恢复可能已损坏的连接：

```go
case err := <-errChan:
    log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
    return
```

5. **资源限制**：通过配置限制单个连接的超时时间，防止资源被耗尽：

```go
var UserPingPongDurTime = config.Config.Env.WssTimeoutDuration // seconds
```

这些策略共同确保系统在高并发情况下仍能保持良好性能，同时避免资源泄漏和过度消耗。特别是在加密货币交易这类对实时性要求高的场景中，这种平衡尤为重要。" 