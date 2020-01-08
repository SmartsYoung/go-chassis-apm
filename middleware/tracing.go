package middleware

import (
	"github.com/go-mesh/openlogging"
	"strconv"
)

// TracingClient for apm interface
type TracingClient interface {
	CreateEntrySpan(sc *SpanContext) (interface{}, error)
	CreateExitSpan(sc *SpanContext) (interface{}, error)
	EndSpan(sp interface{}, statusCode int) error
}

var TC TracingClient

//CreateEntrySpan create entry span
func CreateEntrySpan(s *SpanContext, op TracingOptions) (interface{}, error) {
	if op.APMName != "" {
		openlogging.Debug("CreateEntrySpan:" + op.MicServiceName)
		return TC.CreateEntrySpan(s)
	}
	var spans interface{}
	return spans, nil
}

//CreateEntrySpan create entry span
func CreateExitSpan(s *SpanContext, op TracingOptions) (interface{}, error) {
	if op.APMName != "" {
		openlogging.Debug("CreateEntrySpan:" + op.MicServiceName)
		return TC.CreateExitSpan(s)
	}
	var spans interface{}
	return spans, nil
}

//EndSpan end span
func EndSpan(span interface{}, status int, op TracingOptions) error {
	if op.APMName != "" {
		openlogging.Debug("EndSpan: " + op.MicServiceName + "status: " + strconv.Itoa(status))
		return TC.EndSpan(span, status)
	}
	return nil
}

//Init apm client
func Init(op TracingOptions) {
	openlogging.Info("apm Init " + op.APMName + " " + op.ServerURI)
	TC, _ = NewApmClient(op)
}
