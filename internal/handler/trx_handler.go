package handler

import (
	"net/http"
	"strconv"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type TrxHandler struct {
	trxService service.TrxService
}

func NewTrxHandler(trxService service.TrxService) *TrxHandler {
	return &TrxHandler{trxService: trxService}
}

func (h *TrxHandler) Create(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req dto.CreateTrxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	newID, err := h.trxService.CreateTrx(userID, &req)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to POST data", newID)
}

func (h *TrxHandler) GetAll(c *gin.Context) {
	userID := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))

	res, err := h.trxService.GetAllTrx(userID, page, limit)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", res)
}

func (h *TrxHandler) GetByID(c *gin.Context) {
	userID := c.GetUint("user_id")
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	trx, err := h.trxService.GetTrxByID(uint(id), userID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", "No Data Trx")
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", trx)
}
