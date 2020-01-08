package middleware

// TracingClient for apm interface
type TracingClient interface {
	CreateEntrySpan(sc *SpanContext) (interface{}, error)
	CreateExitSpan(sc *SpanContext) (interface{}, error)
	EndSpan(sp interface{}, statusCode int) error
}

var tc TracingClient

//CreateEntrySpan create entry span
func CreateEntrySpan(s *SpanContext) (interface{}, error) {
	return tc.CreateEntrySpan(s)
}

//CreateEntrySpan create entry span
func CreateExitSpan(s *SpanContext) (interface{}, error) {
	return tc.CreateExitSpan(s)
}

//EndSpan end span
func EndSpan(span interface{}, status int) error {
	return tc.EndSpan(span, status)
}
