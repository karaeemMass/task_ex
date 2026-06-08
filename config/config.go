package config

import (
	"os"
	"strconv"

	"task_ex/config/crypto"

	"github.com/spf13/viper"
)

// Config represents application configuration. It contains the nested
// Database section and a top-level Enc flag that indicates whether the
// database fields are stored encrypted in the config file.
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Enc      bool           `mapstructure:"Enc"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// If the config file indicates values are encrypted, decrypt the values
	// using the CONFIG_KEY environment variable. We read the raw strings
	// from viper so we can decrypt even when the original YAML may contain
	// encrypted strings instead of plain typed values.
	if cfg.Enc {
		key := os.Getenv("CONFIG_KEY")
		if key == "" {
			return nil, os.ErrInvalid
		}

		// Decrypt each value; if any decryption fails we return the error.
		hostEnc := viper.GetString("database.host")
		portEnc := viper.GetString("database.port")
		userEnc := viper.GetString("database.user")
		passEnc := viper.GetString("database.password")
		dbNameEnc := viper.GetString("database.dbname")

		if hostEnc != "" {
			host, err := crypto.Decrypt(hostEnc, key)
			if err != nil {
				return nil, err
			}
			cfg.Database.Host = host
		}

		if portEnc != "" {
			portStr, err := crypto.Decrypt(portEnc, key)
			if err != nil {
				return nil, err
			}
			p, err := strconv.Atoi(portStr)
			if err != nil {
				return nil, err
			}
			cfg.Database.Port = p
		}

		if userEnc != "" {
			user, err := crypto.Decrypt(userEnc, key)
			if err != nil {
				return nil, err
			}
			cfg.Database.User = user
		}

		if passEnc != "" {
			pass, err := crypto.Decrypt(passEnc, key)
			if err != nil {
				return nil, err
			}
			cfg.Database.Password = pass
		}

		if dbNameEnc != "" {
			dbName, err := crypto.Decrypt(dbNameEnc, key)
			if err != nil {
				return nil, err
			}
			cfg.Database.DBName = dbName
		}
	}

	return &cfg, nil
}
