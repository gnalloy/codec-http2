# 用法

[English](usage.md) | [文档索引](README.zh-CN.md)

## 要求

- Go 1.25 或更新版本，并与 module 的 `go` 指令一致。
- 由 Gnalloy 应用、recipe、example 或 benchmark harness 负责生命周期与部署配置。
- 独立模块复验应设置 `GOWORK=off`，确保通过已发布依赖图测试。

## 安装
```bash
go get gnalloy.org/codec-http2@dev
```

## 导入
```go
import "gnalloy.org/codec-http2"
```

## 集成模式
- frame、header、body 与 decoded-content 上限必须由服务的可信边界决定。
- 大 payload 应使用 streaming 或 chunked 模式，避免无界聚合正文。
- 压缩模块必须配置解码后大小限制，防御膨胀攻击。
- ByteBuf 所有权遵守 Gnalloy 消息规则：只有当前组件真正消费消息后才释放。
- HTTP/2 over TLS 要求 TLS 1.2 或更新版本，并使用 ALPN `h2`；h2c 需要作为明文 upgrade 路径显式配置。

## API 选择

通过 API 清单选择当前协议路径需要的具体构造函数或 option 类型：

```bash
go doc gnalloy.org/codec-http2
```

当前常用入口：
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

## 跨模块装配

多个 Gnalloy 仓库一起开发时，在自己选择的 workspace 中创建本地 `go.work` 文件。不要把应用本地 `replace` 指令提交到发布用 library module，除非它是明确的临时变更且不会进入提交。

## 错误处理

网络输入、对端行为、平台能力和超时失败都必须作为普通错误处理。不要用 panic 恢复协议正确性。返回或传播模块错误，并在所有权要求时关闭受影响的 Channel。
