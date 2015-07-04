package data

type Item struct {
	ASIN             string
	
	BrowseNodes      BrowseNodes
	DetailPageURL    string
	EditorialReviews EditorialReviews
	ItemAttributes   ItemAttributes
	OfferSummary     OfferSummary
	ParentASIN       string
}