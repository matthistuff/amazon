package data

type ItemLookupResponse struct {
	Items LookupItems
}

type LookupItems struct {
	Request  Request
	ItemList []Item `xml:"Item"`
}
