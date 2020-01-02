package apm

import (
	//"github.com/go-chassis/go-chassis-apm/tracing"

	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/core/invocation"
	"github.com/go-mesh/openlogging"
	"strconv"
)

//monitoring.yaml
const (
	URI        = "URI"
	ServerType = "serverType"
)

var troption TracingOptions

//CreateEntrySpan use invocation to make spans for apm
func CreateEntrySpan(i *invocation.Invocation) (interface{}, error) {
	openlogging.Debug("CreateEntrySpan:" + i.MicroServiceName)
	spanCtx := SpanContext{Ctx: i.Ctx, OperationName: i.MicroServiceName + i.URLPathFormat, ParTraceCtx: i.Headers(), Method: i.Protocol, URL: i.MicroServiceName + i.URLPathFormat}
	/*span, err := CreateEntrySpan(&spanCtx, troption)
	if err != nil {
		openlogging.Error("CreateEntrySpan err:" + err.Error())
		return nil, err
	}*/
	i.Ctx = spanCtx.Ctx
	return spanCtx, nil
}

//CreateExitSpan use invocation to make spans for apm
func CreateExitSpan(i *invocation.Invocation) (interface{}, error) {
	openlogging.Debug("CreateExitSpan:" + i.MicroServiceName)
	spanCtx := SpanContext{Ctx: i.Ctx, OperationName: i.MicroServiceName + i.URLPathFormat, ParTraceCtx: i.Headers(), Method: i.Protocol, URL: i.MicroServiceName + i.URLPathFormat, Peer: i.Endpoint + i.URLPathFormat, TraceCtx: map[string]string{}}
	/*span, err := CreateExitSpan(&spanCtx, troption)
	if err != nil {
		openlogging.Error("CreateExitSpan err:" + err.Error())
		return nil, err
	}*/
	for k, v := range spanCtx.TraceCtx { //ctx need transfer by header
		i.SetHeader(k, v)
	}
	return spanCtx, nil
}

//EndSpan use invocation to make spans of apm end
func EndSpan(span interface{}, status int) error {
	openlogging.Debug("EndSpan " + strconv.Itoa(status))
	// EndSpan(span, status, troption)
	return nil
}

//Init apm
func Init() error {
	openlogging.Debug("apm Init " + config.GetAPM().Tracing.Tracer)
	if config.GetAPM().Tracing.Tracer != "" && config.GetAPM().Tracing.Settings != nil && config.GetAPM().Tracing.Settings[URI] != "" {
		troption = TracingOptions{APMName: config.GetAPM().Tracing.Tracer, MicServiceName: config.MicroserviceDefinition.ServiceDescription.Name, ServerURI: config.GetAPM().Tracing.Settings["URI"]}
		if serverType, ok := config.GetAPM().Tracing.Settings[ServerType]; ok { //
		    var err error
			troption.MicServiceType, err = strconv.Atoi(serverType)
			if err != nil {
				openlogging.Error("get MicServiceType error:" + err.Error())
				return err
			}
		}
		apm.Init(troption)
		openlogging.Info("apm Init:" + config.GetAPM().Tracing.Tracer + " service:" + config.MicroserviceDefinition.ServiceDescription.Name)
	} else {
		openlogging.Warn("apm Init failed. check apm config " + config.GetAPM().Tracing.Tracer)
	}
	return nil
}
