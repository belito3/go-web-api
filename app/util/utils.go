package util

import (
	"context"
	"encoding/json"
	"github.com/belito3/go-api-codebase/app/config"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"github.com/belito3/go-api-codebase/pkg/unique"
)

var idFunc = func() string {
	return unique.NewSnowflakeID().String()
}

// InitID ...
func InitID() {
	switch config.C.UniqueID.Type {
	case "uuid":
		idFunc = func() string {
			return unique.MustUUID().String()
		}
	case "object":
		idFunc = func() string {
			return unique.NewObjectID().Hex()
		}
	default:
		// Initialize snowflake node
		err := unique.SetSnowflakeNode(config.C.UniqueID.Snowflake.Node, config.C.UniqueID.Snowflake.Epoch)
		if err != nil {
			panic(err)
		}

		logger.SetTraceIDFunc(func() string {
			return unique.NewSnowflakeID().String()
		})

		idFunc = func() string {
			return unique.NewSnowflakeID().String()
		}
	}
}

// NewID Create unique id
func NewID() string {
	return idFunc()
}

func PrintInterface(v interface{}){
	s, err := json.MarshalIndent(v, "", "\t")
	//s, err := json.MarshalIndent(v, "", "")
	if err != nil {
		logger.Errorf(context.Background(),"error print interface")
		return
	} else {
		logger.Infof(context.Background(), string(s))
	}
}