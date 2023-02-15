package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Api struct {
	Port           int      `json:"port"             mapstructure:"APP_PORT"`
	Env            string   `json:"env"              mapstructure:"ENV"`
	JWTSecretKey   string   `json:"-"                mapstructure:"JWT_SECRET_KEY"`
	JWTExpiredTime int      `json:"jwt_expired_time" mapstructure:"JWT_EXPIRED"`
	Database       Database `json:"database"`
	Shipper        Shipper  `json:"shipper"`
}

type Database struct {
	Host     string `json:"host"     mapstructure:"DATABASE_HOST"`
	Port     int    `json:"port"     mapstructure:"DATABASE_PORT"`
	Username string `json:"username" mapstructure:"DATABASE_USERNAME"`
	Password string `json:"password" mapstructure:"DATABASE_PASSWORD"`
	Schema   string `json:"schema"   mapstructure:"DATABASE_SCHEMA"`
	Loc      string `json:"loc"      mapstructure:"DATABASE_LOC"`
}

type Shipper struct {
	BaseUrl string `json:"base_url" mapstructure:"SHIPPER_BASEURL"`
	Key     string `json:"key"      mapstructure:"SHIPPER_API_KEY"`
}

func LoadConfigAPI(path string) *Api {
	if path := strings.TrimSpace(path); path == "" {
		path = "."
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("read config failed:", err.Error())
	}

	viper.SetDefault("ENV", "development")
	viper.SetDefault("APP_PORT", 4001)

	api := &Api{}

	viper.Unmarshal(api)
	viper.Unmarshal(&api.Database)
	viper.Unmarshal(&api.Shipper)

	return api
}
