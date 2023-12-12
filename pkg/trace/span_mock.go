package trace

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

var _ Tracer = (*MockTracer)(nil)

// MockTracer represents Tracer mock.
type MockTracer struct {
	StartSpanFunc func(ctx context.Context, name string, cus SpanCustomiser) (context.Context, trace.Span)
}

// StartSpan calls StartSpanFunc.
func (m *MockTracer) StartSpan(ctx context.Context, name string, cus SpanCustomiser) (context.Context, trace.Span) {
	return m.StartSpanFunc(ctx, name, cus)
}
