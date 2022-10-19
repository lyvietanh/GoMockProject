package customer_api_handlers

import (
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"

	"github.com/gin-gonic/gin"

	customer_api_payloads "anhlv11/go-mock-project/api/customer-api/payloads"
)

func (h *CustomerApiHandler) Login(ctx *gin.Context) {
	request := customer_api_payloads.LoginRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcLoginRequest := &pb.GrpcLoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}

	grpcLoginResponse, err := h.UserGrpcClient.Login(ctx.Request.Context(), grpcLoginRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	response := customer_api_payloads.LoginResponse{
		// RequestData:  request,
		Email:        request.Email,
		AccessToken:  grpcLoginResponse.AccessToken,
		RefreshToken: grpcLoginResponse.RefreshToken,
		ExpireIn:     grpcLoginResponse.ExpireIn,
		LastLoginAt:  grpcLoginResponse.LastLoginAt.AsTime(),
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
