package tracer

import (
	"context"
	"time"
)

func StartSpan(
	service string,
	operation string,
	traceID string,
	parentSpanID string,
) *Span {
	return &Span{
		TraceID: traceID,
		SpanID: GenerateID(),
		ParentSpanID: parentSpanID,
		ServiceName: service,
		Operation: operation,
		StartTime: time.Now(),
		Status: "Running",

	}
}

func (s *Span) End() {
	s.EndTime = time.Now()
	s.Status = "SUCCESS"
	QueueSpan(s)
	
}


func StartChildSpan(
	ctx context.Context,
	service string,
	operation string,
) *Span {
	traceID := GetTraceID(ctx)
	parentSpanID := GetSpanID(ctx)
	return StartSpan(
		service,
		operation,
		traceID,
		parentSpanID,
	)
}

