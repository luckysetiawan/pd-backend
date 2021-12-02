package controllers

import (
	"net/http"
	model "pd-backend/model"

	"github.com/gin-gonic/gin"
)

// Menu Response
func sendMenuSuccessresponse(c *gin.Context, ur model.MenuResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendMenuErrorResponse(c *gin.Context, ur model.MenuResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

// User Response
func sendUserSuccessresponse(c *gin.Context, ur model.UserResponse) {
	c.JSON(http.StatusOK, ur)
}

func sendUserErrorResponse(c *gin.Context, ur model.UserResponse) {
	c.JSON(http.StatusBadRequest, ur)
}

// Payment Response
func sendPaymentSuccessresponse(c *gin.Context, pr model.PaymentResponse) {
	c.JSON(http.StatusOK, pr)
}

func sendPaymentErrorResponse(c *gin.Context, pr model.PaymentResponse) {
	c.JSON(http.StatusBadRequest, pr)
}

func sendGetPendapatanPeriodSuccessResponse(c *gin.Context, pr model.PendapatanResponse) {
	c.JSON(http.StatusOK, pr)
}

func sendGetPendapatanPeriodErrorResponse(c *gin.Context, pr model.PendapatanResponse) {
	c.JSON(http.StatusBadRequest, pr)
}

// Order Response
func sendSuccessCustomerresponse(c *gin.Context, or model.CustomerResponse) {
	c.JSON(http.StatusOK, or)
}

func sendErrorCustomerResponse(c *gin.Context, or model.CustomerResponse) {
	c.JSON(http.StatusBadRequest, or)
}

func sendSuccessOrderDetailresponse(c *gin.Context, or model.OrderDetailResponse) {
	c.JSON(http.StatusOK, or)
}

func sendErrorOrderDetailResponse(c *gin.Context, or model.OrderDetailResponse) {
	c.JSON(http.StatusBadRequest, or)
}

// General Response
func sendSuccessresponse(c *gin.Context, or model.OrderResponse) {
	c.JSON(http.StatusOK, or)
}

func sendErrorResponse(c *gin.Context, or model.OrderResponse) {
	c.JSON(http.StatusBadRequest, or)
}
