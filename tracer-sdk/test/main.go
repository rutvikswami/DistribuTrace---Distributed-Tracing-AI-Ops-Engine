package main

import (
	"tracer-sdk"
)

func main() {
	for i := 0; i < 10; i++ {

		span := tracer.StartSpan(
			"order-service",
			"test",
			tracer.NewTraceID(),
			"",
		)

		span.End()
	}

}