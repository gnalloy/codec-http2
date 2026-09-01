# API Reference

[简体中文](api.zh-CN.md) | [Docs Index](README.md)

This inventory is generated from `go doc -short` for the packages in this repository. It is a quick public-surface map; source files and tests remain the authority for exact semantics.

## Packages

### `gnalloy.org/codec-http2`

Package name: `http2`

```text
const ClientPreface = "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n" ...
const SettingHeaderTableSize uint16 = 0x1 ...
const ErrorCodeCancel uint32 = 0x8
var ErrInvalidFrame = errors.New("gnalloy/codec/http2: invalid frame") ...
func AppendFrameHeader(dst []byte, h FrameHeader) ([]byte, error)
type ConnectionController struct{ ... }
    func NewConnectionController(cfg ConnectionControllerConfig) (*ConnectionController, error)
type ConnectionControllerConfig struct{ ... }
type ContinuationFrame struct{ ... }
type DataFrame struct{ ... }
type Flags uint8
    const FlagEndStream Flags = 0x1 ...
type Frame struct{ ... }
    func EncodeTypedFrame(ctx *channel.HandlerContext, msg any) (Frame, bool, error)
    func SettingsAck() Frame
type FrameDecoder struct{ ... }
    func NewFrameDecoder(maxFrameSize int) (*FrameDecoder, error)
type FrameEncoder struct{ ... }
    func NewFrameEncoder() *FrameEncoder
    func NewFrameEncoderWithConfig(cfg FrameEncoderConfig) *FrameEncoder
type FrameEncoderConfig struct{ ... }
type FrameHeader struct{ ... }
    func ParseFrameHeader(src []byte) (FrameHeader, error)
type FrameType uint8
    const FrameData FrameType = iota ...
type GoAwayFrame struct{ ... }
type HeaderCodecConfig struct{ ... }
type HeaderDecoder struct{ ... }
    func NewHeaderDecoder(cfg HeaderCodecConfig) (*HeaderDecoder, error)
type HeaderEncoder struct{ ... }
    func NewHeaderEncoder(cfg HeaderCodecConfig) (*HeaderEncoder, error)
type HeaderField struct{ ... }
type HeadersBlock struct{ ... }
type HeadersFrame struct{ ... }
type MultiplexerConfig struct{ ... }
type OutboundFlowControlConfig struct{ ... }
type OutboundFlowController struct{ ... }
    func NewOutboundFlowController(cfg OutboundFlowControlConfig) *OutboundFlowController
type PingFrame struct{ ... }
type PrefaceDecoder struct{ ... }
    func NewPrefaceDecoder() *PrefaceDecoder
type PrefaceEncoder struct{}
    func NewPrefaceEncoder() *PrefaceEncoder
type PrefaceReceivedEvent struct{}
type PriorityFrame struct{ ... }
type PriorityParam struct{ ... }
type PushPromiseBlock struct{ ... }
type PushPromiseFrame struct{ ... }
type RSTStreamFrame struct{ ... }
type Setting struct{ ... }
type SettingsAckHandler struct{}
    func NewSettingsAckHandler() *SettingsAckHandler
type SettingsFrame struct{ ... }
type SettingsReceivedEvent struct{ ... }
type SettingsSnapshot struct{ ... }
type Stream struct{ ... }
    func NewStream(id StreamID) Stream
type StreamBufferingEncoder struct{ ... }
    func NewStreamBufferingEncoder(cfg StreamBufferingEncoderConfig) *StreamBufferingEncoder
type StreamBufferingEncoderConfig struct{ ... }
type StreamChannel struct{ ... }
type StreamChildConfig struct{ ... }
type StreamChildHandler struct{ ... }
    func NewStreamChildHandler(cfg StreamChildConfig) (*StreamChildHandler, error)
type StreamChildInitializer func(ch *StreamChannel) error
type StreamEvent struct{ ... }
type StreamEventType uint8
    const StreamEventActive StreamEventType = iota + 1 ...
type StreamID uint32
type StreamMultiplexer struct{ ... }
    func NewStreamMultiplexer(cfg MultiplexerConfig) (*StreamMultiplexer, error)
type StreamState uint8
    const StreamIdle StreamState = iota ...
type TypedFrame interface{ ... }
    func DecodeTypedFrame(frame Frame) (TypedFrame, error)
type TypedFrameDecoder struct{ ... }
    func NewTypedFrameDecoder() *TypedFrameDecoder
type TypedFrameEncoder struct{}
    func NewTypedFrameEncoder() *TypedFrameEncoder
type UnknownFrame struct{ ... }
type WindowUpdateFrame struct{ ... }
```

### `gnalloy.org/codec-http2/chunked`

Package name: `chunked`

```text
type DataChunkedInput struct{ ... }
    func NewDataChunkedInput(streamID http2.StreamID, input codec.ChunkedInput, endStream bool) (*DataChunkedInput, error)
type WriteHandler struct{}
    func NewWriteHandler() *WriteHandler
```

### `gnalloy.org/codec-http2/content`

Package name: `content`

```text
func NewDataCompressingInput(streamID http2.StreamID, input codec.ChunkedInput, coding Coding, ...) (*h2chunked.DataChunkedInput, error)
type Coding string
    const CodingGzip Coding = "gzip" ...
type DataCompressingInputConfig struct{ ... }
type Decompressor struct{ ... }
    func NewDecompressor(maxDecodedBytes int) *Decompressor
type ResponseCompressor struct{ ... }
    func NewResponseCompressor(cfg ResponseCompressorConfig) *ResponseCompressor
type ResponseCompressorConfig struct{ ... }
```

### `gnalloy.org/codec-http2/defense`

Package name: `defense`

```text
var ErrTooManyRSTFrames = errors.New("gnalloy/codec/http2/defense: too many rst_stream frames") ...
type ControlFrameLimitConfig struct{ ... }
type ControlFrameLimitEncoder struct{ ... }
    func NewControlFrameLimitEncoder(cfg ControlFrameLimitConfig) *ControlFrameLimitEncoder
type MaxRstFrameConfig struct{ ... }
type MaxRstFrameDecoder struct{ ... }
    func NewMaxRstFrameDecoder(cfg MaxRstFrameConfig) *MaxRstFrameDecoder
```

### `gnalloy.org/codec-http2/h2c`

Package name: `h2c`

```text
const ProtocolName = "h2c" ...
func ApplyUpgradeHeaders(req http1.Request, settings []http2.Setting) (http1.Request, error)
func DecodeHTTP2Settings(value string) ([]http2.Setting, error)
func EncodeHTTP2Settings(settings []http2.Setting) string
```

### `gnalloy.org/codec-http2/http1bridge`

Package name: `http1bridge`

```text
var ErrInvalidHeadersBlock = errors.New("gnalloy/codec/http2/http1bridge: invalid headers block")
func HeadersBlockFromRequest(streamID http2.StreamID, req http1.Request, endStream bool) http2.HeadersBlock
func HeadersBlockFromResponse(streamID http2.StreamID, resp http1.Response, endStream bool) http2.HeadersBlock
func HeadersBlockFromTrailers(streamID http2.StreamID, trailers http1.Headers, endStream bool) http2.HeadersBlock
func RequestFromHeadersBlock(block http2.HeadersBlock) (http1.Request, error)
func ResponseFromHeadersBlock(block http2.HeadersBlock) (http1.Response, error)
type Config struct{ ... }
type PushPromise struct{ ... }
type StreamFrameToHTTPObjectCodec struct{ ... }
    func NewStreamFrameToHTTPObjectCodec(cfg Config) *StreamFrameToHTTPObjectCodec
```

### `gnalloy.org/codec-http2/scheduler`

Package name: `scheduler`

```text
var ErrInvalidWrite = errors.New("gnalloy/codec/http2/scheduler: invalid write result")
type Config struct{ ... }
type StreamState struct{ ... }
type WeightedFairQueueByteDistributor struct{ ... }
    func NewWeightedFairQueueByteDistributor(cfg Config) *WeightedFairQueueByteDistributor
type WriteFunc func(id http2.StreamID, maxBytes int) (int, error)
```
