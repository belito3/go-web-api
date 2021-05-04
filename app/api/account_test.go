package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/belito3/go-web-api/app/repository/mock"

	"github.com/belito3/go-web-api/app/config"
	"github.com/belito3/go-web-api/app/repository/impl"
	"github.com/belito3/go-web-api/app/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockIStore(ctrl)
	// build stubs
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// Build container
	container := dig.New()
	conf := config.AppConfiguration{}
	//// Inject store to container
	_ = container.Provide(func() impl.IStore {
		return store
	})

	// start test server and send request
	server := NewServer(conf, container)
	// Inject api to container
	_ = server.InitGinEngine()

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/v1/account/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomAccount() impl.Account {
	return impl.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
