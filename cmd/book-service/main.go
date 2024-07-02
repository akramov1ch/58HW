package main

import (
    "log"
    "net"

    pb "58HW/pkg/book"

    "google.golang.org/grpc"
)

const (
    port = ":50051"
)

type server struct {
    pb.UnimplementedBookServiceServer
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterBookServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
