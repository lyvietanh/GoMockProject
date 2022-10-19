package main

import (
	flight_api_handlers "anhlv11/go-mock-project/api/flight-api/handlers"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	apiConfig, err := util.LoadApiConfig("./api/flight-api/")
	if err != nil {
		log.Fatal("Can not load api config with path:", err)
	}

	grpcFlightClientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", apiConfig.Grpc.FlightGrpc.Host, apiConfig.Grpc.FlightGrpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	grpcFlightClient := pb.NewFlightGrpcClient(grpcFlightClientConn)

	grpcUserClientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", apiConfig.Grpc.UserGrpc.Host, apiConfig.Grpc.UserGrpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	grpcUserClient := pb.NewUserGrpcClient(grpcUserClientConn)

	flightApiHandler := flight_api_handlers.NewFlightApiHandler(grpcFlightClient, grpcUserClient)

	engine := gin.Default()

	routerGroup := engine.Group("/api/flight")

	routerGroup.GET("/:originDestination/:departureDate", flightApiHandler.SearchOneWayFlight)
	routerGroup.GET("/:originDestination/:departureDate/:returnDate", flightApiHandler.SearchRoundTripFlight)
	routerGroup.POST("/", flightApiHandler.CreateFlight)
	routerGroup.PUT("/", flightApiHandler.UpdateFlight)

	err = engine.Run(fmt.Sprintf("127.0.0.1:%v", apiConfig.Server.Port))
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
