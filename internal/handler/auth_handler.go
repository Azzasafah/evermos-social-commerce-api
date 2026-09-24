package handler

import (
	"net/http"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	if err := h.authService.Register(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to POST data", "Register Succeed")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	res, err := h.authService.Login(&req)
	if err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, "Failed to POST data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to POST data", res)
}
