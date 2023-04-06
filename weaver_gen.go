package main

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "weaver-start\\Reverser",
		Iface: reflect.TypeOf((*Reverser)(nil)).Elem(),
		New:   func() any { return &reverser{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return reverser_local_stub{impl: impl.(Reverser), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return reverser_client_stub{stub: stub, reverseMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "weaver-start\\Reverser", Method: "Reverse"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return reverser_server_stub{impl: impl.(Reverser), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "weaver-start\\Shuffler",
		Iface: reflect.TypeOf((*Shuffler)(nil)).Elem(),
		New:   func() any { return &shuffler{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return shuffler_local_stub{impl: impl.(Shuffler), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return shuffler_client_stub{stub: stub, shuffleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "weaver-start\\Shuffler", Method: "Shuffle"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return shuffler_server_stub{impl: impl.(Shuffler), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type reverser_local_stub struct {
	impl   Reverser
	tracer trace.Tracer
}

func (s reverser_local_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Reverse(ctx, a0)
}

type shuffler_local_stub struct {
	impl   Shuffler
	tracer trace.Tracer
}

func (s shuffler_local_stub) Shuffle(ctx context.Context, a0 string) (r0 string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.Shuffler.Shuffle", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Shuffle(ctx, a0)
}

// Client stub implementations.

type reverser_client_stub struct {
	stub           codegen.Stub
	reverseMetrics *codegen.MethodMetrics
}

func (s reverser_client_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	start := time.Now()
	s.reverseMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.reverseMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.reverseMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.reverseMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.reverseMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

type shuffler_client_stub struct {
	stub           codegen.Stub
	shuffleMetrics *codegen.MethodMetrics
}

func (s shuffler_client_stub) Shuffle(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	start := time.Now()
	s.shuffleMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.Shuffler.Shuffle", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.shuffleMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.shuffleMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.shuffleMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		return
	}
	s.shuffleMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

// Server stub implementations.

type reverser_server_stub struct {
	impl    Reverser
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s reverser_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Reverse":
		return s.reverse
	default:
		return nil
	}
}

func (s reverser_server_stub) reverse(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Reverse(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type shuffler_server_stub struct {
	impl    Shuffler
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s shuffler_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Shuffle":
		return s.shuffle
	default:
		return nil
	}
}

func (s shuffler_server_stub) shuffle(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Shuffle(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}
