package customer_api_handlers

import (
	customer_api_payloads "anhlv11/go-mock-project/api/customer-api/payloads"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CustomerApiHandler) ChangePassword(ctx *gin.Context) {
	// need valid token to use this feature
	grpcVerifyAuthResponse, err := h.VerifyRolePermission(ctx, nil)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	// binding request
	request := &customer_api_payloads.ChangePasswordRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcUpdateUserPasswordRequest := &pb.GrpcUpdateUserPasswordRequest{
		Email:           grpcVerifyAuthResponse.Email,
		CurrentPassword: request.CurrentPassword,
		NewPassword:     request.NewPassword,
	}

	grpcUpdateUserPasswordResponse, err := h.UserGrpcClient.UpdateUserPassword(ctx.Request.Context(), grpcUpdateUserPasswordRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	response := customer_api_payloads.ChangePasswordResponse{
		UpdateDate: util.JsonDateTime(grpcUpdateUserPasswordResponse.UpdateDate.AsTime()),
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
