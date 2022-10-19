package customer_api_handlers

import (
	customer_api_payloads "anhlv11/go-mock-project/api/customer-api/payloads"
	"anhlv11/go-mock-project/constant"
	"anhlv11/go-mock-project/pb"
	"anhlv11/go-mock-project/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *CustomerApiHandler) RegisterCustomer(ctx *gin.Context) {
	request := &customer_api_payloads.RegisterCustomerRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusOK, util.CreateApiErrorResponse(BAD_REQUEST_ERROR, err.Error()))
		return
	}

	email := strings.ToLower(strings.TrimSpace(request.Email))
	grpcRegisterUserRequest := &pb.GrpcRegisterUserRequest{
		Email:           email,
		RoleId:          constant.ROLE_ID_SUBSCRIBER,
		Title:           request.Title,
		LastName:        request.LastName,
		FirstName:       request.FirstName,
		DateOfBirth:     timestamppb.New(time.Time(request.DateOfBirth)),
		Phone:           request.Phone,
		TravelDocType:   request.TravelDocType,
		TravelDocNumber: request.TravelDocNumber,
		FfpNumber:       request.FfpNumber,
		Password:        request.Password,
	}

	grpcRegisterUserResponse, err := h.UserGrpcClient.RegisterUser(ctx.Request.Context(), grpcRegisterUserRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, util.CreateCustomApiErrorResponse(BAD_REQUEST_ERROR, err.Error(), nil))
		return
	}

	response := customer_api_payloads.RegisterCustomerResponse{
		RequestData:  request,
		RegisterDate: util.JsonDateTime(grpcRegisterUserResponse.RegisterDate.AsTime()),
	}
	ctx.JSON(http.StatusOK, util.CreateSuccessResponse(response))
}
