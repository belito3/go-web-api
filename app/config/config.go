package config

import (
	"fmt"
	"github.com/belito3/go-api-codebase/pkg/logger"
	"github.com/belito3/go-api-codebase/pkg/util"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Configuration ... The configuration of system
var C = new(AppConfiguration)

func Init(fileConf string){
	// fileConf: file config path
	//loadConfiguration
	yamlFile, err := ioutil.ReadFile(fileConf)
	if err != nil {
		logger.Errorf(nil,"Can not load configuration file %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &C)
	if err != nil {
		logger.Errorf(nil, "Can not load App Configuration %v", err)
	}

	return
}

// Print With JSON
func PrintWithJSON() {
	if C.PrintConfig {
		b, err := util.JSONMarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

// Configuration ... The configuration of system
type AppConfiguration struct {
	HTTP					HTTP			`yaml:"http"`
	PrintConfig  			bool			`yaml:"print_config"`
	RunMode					string			`yaml:"run_mode"`
	CORS					CORS			`yaml:"cors"`
	Log						Log				`yaml:"log"`
	UniqueID				UniqueID		`yaml:"unique_id"`
	JWTSecretKey            string 			`yaml:"jwt_secret_key"`
	ARateLimiter			ARateLimiter	`yaml:"app_rate_limiter"`
	CRateLimiter			CRateLimiter	`yaml:"client_rate_limiter"`
	Redis					Redis			`yaml:"redis"`
	Postgres				Postgres		`yaml:"postgres"`
}

// HTTP http
type HTTP struct {
	Host				string		`yaml:"host"`
	Port				int			`yaml:"port"`
	ShutdownTimeout		int			`yaml:"shutdown_timeout"`
	MaxContentLength	int64		`yaml:"max_content_length"`
}

// CORS Cross-domain request configuration parameters
type CORS struct {
	Enable					bool		`yaml:"enable"`
	AllowOrigins			string		`yaml:"allow_origins"`
	AllowMethods			string		`yaml:"allow_methods"`
	AllowHeaders			string		`yaml:"allow_headers"`
	AllowCredentials		string		`yaml:"allow_credentials"`
	MaxAge					string		`yaml:"max_age"`
}

// Log
type Log struct {
	Level		int			`yaml:"level"`
	Format 		string		`yaml:"format"`
}

// UniqueID
type UniqueID struct {
	Type		string		`yaml:"type"`
	Snowflake	struct {
		Node	int64		`yaml:"node"`
		Epoch	int64		`yaml:"epoch"`
	}
}

// App RateLimiter
type ARateLimiter struct {
	Enable		bool	`yaml:"enable"`
	Count		int		`yaml:"count"`
}

// Client RateLimiter Request frequency limit configuration parameters
type CRateLimiter struct {
	Enable		bool	`yaml:"enable"`
	Count		int		`yaml:"count"`
	RedisDB 	int		`yaml:"redis_db"`
}

// Redis redis Configuration parameter
type Redis struct {
	Addr		string		`yaml:"addr"`
	Password	string		`yaml:"password"`
}

// PostgreSQL Postgres config
type Postgres struct {
	Username		string		`yaml:"username"`
	Password		string		`yaml:"password"`
	Host			string		`yaml:"host"`
	Port			int			`yaml:"port"`
	DatabaseName	string		`yaml:"database_name"`
	MaxLifeTime		int			`yaml:"max_life_time"`
	MaxOpenConns	int			`yaml:"max_open_conns"`
	MaxIdleConns	int			`yaml:"max_idle_conns"`
}

// DSN connect Postgres
func (a Postgres) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable", a.Username, a.Password, a.Host, a.Port, a.DatabaseName)
}
