package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		h.writeErrorResponse(c, constants.ErrorFailedToGetUserData, nil)
		return
	}

	data, ok := userData.(models.TokenData)
	if !ok {
		h.writeErrorResponse(c, constants.ErrorFailedToParseToken, nil)
		return
	}

	var request models.TransactionRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		h.writeErrorResponse(c, constants.ErrorBadRequest, nil)
		return
	}

	response, err := h.service.CreateTransaction(ctx, data.UserID, &request)
	if err != nil {
		h.writeErrorResponse(c, constants.ErrorTransactionFailed, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.TransactionSuccess, response)
}

func (h *TransactionHandler) writeErrorResponse(c *gin.Context, err error, data any) {
	var appErr *constants.AppError
	var valErrs validator.ValidationErrors

	if errors.As(err, &appErr) {
		helpers.SendResponseHTTP(c, appErr.StatusCode, appErr.Message, data)
		return
	}

	if errors.As(err, &valErrs) {
		errStr := helpers.ConstructErrString(valErrs)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, errStr, data)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusInternalServerError, err.Error(), nil)
}
