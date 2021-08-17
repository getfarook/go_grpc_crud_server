package main

import (
	"fmt"
	"log"
	"net"

	"github.com/getfarook/go_grpc_crud_server/prodpb"
	"github.com/getfarook/go_grpc_crud_server/salespro/serves"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting the SalesPro Server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to open listner : %v", err)
	}

	s := grpc.NewServer()
	prodpb.RegisterProductServiceServer(s, &serves.Prodserver{})

	prodpb.RegisterPartnerServiceServer(s, &serves.Pserver{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

	fmt.Println("SalesPro Server Started Successfully...Ready to serve...")
}
