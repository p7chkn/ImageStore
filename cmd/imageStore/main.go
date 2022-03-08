package main

import (
	"context"
	"fmt"
	configurations "goImageStore/cmd/imageStore/configuration"
	"goImageStore/cmd/imageStore/services"
	"goImageStore/iternal/filestorage"
	"goImageStore/iternal/handlers"
	"goImageStore/iternal/logger"
	"goImageStore/pb"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	//"database/sql"
	//"goImageStore/cmd/imageStore/configuration"
	//"goImageStore/cmd/imageStore/services"

	_ "github.com/lib/pq"
)

var (
	grpcServer *grpc.Server
	httpServer *http.Server
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	log := logger.InitLogger()
	cfg := configurations.New()
	repo := filestorage.New(cfg.PathToFile)

	g, ctx := errgroup.WithContext(ctx)

	httpHandler := services.SetUpRouter(repo, log, cfg.ServerAddress, cfg.PathToFile)
	grpcHandler := handlers.GrpcHandlerNew(repo, log, cfg.ServerAddress)

	g.Go(func() error {
		httpServer = &http.Server{
			Addr:    cfg.ServerAddress,
			Handler: httpHandler,
		}
		log.Infof("httpServer starting at: %v", cfg.ServerAddress)
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcPort))
		if err != nil {
			log.Errorf("gRPC server failed to listen: %v", err.Error())
			os.Exit(2)
		}
		grpcServer = grpc.NewServer()
		pb.RegisterFileServer(grpcServer, grpcHandler)
		log.Infof("server listening at %v", lis.Addr())
		return grpcServer.Serve(lis)
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Warn("Receive shutdown signal")

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer shutdownCancel()

	if httpServer != nil {
		_ = httpServer.Shutdown(shutdownCtx)
	}
	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	err := g.Wait()
	if err != nil {
		log.Errorf("server returning an error: %v", err)
		os.Exit(2)
	}
}
