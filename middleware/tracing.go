package middleware

//monitoring.yaml
const (
	URI        = "URI"
	ServerType = "serverType"
)

var troption TracingOptions

// TracingClient for apm interface
type TracingClient interface {
	CreateEntrySpan(sc SpanContext) (interface{}, error)
	CreateExitSpan(sc SpanContext) (interface{}, error)
	EndSpan(sp interface{}, statusCode int) error
}

var apmClients = make(map[string]TracingClient)

/*func Init() error {
	openlogging.Debug("apm Init " + config.GetAPM().Tracing.Tracer)
	if config.GetAPM().Tracing.Tracer != "" && config.GetAPM().Tracing.Settings != nil && config.GetAPM().Tracing.Settings[URI] != "" {
		troption = TracingOptions{APMName: config.GetAPM().Tracing.Tracer, MicServiceName: config.MicroserviceDefinition.ServiceDescription.Name, ServerURI: config.GetAPM().Tracing.Settings["URI"]}
		if serverType, ok := config.GetAPM().Tracing.Settings[ServerType]; ok { //
			var err error
			troption.MicServiceType, err = strconv.Atoi(serverType)
			if err != nil {
				openlogging.Error("get MicServiceType error:" + err.Error())
				return nil
			}
		}
		//apm.Init(troption)
		openlogging.Info("apm Init:" + config.GetAPM().Tracing.Tracer + " service:" + config.MicroserviceDefinition.ServiceDescription.Name)
	} else {
		openlogging.Warn("apm Init failed. check apm config " + config.GetAPM().Tracing.Tracer)
	}
	return nil
}

func init(){
	if err := Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
}*/
