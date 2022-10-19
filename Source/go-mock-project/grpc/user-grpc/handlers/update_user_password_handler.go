package user_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGrpcHandler) UpdateUserPassword(ctx context.Context, in *pb.GrpcUpdateUserPasswordRequest) (*pb.GrpcUpdateUserPasswordResponse, error) {
	email := strings.ToLower(strings.TrimSpace(in.Email))

	user, err := h.store.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not found any users with email=%v|%v", in.Email, err))
	}

	requestCurrentSecurePassword := util.GenerateSha256SecurePassword(email, in.CurrentPassword, user.PasswordSalt.String)
	userCurrentSecurePassword := user.SecurePassword.String
	log.Printf("requestCurrentSecurePassword=%v\n", requestCurrentSecurePassword)
	log.Printf("userCurrentSecurePassword=%v\n", userCurrentSecurePassword)
	if !strings.EqualFold(requestCurrentSecurePassword, userCurrentSecurePassword) {
		return nil, status.Error(codes.Internal, "Wrong current password")
	}

	requestNewSecurePassword := util.GenerateSha256SecurePassword(email, in.NewPassword, user.PasswordSalt.String)
	if strings.EqualFold(requestNewSecurePassword, userCurrentSecurePassword) {
		return nil, status.Error(codes.Internal, "New password must be difference from old password")
	}

	newPasswordSalt := util.GenerateRandomString(16)
	updateUserPasswordParams := db.UpdateUserPasswordParams{
		Email:          email,
		ModifyDate:     sql.NullTime{Time: time.Now(), Valid: true},
		PasswordSalt:   sql.NullString{String: newPasswordSalt, Valid: true},
		SecurePassword: sql.NullString{String: util.GenerateSha256SecurePassword(email, in.NewPassword, newPasswordSalt), Valid: true},
	}

	user, err = h.store.Queries.UpdateUserPassword(ctx, updateUserPasswordParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not change password for user with email=%v|%v", in.Email, err))
	}

	out := &pb.GrpcUpdateUserPasswordResponse{
		UpdateDate: timestamppb.New(user.ModifyDate.Time),
	}
	return out, nil
}
