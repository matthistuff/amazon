package data

type ItemLookupResponse struct {
	Items LookupItems
}

type LookupItems struct {
	Request  LookupRequest
	ItemList []Item `xml:"Item"`
}

type LookupRequest struct {
	IsValid bool
}
