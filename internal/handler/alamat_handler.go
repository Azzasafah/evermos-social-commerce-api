package handler

import (
	"net/http"
	"strconv"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AlamatHandler struct {
	alamatService service.AlamatService
}

func NewAlamatHandler(alamatService service.AlamatService) *AlamatHandler {
	return &AlamatHandler{alamatService: alamatService}
}

func (h *AlamatHandler) GetMyAlamat(c *gin.Context) {
	userID := c.GetUint("user_id")
	alamats, err := h.alamatService.GetMyAlamat(userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", alamats)
}

func (h *AlamatHandler) GetAlamatByID(c *gin.Context) {
	userID := c.GetUint("user_id")
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	alamat, err := h.alamatService.GetAlamatByID(uint(id), userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", alamat)
}

func (h *AlamatHandler) Create(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req dto.CreateAlamatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	newID, err := h.alamatService.CreateAlamat(userID, &req)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to POST data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to POST data", newID)
}

func (h *AlamatHandler) Update(c *gin.Context) {
	userID := c.GetUint("user_id")
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	var req dto.UpdateAlamatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", err.Error())
		return
	}

	if err := h.alamatService.UpdateAlamat(uint(id), userID, &req); err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", "")
}

func (h *AlamatHandler) Delete(c *gin.Context) {
	userID := c.GetUint("user_id")
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	if err := h.alamatService.DeleteAlamat(uint(id), userID); err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", "record not found")
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", "")
}
