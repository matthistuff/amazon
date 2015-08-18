package config

import (
	"github.com/mitchellh/go-homedir"
	"path"
	"os"
	"github.com/BurntSushi/toml"
	"github.com/matthistuff/amazon/helper"
	"time"
)

type Config struct {
	Version     string
	Locale      string
	ResultCache map[string]map[string]string
	Carts       map[string]*Cart
}

type Cart struct {
	Name    string
	Created time.Time
	CartId  string
	HMAC    string
}

func (c *Config) Flush() {
	file, _ := os.Create(confPath)
	encoder := toml.NewEncoder(file)

	if err := encoder.Encode(c); err != nil {
		panic(err)
	}
}

func (c Config) FromCache(cache string, key string) string {
	return c.ResultCache[cache][key]
}

func (c Config) ASINFromCache(cache string, item string) string {
	if helper.IsASIN(item) {
		return item
	}

	return c.FromCache(cache, item)
}

func (c Config) NumericFromCache(cache string, index string) string {
	if helper.IsNumeric(index) {
		return c.FromCache(cache, index)
	}

	return index
}

var (
	confPath string
	conf *Config
)

func LoadConfig() error {
	err := ensureConfig()

	if err != nil {
		return err
	}

	_, err = toml.DecodeFile(confPath, &conf)

	if err != nil {
		return err
	}

	return nil
}

func ensureConfig() error {
	home, err := homedir.Dir()

	if err == nil {
		confPath = path.Join(home, ".amazon-cli")

		if _, err := os.Stat(confPath); os.IsNotExist(err) {
			conf = &Config{
				Version: "0.0.1",
				Locale: "US",
				ResultCache: make(map[string]map[string]string),
				Carts: make(map[string]*Cart),
			}

			conf.Flush()
		}
	}

	return err
}

func GetConfig() *Config {
	return conf
}