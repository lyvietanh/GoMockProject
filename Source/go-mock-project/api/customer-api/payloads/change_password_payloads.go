package customer_api_payloads

import "anhlv11/go-mock-project/util"

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"min=6,max=25"`
	NewPassword     string `json:"newPassword" binding:"min=6,max=25"`
}

type ChangePasswordResponse struct {
	UpdateDate util.JsonDateTime `json:"updateDate"`
}
