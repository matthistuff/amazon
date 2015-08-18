package data

type ItemAttributes struct {
	Binding   string
	Brand     string
	EAN       string
	Languages Languages
	ListPrice Price
	Title     string
	UPC       string
}

type Languages struct {
	LanguageList []Language `xml:"Language"`
}

type Language struct {
	Name string
	Type string
}