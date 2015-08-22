package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var ratingReg, _ = regexp.Compile("([0-9.]+) [^0-9]+ 5")

var amazons = map[string]string{
	"BR": "amazon.com.br",
	"CA": "amazon.ca",
	"CN": "amazon.cn",
	"DE": "amazon.de",
	"ES": "amazon.es",
	"FR": "amazon.fr",
	"IN": "amazon.in",
	"IT": "amazon.it",
	"JP": "amazon.jp",
	"UK": "amazon.uk",
	"US": "amazon.com",
}

func Rating(asin string, locale string) (float64, error) {
	url := fmt.Sprintf(
		"https://www.%s/gp/customer-reviews/widgets/average-customer-review/popover/ref=dpx_acr_pop_?contextId=dpx&asin=%s",
		amazons[locale],
		asin)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	regMatch := ratingReg.FindStringSubmatch(string(body[:]))
	if len(regMatch) < 2 {
		return 0, nil
	}

	i, err := strconv.ParseFloat(regMatch[1], 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func FormatRating(rating float64) string {
	base := int(rating)
	fraction := rating - float64(base)

	formatted := strings.Repeat("★", base)

	if fraction >= 0.5 {
		formatted = formatted + "½"
	}

	return formatted
}
