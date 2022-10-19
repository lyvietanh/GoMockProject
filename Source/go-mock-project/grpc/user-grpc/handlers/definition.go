package user_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
)

type UserGrpcHandler struct {
	pb.UnimplementedUserGrpcServer
	config util.GrpcConfig
	store  *db.Store
}

func NewUserGrpcHandler(config util.GrpcConfig, store *db.Store) (*UserGrpcHandler, error) {
	return &UserGrpcHandler{
		config: config,
		store:  store,
	}, nil
}

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var (
	HMAC_SECRET_KEY              = []byte("GO MOCK PROJECT - HMAC KEY")
	ACCESS_TOKEN_EXPIRE_DURATION = 15
)
