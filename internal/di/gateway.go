package di

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	rpc "github.com/yonisaka/urlshortener/api/go/grpc"
	"github.com/yonisaka/urlshortener/pkg/di"
	"github.com/yonisaka/urlshortener/pkg/server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type httpServer struct {
	mux *runtime.ServeMux
	srv *http.Server
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")

		if r.Method == http.MethodOptions {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// GetURLShortenerGatewayServer returns gRPC Gateway server instance for URLShortener service.
func GetURLShortenerGatewayServer() server.Server {
	mux := runtime.NewServeMux()

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port+1),
		Handler: cors(mux),
	}

	return &httpServer{
		mux: mux,
		srv: srv,
	}
}

func (h *httpServer) Run() error {
	err := rpc.RegisterURLShortenerServiceHandlerFromEndpoint(
		context.Background(),
		h.mux,
		"0.0.0.0:"+os.Getenv("SERVER_PORT"),
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	errorChan := make(chan error)

	go func() {
		if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errorChan <- err
		}
	}()

	di.RegisterCloser("Gateway server", di.NewCloser(h.GracefulStop))

	return <-errorChan
}

func (h *httpServer) GracefulStop() {
	_ = h.srv.Shutdown(context.Background())
}
