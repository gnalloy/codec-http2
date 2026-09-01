# Overview

[简体中文](overview.zh-CN.md) | [Docs Index](README.md)

## Purpose

HTTP/2 frame, HPACK, stream flow-control, h2c, HTTP/1 bridge, and defensive codecs for Gnalloy.

This module sits above transports and below application handlers. It translates bytes or Gnalloy messages into protocol objects, and translates outbound protocol objects back to bytes. It does not open sockets or own EventLoops.

## Repository Identity

- Module path: `gnalloy.org/codec-http2`
- GitHub repository: `github.com/gnalloy/codec-http2`
- Default branch: `dev`
- License: Apache-2.0

## Package Map
- `gnalloy.org/codec-http2` (`http2`)
- `gnalloy.org/codec-http2/chunked` (`chunked`)
- `gnalloy.org/codec-http2/content` (`content`)
- `gnalloy.org/codec-http2/defense` (`defense`)
- `gnalloy.org/codec-http2/h2c` (`h2c`)
- `gnalloy.org/codec-http2/http1bridge` (`http1bridge`)
- `gnalloy.org/codec-http2/scheduler` (`scheduler`)

## Direct Gnalloy Dependencies
- `gnalloy.org/gnalloy`
- `gnalloy.org/codec-http1`
- `gnalloy.org/codec-compression`

## Direct Dependents in the Current Module Plan
- `gnalloy.org/recipes`

## Architecture Position

Gnalloy keeps the core small and dependency-light. This repository is a replaceable module around one responsibility, connected through explicit Go packages instead of runtime discovery.

## Compatibility

The public import path is `gnalloy.org/codec-http2`. Until the first stable tag is published, use `@dev` or an explicit pseudo-version selected by your dependency policy.
