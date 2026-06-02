package tracer

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ExportSpan(
	span *Span,
) error {
	body ,err := json.Marshal(span)
	if err != nil {
		return err
	}
	_,err = http.Post(
		CollectorUrl,
		"application/json",
		bytes.NewBuffer(body),
	)
	return err
}

func ExportBatch(
	spans []*Span,
) error {
	body,err := json.Marshal(
		spans,
	)
	if err != nil {
		return err
	}
	_, err = http.Post(
		CollectorUrl,
		"application/json",
		bytes.NewBuffer(body),
	)
	return err
}