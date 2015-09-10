package data

// Price generic
type Price struct {
	Amount         int
	CurrencyCode   string
	FormattedPrice string
}

// Item information
type Item struct {
	ASIN string

	BrowseNodes      BrowseNodes
	DetailPageURL    string
	EditorialReviews EditorialReviews
	ItemAttributes   ItemAttributes
	OfferSummary     OfferSummary
	Offers           Offers
}
