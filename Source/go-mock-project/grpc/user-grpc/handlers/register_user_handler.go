package user_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGrpcHandler) RegisterUser(ctx context.Context, in *pb.GrpcRegisterUserRequest) (*pb.GrpcRegisterUserResponse, error) {
	email := strings.ToLower(strings.TrimSpace(in.Email))

	_, err := h.store.Queries.GetUserByEmail(ctx, email)
	if err == nil {
		return nil, status.Error(codes.Internal, "Email is already in used by other")
	}

	passwordSalt := util.GenerateRandomString(16)
	createUserParams := db.CreateUserParams{
		Email:           email,
		RoleID:          sql.NullString{String: in.RoleId, Valid: true},
		Title:           sql.NullString{String: in.Title, Valid: true},
		LastName:        sql.NullString{String: in.LastName, Valid: true},
		FirstName:       sql.NullString{String: in.FirstName, Valid: true},
		DateOfBirth:     sql.NullTime{Time: in.DateOfBirth.AsTime(), Valid: true},
		Phone:           sql.NullString{String: in.Phone, Valid: true},
		TravelDocType:   sql.NullString{String: in.TravelDocType, Valid: true},
		TravelDocNumber: sql.NullString{String: in.TravelDocNumber, Valid: true},
		FfpNumber:       sql.NullString{String: in.FfpNumber, Valid: true},
		Enabled:         true,
		CreateDate:      sql.NullTime{Time: time.Now(), Valid: true},
		PasswordSalt:    sql.NullString{String: passwordSalt, Valid: true},
		SecurePassword:  sql.NullString{String: util.GenerateSha256SecurePassword(email, in.Password, passwordSalt), Valid: true},
	}

	user, err := h.store.Queries.CreateUser(ctx, createUserParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not create user with email=%v|%v", in.Email, err))
	}

	out := &pb.GrpcRegisterUserResponse{
		RegisterDate: timestamppb.New(user.CreateDate.Time),
	}
	return out, nil
}
