package handler

import (
	"net/http"
	"strconv"

	"github.com/YusronAbi/ecomerce-api/helper"
	"github.com/YusronAbi/ecomerce-api/models"
	"github.com/gin-gonic/gin"
)

type HandlerTransaction struct {
	repositoryTransaction models.TransactionRepository
	repositoryProduct     models.ProductRepository
}

func NewHandlerTransaction(repositoryTransaction models.TransactionRepository, repositoryProduct models.ProductRepository) HandlerTransaction {
	return HandlerTransaction{repositoryTransaction: repositoryTransaction, repositoryProduct: repositoryProduct}
}

func (h *HandlerTransaction) GetAllTransaction(ctx *gin.Context) {
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "User not found",
		})
		return
	}
	transaction, err := h.repositoryTransaction.GetAllTransaction(ctx, int64(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}

	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", transaction))
}

func (h *HandlerTransaction) CreateTransaction(ctx *gin.Context) {
	transaction := &models.Transaction{}
	if err := ctx.ShouldBind(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Payload invalid"))
		return
	}
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "User not found",
		})
		return
	}
	product, err := h.repositoryProduct.GetProductByID(ctx, int64(transaction.ProductID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	if transaction.Quantity > product.Stock {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Stock not available"))
		return
	}
	updateData := map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"stock":       product.Stock - transaction.Quantity,
		"category_id": product.CategoryID,
	}
	_, err = h.repositoryProduct.UpdateProductByID(ctx, int64(transaction.ProductID), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update product"))
		return
	}
	transaction.UserID = int64(userID)
	transaction.Price = product.Price
	transaction.SubTotal = float64(transaction.Quantity) * transaction.Price
	transaction, err = h.repositoryTransaction.CreateTransaction(ctx, transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	ctx.JSON(http.StatusCreated, helper.ResponseSuccess("Create data successfuly", transaction))
}

func (h *HandlerTransaction) GetTransactionByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "User not found",
		})
		return
	}
	transaction, err := h.repositoryTransaction.GetTransactionByID(ctx, int64(userID), int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", transaction))
}

func (h *HandlerTransaction) GetTotalTransactionReport(ctx *gin.Context) {
	totalSales, err := h.repositoryTransaction.GetTransactionReport(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to calculate total amount"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data success", totalSales))
}

func (h *HandlerTransaction) UpdateTransactionByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "User not found",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	transaction, err := h.repositoryTransaction.GetTransactionByID(ctx, int64(userID), int64(id))
	updateData := models.Transaction{}
	if err := ctx.ShouldBind(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	data := map[string]interface{}{
		"id":         transaction.ID,
		"user_id":    transaction.UserID,
		"product_id": transaction.ProductID,
		"quantity":   updateData.Quantity,
		"price":      transaction.Price,
		"sub_total":  float64(updateData.Quantity) * transaction.Price,
	}
	updateTransaction, err := h.repositoryTransaction.UpdateTransactionByID(ctx, int64(userID), int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfuly", updateTransaction))
}
func (h *HandlerTransaction) UpdateTransactionPayment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "User not found",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}

	updateTransaction, err := h.repositoryTransaction.UpdateTransactionPayment(ctx, int64(userID), int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfuly", updateTransaction))
}

func (h *HandlerTransaction) DeleteTransactionByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "User not found",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	err = h.repositoryTransaction.DeleteTransactionByID(ctx, int64(userID), int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfuly", nil))
}
