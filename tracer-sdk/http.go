package tracer

import "net/http"

func InjectTraceHeaders(
	r *http.Request,
	span *Span,
) {
	r.Header.Set(
		HeaderTraceID,
		span.TraceID,
	)
	r.Header.Set(
		HeaderParentID,
		span.SpanID,
	)
}


func ExtractTraceHeaders(
	r *http.Request,
) (string, string) {
	traceID := r.Header.Get(
		HeaderTraceID,
	)
	parentSpanID := r.Header.Get(
		HeaderParentID,
	)
	return traceID, parentSpanID
}


func StartSpanFromRequest(
	service string,
	operation string,
	req *http.Request,
) *Span {
	traceID, parentSpanID := ExtractTraceHeaders(req)
	return StartSpan(
		service,
		operation,
		traceID,
		parentSpanID,
	)
}