package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type ApiConfig struct {
	Server ApiCfgServer `mapstructure:"server"`
	Grpc   ApiCfgGrpc   `mapstructure:"grpc"`
}

type ApiCfgServer struct {
	Port uint `mapstructure:"port"`
}

type ApiCfgGrpc struct {
	FlightGrpc  ApiCfgGrpcItem `mapstructure:"flight_grpc"`
	BookingGrpc ApiCfgGrpcItem `mapstructure:"booking_grpc"`
	UserGrpc    ApiCfgGrpcItem `mapstructure:"user_grpc"`
}

type ApiCfgGrpcItem struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
}

func LoadApiConfig(path string) (config ApiConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("api_config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Printf("config is: %v\n", config)
	return
}
