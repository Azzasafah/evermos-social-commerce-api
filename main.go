package main

import (
	"log"
	"os"

	"evermos-backend/config"
	"evermos-backend/internal/handler"
	"evermos-backend/internal/middleware"
	"evermos-backend/internal/repository"
	"evermos-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// 1. Initialize Database
	db := config.ConnectDB()

	// 2. Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	tokoRepo := repository.NewTokoRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	alamatRepo := repository.NewAlamatRepository(db)
	productRepo := repository.NewProductRepository(db)
	trxRepo := repository.NewTrxRepository(db)

	// 3. Initialize Services
	provCityService := service.NewProvCityService()
	authService := service.NewAuthService(userRepo, tokoRepo, provCityService)
	userService := service.NewUserService(userRepo, provCityService)
	tokoService := service.NewTokoService(tokoRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	alamatService := service.NewAlamatService(alamatRepo)
	productService := service.NewProductService(productRepo, tokoRepo)
	trxService := service.NewTrxService(trxRepo, alamatRepo, productRepo)

	// 4. Initialize Handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	tokoHandler := handler.NewTokoHandler(tokoService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	alamatHandler := handler.NewAlamatHandler(alamatService)
	productHandler := handler.NewProductHandler(productService)
	trxHandler := handler.NewTrxHandler(trxService)
	provCityHandler := handler.NewProvCityHandler(provCityService)

	// 5. Setup Router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Static route for serving uploaded images
	r.Static("/uploads", "./uploads")

	// Public Routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	provcity := r.Group("/provcity")
	{
		provcity.GET("/listprovincies", provCityHandler.GetListProvince)
		provcity.GET("/listcities/:prov_id", provCityHandler.GetListCities)
		provcity.GET("/detailprovince/:prov_id", provCityHandler.GetDetailProvince)
		provcity.GET("/detailcity/:city_id", provCityHandler.GetDetailCity)
	}

	// Protected Routes (Authentication required)
	api := r.Group("")
	api.Use(middleware.AuthMiddleware(db))
	{
		// Category (Admin only for POST, PUT, DELETE)
		category := api.Group("/category")
		{
			category.GET("", categoryHandler.GetAll)
			category.GET("/:id", categoryHandler.GetByID)

			adminCategory := category.Group("")
			adminCategory.Use(middleware.AdminMiddleware())
			{
				adminCategory.POST("", categoryHandler.Create)
				adminCategory.PUT("/:id", categoryHandler.Update)
				adminCategory.DELETE("/:id", categoryHandler.Delete)
			}
		}

		// User
		user := api.Group("/user")
		{
			user.GET("", userHandler.GetProfile)
			user.PUT("", userHandler.UpdateProfile)

			// Alamat
			alamat := user.Group("/alamat")
			{
				alamat.GET("", alamatHandler.GetMyAlamat)
				alamat.GET("/:id", alamatHandler.GetAlamatByID)
				alamat.POST("", alamatHandler.Create)
				alamat.PUT("/:id", alamatHandler.Update)
				alamat.DELETE("/:id", alamatHandler.Delete)
			}
		}

		// Toko
		toko := api.Group("/toko")
		{
			toko.GET("/my", tokoHandler.GetMyToko)
			toko.PUT("/:id_toko", tokoHandler.UpdateToko)
			toko.GET("/:id_toko", tokoHandler.GetTokoByID)
			toko.GET("", tokoHandler.GetAllToko)
		}

		// Product
		product := api.Group("/product")
		{
			product.GET("", productHandler.GetAll)
			product.GET("/:id", productHandler.GetByID)
			product.POST("", productHandler.Create)
			product.PUT("/:id", productHandler.Update)
			product.DELETE("/:id", productHandler.Delete)
		}

		// Transaction (Trx)
		trx := api.Group("/trx")
		{
			trx.GET("", trxHandler.GetAll)
			trx.GET("/:id", trxHandler.GetByID)
			trx.POST("", trxHandler.Create)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
