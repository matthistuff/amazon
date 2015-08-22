package data

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

type Offers struct {
	TotalOffers     int
	TotalOfferPages int
	OfferList       []Offer `xml:"Offer"`
}

type Offer struct {
	Merchant        Merchant
	OfferAttributes OfferAttributes
	OfferListing    OfferListing
}

type Merchant struct {
	Name string
}

type OfferAttributes struct {
	Condition string
}

type OfferListing struct {
	OfferListingId                  string
	Price                           Price
	AmountSaved                     Price
	PercentageSaved                 int
	Availability                    string
	AvailabilityAttributes          AvailabilityAttributes
	IsEligibleForSuperSaverShipping bool
}

type AvailabilityAttributes struct {
	AvailabilityType string
	MinimumHours     string
	MaximumHours     string
}
