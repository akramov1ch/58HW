package main

import (
    "log"
    "net"

    pb "58HW/pkg/loan"

    "google.golang.org/grpc"
)

const (
    port = ":50052"
)

type server struct {
    pb.UnimplementedLoanServiceServer
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterLoanServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
