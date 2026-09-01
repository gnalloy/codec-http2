# codec-http2

[English](README.md) | [文档](docs/README.zh-CN.md)

Gnalloy HTTP/2 帧、HPACK、stream 流控、h2c、HTTP/1 桥接与防护编解码器。

该模块位于 transport 之上、业务 handler 之下，负责把字节或 Gnalloy 消息转换成协议对象，并把出站协议对象转换回字节。它不打开 socket，也不拥有 EventLoop。

## 状态

- 导入路径：`gnalloy.org/codec-http2`
- 仓库：`github.com/gnalloy/codec-http2`
- 默认分支：`dev`
- 预览安装：`go get gnalloy.org/codec-http2@dev`
- 许可证：Apache-2.0

## 安装
```bash
go get gnalloy.org/codec-http2@dev
go doc gnalloy.org/codec-http2
GOWORK=off GOTOOLCHAIN=local go test ./... -count=1
```

## 文档
- [概览](docs/overview.zh-CN.md) ([English](docs/overview.md))
- [用法](docs/usage.zh-CN.md) ([English](docs/usage.md))
- [案例](docs/examples.zh-CN.md) ([English](docs/examples.md))
- [配置说明](docs/configuration.zh-CN.md) ([English](docs/configuration.md))
- [测试与性能](docs/testing.zh-CN.md) ([English](docs/testing.md))
- [API 参考](docs/api.zh-CN.md) ([English](docs/api.md))
- [注意事项与备注](docs/notes.zh-CN.md) ([English](docs/notes.md))
- [ADR-001 模块边界](docs/decisions/0001-module-boundary.zh-CN.md) ([English](docs/decisions/0001-module-boundary.md))

## 模块边界

本仓库负责：Gnalloy HTTP/2 帧、HPACK、stream 流控、h2c、HTTP/1 桥接与防护编解码器。

它不吸收相邻模块职责。核心基础能力保留在 `gnalloy.org/gnalloy`；协议 codec、transport、handler、resolver、examples 与 benchmarks 分别由独立仓库负责。

## 包结构
- `gnalloy.org/codec-http2`（`http2`）
- `gnalloy.org/codec-http2/chunked`（`chunked`）
- `gnalloy.org/codec-http2/content`（`content`）
- `gnalloy.org/codec-http2/defense`（`defense`）
- `gnalloy.org/codec-http2/h2c`（`h2c`）
- `gnalloy.org/codec-http2/http1bridge`（`http1bridge`）
- `gnalloy.org/codec-http2/scheduler`（`scheduler`）

## Gnalloy 依赖
- `gnalloy.org/gnalloy`
- `gnalloy.org/codec-http1`
- `gnalloy.org/codec-compression`

## 常见集成方式
- frame、header、body 与 decoded-content 上限必须由服务的可信边界决定。
- 大 payload 应使用 streaming 或 chunked 模式，避免无界聚合正文。
- 压缩模块必须配置解码后大小限制，防御膨胀攻击。
- ByteBuf 所有权遵守 Gnalloy 消息规则：只有当前组件真正消费消息后才释放。
- HTTP/2 over TLS 要求 TLS 1.2 或更新版本，并使用 ALPN `h2`；h2c 需要作为明文 upgrade 路径显式配置。

## 当前公共入口

生成的 API 参考列出了完整公共面。当前常用构造函数或 option 类型包括：
- `const ClientPreface = "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n" ...`
- `const SettingHeaderTableSize uint16 = 0x1 ...`
- `const ErrorCodeCancel uint32 = 0x8`
- `var ErrInvalidFrame = errors.New("gnalloy/codec/http2: invalid frame") ...`
- `type ConnectionControllerConfig struct{ ... }`
- `type FrameEncoderConfig struct{ ... }`
- `type HeaderCodecConfig struct{ ... }`
- `type MultiplexerConfig struct{ ... }`
- `type OutboundFlowControlConfig struct{ ... }`
- `type StreamBufferingEncoderConfig struct{ ... }`
- `type StreamChildConfig struct{ ... }`
- `func NewDataCompressingInput(streamID http2.StreamID, input codec.ChunkedInput, coding Coding, ...) (*h2chunked.DataChunkedInput, error)`
- `type DataCompressingInputConfig struct{ ... }`
- `type ResponseCompressorConfig struct{ ... }`
- `var ErrTooManyRSTFrames = errors.New("gnalloy/codec/http2/defense: too many rst_stream frames") ...`
- `type ControlFrameLimitConfig struct{ ... }`

## 验证

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -count=1
GOWORK=off GOTOOLCHAIN=local go vet ./...
GOWORK=off GOTOOLCHAIN=local go test ./... -run '^$' -bench . -benchmem -count=1
```

压测时，将该模块和相应 transport、codec、handler 栈装配后，使用 `gnalloy.org/benchmarks` 或 `gnalloy.org/examples` 中的场景运行。报告必须保留主机、操作系统、payload、并发度、warmup 和 repetition。

## 注意事项
- 本仓库保持窄边界。跨模块行为应在应用、recipes、examples 或 benchmark harness 中装配。
- 公共 API 必须保持 Go 原生和显式；热路径避免运行时扫描、隐藏全局注册表和重反射。
- 网络输入一律视为不可信。配置解析上限，返回类型化错误，不使用 panic 处理输入错误。
- 性能结论必须绑定具体主机、操作系统、协议、payload、并发度、warmup 和 repetition。
- codec 模块本身不提供网络服务器；需要和 transport 模块及应用 handler 组合。
