# 概览

[English](overview.md) | [文档索引](README.zh-CN.md)

## 目标

Gnalloy HTTP/2 帧、HPACK、stream 流控、h2c、HTTP/1 桥接与防护编解码器。

该模块位于 transport 之上、业务 handler 之下，负责把字节或 Gnalloy 消息转换成协议对象，并把出站协议对象转换回字节。它不打开 socket，也不拥有 EventLoop。

## 仓库身份

- 模块路径：`gnalloy.org/codec-http2`
- GitHub 仓库：`github.com/gnalloy/codec-http2`
- 默认分支：`dev`
- 许可证：Apache-2.0

## 包结构
- `gnalloy.org/codec-http2`（`http2`）
- `gnalloy.org/codec-http2/chunked`（`chunked`）
- `gnalloy.org/codec-http2/content`（`content`）
- `gnalloy.org/codec-http2/defense`（`defense`）
- `gnalloy.org/codec-http2/h2c`（`h2c`）
- `gnalloy.org/codec-http2/http1bridge`（`http1bridge`）
- `gnalloy.org/codec-http2/scheduler`（`scheduler`）

## 直接 Gnalloy 依赖
- `gnalloy.org/gnalloy`
- `gnalloy.org/codec-http1`
- `gnalloy.org/codec-compression`

## 当前模块规划中的直接下游
- `gnalloy.org/recipes`

## 架构位置

Gnalloy 保持核心小而轻依赖。本仓库围绕单一职责形成可替换模块，通过显式 Go package 连接，而不是依靠运行时发现。

## 兼容性

公共导入路径是 `gnalloy.org/codec-http2`。首个稳定 tag 发布前，请按依赖策略使用 `@dev` 或明确的 pseudo-version。
