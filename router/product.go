package router

import (
	"github.com/YusronAbi/ecomerce-api/handler"
	"github.com/YusronAbi/ecomerce-api/middleware"
	"github.com/YusronAbi/ecomerce-api/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRouter(r *gin.Engine, db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)
	productHandler := handler.NewHandlerProduct(productRepository)

	product := r.Group("/product")
	product.Use(middleware.AuthProtected(db))
	{
		product.GET("", productHandler.GetAllProduct)
		product.POST("", middleware.RoleRequired("admin"), productHandler.CreateProduct)
		product.GET("/:id", productHandler.GetProductByID)
		product.PUT("/:id", middleware.RoleRequired("admin"), productHandler.UpdateProductByID)
		product.DELETE("/:id", middleware.RoleRequired("admin"), productHandler.DeleteProductByID)
	}
}
