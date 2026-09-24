package handler

import (
	"net/http"
	"strconv"

	"evermos-backend/internal/dto"
	"evermos-backend/internal/helper"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	nama := c.Query("nama_produk")
	if nama == "" {
		nama = c.Query("nama")
	}
	categoryID := c.Query("category_id")
	minHarga, _ := strconv.Atoi(c.Query("min_harga"))
	maxHarga, _ := strconv.Atoi(c.Query("max_harga"))

	res, err := h.productService.GetAllProducts(page, limit, nama, categoryID, minHarga, maxHarga)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", res)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", "No Data Product")
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", product)
}

func (h *ProductHandler) Create(c *gin.Context) {
	userID := c.GetUint("user_id")

	namaProduk := c.PostForm("nama_produk")
	categoryID, _ := strconv.ParseUint(c.PostForm("category_id"), 10, 32)
	hargaReseller, _ := strconv.Atoi(c.PostForm("harga_reseller"))
	hargaKonsumen, _ := strconv.Atoi(c.PostForm("harga_konsumen"))
	stok, _ := strconv.Atoi(c.PostForm("stok"))
	deskripsi := c.PostForm("deskripsi")

	if namaProduk == "" {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", "Nama produk tidak boleh kosong")
		return
	}

	req := dto.CreateProductRequest{
		NamaProduk:    namaProduk,
		CategoryID:    uint(categoryID),
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
	}

	form, _ := c.MultipartForm()
	var photos = form.File["photos"]
	if len(photos) == 0 {
		photos = form.File["photo"]
	}

	newID, err := h.productService.CreateProduct(userID, &req, photos)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to POST data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to POST data", newID)
}

func (h *ProductHandler) Update(c *gin.Context) {
	userID := c.GetUint("user_id")
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	namaProduk := c.PostForm("nama_produk")
	categoryID, _ := strconv.ParseUint(c.PostForm("category_id"), 10, 32)
	hargaReseller, _ := strconv.Atoi(c.PostForm("harga_reseller"))
	hargaKonsumen, _ := strconv.Atoi(c.PostForm("harga_konsumen"))
	stok, _ := strconv.Atoi(c.PostForm("stok"))
	deskripsi := c.PostForm("deskripsi")

	req := dto.UpdateProductRequest{
		NamaProduk:    namaProduk,
		CategoryID:    uint(categoryID),
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
	}

	form, _ := c.MultipartForm()
	var photos = form.File["photos"]
	if len(photos) == 0 && form != nil {
		photos = form.File["photo"]
	}

	if err := h.productService.UpdateProduct(uint(id), userID, &req, photos); err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", "")
}

func (h *ProductHandler) Delete(c *gin.Context) {
	userID := c.GetUint("user_id")
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "Failed to GET data", "Invalid ID")
		return
	}

	if err := h.productService.DeleteProduct(uint(id), userID); err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "Failed to GET data", "record not found")
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "Succeed to GET data", "")
}
