package http2

import "gnalloy.org/gnalloy/channel"

// PingAckHandler 为非 ACK PING 帧写出 PING ACK，并保留原帧继续向后传播。
type PingAckHandler struct{}

// NewPingAckHandler 创建 HTTP/2 PING ACK 生命周期处理器。
func NewPingAckHandler() *PingAckHandler {
	return &PingAckHandler{}
}

// ChannelRead 对齐 RFC 7540：PING ACK 回显原始八字节载荷，且不能再次确认 ACK。
func (h *PingAckHandler) ChannelRead(ctx *channel.HandlerContext, msg any) {
	ping, ok := msg.(PingFrame)
	if !ok || ping.Ack {
		ctx.FireChannelRead(msg)
		return
	}
	ack := ping
	ack.Ack = true
	if err := ctx.WriteAndFlush(ack); err != nil {
		ctx.FireExceptionCaught(err)
		return
	}
	ctx.FireChannelRead(msg)
}
