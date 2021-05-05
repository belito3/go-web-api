package api

import (
	"fmt"
	"net/http"

	"github.com/belito3/go-web-api/app/repository/impl"
	"github.com/belito3/go-web-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Account struct {
	store impl.IStore
}

func NewAccount(store impl.IStore) *Account {
	return &Account{store: store}
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (s *Account) createAccount(c *gin.Context) {
	// Add account
	ctx := c.Request.Context()
	var req createAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Errorf(ctx, "Invalid input parameter")
		responseError(c, http.StatusBadRequest, fmt.Sprintf("Invalid input parameter: %v", err))
		return
	}
	arg := impl.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		logger.Errorf(ctx, "Add account error %v\n", err)
		responseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	r := map[string]interface{}{
		"account": account,
	}
	responseSuccess(c, http.StatusOK, r)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Account) getAccount(c *gin.Context) {
	ctx := c.Request.Context()
	var req getAccountRequest
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Errorf(ctx, "Invalid input parameter")
		responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		logger.Errorf(ctx, "GetAccount error: %v\n", err)
		responseError(c, http.StatusInternalServerError, err.Error())
		return

	}

	r := map[string]interface{}{
		"account": account,
	}
	responseSuccess(c, http.StatusOK, r)
}
