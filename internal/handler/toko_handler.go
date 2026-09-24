package handler

import (
	"net/http"
	"strconv"

	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type TokoHandler struct {
	tokoService service.TokoService
}

func NewTokoHandler(tokoService service.TokoService) *TokoHandler {
	return &TokoHandler{tokoService: tokoService}
}

func (h *TokoHandler) GetMyToko(c *gin.Context) {
	userID := c.GetUint("user_id")
	toko, err := h.tokoService.GetMyToko(userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", "Toko tidak ditemukan")
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", toko)
}

func (h *TokoHandler) UpdateToko(c *gin.Context) {
	userID := c.GetUint("user_id")
	tokoIDParam := c.Param("id_toko")
	tokoID, err := strconv.ParseUint(tokoIDParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to UPDATE data", "Invalid ID toko")
		return
	}

	namaToko := c.PostForm("nama_toko")
	photoHeader, _ := c.FormFile("photo")

	if err := h.tokoService.UpdateToko(userID, uint(tokoID), namaToko, photoHeader); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to UPDATE data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to UPDATE data", "Update toko succeed")
}

func (h *TokoHandler) GetTokoByID(c *gin.Context) {
	tokoIDParam := c.Param("id_toko")
	tokoID, err := strconv.ParseUint(tokoIDParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID toko")
		return
	}

	toko, err := h.tokoService.GetTokoByID(uint(tokoID))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", toko)
}

func (h *TokoHandler) GetAllToko(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	nama := c.Query("nama")

	res, err := h.tokoService.GetAllToko(page, limit, nama)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", res)
}
