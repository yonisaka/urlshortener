package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"

	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Provider is an interface for trace Provider.
type Provider interface {
	// Tracer creates a named tracer that implements Tracer interface.
	// If the name is an empty string then provider uses default name.
	Tracer() Tracer
	// Close closes the tracer provider.
	Close(ctx context.Context) error
}

// ProviderConfig represents the provider configuration and used to create a new
// `Provider` type.
type ProviderConfig struct {
	JaegerEndpoint string
	ServiceName    string
	ServiceVersion string
	Environment    string
	// Set this to `true` if you want to disable tracing completely.
	Disabled bool
}

// Provider represents the tracer provider. Depending on the `config.Disabled`
// parameter, it will either use a "live" provider or a "no operations" version.
// The "no operations" means, tracing will be globally disabled.
type tracerProvider struct {
	provider trace.TracerProvider
}

// NewProvider returns a new `Provider` type. It uses Jaeger exporter and globally sets
// the tracer provider as well as the global tracer for spans.
func NewProvider(config ProviderConfig) (Provider, error) {
	if config.Disabled {
		return &tracerProvider{provider: trace.NewNoopTracerProvider()}, nil
	}

	exp, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.JaegerEndpoint)),
	)
	if err != nil {
		return &tracerProvider{}, err
	}

	prv := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(sdkresource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String(config.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(config.Environment),
		)),
	)

	otel.SetTracerProvider(prv)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return &tracerProvider{provider: prv}, nil
}

// Tracer implements trace.Provider.
func (t *tracerProvider) Tracer() Tracer {
	return &otelTracer{}
}

// Close shuts down the tracer provider only if it was not "no operations"
// version.
func (t tracerProvider) Close(ctx context.Context) error {
	if prv, ok := t.provider.(*sdktrace.TracerProvider); ok {
		return prv.Shutdown(ctx)
	}

	return nil
}
