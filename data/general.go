package data

// ItemResponse of an item request
type ItemResponse struct {
	Items ItemResponseBody
}

// ItemResponseBody of a ItemResponse
type ItemResponseBody struct {
	Request  Request
	ItemList []Item `xml:"Item"`
}

// Request information
type Request struct {
	IsValid bool
	Errors  Errors
}

// Errors of a request
type Errors struct {
	ErrorList []Error `xml:"Error"`
}

// Error of Errors
type Error struct {
	Code    string
	Message string
}
