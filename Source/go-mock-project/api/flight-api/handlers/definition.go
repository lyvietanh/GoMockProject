package flight_api_handlers

import (
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type FlightApiHandler struct {
	FlightGrpcClient pb.FlightGrpcClient
	UserGrpcClient   pb.UserGrpcClient
}

func NewFlightApiHandler(flightGrpcClient pb.FlightGrpcClient, userGrpcClient pb.UserGrpcClient) FlightApiHandler {
	return FlightApiHandler{
		FlightGrpcClient: flightGrpcClient,
		UserGrpcClient:   userGrpcClient,
	}
}

func (h *FlightApiHandler) VerifyRolePermission(ctx *gin.Context, roleIds *[]string) (*pb.GrpcVerifyAuthResponse, error) {
	securedHeader := util.SecuredHeader{}

	if err := ctx.ShouldBindHeader(&securedHeader); err != nil {
		return nil, err
	}

	grpcVerifyAuthRequest := &pb.GrpcVerifyAuthRequest{
		AccessToken: securedHeader.AccessToken,
	}
	log.Printf("AccessToken=%v\n", securedHeader.AccessToken)
	grpcVerifyAuthResponse, err := h.UserGrpcClient.VerifyAuth(ctx.Request.Context(), grpcVerifyAuthRequest)
	if err != nil {
		return nil, err
	}

	if grpcVerifyAuthResponse.ExpireIn <= 0 {
		return nil, errors.New("token is expired")
	}

	if roleIds != nil && len(*roleIds) > 0 {
		hasPermission := false
		for _, roleId := range *roleIds {
			if strings.EqualFold(grpcVerifyAuthResponse.RoleId, roleId) {
				hasPermission = true
				break
			}
		}
		if !hasPermission {
			return nil, errors.New("no permission")
		}
	}

	return grpcVerifyAuthResponse, nil
}

var (
	BAD_REQUEST_ERROR    = util.NewApiError(400, "Bad Request")
	BUSINESS_LOGIC_ERROR = util.NewApiError(500, "Business Logic Error")
)
