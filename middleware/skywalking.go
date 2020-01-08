package middleware

import (
	"github.com/go-mesh/openlogging"
	"github.com/tetratelabs/go2sky"
	"github.com/tetratelabs/go2sky/reporter"
	skycom "github.com/tetratelabs/go2sky/reporter/grpc/common"
	"strconv"
)

//for skywalkinng use
const (
	HTTPPrefix             = "http://"
	CrossProcessProtocolV2 = "Sw6"
	SkyName                = "skywalking"
	DefaultTraceContext    = ""
)

//SkyWalkingClient for connecting and reporting to skywalking server
type SkyWalkingClient struct {
	reporter    go2sky.Reporter
	tracer      *go2sky.Tracer
	ServiceType int32
}

//CreateEntrySpan create entry span
func (s *SkyWalkingClient) CreateEntrySpan(sc *SpanContext) (interface{}, error) {
	openlogging.Debug("CreateEntrySpan begin. span" + sc.OperationName)
	span, ctx, err := s.tracer.CreateEntrySpan(sc.Ctx, sc.OperationName, func() (string, error) {
		if sc.ParTraceCtx != nil {
			return sc.ParTraceCtx[CrossProcessProtocolV2], nil
		}
		return DefaultTraceContext, nil
	})
	if err != nil {
		openlogging.Error("CreateEntrySpan error:" + err.Error())
		return &span, err
	}
	span.Tag(go2sky.TagHTTPMethod, sc.Method)
	span.Tag(go2sky.TagURL, sc.URL)
	span.SetSpanLayer(skycom.SpanLayer_Http)
	span.SetComponent(s.ServiceType)
	sc.Ctx = ctx
	return &span, err
}

//CreateExitSpan create end span
func (s *SkyWalkingClient) CreateExitSpan(sc *SpanContext) (interface{}, error) {
	openlogging.Debug("CreateExitSpan begin. span:" + sc.OperationName)
	span, err := s.tracer.CreateExitSpan(sc.Ctx, sc.OperationName, sc.Peer, func(header string) error {
		sc.TraceCtx[CrossProcessProtocolV2] = header
		return nil
	})
	if err != nil {
		openlogging.Error("CreateExitSpan error:" + err.Error())
		return nil, err
	}
	span.Tag(go2sky.TagHTTPMethod, sc.Method)
	span.Tag(go2sky.TagURL, sc.URL)
	span.SetSpanLayer(skycom.SpanLayer_Http)
	span.SetComponent(s.ServiceType)
	return &span, err
}

//EndSpan make span end and report to skywalking
func (s *SkyWalkingClient) EndSpan(sp interface{}, statusCode int) error {
	openlogging.Debug("EndSpan status:" + strconv.Itoa(statusCode))
	span, ok := (sp).(*go2sky.Span)
	if !ok || span == nil {
		openlogging.Error("EndSpan failed.")
		return nil
	}
	(*span).Tag(go2sky.TagStatusCode, strconv.Itoa(statusCode))
	(*span).End()
	return nil
}

//NewApmClient init report and tracer for connecting and sending messages to skywalking server
func NewApmClient(op TracingOptions) (TracingClient, error) {
	var (
		err    error
		client SkyWalkingClient
	)
	client.reporter, err = reporter.NewGRPCReporter(op.ServerURI)
	if err != nil {
		openlogging.Error("NewGRPCReporter error:" + err.Error())
		return &client, err
	}
	client.tracer, err = go2sky.NewTracer(op.MicServiceName, go2sky.WithReporter(client.reporter))
	//not wait for register here
	//t.WaitUntilRegister()
	if err != nil {
		openlogging.Error("NewTracer error:" + err.Error())
		return &client, err
	}
	client.ServiceType = int32(op.MicServiceType)
	openlogging.Debug("NewApmClient succ. name:" + op.APMName + "uri:" + op.ServerURI)
	return &client, err
}