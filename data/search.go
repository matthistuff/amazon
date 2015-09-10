package data

// ItemSearchResponse of a search request
type ItemSearchResponse struct {
	Items ItemSearchBody
}

// ItemSearchBody of the ItemSearchResponse
type ItemSearchBody struct {
	Request      Request
	TotalResults int
	TotalPages   int
	ItemList     []Item `xml:"Item"`
}
