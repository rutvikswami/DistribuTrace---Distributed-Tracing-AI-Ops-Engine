package tracer

import (
	"sync"
	"time"
)

var spanQueue []*Span
var queueMutex sync.Mutex

func QueueSpan(
	span *Span,
) {
	queueMutex.Lock()
	spanQueue = append(spanQueue, span,)
	queueMutex.Unlock()
}

func FlushBatch() {
	queueMutex.Lock()
	if len(spanQueue) == 0 {
		queueMutex.Unlock()
		return 
	}
	batch := spanQueue
	spanQueue = nil
	queueMutex.Unlock()
	ExportBatch(batch)
}

func StartBatchWorker() {
	ticker := time.NewTicker(
		2 * time.Second,
	)
	go func ()  {
		for range ticker.C {
			FlushBatch()
		}
	}()
}