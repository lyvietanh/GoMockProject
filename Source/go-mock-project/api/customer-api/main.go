package main

import (
	customer_api_handlers "anhlv11/go-mock-project/api/customer-api/handlers"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	apiConfig, err := util.LoadApiConfig("./api/customer-api/")
	if err != nil {
		log.Fatal("Can not load api config with path:", err)
	}

	userGrpcClientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", apiConfig.Grpc.UserGrpc.Host, apiConfig.Grpc.UserGrpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userGrpcClient := pb.NewUserGrpcClient(userGrpcClientConn)

	customerApiHandler := customer_api_handlers.NewCustomerApiHandler(userGrpcClient)

	engine := gin.Default()

	routerAuthGroup := engine.Group("/api/auth")
	routerAuthGroup.POST("/login", customerApiHandler.Login)
	routerAuthGroup.PUT("/change-password", customerApiHandler.ChangePassword)

	routerCustomerGroup := engine.Group("/api/customer")
	routerCustomerGroup.POST("/", customerApiHandler.RegisterCustomer)
	routerCustomerGroup.PUT("/", customerApiHandler.UpdateCustomerProfile)
	routerCustomerGroup.GET("/booking-history", customerApiHandler.ViewBookingHistory)

	err = engine.Run(fmt.Sprintf("127.0.0.1:%v", apiConfig.Server.Port))
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
