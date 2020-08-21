// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"flag"

	env "github.com/Netflix/go-env"
	"github.com/jinzhu/gorm"
)

//DB - Database connection
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	SSL      string
}

//SystemConfig represents system service configuration
type SystemConfig struct {
	ProdDBHost       string `env:"ENV_PROD_DB_HOST"`
	ProdDBPort       string `env:"ENV_PROD_DB_PORT"`
	ProdDBUser       string `env:"ENV_PROD_DB_USER"`
	ProdDBPass       string `env:"ENV_PROD_DB_PASS"`
	ProdDBDatabase   string `env:"ENV_PROD_DB_DATABASE"`
	ProdDBSSL        string `env:"ENV_PROD_DB_SSL"`
	DevDBHost        string `env:"ENV_DEV_DB_HOST"`
	DevDBPort        string `env:"ENV_DEV_DB_PORT"`
	DevDBUser        string `env:"ENV_DEV_DB_USER"`
	DevDBPass        string `env:"ENV_DEV_DB_PASS"`
	DevDBDatabase    string `env:"ENV_DEV_DB_DATABASE"`
	DevDBSSL         string `env:"ENV_DEV_DB_SSL"`
	TwilioAccountSid string `env:"ENV_TWILIO_ACCOUNT_SID"`
	TwilioAuthToken  string `env:"ENV_TWILIO_AUTH_TOKEN"`
	SenderPhone      string `env:"ENV_SENDER_PHONE"`
	OneSignalAppKey  string `env:"ENV_ONE_SIGNAL_APP_KEY"`
	AppID            string `env:"ENV_ONE_SIGNAL_APP_ID"`
	SendgridAPIKey   string `env:"ENV_SENDGRID_API_KEY"`
	BaseURL          string `env:"ENV_BASE_URL"`
	SenderEmail      string `env:"ENV_SENDER_EMAIL"`
	AWSS3Bucket      string `env:"ENV_AWS_S3_BUCKET"`
	AWSRegion        string `env:"ENV_AWS_REGION"`
	AWSAccessKeyID   string `env:"ENV_AWS_ACCESS_KEY_ID"`
	AWSSecretKey     string `env:"ENV_AWS_SECRET_KEY"`
	AWSSessionToken  string `env:"ENV_AWS_SESSION_TOKEN"`
}

//GetConnectionContext - returns database connection string based on environment
func GetConnectionContext() string {
	dbContext := flag.Bool("isDev", false, "a bool")
	if *dbContext {
		return DevDbURL(BuildProdDBConfig())
	}
	return ProdDbURL(BuildDevDBConfig())
}

//InitConfig - initial the configuration struct with environment variables
func InitConfig(cfg interface{}) error {
	_, err := env.UnmarshalFromEnviron(cfg)
	if err != nil {
		return err
	}
	return nil
}
