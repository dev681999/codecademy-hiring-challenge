package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// A Source is a way to load config
type Source interface {
	apply(interface{}) error
}

type funcOption struct {
	f func(interface{}) error
}

func (fdo *funcOption) apply(do interface{}) error {
	return fdo.f(do)
}

func newFuncServerOption(f func(interface{}) error) *funcOption {
	return &funcOption{
		f: f,
	}
}

// FromENV reads config from env
func FromENV(prefix string) Source {
	return newFuncServerOption(func(i interface{}) error {
		return envconfig.Process("", i)
	})
}

// FromFile reads config from file
func FromFile(filename string) Source {
	return newFuncServerOption(func(i interface{}) error {
		if len(filename) == 0 {
			return fmt.Errorf("invalid config file %s", filename)
		}

		extension := filepath.Ext(filename)
		if extension == "" || (extension != ".yml" && extension != ".yaml") {
			return fmt.Errorf("invalid file extension for file %s extension %s", filename, extension)
		}

		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("file error: %v", err)
		}

		err = yaml.NewDecoder(file).Decode(i)
		if err != nil {
			return fmt.Errorf("yaml decoder error : %v", err)
		}

		return nil
	})
}

// New reads the config with given options into the `value`
func New(value *Config, options ...Source) error {
	if value == nil {
		return nil
	}
	for _, option := range options {
		err := option.apply(value)
		if err != nil {
			return err
		}
	}
	return nil
}

type Server struct {
	Host                string `yaml:"host"`
	Port                string `yaml:"port"`
	Debug               bool   `yaml:"debug"`
	PublicStorageFolder string `yaml:"public_storage_folder"`
	SwaggerUIFolder     string `yaml:"swagger_ui_folder"`
	SwaggerAPIFilePath  string `yaml:"swagger_api_file_path"`
	JWTSecret           string `yaml:"jwt_secret"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"db_name"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode"`
	Debug    bool   `yaml:"debug"`
}

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}
