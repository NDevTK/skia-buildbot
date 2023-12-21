package grpclogging

import (
	"bytes"
	"context"
	"testing"
	"time"

	"go.opencensus.io/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	pb "go.skia.org/infra/go/grpclogging/proto"
	tpb "go.skia.org/infra/go/grpclogging/testproto"
	"go.skia.org/infra/go/now"
	"go.skia.org/infra/go/tracing/tracingtest"
	"go.skia.org/infra/kube/go/authproxy"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	startTime = time.Date(2022, time.January, 31, 2, 2, 3, 0, time.FixedZone("UTC+1", 60*60))
)

func testSetupLogger(t *testing.T) (*now.TimeTravelCtx, *GRPCLogger, *bytes.Buffer) {
	ttCtx := now.TimeTravelingContext(startTime)
	buf := &bytes.Buffer{}
	l := New(buf)

	return ttCtx, l, buf
}

func entryFromBuf(t *testing.T, buf *bytes.Buffer) *pb.Entry {
	entry := &pb.Entry{}
	err := protojson.Unmarshal(buf.Bytes(), entry)
	require.NoError(t, err)
	return entry
}

func assertLoggedServerUnary(t *testing.T, entry *pb.Entry, req *tpb.GetSomethingRequest) {
	loggedReq := &tpb.GetSomethingRequest{}
	err := entry.ServerUnary.Request.UnmarshalTo(loggedReq)
	require.NoError(t, err)
	assert.Equal(t, loggedReq.SomethingId, req.SomethingId)

	if entry.StatusCode == int32(codes.OK) {
		loggedResp := &tpb.GetSomethingResponse{}
		err = entry.ServerUnary.Response.UnmarshalTo(loggedResp)
		require.NoError(t, err)
	}
}

func TestServerUnaryLoggingInterceptor_noError(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	req := &tpb.GetSomethingRequest{
		SomethingId: "d3c4f84d",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		ttCtx.SetTime(startTime.Add(3 * time.Second))
		return &tpb.GetSomethingResponse{}, nil
	}
	resp, err := l.ServerUnaryLoggingInterceptor(ttCtx, req, &grpc.UnaryServerInfo{FullMethod: "test.service/TestMethod"}, handler)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	entry := entryFromBuf(t, buf)

	assert.Equal(t, entry.ServerUnary.FullMethod, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
	assertLoggedServerUnary(t, entry, req)

	// Make sure that no trace fields are set if tracing isn't active.
	assert.Equal(t, "", entry.SpanId)
	assert.Equal(t, "", entry.TraceId)
	assert.False(t, entry.TraceSampled)
}

func TestServerUnaryLoggingInterceptor_withNaNs(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	req := &tpb.GetSomethingRequest{
		SomethingId: "d3c4f84d",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		ttCtx.SetTime(startTime.Add(3 * time.Second))
		return &tpb.GetSomethingResponse{
			SomethingContents: "contents",
		}, nil
	}
	resp, err := l.ServerUnaryLoggingInterceptor(ttCtx, req, &grpc.UnaryServerInfo{FullMethod: "test.service/TestMethod"}, handler)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	entry := entryFromBuf(t, buf)

	assert.Equal(t, entry.ServerUnary.FullMethod, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
	assertLoggedServerUnary(t, entry, req)
}

func TestServerUnaryLoggingInterceptor_withAuthProxyUser(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	req := &tpb.GetSomethingRequest{
		SomethingId: "d3c4f84d",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		ttCtx.SetTime(startTime.Add(3 * time.Second))
		return &tpb.GetSomethingResponse{}, nil
	}
	md := metadata.New(map[string]string{
		authproxy.WebAuthHeaderName: "user@domain.com",
	})

	resp, err := l.ServerUnaryLoggingInterceptor(
		metadata.NewIncomingContext(ttCtx.Context, md),
		req, &grpc.UnaryServerInfo{FullMethod: "test.service/TestMethod"}, handler)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	entry := entryFromBuf(t, buf)

	assert.Equal(t, entry.ServerUnary.FullMethod, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
	assert.Equal(t, "user@domain.com", entry.ServerUnary.User)
	assertLoggedServerUnary(t, entry, req)
}

func TestServerUnaryLoggingInterceptor_error(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	req := &tpb.GetSomethingRequest{
		SomethingId: "d3c4f84d",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		ttCtx.SetTime(startTime.Add(1 * time.Second))
		return nil, status.Errorf(codes.InvalidArgument, "forced error response")
	}
	resp, err := l.ServerUnaryLoggingInterceptor(ttCtx, req, &grpc.UnaryServerInfo{FullMethod: "test.service/TestMethod"}, handler)
	require.Error(t, err)
	assert.Nil(t, resp)
	entry := entryFromBuf(t, buf)

	assert.Equal(t, int32(codes.InvalidArgument), entry.StatusCode)
	assert.Equal(t, int64(1), entry.Elapsed.Seconds)
	assertLoggedServerUnary(t, entry, req)
}

func TestServerUnaryLoggingInterceptor_tracing(t *testing.T) {
	exporter := &tracingtest.Exporter{}
	trace.RegisterExporter(exporter)
	defer trace.UnregisterExporter(exporter)
	trace.ApplyConfig(trace.Config{

		DefaultSampler: trace.AlwaysSample()})

	ttCtx, l, buf := testSetupLogger(t)
	req := &tpb.GetSomethingRequest{
		SomethingId: "d3c4f84d",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		ctx, span := trace.StartSpan(ctx, "handler")
		defer span.End()
		ttCtx.SetTime(startTime.Add(3 * time.Second))
		return &tpb.GetSomethingResponse{}, nil
	}

	ctx, span := trace.StartSpan(ttCtx, "caller")
	resp, err := l.ServerUnaryLoggingInterceptor(ctx, req, &grpc.UnaryServerInfo{FullMethod: "test.service/TestMethod"}, handler)
	span.End()

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(exporter.SpanData()))
	entry := entryFromBuf(t, buf)

	assert.Equal(t, entry.ServerUnary.FullMethod, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
	assertLoggedServerUnary(t, entry, req)
	assert.NotEqual(t, "", entry.SpanId)
	assert.NotEqual(t, "", entry.TraceId)
	assert.True(t, entry.TraceSampled)
}

func TestClientUnaryLoggingInterceptor_noError(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	req := &tpb.GetSomethingRequest{
		SomethingId: "d3c4f84d",
	}
	invoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		ttCtx.SetTime(startTime.Add(3 * time.Second))
		pb, _ := reply.(*tpb.GetSomethingResponse)
		pb.SomethingContents = "something"

		return nil
	}
	resp := &tpb.GetSomethingResponse{}
	err := l.ClientUnaryLoggingInterceptor(ttCtx, "test.service/TestMethod", req, resp, nil, invoker, nil)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	entry := entryFromBuf(t, buf)
	assert.Equal(t, entry.ClientUnary.Method, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
}

func TestClientUnaryLoggingInterceptor_error(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	req := &tpb.GetSomethingRequest{
		SomethingId: "d3c4f84d",
	}
	invoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		ttCtx.SetTime(startTime.Add(3 * time.Second))

		return status.Errorf(codes.InvalidArgument, "forced error response")
	}
	resp := &tpb.GetSomethingResponse{}
	err := l.ClientUnaryLoggingInterceptor(ttCtx, "test.service/TestMethod", req, resp, nil, invoker, nil)
	require.Error(t, err)

	entry := entryFromBuf(t, buf)
	assert.Equal(t, entry.ClientUnary.Method, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
	assert.Equal(t, int32(codes.InvalidArgument), entry.StatusCode)
}

func TestClientStreamLoggingInterceptor_noError(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	streamer := func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ttCtx.SetTime(startTime.Add(3 * time.Second))

		return nil, nil
	}
	desc := &grpc.StreamDesc{
		StreamName: "test/stream",
	}
	_, err := l.ClientStreamLoggingInterceptor(ttCtx, desc, nil, "test.service/TestMethod", streamer)
	require.NoError(t, err)
	entry := entryFromBuf(t, buf)
	assert.Equal(t, entry.ClientStream.Method, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
}

func TestClientStreamLoggingInterceptor_error(t *testing.T) {
	ttCtx, l, buf := testSetupLogger(t)
	streamer := func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ttCtx.SetTime(startTime.Add(3 * time.Second))

		return nil, status.Errorf(codes.InvalidArgument, "forced error response")
	}
	desc := &grpc.StreamDesc{
		StreamName: "test/stream",
	}
	_, err := l.ClientStreamLoggingInterceptor(ttCtx, desc, nil, "test.service/TestMethod", streamer)
	require.Error(t, err)

	entry := entryFromBuf(t, buf)
	assert.Equal(t, entry.ClientStream.Method, "test.service/TestMethod")
	assert.Equal(t, int64(3), entry.Elapsed.Seconds)
	assert.Equal(t, int32(codes.InvalidArgument), entry.StatusCode)
}
