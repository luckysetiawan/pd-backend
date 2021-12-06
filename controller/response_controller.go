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

func sendGetPizzaTerjualPeriodSuccessResponse(c *gin.Context, pr model.PizzaTerjualResponse) {
	c.JSON(http.StatusOK, pr)
}

func sendGetPizzaTerjualPeriodErrorResponse(c *gin.Context, pr model.PizzaTerjualResponse) {
	c.JSON(http.StatusBadRequest, pr)
}

// Order Response
func sendOrderSuccessResponse(c *gin.Context, or model.OrderResponse) {
	c.JSON(http.StatusOK, or)
}

func sendOrderErrorOResponse(c *gin.Context, or model.OrderResponse) {
	c.JSON(http.StatusBadRequest, or)
}

func sendStatusSuccessResponse(c *gin.Context, or model.StatusResponse) {
	c.JSON(http.StatusOK, or)
}

func sendStatusErrorResponse(c *gin.Context, or model.StatusResponse) {
	c.JSON(http.StatusBadRequest, or)
}

func sendActiveOrderSuccessResponse(c *gin.Context, or model.ActiveOrderResponse) {
	c.JSON(http.StatusOK, or)
}

func sendActiveOrderErrorResponse(c *gin.Context, or model.ActiveOrderResponse) {
	c.JSON(http.StatusBadRequest, or)
}

// Histories
func sendHistorySuccessResponse(c *gin.Context, tr model.TransactionResponse) {
	c.JSON(http.StatusOK, tr)
}

func sendHistoryErrorResponse(c *gin.Context, tr model.TransactionResponse) {
	c.JSON(http.StatusBadRequest, tr)
}
