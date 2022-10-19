package user_grpc_handlers

import (
	db "anhlv11/go-mock-project/db/sqlc"
	"anhlv11/go-mock-project/pb"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGrpcHandler) UpdateUserProfile(ctx context.Context, in *pb.GrpcUpdateUserProfileRequest) (*pb.GrpcUpdateUserProfileResponse, error) {
	email := strings.ToLower(strings.TrimSpace(in.Email))
	updateUserProfileParams := db.UpdateUserProfileParams{
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
		ModifyDate:      sql.NullTime{Time: time.Now(), Valid: true},
	}

	user, err := h.store.Queries.UpdateUserProfile(ctx, updateUserProfileParams)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not update user with email=%v|%v", in.Email, err))
	}

	out := &pb.GrpcUpdateUserProfileResponse{
		UpdateDate: timestamppb.New(user.ModifyDate.Time),
	}
	return out, nil
}
