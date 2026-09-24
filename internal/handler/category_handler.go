package handler

import (
	"net/http"
	"strconv"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", categories)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	cat, err := h.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", "Category not found")
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", cat)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	if err := h.categoryService.CreateCategory(&req); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to POST data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to POST data", "")
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", err.Error())
		return
	}

	if err := h.categoryService.UpdateCategory(uint(id), &req); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", "")
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	if err := h.categoryService.DeleteCategory(uint(id)); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", "")
}
