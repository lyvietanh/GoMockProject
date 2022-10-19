package user_grpc_handlers

import (
	user_grpc_payloads "anhlv11/go-mock-project/grpc/user-grpc/payloads"
	"anhlv11/go-mock-project/pb"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func (h *UserGrpcHandler) VerifyAuth(ctx context.Context, in *pb.GrpcVerifyAuthRequest) (*pb.GrpcVerifyAuthResponse, error) {
	token, err := jwt.ParseWithClaims(in.AccessToken, &user_grpc_payloads.UserAuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return HMAC_SECRET_KEY, nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not parse token|%v", err)
	}

	if userAuthClaims, ok := token.Claims.(*user_grpc_payloads.UserAuthClaims); ok && token.Valid {
		fmt.Printf("%v %v\n", userAuthClaims.Email, userAuthClaims.StandardClaims.ExpiresAt)

		user, err := h.store.Queries.GetUserByEmail(ctx, userAuthClaims.Email)
		if err != nil {
			return nil, fmt.Errorf("can not found any users with email=%v|%v", userAuthClaims.Email, err)
		}

		out := &pb.GrpcVerifyAuthResponse{
			Email:    user.Email,
			RoleId:   user.RoleID.String,
			ExpireIn: userAuthClaims.ExpiresAt - time.Now().Unix(),
		}
		return out, nil
	}
	return nil, errors.New("access token is not valid")
}
