package data

type ItemSearchResponse struct {
	Items SearchItems
}

type SearchItems struct {
	Request      SearchRequest
	TotalResults int
	TotalPages   int
	ItemList     []Item `xml:"Item"`
}

type SearchRequest struct {
	IsValid           bool
	ItemSearchRequest ItemSearchRequest
}

type ItemSearchRequest struct {
	ItemPage          int
	Keywords          string
	ResponseGroupList []string `xml:"ResponseGroup"`
	SearchIndex       string
}