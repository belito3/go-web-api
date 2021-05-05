package api

import (
	"os"
	"testing"

	"github.com/belito3/go-web-api/app/config"
	"github.com/belito3/go-web-api/app/repository/impl"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func newTestServer(t *testing.T, store impl.IStore) *Server {
	// Build container
	container := dig.New()
	conf := config.AppConfiguration{RunMode: gin.TestMode}
	// Inject store to container
	_ = container.Provide(func() impl.IStore {
		return store
	})

	// start test server and send request
	server := NewServer(conf, container)
	// Inject api to container
	_ = server.InitGinEngine()
	gin.SetMode(gin.TestMode)
	return server
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
