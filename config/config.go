package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/matthistuff/amazon/helper"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
	"strings"
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

func (c Config) FromCache(cache string, key string) (string, bool) {
	cached, exists := c.ResultCache[cache]

	if !exists {
		return "", exists
	}

	cachedValue, exists := cached[key]

	return cachedValue, exists
}

func (c Config) CartNameFromCache(indexOrName string) string {
	cached, exists := c.FromCache("Carts", indexOrName)

	if !exists {
		if indexOrName == "" {
			return "default"
		}

		return indexOrName
	}

	return cached
}

func (c Config) ASINFromCache(cache string, item string) (string, bool) {
	if helper.IsASIN(item) {
		return item, true
	}

	return c.FromCache(cache, item)
}

func (c Config) CartItemIdFromCache(cartName string, item string) (string, bool) {
	if helper.IsCartItemId(item) {
		return item, true
	}

	return c.FromCache(fmt.Sprintf("Cart%sItems", strings.Title(cartName)), item)
}

func (c Config) NumericFromCache(cache string, index string) string {
	if helper.IsNumeric(index) {
		value, exists := c.FromCache(cache, index)

		if exists {
			return value
		}
	}

	return index
}

var (
	confPath string
	conf     *Config
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
				Version:     "0.1.1",
				Locale:      "US",
				ResultCache: make(map[string]map[string]string),
				Carts:       make(map[string]*Cart),
			}

			conf.Flush()
		}
	}

	return err
}

func GetConfig() *Config {
	return conf
}
