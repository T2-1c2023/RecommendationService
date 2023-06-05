package model

type UserServiceErrorResponse struct {
	Message string `json:"message" binding:"required"`
}
