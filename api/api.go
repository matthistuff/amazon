package api

import (
	"github.com/matthistuff/go-amazon-product-api"
	"os"
	"github.com/matthistuff/amazon/config"
	"strings"
	"github.com/matthistuff/amazon/data"
	"encoding/xml"
	"fmt"
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

func (a API) checkSanity(request data.Request) {
	if len(request.Errors.ErrorList) > 0 {
		for _, err := range request.Errors.ErrorList {
			fmt.Fprintf(os.Stderr, "%s\n", err.Message)
		}

		os.Exit(1)
	}
}

func (a API) ItemLookup(ASIN string) (data.ItemResponse, error) {
	var lookupResult data.ItemResponse

	response, err := a.ProductAPI.ItemLookup(ASIN)
	if err != nil {
		return lookupResult, err
	}

	if err := xml.Unmarshal([]byte(response), &lookupResult); err != nil {
		return lookupResult, err
	}

	a.checkSanity(lookupResult.Items.Request)

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

	a.checkSanity(searchResult.Items.Request)

	return searchResult, nil
}

func (a API) CartGet(CartId, HMAC string) (data.CartResponse, error) {
	var cartGetResult data.CartResponse

	response, err := a.ProductAPI.CartGet(CartId, HMAC)

	if err != nil {
		return cartGetResult, err
	}

	if err := xml.Unmarshal([]byte(response), &cartGetResult); err != nil {
		return cartGetResult, err
	}

	a.checkSanity(cartGetResult.Cart.Request)

	return cartGetResult, nil
}

func (a API) CartCreate(Items map[string]int) (data.CartResponse, error) {
	var cartCreateResult data.CartResponse

	response, err := a.ProductAPI.CartCreate(Items)

	if err != nil {
		return cartCreateResult, err
	}

	if err := xml.Unmarshal([]byte(response), &cartCreateResult); err != nil {
		return cartCreateResult, err
	}

	a.checkSanity(cartCreateResult.Cart.Request)

	return cartCreateResult, nil
}

func (a API) CartAdd(CartId, HMAC string, Items map[string]int) (data.CartResponse, error) {
	var cartAddResult data.CartResponse

	response, err := a.ProductAPI.CartAdd(Items, CartId, HMAC)

	if err != nil {
		return cartAddResult, err
	}

	if err := xml.Unmarshal([]byte(response), &cartAddResult); err != nil {
		return cartAddResult, err
	}

	a.checkSanity(cartAddResult.Cart.Request)

	return cartAddResult, nil
}

func (a API) CartModify(CartId, HMAC string, CartItems map[string]int) (data.CartResponse, error) {
	var cartModifyResult data.CartResponse

	response, err := a.ProductAPI.CartModify(CartItems, CartId, HMAC)

	if err != nil {
		return cartModifyResult, err
	}

	if err := xml.Unmarshal([]byte(response), &cartModifyResult); err != nil {
		return cartModifyResult, err
	}

	a.checkSanity(cartModifyResult.Cart.Request)

	return cartModifyResult, nil
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