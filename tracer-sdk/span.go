package tracer
import "time"

type Span struct {
	TraceID string
	SpanID string
	ParentSpanID string
	ServiceName string
	Operation string
	StartTime time.Time
	EndTime time.Time
	Status string
}
