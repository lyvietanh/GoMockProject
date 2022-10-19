package customer_api_handlers

import (
	customer_api_payloads "anhlv11/go-mock-project/api/customer-api/payloads"
	"anhlv11/go-mock-project/constant"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *CustomerApiHandler) UpdateCustomerProfile(ctx *gin.Context) {
	// need valid token to use this feature
	grpcVerifyAuthResponse, err := h.VerifyRolePermission(ctx, nil)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	// binding request
	request := &customer_api_payloads.UpdateCustomerProfileRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	grpcUpdateUserProfileRequest := &pb.GrpcUpdateUserProfileRequest{
		Email:           grpcVerifyAuthResponse.Email,
		RoleId:          constant.ROLE_ID_SUBSCRIBER,
		Title:           request.Title,
		LastName:        request.LastName,
		FirstName:       request.FirstName,
		DateOfBirth:     timestamppb.New(time.Time(request.DateOfBirth)),
		Phone:           request.Phone,
		TravelDocType:   request.TravelDocType,
		TravelDocNumber: request.TravelDocNumber,
		FfpNumber:       request.FfpNumber,
	}

	grpcUpdateUserProfileResponse, err := h.UserGrpcClient.UpdateUserProfile(ctx.Request.Context(), grpcUpdateUserProfileRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	response := customer_api_payloads.UpdateCustomerProfileResponse{
		RequestData: request,
		UpdateDate:  util.JsonDateTime(grpcUpdateUserProfileResponse.UpdateDate.AsTime()),
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
