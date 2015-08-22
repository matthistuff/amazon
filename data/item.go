package data

type Price struct {
	Amount         int
	CurrencyCode   string
	FormattedPrice string
}

type Item struct {
	ASIN string

	BrowseNodes      BrowseNodes
	DetailPageURL    string
	EditorialReviews EditorialReviews
	ItemAttributes   ItemAttributes
	OfferSummary     OfferSummary
	Offers           Offers
}
