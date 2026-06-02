package tracer

import "context"

type traceContextKey string

const (
	TraceIDKey traceContextKey = "trace_id"
	SpanIDKey traceContextKey = "span_id"
)


func WithTrace (
	ctx context.Context,
	traceID string,
	spanID string,
) context.Context {
	ctx = context.WithValue(
		ctx,
		TraceIDKey,
		traceID,
	)
	ctx = context.WithValue(
		ctx,
		SpanIDKey,
		spanID,
	)
	return ctx	
}

func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(
		TraceIDKey,
	).(string)

	if !ok {
		return ""
	}

	return traceID
}

func GetSpanID(ctx context.Context) string {
	spanID, ok := ctx.Value(
		SpanIDKey,
	).(string)

	if !ok {
		return ""
	}

	return spanID
}