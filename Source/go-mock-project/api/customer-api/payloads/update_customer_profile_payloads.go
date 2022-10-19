package customer_api_payloads

import "anhlv11/go-mock-project/util"

type UpdateCustomerProfileRequest struct {
	Title           string        `json:"title" binding:"required,alpha,min=2,max=8"`
	LastName        string        `json:"lastName" binding:"required,alpha,min=1,max=25"`
	FirstName       string        `json:"firstName" binding:"required,alpha,min=1,max=25"`
	DateOfBirth     util.JsonDate `json:"dateOfBirth" binding:"required"`
	TravelDocType   string        `json:"travelDocType" binding:"alpha,len=1,oneof=P I"`
	TravelDocNumber string        `json:"travelDocNumber" binding:"min=1,max=25"`
	FfpNumber       string        `json:"ffpNumber" binding:"min=1,max=20"`
	Phone           string        `json:"phone" binding:"min=1,max=25"`
}

type UpdateCustomerProfileResponse struct {
	UpdateDate  util.JsonDateTime `json:"updateDate"`
	RequestData interface{}       `json:"requestData"`
}
