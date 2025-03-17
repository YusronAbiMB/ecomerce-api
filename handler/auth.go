package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/YusronAbi/ecomerce-api/helper"
	"github.com/YusronAbi/ecomerce-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AuthHandler struct {
	service models.AuthService
}

func NewAuthHandler(
	service models.AuthService,
) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	creds := &models.AuthCredentials{}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid payload"))
		return
	}

	if err := validate.Struct(creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed models"))
		return
	}

	token, _, err := h.service.Login(ctxTimeout, creds)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		return
	}
	ctx.SetCookie("token", token, 3600*24*1, "/", "", false, true)
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Login success", token))
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	creds := &models.User{}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid payload"))
		return
	}

	if !helper.IsValidEmail(creds.Email) {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Email not valid"))
		return
	}
	if !helper.IsValidPassword(creds.Password) {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Password must be 8 and combine char"))
		return
	}

	if err := validate.Struct(creds); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed models"))
		return
	}

	token, _, err := h.service.Register(ctxTimeout, creds)
	ctx.SetCookie("token", token, 3600*24*1, "/", "", false, true)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Bad Request"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.ResponseSuccess("Register success", token))
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Logout success", nil))
}
