package middleware

// TracingClient for apm interface
type TracingClient interface {
	CreateEntrySpan(sc SpanContext) (interface{}, error)
	CreateExitSpan(sc SpanContext) (interface{}, error)
	EndSpan(sp interface{}, statusCode int) error
}
