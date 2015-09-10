package data

// OfferSummary response group
type OfferSummary struct {
	LowestNewPrice         Price
	LowestUsedPrice        Price
	LowestCollectiblePrice Price
	LowestRefurbishedPrice Price
	TotalNew               int
	TotalUsed              int
	TotalCollectible       int
	TotalRefurbished       int
}

// Offers response group
type Offers struct {
	TotalOffers     int
	TotalOfferPages int
	OfferList       []Offer `xml:"Offer"`
}

// A Offer of an OfferList
type Offer struct {
	Merchant        Merchant
	OfferAttributes OfferAttributes
	OfferListing    OfferListing
}

// Merchant of the product
type Merchant struct {
	Name string
}

// OfferAttributes of an Offer
type OfferAttributes struct {
	Condition string
}

// The OfferListing holds information about an Offer
type OfferListing struct {
	OfferListingID                  string `xml:"OfferListingId"`
	Price                           Price
	AmountSaved                     Price
	PercentageSaved                 int
	Availability                    string
	AvailabilityAttributes          AvailabilityAttributes
	IsEligibleForSuperSaverShipping bool
}

// AvailabilityAttributes holds information about an Offer availability
type AvailabilityAttributes struct {
	AvailabilityType string
	MinimumHours     string
	MaximumHours     string
}
