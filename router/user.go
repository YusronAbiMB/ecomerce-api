package router

import (
	"github.com/YusronAbi/ecomerce-api/handler"
	"github.com/YusronAbi/ecomerce-api/middleware"
	"github.com/YusronAbi/ecomerce-api/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) {
	UserRepository := repository.NewUserRepository(db)
	UserHandler := handler.NewHandlerUser(UserRepository)

	user := r.Group("/user")
	user.Use(middleware.AuthProtected(db), middleware.RoleRequired("admin"))
	{
		user.GET("", UserHandler.GetAllUser)
		user.GET("/:id", UserHandler.GetUserByID)
		user.PUT("/:id", UserHandler.UpdateUserByID)
		user.DELETE("/:id", UserHandler.DeleteUserByID)
	}
}
