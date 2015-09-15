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

// Config is the top level config structure
type Config struct {
	Version     string
	Locale      string
	ResultCache map[string]map[string]string
	Carts       map[string]*Cart
}

// Cart represents an active vart in the config
type Cart struct {
	Name    string
	Created time.Time
	CartID  string
	HMAC    string
}

// Flush writes the config to disk
func (c *Config) Flush() {
	file, _ := os.Create(confPath)
	encoder := toml.NewEncoder(file)

	if err := encoder.Encode(c); err != nil {
		panic(err)
	}
}

// FromCache returns a key from a named cache
func (c Config) FromCache(cache string, key string) (string, bool) {
	cached, exists := c.ResultCache[cache]

	if !exists {
		return "", exists
	}

	cachedValue, exists := cached[key]

	return cachedValue, exists
}

// CartNameFromCache returns a cart name from the cart cache
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

// ASINFromCache returns a asin from a named cache
func (c Config) ASINFromCache(cache string, item string) (string, bool) {
	if helper.IsASIN(item) {
		return item, true
	}

	return c.FromCache(cache, item)
}

// CartItemIDFromCache returns a cart item ID from a named cart
func (c Config) CartItemIDFromCache(cartName string, item string) (string, bool) {
	if helper.IsCartItemID(item) {
		return item, true
	}

	return c.FromCache(fmt.Sprintf("Cart%sItems", strings.Title(cartName)), item)
}

// NumericFromCache returns a numeric value from a named cache
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

// LoadConfig loads the config from the disk
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
				Version:     "0.1.3",
				Locale:      "US",
				ResultCache: make(map[string]map[string]string),
				Carts:       make(map[string]*Cart),
			}

			conf.Flush()
		}
	}

	return err
}

// GetConfig returns the current config
func GetConfig() *Config {
	return conf
}
