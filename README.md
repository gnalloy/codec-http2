# codec-http2

[简体中文](README.zh-CN.md) | [Documentation](docs/README.md)

HTTP/2 frame, HPACK, stream flow-control, h2c, HTTP/1 bridge, and defensive codecs for Gnalloy.

This module sits above transports and below application handlers. It translates bytes or Gnalloy messages into protocol objects, and translates outbound protocol objects back to bytes. It does not open sockets or own EventLoops.

## Status

- Import path: `gnalloy.org/codec-http2`
- Repository: `github.com/gnalloy/codec-http2`
- Default branch: `dev`
- Preview install: `go get gnalloy.org/codec-http2@dev`
- License: Apache-2.0

## Install
```bash
go get gnalloy.org/codec-http2@dev
go doc gnalloy.org/codec-http2
GOWORK=off GOTOOLCHAIN=local go test ./... -count=1
```

## Documentation
- [Overview](docs/overview.md) ([中文](docs/overview.zh-CN.md))
- [Usage](docs/usage.md) ([中文](docs/usage.zh-CN.md))
- [Examples](docs/examples.md) ([中文](docs/examples.zh-CN.md))
- [Configuration](docs/configuration.md) ([中文](docs/configuration.zh-CN.md))
- [Testing and Performance](docs/testing.md) ([中文](docs/testing.zh-CN.md))
- [API Reference](docs/api.md) ([中文](docs/api.zh-CN.md))
- [Notes and Caveats](docs/notes.md) ([中文](docs/notes.zh-CN.md))
- [ADR-001 Module Boundary](docs/decisions/0001-module-boundary.md) ([中文](docs/decisions/0001-module-boundary.zh-CN.md))

## Module Boundary

This repository owns: HTTP/2 frame, HPACK, stream flow-control, h2c, HTTP/1 bridge, and defensive codecs for Gnalloy.

It does not absorb neighboring module responsibilities. Core primitives stay in `gnalloy.org/gnalloy`; protocol codecs, transports, handlers, resolvers, examples, and benchmarks stay in their own repositories.

## Packages
- `gnalloy.org/codec-http2` (`http2`)
- `gnalloy.org/codec-http2/chunked` (`chunked`)
- `gnalloy.org/codec-http2/content` (`content`)
- `gnalloy.org/codec-http2/defense` (`defense`)
- `gnalloy.org/codec-http2/h2c` (`h2c`)
- `gnalloy.org/codec-http2/http1bridge` (`http1bridge`)
- `gnalloy.org/codec-http2/scheduler` (`scheduler`)

## Gnalloy Dependencies
- `gnalloy.org/gnalloy`
- `gnalloy.org/codec-http1`
- `gnalloy.org/codec-compression`

## Common Integration Pattern
- Frame, header, body, and decoded-content limits must be selected from the trusted boundary of the service.
- Streaming or chunked modes should be used for large payloads instead of materializing unbounded bodies.
- Compression modules must set decoded-size limits to defend against expansion attacks.
- ByteBuf ownership follows Gnalloy message rules: release only after the current component consumes the message.
- HTTP/2 over TLS requires TLS 1.2 or newer and ALPN `h2`; h2c must be configured as a cleartext upgrade path.

## Current Public Entry Points

The generated API reference lists the full public surface. Common constructors or option types currently include:
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

## Verification

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -count=1
GOWORK=off GOTOOLCHAIN=local go vet ./...
GOWORK=off GOTOOLCHAIN=local go test ./... -run '^$' -bench . -benchmem -count=1
```

For pressure tests, assemble this module with the relevant transport, codec, and handler stack and run the scenario from `gnalloy.org/benchmarks` or `gnalloy.org/examples`. Keep host, operating system, payload, concurrency, warmup, and repetitions in the report.

## Caveats
- This repository is intentionally narrow. Cross-module behavior should be assembled in applications, recipes, examples, or benchmark harnesses.
- Public APIs should remain Go-native and explicit; avoid runtime scanning, hidden global registries, and reflection-heavy behavior in hot paths.
- Treat network input as untrusted. Configure parser limits and return typed errors instead of panics.
- Keep benchmark claims tied to a concrete host, operating system, protocol, payload, concurrency, warmup, and repetition count.
- Codec modules do not provide a network server by themselves; combine them with a transport module and application handlers.
