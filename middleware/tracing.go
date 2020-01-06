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

var apmClientPlugins = make(map[string]func(TracingOptions) (TracingClient, error))
var apmClients = make(map[string]TracingClient)

//InstallClientPlugins register TracingClient create func
func InstallClientPlugins(name string, f func(TracingOptions) (TracingClient, error)) {
	apmClientPlugins[name] = f
	openlogging.Info("Install apm client: " + name)
}

//CreateEntrySpan create entry span
func CreateEntrySpan(s *SpanContext, op TracingOptions) (interface{}, error) {
	if client, ok := apmClients[op.APMName]; ok {
		openlogging.Debug("CreateEntrySpan:" + op.MicServiceName)
		return client.CreateEntrySpan(s)
	}
	var spans interface{}
	return spans, nil
}

//CreateExitSpan create exit span
func CreateExitSpan(s *SpanContext, op TracingOptions) (interface{}, error) {
	if client, ok := apmClients[op.APMName]; ok {
		openlogging.Debug("CreateExitSpan:" + op.MicServiceName)
		return client.CreateExitSpan(s)
	}
	var span interface{}
	return span, nil
}

//EndSpan end span
func EndSpan(span interface{}, status int, op TracingOptions) error {
	if client, ok := apmClients[op.APMName]; ok {
		openlogging.Debug("EndSpan: " + op.MicServiceName + "status: " + strconv.Itoa(status))
		return client.EndSpan(span, status)
	}
	return nil
}

//Init apm client
func Init(op TracingOptions) {
	openlogging.Info("apm Init " + op.APMName + " " + op.ServerURI)
	f, ok := apmClientPlugins[op.APMName]
	if ok {
		client, err := f(op)
		if err == nil {
			apmClients[op.APMName] = client
		} else {
			openlogging.Error("apmClients init failed. " + err.Error())
		}
	}
}

