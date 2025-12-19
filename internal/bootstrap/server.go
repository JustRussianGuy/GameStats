package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	server "github.com/JustRussianGuy/GameStats/internal/api/gamestats_api"
	"github.com/JustRussianGuy/GameStats/internal/consumer/eventsconsumer"
	
	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppRun(api server.GameStatsAPI, gameEventsConsumer *eventsconsumer.GameEventsConsumer) {
	// Запускаем консьюмера Kafka в отдельной горутине
	go gameEventsConsumer.Consume(context.Background())

	// Запускаем gRPC сервер
	go func() {
		if err := runGRPCServer(api); err != nil {
			panic(fmt.Errorf("failed to run gRPC server: %v", err))
		}
	}()

	// Запускаем HTTP Gateway
	if err := runGatewayServer(); err != nil {
		panic(fmt.Errorf("failed to run gateway server: %v", err))
	}
}

func runGRPCServer(api server.GameStatsAPI) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	gamestats_api.RegisterGameStatsServiceServer(s, &api)

	slog.Info("gRPC server listening on :50051")
	return s.Serve(lis)
}

func runGatewayServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	swaggerPath := os.Getenv("SWAGGER_PATH")
	if _, err := os.Stat(swaggerPath); os.IsNotExist(err) {
		panic(fmt.Errorf("swagger file not found: %s", swaggerPath))
	}

	r := chi.NewRouter()
	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerPath)
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := gamestats_api.RegisterGameStatsServiceHandlerFromEndpoint(ctx, mux, ":50051", opts)
	if err != nil {
		panic(err)
	}

	r.Mount("/", mux)

	slog.Info("gRPC-Gateway server listening on :8080")
	return http.ListenAndServe(":8080", r)
}
