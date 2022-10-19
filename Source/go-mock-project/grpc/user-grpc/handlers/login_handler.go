package user_grpc_handlers

import (
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	user_grpc_payloads "anhlv11/go-mock-project/grpc/user-grpc/payloads"

	"github.com/golang-jwt/jwt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGrpcHandler) Login(ctx context.Context, in *pb.GrpcLoginRequest) (*pb.GrpcLoginResponse, error) {
	email := strings.ToLower(strings.TrimSpace(in.Email))
	user, err := h.store.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("could not found any user with email=%v|%v", email, err)
	}

	requestSecurePassword := util.GenerateSha256SecurePassword(email, in.Password, user.PasswordSalt.String)
	if !user.SecurePassword.Valid || !strings.EqualFold(user.SecurePassword.String, requestSecurePassword) {
		return nil, errors.New("your email or password are not valid")
	}

	issuedAt := time.Now()
	expireIn := ACCESS_TOKEN_EXPIRE_DURATION
	expiresAt := issuedAt.Add(time.Duration(expireIn) * time.Minute)
	userAuthClaims := user_grpc_payloads.UserAuthClaims{
		Email:    user.Email,
		FullName: fmt.Sprintf("%v %v", user.LastName, user.FirstName),
		StandardClaims: jwt.StandardClaims{
			Subject:   "ACCESS_TOKEN",
			Issuer:    "GO_MOCK_PROJECT",
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userAuthClaims)

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString(HMAC_SECRET_KEY)
	lastLoginDate := time.Now()
	if user.LastLoginDate.Valid {
		lastLoginDate = user.LastLoginDate.Time
	}
	out := &pb.GrpcLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: "(TBD)",
		ExpireIn:     expiresAt.Unix() - issuedAt.Unix(),
		LastLoginAt:  timestamppb.New(lastLoginDate)}
	fmt.Println(accessToken, err)
	return out, nil
}
