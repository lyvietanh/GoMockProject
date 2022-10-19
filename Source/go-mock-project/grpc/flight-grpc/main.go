package main

import (
	db "anhlv11/go-mock-project/db/sqlc"
	flightgrpc_handlers "anhlv11/go-mock-project/grpc/flight-grpc/handlers"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadGrpcConfig("./grpc/flight-grpc/")
	if err != nil {
		log.Fatal("Can not load config with path:", err)
	}
	conn, err := sql.Open(config.Database.DriverName, util.GetDatabaseSourceName(config.Database))
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}

	store := db.NewStore(conn)

	grpcServer := grpc.NewServer()
	grpcHandler, err := flightgrpc_handlers.NewFlightGrpcHandler(config, store)
	if err != nil {
		panic(err)
	}

	reflection.Register(grpcServer)
	pb.RegisterFlightGrpcServer(grpcServer, grpcHandler)

	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", config.Server.Port))
	if err != nil {
		panic(err)
	}
	log.Printf("Listening at port: %v\n", config.Server.Port)

	grpcServer.Serve(listen)
}
