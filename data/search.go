package data

type ItemSearchResponse struct {
	Items ItemSearchBody
}

type ItemSearchBody struct {
	Request      Request
	TotalResults int
	TotalPages   int
	ItemList     []Item `xml:"Item"`
}
