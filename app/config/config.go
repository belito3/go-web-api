package config

import "C"
import (
	"fmt"
	"github.com/belito3/go-api-codebase/pkg/util"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Configuration ... The configuration of system
func LoadConfig(filePath string) (config AppConfiguration, err error){
	viper.AddConfigPath(filePath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// It will check for an environment variable with a name matching
	//the key uppercased and prefixed with the EnvPrefix if set
	// TODO: Auto overwrite config from env variable (if existed) corresponding with yaml config default
	// example yaml: http.port -> ENV: HTTP_PORT, http.max_content_length -> HTTP_MAX_CONTENT_LENGTH
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Println("port: ", viper.GetString("HTTP_MAX_CONTENT_LENGTH"))

	return
}

// Print With JSON
func PrintWithJSON(config AppConfiguration) {
	if config.PrintConfig {
		b, err := util.JSONMarshalIndent(config, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

// Configuration ... The configuration of system
type AppConfiguration struct {
	HTTP					HTTP			`mapstructure:"http"`
	PrintConfig  			bool			`mapstructure:"print_config"`
	RunMode					string			`mapstructure:"run_mode"`
	Log						Log				`mapstructure:"log"`
	UniqueID				UniqueID		`mapstructure:"unique_id"`
	JWTSecretKey            string 			`mapstructure:"jwt_secret_key"`
	ARateLimiter			ARateLimiter	`mapstructure:"app_rate_limiter"`
	CRateLimiter			CRateLimiter	`mapstructure:"client_rate_limiter"`
	Redis					Redis			`mapstructure:"redis"`
	DBSQL					DBSQL			`mapstructure:"dbsql"`
}


// HTTP http
type HTTP struct {
	Host				string		`mapstructure:"host"`
	Port				int			`mapstructure:"port"`
	ShutdownTimeout		int			`mapstructure:"shutdown_timeout"`
	MaxContentLength	int64		`mapstructure:"max_content_length"`
}


// Log
type Log struct {
	Level		int			`mapstructure:"level"`
	Format 		string		`mapstructure:"format"`
}

// UniqueID
type UniqueID struct {
	Type		string		`mapstructure:"type"`
	Snowflake	struct {
		Node	int64		`mapstructure:"node"`
		Epoch	int64		`mapstructure:"epoch"`
	}
}

// App RateLimiter
type ARateLimiter struct {
	Enable		bool	`mapstructure:"enable"`
	Count		int		`mapstructure:"count"`
}

// Client RateLimiter Request frequency limit configuration parameters
type CRateLimiter struct {
	Enable		bool	`mapstructure:"enable"`
	Count		int		`mapstructure:"count"`
	RedisDB 	int		`mapstructure:"redis_db"`
}

// Redis redis Configuration parameter
type Redis struct {
	Addr		string		`mapstructure:"addr"`
	Password	string		`mapstructure:"password"`
}

// DBSQL  config
type DBSQL struct {
	DriverName		string		`mapstructure:"driver_name"`
	Username		string		`mapstructure:"username"`
	Password		string		`mapstructure:"password"`
	Host			string		`mapstructure:"host"`
	Port			int			`mapstructure:"port"`
	DatabaseName	string		`mapstructure:"database_name"`
	MaxLifeTime		int			`mapstructure:"max_life_time"`
	MaxOpenConns	int			`mapstructure:"max_open_conns"`
	MaxIdleConns	int			`mapstructure:"max_idle_conns"`
}

// DSN connect
func (a DBSQL) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable", a.Username, a.Password, a.Host, a.Port, a.DatabaseName)
}
