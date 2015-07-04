package data

type Price struct {
	Amount         int
	CurrencyCode   string
	FormattedPrice string
}

type Errors struct {
	ErrorList []Error `xml:"Error"`
}

type Error struct {
	Code    string
	Message string
}

type Request struct {
	IsValid bool
	Errors  Errors
}