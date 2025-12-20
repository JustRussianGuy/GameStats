package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	server "github.com/JustRussianGuy/GameStats/internal/api/gamestats_api"
	"github.com/JustRussianGuy/GameStats/internal/consumer/eventsconsumer"

	"github.com/JustRussianGuy/GameStats/internal/pb/gamestats_api"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	r := chi.NewRouter()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := gamestats_api.RegisterGameStatsServiceHandlerFromEndpoint(
		ctx,
		mux,
		":50051",
		opts,
	); err != nil {
		return err
	}

	r.Mount("/", mux)

	slog.Info("gRPC-Gateway server listening on :8080")
	return http.ListenAndServe(":8080", r)
}
