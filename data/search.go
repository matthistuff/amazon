package data

type ItemSearchResponse struct {
	Items SearchItems
}

type SearchItems struct {
	Request      Request
	TotalResults int
	TotalPages   int
	ItemList     []Item `xml:"Item"`
}

type ItemSearchRequest struct {
	ItemPage          int
	Keywords          string
	ResponseGroupList []string `xml:"ResponseGroup"`
	SearchIndex       string
}