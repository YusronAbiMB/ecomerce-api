package handler

import (
	"net/http"
	"strconv"

	"github.com/YusronAbi/ecomerce-api/helper"
	"github.com/YusronAbi/ecomerce-api/models"
	"github.com/gin-gonic/gin"
)

type HandlerCategory struct {
	repositoryCategory models.CategoryRepository
}

func NewHandlerCategory(repositoryCategory models.CategoryRepository) HandlerCategory {
	return HandlerCategory{repositoryCategory: repositoryCategory}
}

func (h *HandlerCategory) GetAllCategory(c *gin.Context) {
	categories, err := h.repositoryCategory.GetAllCategory(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", categories))
}

func (h *HandlerCategory) CreateCategory(ctx *gin.Context) {
	category := &models.Category{}
	if err := ctx.ShouldBind(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Payload invalid"))
		return
	}
	category, err := h.repositoryCategory.CreateCategory(ctx, category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	ctx.JSON(http.StatusCreated, helper.ResponseSuccess("Create data successfuly", category))
}

func (h *HandlerCategory) GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	category, err := h.repositoryCategory.GetCategoryByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", category))
}

func (h *HandlerCategory) UpdateCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	category, err := h.repositoryCategory.GetCategoryByID(ctx, int64(id))
	updateData := models.Category{}
	if err := ctx.ShouldBind(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	data := map[string]interface{}{
		"id":   category.ID,
		"name": updateData.Name,
	}
	updateCategory, err := h.repositoryCategory.UpdateCategoryByID(ctx, int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfuly", updateCategory))
}

func (h *HandlerCategory) DeleteCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	err = h.repositoryCategory.DeleteCategoryByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfuly", nil))
}
