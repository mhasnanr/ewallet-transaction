package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhasnanr/ewallet-transaction/constants"
	"github.com/mhasnanr/ewallet-transaction/helpers"
	"github.com/mhasnanr/ewallet-transaction/internal/models"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, userID int, transaction *models.TransactionRequest) (*models.TransactionResponse, error)
}

type AuthMiddleware interface {
	MiddlewareAccessToken(c *gin.Context)
}

type TransactionHandler struct {
	service        TransactionService
	authMiddleware AuthMiddleware
}

func NewTransactionHandler(service TransactionService, authMiddleware AuthMiddleware) *TransactionHandler {
	return &TransactionHandler{service, authMiddleware}
}

func (h *TransactionHandler) RegisterRoute(r *gin.Engine) {
	walletV1 := r.Group("/transactions/v1")
	walletV1.POST("/", h.authMiddleware.MiddlewareAccessToken, h.createTransaction)
}

func (h *TransactionHandler) createTransaction(c *gin.Context) {
	ctx := c.Request.Context()

	userData, ok := c.Get("tokenData")

	if !ok {
		c.Error(constants.ErrorFailedToGetUserData)
		return
	}

	data, ok := userData.(models.TokenData)
	if !ok {
		c.Error(constants.ErrorFailedToParseToken)
		return
	}

	var request models.TransactionRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.Error(err)
		return
	}

	response, err := h.service.CreateTransaction(ctx, data.UserID, &request)
	if err != nil {
		c.Error(constants.ErrorTransactionFailed)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.TransactionSuccess, response)
}
