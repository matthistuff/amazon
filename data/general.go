package data

type ItemResponse struct {
	Items ItemResponseBody
}

type ItemResponseBody struct {
	Request  Request
	ItemList []Item `xml:"Item"`
}

type Request struct {
	IsValid bool
	Errors  Errors
}

type Errors struct {
	ErrorList []Error `xml:"Error"`
}

type Error struct {
	Code    string
	Message string
}
