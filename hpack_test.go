package http2

import (
	"errors"
	"testing"

	"gnalloy.org/gnalloy/buffer"
	"gnalloy.org/gnalloy/channel"
	"gnalloy.org/gnalloy/channel/embedded"
)

func TestHeaderCodecRoundTripWithContinuation(t *testing.T) {
	encoder, err := NewHeaderEncoder(HeaderCodecConfig{MaxFrameSize: 8})
	if err != nil {
		t.Fatal(err)
	}
	outbound, err := embedded.New(encoder)
	if err != nil {
		t.Fatal(err)
	}
	defer outbound.FinishAndReleaseAll()

	wrote, err := outbound.WriteOutbound(HeadersBlock{
		StreamID:  1,
		EndStream: true,
		Fields: []HeaderField{
			{Name: ":method", Value: "GET"},
			{Name: ":path", Value: "/resource"},
			{Name: "x-large", Value: "abcdefghijklmnopqrstuvwxyz"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !wrote {
		t.Fatal("headers frame was not emitted")
	}

	decoder, err := NewHeaderDecoder(HeaderCodecConfig{})
	if err != nil {
		t.Fatal(err)
	}
	inbound, err := embedded.New(decoder)
	if err != nil {
		t.Fatal(err)
	}
	defer inbound.FinishAndReleaseAll()

	frames := 0
	for {
		msg, ok := outbound.ReadOutbound()
		if !ok {
			break
		}
		frames++
		if _, err := inbound.WriteInbound(msg); err != nil {
			t.Fatal(err)
		}
	}
	if frames < 2 {
		t.Fatalf("frames=%d, want continuation split", frames)
	}
	msg, ok := inbound.ReadInbound()
	if !ok {
		t.Fatal("missing decoded headers")
	}
	headers := msg.(HeadersBlock)
	if !headers.EndStream || len(headers.Fields) != 3 || headers.Fields[2].Value != "abcdefghijklmnopqrstuvwxyz" {
		t.Fatalf("headers=%+v", headers)
	}
}

func TestHeaderEncoderTransfersFrameOwnershipOnWriteError(t *testing.T) {
	wantErr := errors.New("write failed")
	encoder, err := NewHeaderEncoder(HeaderCodecConfig{})
	if err != nil {
		t.Fatal(err)
	}
	ch := channel.NewLocalChannel(1, buffer.NewHeapAllocator(), failingSink{err: wantErr})
	if err := ch.Pipeline().AddLast("headers", encoder); err != nil {
		t.Fatal(err)
	}

	err = ch.Write(HeadersBlock{
		StreamID: 1,
		Fields:   []HeaderField{{Name: ":status", Value: "200"}},
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("err=%v, want %v", err, wantErr)
	}
}

func BenchmarkHeaderEncoderEncodeFieldsSteadyState(b *testing.B) {
	encoder, err := NewHeaderEncoder(HeaderCodecConfig{})
	if err != nil {
		b.Fatal(err)
	}
	fields := []HeaderField{
		{Name: ":status", Value: "200"},
		{Name: "content-type", Value: "application/octet-stream"},
		{Name: "content-length", Value: "128"},
	}
	if _, err := encoder.encodeFields(fields); err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, err := encoder.encodeFields(fields); err != nil {
			b.Fatal(err)
		}
	}
}

func TestHeaderDecoderDecodeFieldsAllocationBudget(t *testing.T) {
	encoder, err := NewHeaderEncoder(HeaderCodecConfig{})
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := encoder.encodeFields([]HeaderField{
		{Name: ":method", Value: "GET"},
		{Name: ":path", Value: "/"},
		{Name: ":scheme", Value: "https"},
		{Name: ":authority", Value: "localhost"},
	})
	if err != nil {
		t.Fatal(err)
	}
	block := buffer.NewSharedBuffer(append([]byte(nil), encoded...))
	decoder, err := NewHeaderDecoder(HeaderCodecConfig{})
	if err != nil {
		t.Fatal(err)
	}
	allocs := testing.AllocsPerRun(1000, func() {
		fields, decodeErr := decoder.decodeFields(block)
		if decodeErr != nil {
			t.Fatal(decodeErr)
		}
		decodedHeaderFieldsSink = fields
	})
	if allocs > 2 {
		t.Fatalf("allocs=%f, want <= 2", allocs)
	}
}

func BenchmarkHeaderDecoderDecodeFieldsSteadyState(b *testing.B) {
	encoder, err := NewHeaderEncoder(HeaderCodecConfig{})
	if err != nil {
		b.Fatal(err)
	}
	encoded, err := encoder.encodeFields([]HeaderField{
		{Name: ":method", Value: "GET"},
		{Name: ":path", Value: "/"},
		{Name: ":scheme", Value: "https"},
		{Name: ":authority", Value: "localhost"},
	})
	if err != nil {
		b.Fatal(err)
	}
	block := buffer.NewSharedBuffer(append([]byte(nil), encoded...))
	decoder, err := NewHeaderDecoder(HeaderCodecConfig{})
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		fields, decodeErr := decoder.decodeFields(block)
		if decodeErr != nil {
			b.Fatal(decodeErr)
		}
		decodedHeaderFieldsSink = fields
	}
}

var decodedHeaderFieldsSink []HeaderField

func TestHeaderDecoderRemainsSynchronizedAfterHeaderListLimit(t *testing.T) {
	encoder, err := NewHeaderEncoder(HeaderCodecConfig{})
	if err != nil {
		t.Fatal(err)
	}
	oversized, err := encoder.encodeFields([]HeaderField{{Name: "x-large", Value: "0123456789012345678901234567890123456789"}})
	if err != nil {
		t.Fatal(err)
	}
	oversized = append([]byte(nil), oversized...)
	valid, err := encoder.encodeFields([]HeaderField{{Name: ":method", Value: "GET"}})
	if err != nil {
		t.Fatal(err)
	}
	valid = append([]byte(nil), valid...)

	decoder, err := NewHeaderDecoder(HeaderCodecConfig{MaxHeaderListSize: 64})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decoder.decodeFields(buffer.NewSharedBuffer(oversized)); !errors.Is(err, ErrHeaderListTooLarge) {
		t.Fatalf("err=%v, want ErrHeaderListTooLarge", err)
	}
	fields, err := decoder.decodeFields(buffer.NewSharedBuffer(valid))
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 || fields[0].Name != ":method" || fields[0].Value != "GET" {
		t.Fatalf("fields=%+v", fields)
	}
}

func TestHeaderDecoderRejectsContinuationWithoutHeaders(t *testing.T) {
	decoder, err := NewHeaderDecoder(HeaderCodecConfig{})
	if err != nil {
		t.Fatal(err)
	}
	ch, err := embedded.New(decoder, exceptionCapture{})
	if err != nil {
		t.Fatal(err)
	}
	defer ch.FinishAndReleaseAll()

	block := buffer.NewHeapBuffer(1)
	if _, err := block.WriteBytes([]byte{0}); err != nil {
		t.Fatal(err)
	}
	if _, err := ch.WriteInbound(ContinuationFrame{StreamID: 1, Flags: FlagEndHeaders, HeaderBlock: block}); err != nil {
		t.Fatal(err)
	}
	msg, ok := ch.ReadInbound()
	if !ok {
		t.Fatal("missing exception")
	}
	if err, ok := msg.(error); !ok || !errors.Is(err, ErrHeaderBlock) {
		t.Fatalf("msg=%v, want ErrHeaderBlock", msg)
	}
}

func TestStreamMultiplexerAcceptsDecodedHeadersBlock(t *testing.T) {
	mux, err := NewStreamMultiplexer(MultiplexerConfig{Server: true})
	if err != nil {
		t.Fatal(err)
	}
	ch, err := embedded.New(mux)
	if err != nil {
		t.Fatal(err)
	}
	defer ch.FinishAndReleaseAll()

	if _, err := ch.WriteInbound(HeadersBlock{StreamID: 1, Fields: []HeaderField{{Name: ":method", Value: "GET"}}}); err != nil {
		t.Fatal(err)
	}
	msg, ok := ch.ReadInbound()
	if !ok {
		t.Fatal("missing active event")
	}
	if event := msg.(StreamEvent); event.Type != StreamEventActive || event.StreamID != 1 {
		t.Fatalf("event=%+v", event)
	}
	msg, ok = ch.ReadInbound()
	if !ok {
		t.Fatal("missing read event")
	}
	event := msg.(StreamEvent)
	if event.Type != StreamEventRead {
		t.Fatalf("event=%+v", event)
	}
	if _, ok := event.Frame.(HeadersBlock); !ok {
		t.Fatalf("frame=%T, want HeadersBlock", event.Frame)
	}
}

type exceptionCapture struct{}

func (exceptionCapture) ExceptionCaught(ctx *channel.HandlerContext, err error) {
	ctx.FireChannelRead(err)
}
