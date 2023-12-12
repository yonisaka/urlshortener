package trace

import (
	"context"
)

var _ Provider = (*MockProvider)(nil)

// MockProvider represents Provider mock.
type MockProvider struct {
	TracerFunc func() Tracer
	CloseFunc  func(ctx context.Context) error
}

// Tracer calls TracerFunc.
func (m *MockProvider) Tracer() Tracer {
	return m.TracerFunc()
}

// Close calls CloseFunc.
func (m *MockProvider) Close(ctx context.Context) error {
	return m.CloseFunc(ctx)
}
