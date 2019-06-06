package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitstored/crypto-service/pb"
	"github.com/bitstored/crypto-service/pkg/server"
	"github.com/bitstored/crypto-service/pkg/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	ServiceName = "crypto"
)

var (
	grpcAddr = flag.String("grpc", "localhost:4004", "gRPC API address")
	// cert     = flag.String("cert", "scripts/localhost.pem", "certificate pathname")
	// certKey  = flag.String("certkey", "scripts/localhost.key", "private key pathname")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	fmt.Println(os.Args)

	service := service.NewCryptoService()
	gRPCListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %s", *grpcAddr, err)
	}

	devServer := server.NewCryptoServer(service)

	// Register standard server metrics and customized metrics to registry.
	grpcMetrics := grpc_prometheus.NewServerMetrics()

	gRPCServer := grpc.NewServer()

	pb.RegisterCryptoServer(gRPCServer, devServer)
	reflection.Register(gRPCServer)
	grpc_prometheus.Register(gRPCServer)
	grpcMetrics.InitializeMetrics(gRPCServer)

	reg := prometheus.NewRegistry()
	reg.MustRegister(grpcMetrics)

	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	go func() {
		if err := gRPCServer.Serve(gRPCListener); err != nil {
			log.Fatalf("Failed to serve gRPC: %s", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, *grpcAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	fmt.Printf("Crypto server listening on  %s for gRPC\n", *grpcAddr)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()
	// Wait for signal
	<-done
}
