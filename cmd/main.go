package main

import (
	"crud/internal/database"
	grpc_svc "crud/proto/grpc_svc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	database.StartDB()
	/*
		--Without Http--
		server := server.NewServer()
		server.Run()
	*/
	grpcServer := grpc.NewServer()
	grpc_svc.RegisterService(*grpcServer)
	port := ":5000"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	err_grpc := grpcServer.Serve(listener)
	if err_grpc != nil {
		log.Fatal(err_grpc)
	}
}
