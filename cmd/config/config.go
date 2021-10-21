package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config has all configuration for the project
type Config struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	DatabaseConfig DatabaseConfig `mapstructure:"database"`
	LoggerConfig   LoggerConfig   `mapstructure:"logger"`
	RedisConfig    RedisConfig    `mapstructure:"redis"`
	AuthConfig     AuthConfig     `mapstructure:"auth"`
	MailConfig     MailConfig     `mapstructure:"mail"`
	SessionConfig  SessionConfig  `mapstructure:"session"`
	CORSConfig     CORSConfig     `mapstructure:"cors"`
	ShopifyConfig  ShopifyConfig  `mapstructure:"shopify"`
}

// ServerConfig has only server specific configuration
type ServerConfig struct {
	Env          string        `mapstructure:"env"`
	Port         string        `mapstructure:"port"`
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	CloseTimeout time.Duration `mapstructure:"closeTimeout"`
	CSRFProtect  string        `mapstructure:"csrf_protect"`
	StaticDir    string        `mapstructure:"static_dir"`
}

// ShopifyConfig contains shopify related configurations
type ShopifyConfig struct {
	APIKey       string `mapstructure:"apiKey"`
	APISecretKey string `mapstructure:"apiSecretKey"`
}

// LoggerConfig has logger related configuration.
type LoggerConfig struct {
	LogFilePath string `mapstructure:"file"`
}

// RedisConfig has cache related configuration.
type RedisConfig struct {
	Host             string `mapstructure:"host"`
	Port             string `mapstructure:"port"`
	Password         string `mapstructure:"password"`
	ConnectionString string `mapstructure:"connectionString"`
}

// AuthConfig has authentication related configuration
type AuthConfig struct {
	PasswordHashSecretKey string `mapstructure:"password_hash_secret_key"`
	HmacSecret            string `mapstructure:"hmacSecret"`
	SessionSecretKey      string `mapstructure:"session_secret_key"`
	EmailVerificationKey  string `mapstructure:"email_verification_key"`
	PasswordResetKey      string `mapstructure:"password_reset_key"`
}

// MailConfig has authentication related configuration
type MailConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	TLS      bool   `mapstructure:"tls"`
	SSL      bool   `mapstructure:"ssl"`
}

// SessionConfig has session related configuration
type SessionConfig struct {
	CookieName string `mapstructure:"cookie_name"`
	MaxAge     int    `mapstructure:"max_age"`
	Domain     string `mapstructure:"domain"`
}

// CORSConfig contains CORS related configurations
type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
}

// DatabaseConfig contains mongodb related configuration
type DatabaseConfig struct {
	Scheme     string `mapstructure:"scheme"`
	Host       string `mapstructure:"host"`
	DBName     string `mapstructure:"name"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	ReplicaSet string `mapstructure:"replicaSet"`
}

// ConnectionURL returns connection string to of mongodb storage
func (d *DatabaseConfig) ConnectionURL() string {
	url := fmt.Sprintf("%s://", d.Scheme)
	if d.Username != "" && d.Password != "" {
		url += fmt.Sprintf("%s:%s@", d.Username, d.Password)
	}
	url += d.Host
	if d.ReplicaSet != "" {
		url += fmt.Sprintf("/?replicaSet=%s", d.ReplicaSet)
	}
	return url
}

// MiddlewareConfig has middlewares related configuration
type MiddlewareConfig struct {
	EnableRequestLog bool `mapstructure:"enableRequestLog"`
}

// GetConfig returns entire project configuration
func GetConfig() *Config {
	return GetConfigFromFile("default")
}

// GetConfigFromFile returns configuration from specific file object
func GetConfigFromFile(fileName string) *Config {
	if fileName == "" {
		fileName = "default"
	}

	// looking for filename `default` inside `src/server` dir with `.toml` extension
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath("../../conf/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't load config: %s", err)
		os.Exit(1)
	}
	config := &Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	return config
}
