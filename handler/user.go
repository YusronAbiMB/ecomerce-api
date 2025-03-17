package handler

import (
	"net/http"
	"strconv"

	"github.com/YusronAbi/ecomerce-api/helper"
	"github.com/YusronAbi/ecomerce-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type HandlerUser struct {
	repositoryUser models.UserRepository
}

func NewHandlerUser(repositoryUser models.UserRepository) HandlerUser {
	return HandlerUser{repositoryUser: repositoryUser}
}

func (h *HandlerUser) GetAllUser(c *gin.Context) {
	categories, err := h.repositoryUser.GetAllUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", categories))
}

func (h *HandlerUser) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	User, err := h.repositoryUser.GetUserByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfuly", User))
}

func (h *HandlerUser) UpdateUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	user, err := h.repositoryUser.GetUserByID(ctx, int64(id))
	updateData := models.User{}
	if err := ctx.ShouldBind(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to hashed"))
		return
	}

	updateData.Password = string(hashedPassword)
	data := map[string]interface{}{
		"id":       user.ID,
		"name":     updateData.Name,
		"email":    user.Email,
		"password": hashedPassword,
		"role":     user.Role,
	}
	updateUser, err := h.repositoryUser.UpdateUserByID(ctx, int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfuly", updateUser))
}

func (h *HandlerUser) DeleteUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to get data"))
		return
	}
	err = h.repositoryUser.DeleteUserByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfuly", nil))
}
