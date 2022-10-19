package main

import (
	booking_api_handlers "anhlv11/go-mock-project/api/booking-api/handlers"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	apiConfig, err := util.LoadApiConfig("./api/booking-api/")
	if err != nil {
		log.Fatal("Can not load api config with path:", err)
	}

	grpcClient, err := grpc.Dial(fmt.Sprintf("%v:%v", apiConfig.Grpc.BookingGrpc.Host, apiConfig.Grpc.BookingGrpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcBookingClient := pb.NewBookingGrpcClient(grpcClient)

	bookingApiHandler := booking_api_handlers.NewBookingApiHandler(grpcBookingClient)

	engine := gin.Default()

	routerGroup := engine.Group("/api/book")

	routerGroup.POST("/", bookingApiHandler.BookFlight)
	routerGroup.PUT("/", bookingApiHandler.CancelFlight)
	routerGroup.GET("/:reservationCode", bookingApiHandler.ViewBooking)

	err = engine.Run(fmt.Sprintf("127.0.0.1:%v", apiConfig.Server.Port))
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
