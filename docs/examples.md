# Examples

[简体中文](examples.zh-CN.md) | [Docs Index](README.md)

## Example 1: Add the Module to an Application

```bash
mkdir gnalloy-app && cd gnalloy-app
go mod init example.com/gnalloy-app
go get gnalloy.org/codec-http2@dev
go doc gnalloy.org/codec-http2
```

## Example 2: Inspect Current Packages

The current source tree exposes these package import paths:
- `gnalloy.org/codec-http2`
- `gnalloy.org/codec-http2/chunked`
- `gnalloy.org/codec-http2/content`
- `gnalloy.org/codec-http2/defense`
- `gnalloy.org/codec-http2/h2c`
- `gnalloy.org/codec-http2/http1bridge`
- `gnalloy.org/codec-http2/scheduler`

Use `go doc` against the package that matches the behavior you need:

```bash
go doc gnalloy.org/codec-http2
go doc gnalloy.org/codec-http2/chunked
go doc gnalloy.org/codec-http2/content
go doc gnalloy.org/codec-http2/defense
go doc gnalloy.org/codec-http2/h2c
go doc gnalloy.org/codec-http2/http1bridge
go doc gnalloy.org/codec-http2/scheduler
```

Selected current exported entry points:
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

## Example 3: Use Executable Tests as Behavioral Examples

Repository tests are executable examples of supported behavior. Start with the selected names below, then read the matching `_test.go` files for complete setup and assertions. See [Testing and Performance](testing.md) for the complete discovered list.

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -run Test -count=1
```

Selected current test, benchmark, fuzz, and example entry points:
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

## Example 4: Cross-Module Assembly

Direct Gnalloy dependencies for this module:
- `gnalloy.org/codec-compression`
- `gnalloy.org/codec-http1`
- `gnalloy.org/gnalloy`

Assembly guidance:
- Use this codec above a byte-oriented or datagram transport and below application handlers.
- The codec converts bytes or Gnalloy messages into protocol objects and converts outbound protocol objects back to bytes.
- It does not open sockets, own EventLoops, or define application lifecycle.

## Example 5: Pressure-Test Harness

For sustained load, wire this module into a scenario under `gnalloy.org/benchmarks` or a runnable client under `gnalloy.org/examples` when the module participates in network traffic. Record host, OS, CPU, Go version, protocol, payload, concurrency, warmup, repetitions, throughput, and p99 latency in the report.
