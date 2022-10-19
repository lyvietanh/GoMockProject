package customer_api_payloads

import "anhlv11/go-mock-project/util"

type RegisterCustomerRequest struct {
	Email           string        `json:"email" binding:"required"`
	Title           string        `json:"title" binding:"required,min=2,max=8"`
	LastName        string        `json:"lastName" binding:"required,min=1,max=25"`
	FirstName       string        `json:"firstName" binding:"required,min=1,max=25"`
	DateOfBirth     util.JsonDate `json:"dateOfBirth" binding:"required"`
	TravelDocType   string        `json:"travelDocType" binding:"alpha,len=1,oneof=P I"`
	TravelDocNumber string        `json:"travelDocNumber" binding:"min=1,max=25"`
	FfpNumber       string        `json:"ffpNumber" binding:"min=1,max=20"`
	Phone           string        `json:"phone" binding:"min=1,max=25"`
	Password        string        `json:"password" binding:"min=6,max=25"`
}

type RegisterCustomerResponse struct {
	RegisterDate util.JsonDateTime `json:"registerDate"`
	RequestData  interface{}       `json:"requestData"`
}
