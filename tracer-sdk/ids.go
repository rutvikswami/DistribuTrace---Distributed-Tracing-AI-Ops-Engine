package tracer

import "github.com/google/uuid"

func GenerateID() string {
	return uuid.New().String()
}
func NewTraceID() string {
	return GenerateID()
}
