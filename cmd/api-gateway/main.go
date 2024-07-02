package main

import (
    "context"
    "log"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "58HW/pkg/book/service"
    "58HW/pkg/loan"
)

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}

    err := book.RegisterBookServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
    if err != nil {
        log.Fatalf("failed to register book service: %v", err)
    }

    err = loan.RegisterLoanServiceHandlerFromEndpoint(ctx, mux, "localhost:50052", opts)
    if err != nil {
        log.Fatalf("failed to register loan service: %v", err)
    }

    log.Println("Serving gRPC-Gateway on http://localhost:8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
