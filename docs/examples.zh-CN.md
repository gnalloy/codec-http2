# 案例

[English](examples.md) | [文档索引](README.zh-CN.md)

## 案例 1：将模块加入应用

```bash
mkdir gnalloy-app && cd gnalloy-app
go mod init example.com/gnalloy-app
go get gnalloy.org/codec-http2@dev
go doc gnalloy.org/codec-http2
```

## 案例 2：查看当前包

当前源码树暴露这些 package 导入路径：
- `gnalloy.org/codec-http2`
- `gnalloy.org/codec-http2/chunked`
- `gnalloy.org/codec-http2/content`
- `gnalloy.org/codec-http2/defense`
- `gnalloy.org/codec-http2/h2c`
- `gnalloy.org/codec-http2/http1bridge`
- `gnalloy.org/codec-http2/scheduler`

按需要的行为对对应 package 执行 `go doc`：

```bash
go doc gnalloy.org/codec-http2
go doc gnalloy.org/codec-http2/chunked
go doc gnalloy.org/codec-http2/content
go doc gnalloy.org/codec-http2/defense
go doc gnalloy.org/codec-http2/h2c
go doc gnalloy.org/codec-http2/http1bridge
go doc gnalloy.org/codec-http2/scheduler
```

精选当前导出入口：
- `const ClientPreface = "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n" ...`
- `const SettingHeaderTableSize uint16 = 0x1 ...`
- `const ErrorCodeCancel uint32 = 0x8`
- `var ErrInvalidFrame = errors.New("gnalloy/codec/http2: invalid frame") ...`
- `func AppendFrameHeader(dst []byte, h FrameHeader) ([]byte, error)`
- `type ConnectionController struct{ ... }`
- `type DataChunkedInput struct{ ... }`
- `type WriteHandler struct{}`
- `func NewDataCompressingInput(streamID http2.StreamID, input codec.ChunkedInput, coding Coding, ...) (*h2chunked.DataChunkedInput, error)`
- `type Coding string`
- `type DataCompressingInputConfig struct{ ... }`
- `type Decompressor struct{ ... }`
- `type ResponseCompressor struct{ ... }`
- `type ResponseCompressorConfig struct{ ... }`
- `var ErrTooManyRSTFrames = errors.New("gnalloy/codec/http2/defense: too many rst_stream frames") ...`
- `type ControlFrameLimitConfig struct{ ... }`
- `type ControlFrameLimitEncoder struct{ ... }`
- `type MaxRstFrameConfig struct{ ... }`

## 案例 3：将可执行测试作为行为示例

仓库测试是受支持行为的可执行示例。先从下面的精选名称开始，再阅读对应 `_test.go` 文件中的完整 setup 和断言。完整发现列表见 [测试与性能](testing.zh-CN.md)。

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -run Test -count=1
```

精选当前 test、benchmark、fuzz 与 example 入口：
- `BenchmarkStreamMultiplexerReadData`
- `FuzzHTTP2FramePipeline`
- `TestApplyUpgradeHeadersAddsH2CHeaders`
- `TestCompressorCompressesAcceptedResponseStream`
- `TestConnectionControllerAppliesAndValidatesSettings`
- `TestConnectionControllerAppliesReceiveWindowUpdate`
- `TestConnectionControllerEnforcesMaxConcurrentRemoteStreams`
- `TestConnectionControllerEnforcesReceiveConnectionWindow`
- `TestConnectionControllerRejectsNewLocalStreamAfterGoAway`
- `TestConnectionControllerRejectsPushWhenPeerDisabled`
- `TestControlFrameLimitEncoderRejectsUnflushedControlFrames`
- `TestDataChunkedInputEmitsEmptyEndStream`
- `TestDataChunkedInputSplitsByteBufIntoDataFrames`
- `TestDataCompressingInputStreamsHTTP2DataFrames`
- `TestDecompressorDecodesGzipDataFrame`
- `TestEncodeDecodeHTTP2SettingsHeader`
- `TestFrameDecoderDecodesPayloadZeroCopy`
- `TestFrameEncoderCoalescesHeaderAndPayload`

## 案例 4：跨模块装配

本模块的直接 Gnalloy 依赖：
- `gnalloy.org/codec-compression`
- `gnalloy.org/codec-http1`
- `gnalloy.org/gnalloy`

装配说明：
- codec 位于面向字节或 datagram 的 transport 之上、应用 handler 之下。
- 它负责把字节或 Gnalloy 消息转换成协议对象，并把出站协议对象转换回字节。
- 它不打开 socket，不拥有 EventLoop，也不定义应用生命周期。

## 案例 5：压测 Harness

持续负载测试时，如果该模块参与网络流量路径，将它接入 `gnalloy.org/benchmarks` 的场景，或接入 `gnalloy.org/examples` 的可运行客户端。报告中记录 host、OS、CPU、Go version、protocol、payload、concurrency、warmup、repetitions、throughput 和 p99 latency。
