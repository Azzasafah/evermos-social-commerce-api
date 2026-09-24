package handler

import (
	"net/http"

	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ProvCityHandler struct {
	provCityService service.ProvCityService
}

func NewProvCityHandler(provCityService service.ProvCityService) *ProvCityHandler {
	return &ProvCityHandler{provCityService: provCityService}
}

func (h *ProvCityHandler) GetListProvince(c *gin.Context) {
	provinces, err := h.provCityService.GetProvinces()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to get data", err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "Succeed to get data", provinces)
}

func (h *ProvCityHandler) GetListCities(c *gin.Context) {
	provID := c.Param("prov_id")
	cities, err := h.provCityService.GetCities(provID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to get data", err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "Succeed to get data", cities)
}

func (h *ProvCityHandler) GetDetailProvince(c *gin.Context) {
	provID := c.Param("prov_id")
	prov, err := h.provCityService.GetDetailProvince(provID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to get data", err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "Succeed to get data", prov)
}

func (h *ProvCityHandler) GetDetailCity(c *gin.Context) {
	cityID := c.Param("city_id")
	city, err := h.provCityService.GetDetailCity(cityID)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to get data", err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "Succeed to get data", city)
}
