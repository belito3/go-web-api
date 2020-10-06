package service

import (
	"fmt"
	repo "github.com/belito3/go-api-codebase/app/repository"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountService struct {
	AccRepo repo.IAccount
}

func NewAccountService(accRepo repo.IAccount) *AccountService {
	return &AccountService{AccRepo: accRepo}
}

func (s *AccountService) Add(c *gin.Context) {
	// Add account
	ctx := c.Request.Context()
	var arg repo.CreateAccountParams
	if err := c.ShouldBindJSON(&arg); err != nil {
		logger.Errorf(ctx, "Invalid input parameter")
		ResponseError(c, http.StatusBadRequest, fmt.Sprintf("Invalid input parameter: %v", err))
		return
	}
	account, err := s.AccRepo.CreateAccount(ctx, arg)
	if err != nil {
		logger.Errorf(ctx, "Add account error %v\n", err)
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	r := map[string]interface{}{
		"account": account,
	}
	ResponseSuccess(c, http.StatusOK, r)
}