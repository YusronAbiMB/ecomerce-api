package handler

import (
	"net/http"
	"strconv"

	"github.com/YusronAbi/ecomerce-api/helper"
	"github.com/YusronAbi/ecomerce-api/models"
	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	repository models.ProductRepository
}

func NewHandlerProduct(repository models.ProductRepository) HandlerProduct {
	return HandlerProduct{repository: repository}
}

func (h *HandlerProduct) GetAllProduct(c *gin.Context) {
	products, err := h.repository.GetAllProduct(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", products))
}

func (h *HandlerProduct) CreateProduct(ctx *gin.Context) {
	product := &models.Product{}
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Payload invalid"))
		return
	}
	product, err := h.repository.CreateProduct(ctx, product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	ctx.JSON(http.StatusCreated, helper.ResponseSuccess("Create data successfuly", product))
}

func (h *HandlerProduct) GetProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	product, err := h.repository.GetProductByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", product))
}

func (h *HandlerProduct) UpdateProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	product, err := h.repository.GetProductByID(ctx, int64(id))
	updateData := models.Product{}
	if err := ctx.ShouldBind(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	data := map[string]interface{}{
		"id":          product.ID,
		"name":        updateData.Name,
		"description": updateData.Description,
		"price":       updateData.Price,
		"stock":       updateData.Stock,
		"category_id": updateData.CategoryID,
	}
	updateCategory, err := h.repository.UpdateProductByID(ctx, int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfuly", updateCategory))
}

func (h *HandlerProduct) DeleteProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	err = h.repository.DeleteProductByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfuly", nil))
}
