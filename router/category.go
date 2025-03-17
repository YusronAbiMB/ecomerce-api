package router

import (
	"github.com/YusronAbi/ecomerce-api/handler"
	"github.com/YusronAbi/ecomerce-api/middleware"
	"github.com/YusronAbi/ecomerce-api/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCategoryRouter(r *gin.Engine, db *gorm.DB) {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryHandler := handler.NewHandlerCategory(categoryRepository)

	category := r.Group("/category")
	category.Use(middleware.AuthProtected(db))
	{
		category.GET("", categoryHandler.GetAllCategory)
		category.POST("", middleware.RoleRequired("admin"), categoryHandler.CreateCategory)
		category.GET("/:id", categoryHandler.GetCategoryByID)
		category.PUT("/:id", middleware.RoleRequired("admin"), categoryHandler.UpdateCategoryByID)
		category.DELETE("/:id", middleware.RoleRequired("admin"), categoryHandler.DeleteCategoryByID)
	}
}
