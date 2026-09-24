package handler

import (
	"net/http"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	res, err := h.userService.GetProfile(userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", res)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", err.Error())
		return
	}

	if err := h.userService.UpdateProfile(userID, &req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", "")
}
