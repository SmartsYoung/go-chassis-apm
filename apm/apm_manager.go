package apm

import (
	"github.com/SkyAPM/go2sky"
	"github.com/go-chassis/go-chassis-apm/tracing"
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

var troption tracing.TracingOptions

//CreateEntrySpan use invocation to make spans for apm
func CreateEntrySpan(i *invocation.Invocation) (go2sky.Span, error) {
	openlogging.Debug("CreateEntrySpan:" + i.MicroServiceName)
	spanCtx := tracing.SpanContext{Ctx: i.Ctx, OperationName: i.MicroServiceName + i.URLPathFormat, ParTraceCtx: i.Headers(), Method: i.Protocol, URL: i.MicroServiceName + i.URLPathFormat}
	span, err := tracing.CreateEntrySpan(&spanCtx)
	if err != nil {
		openlogging.Error("CreateEntrySpan err:" + err.Error())
		return nil, err
	}
	i.Ctx = spanCtx.Ctx
	return span, nil
}

//CreateExitSpan use invocation to make spans for apm
func CreateExitSpan(i *invocation.Invocation) (go2sky.Span, error) {
	openlogging.Debug("CreateExitSpan:" + i.MicroServiceName)
	spanCtx := tracing.SpanContext{Ctx: i.Ctx, OperationName: i.MicroServiceName + i.URLPathFormat, ParTraceCtx: i.Headers(), Method: i.Protocol, URL: i.MicroServiceName + i.URLPathFormat, Peer: i.Endpoint + i.URLPathFormat, TraceCtx: map[string]string{}}
	span, err := tracing.CreateExitSpan(&spanCtx)
	if err != nil {
		openlogging.Error("CreateExitSpan err:" + err.Error())
		return nil, err
	}
	for k, v := range spanCtx.TraceCtx { //ctx need transfer by header
		i.SetHeader(k, v)
	}
	return span, nil
}

//EndSpan use invocation to make spans of apm end
func EndSpan(span go2sky.Span, status int) error {
	openlogging.Debug("EndSpan " + strconv.Itoa(status))
	err := tracing.EndSpan(span, status)
	if err != nil {
		openlogging.Error("EndSpan err:" + err.Error())
		return err
	}
	return nil
}

//Init apm
func Init() error {
	openlogging.Debug("apm Init " + config.GetAPM().Tracing.Tracer)
	if config.GetAPM().Tracing.Settings != nil && config.GetAPM().Tracing.Settings[URI] != "" {
		troption = tracing.TracingOptions{MicServiceName: config.MicroserviceDefinition.ServiceDescription.Name, ServerURI: config.GetAPM().Tracing.Settings["URI"]}
		if serverType, ok := config.GetAPM().Tracing.Settings[ServerType]; ok { //
			var err error
			troption.MicServiceType, err = strconv.Atoi(serverType)
			if err != nil {
				openlogging.Error("get MicServiceType error:" + err.Error())
				return err
			}
		}
		tracing.Init(troption)
		openlogging.Info("apm init: service:" + config.MicroserviceDefinition.ServiceDescription.Name)

	} else {
		openlogging.Warn("apm Init failed. check apm config ")
	}

	return nil
}
