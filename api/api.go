package api

import (
	"github.com/DDRBoxman/go-amazon-product-api"
	"os"
	"github.com/matthistuff/amazon/config"
	"strings"
	"github.com/matthistuff/amazon/data"
	"encoding/xml"
)

var hosts = map[string]string{
	"BR": "webservices.amazon.com.br",
	"CA": "webservices.amazon.ca",
	"CN": "webservices.amazon.cn",
	"DE": "webservices.amazon.de",
	"ES": "webservices.amazon.es",
	"FR": "webservices.amazon.fr",
	"IN": "webservices.amazon.in",
	"IT": "webservices.amazon.it",
	"JP": "webservices.amazon.jp",
	"UK": "webservices.amazon.uk",
	"US": "webservices.amazon.com",
}

type API struct {
	ProductAPI *amazonproduct.AmazonProductAPI
}

func (a API) ItemLookup(ASIN string) (data.ItemLookupResponse, error) {
	var lookupResult data.ItemLookupResponse

	response, err := a.ProductAPI.ItemLookup(ASIN)
	if err != nil {
		return lookupResult, err
	}

	if err := xml.Unmarshal([]byte(response), &lookupResult); err != nil {
		return lookupResult, err
	}

	return lookupResult, nil
}

func (a API) ItemSearch(SearchIndex string, Parameters map[string]string) (data.ItemSearchResponse, error) {
	var searchResult data.ItemSearchResponse

	response, err := a.ProductAPI.ItemSearch(SearchIndex, Parameters)

	if err != nil {
		return searchResult, err
	}

	if err := xml.Unmarshal([]byte(response), &searchResult); err != nil {
		return searchResult, err
	}

	return searchResult, nil
}

func (a API) CartGet(CartId, HMAC string) (data.CartGetResponse, error) {
	var cartGetResult data.CartGetResponse

	response, err := a.ProductAPI.CartGet(CartId, HMAC)

	if err != nil {
		return cartGetResult, err
	}

	if err := xml.Unmarshal([]byte(response), &cartGetResult); err != nil {
		return cartGetResult, err
	}

	return cartGetResult, nil
}

func (a API) CartCreate(Items map[string]int) (data.CartCreateResponse, error) {
	var cartCreateResult data.CartCreateResponse

	response, err := a.ProductAPI.CartCreate(Items)

	if err != nil {
		return cartCreateResult, err
	}

	if err := xml.Unmarshal([]byte(response), &cartCreateResult); err != nil {
		return cartCreateResult, err
	}

	return cartCreateResult, nil
}

func (a API) CartAdd(CartId, HMAC string, Items map[string]int) (data.CartAddResponse, error) {
	var cartAddResult data.CartAddResponse

	response, err := a.ProductAPI.CartAdd(CartId, HMAC, Items)

	if err != nil {
		return cartAddResult, err
	}

	if err := xml.Unmarshal([]byte(response), &cartAddResult); err != nil {
		return cartAddResult, err
	}

	return cartAddResult, nil
}

func Create(locale string) API {
	conf := config.GetConfig()

	if locale == "" {
		locale = conf.Locale
	}

	return API{
		ProductAPI: &amazonproduct.AmazonProductAPI{
			AccessKey: os.Getenv("AMAZON_ACCESS_KEY"),
			SecretKey: os.Getenv("AMAZON_SECRET_KEY"),
			Host: hosts[strings.ToUpper(locale)],
			AssociateTag: "matthi-20",
		},
	}
}

func GetLocales() []string {
	locales := make([]string, len(hosts))

	i := 0
	for locale := range hosts {
		locales[i] = locale
		i += 1
	}

	return locales
}