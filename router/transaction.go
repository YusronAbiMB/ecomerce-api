package router

import (
	"github.com/YusronAbi/ecomerce-api/handler"
	"github.com/YusronAbi/ecomerce-api/middleware"
	"github.com/YusronAbi/ecomerce-api/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTransactionRouter(r *gin.Engine, db *gorm.DB) {
	transactionRepository := repository.NewTransactionRepository(db)
	productRepository := repository.NewProductRepository(db)
	transactionHandler := handler.NewHandlerTransaction(transactionRepository, productRepository)

	transaction := r.Group("/transaction")
	transaction.GET("/reports", middleware.RoleRequired("admin"), transactionHandler.GetTotalTransactionReport)
	transaction.Use(middleware.AuthProtected(db), middleware.RoleRequired("user"))
	{
		transaction.GET("", transactionHandler.GetAllTransaction)
		transaction.POST("", transactionHandler.CreateTransaction)
		transaction.GET("/:id", transactionHandler.GetTransactionByID)
		transaction.PUT("/:id", transactionHandler.UpdateTransactionByID)
		transaction.PATCH("/payment/:id", transactionHandler.UpdateTransactionPayment)
		transaction.DELETE("/:id", transactionHandler.DeleteTransactionByID)
	}
}
