package di

import (
	"context"
	"io"
	"log"
	"os"
	"sync"

	"github.com/yonisaka/urlshortener/pkg/di"
	"github.com/yonisaka/urlshortener/pkg/trace"

	otelTrace "go.opentelemetry.io/otel/trace"
)

var (
	traceProviderOnce sync.Once
	traceProvider     trace.Provider
)

type wrapProvider struct {
	provider trace.Provider
}

func (w *wrapProvider) Close() error {
	return w.provider.Close(context.Background())
}

func GetTracer() trace.Provider {
	traceProviderOnce.Do(func() {
		var err error

		if os.Getenv("OTEL_AGENT") != "" {
			traceProvider, err = trace.NewProvider(trace.ProviderConfig{
				JaegerEndpoint: os.Getenv("OTEL_AGENT"),
				ServiceName:    "URL Shortener Server",
				ServiceVersion: "1.0.0",
				Environment:    os.Getenv("APP_ENV"),
				Disabled:       false,
			})
			if err != nil {
				log.Fatal("trace new provider", err)
			}
		} else {
			traceProvider = &trace.MockProvider{
				CloseFunc: func(ctx context.Context) error {
					return nil
				},
				TracerFunc: func() trace.Tracer {
					return &trace.MockTracer{
						StartSpanFunc: func(ctx context.Context, name string, cus trace.SpanCustomiser) (context.Context, otelTrace.Span) {
							return ctx, otelTrace.SpanFromContext(ctx)
						},
					}
				},
			}
		}

		var c io.Closer = &wrapProvider{
			provider: traceProvider,
		}

		di.RegisterCloser("Trace Provider", c)
	})

	return traceProvider
}
