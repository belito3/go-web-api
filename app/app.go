package app

import (
	"context"
	"github.com/belito3/go-api-codebase/app/config"
	"github.com/belito3/go-api-codebase/app/service"
	"github.com/belito3/go-api-codebase/app/util"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type options struct {
	ConfigFile	string
	Version		string
}

// Option
type Option func(*options)

// SetConfigFile
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetVersion
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// Run server
func Run(ctx context.Context, opts ...Option) error {
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	// SIGHUP: (signal hang up) sent to a process when its controlling terminal is closed, such as daemons
	// SIGINT: Ctrl-C sends an INT signal ("interrupt")
	// SIGTERM: signal is sent to a proc ess to request its  termination, allows process releasing releasing resources and saving state
	// SIGKILL: sent to a process to cause it to terminate immediately (kill), can't perform any clean-up upon receiving this signal
	// SIGQUIT: when user requests that the process quit and perform a core dump
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

	EXIT:
		for {
			sig := <- sc
			logger.Printf(ctx, "Received a signal[%s]", sig.String())
			switch sig {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				atomic.CompareAndSwapInt32(&state, 1, 0)
				break EXIT
			case syscall.SIGHUP:
			default:
				break EXIT
			}
		}

		cleanFunc()
		logger.Printf(ctx, "Service exit")
		time.Sleep(time.Second)
		os.Exit(int(atomic.LoadInt32(&state)))
		return nil
}

// Init
func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	// Init global config
	config.Init(o.ConfigFile)
	config.PrintWithJSON()
	logger.Printf(ctx, "Service started, running mode：%s，version number：%s，process number：%d", config.C.RunMode, o.Version, os.Getpid())

	// Initialize trace_id for node that app is running
	// TODO: uuid, object, snowflake
	util.InitID()

	// Init logger
	setupLogger()

	container, containerCall := BuildContainer()

	httpServerCleanFunc := InitHTTPServer(ctx, container)

	return func() {
		httpServerCleanFunc()
		containerCall()
	}, nil
}


func BuildContainer() (*dig.Container, func()) {
	container := dig.New()

	// store DB
	storeCall, err := InitStore(container)
	handleError(err)

	// register service
	err = service.Inject(container)
	handleError(err)

	return container, func() {
		if storeCall != nil {
			storeCall()
		}
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

