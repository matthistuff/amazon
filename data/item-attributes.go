package data

type ItemAttributes struct {
	Authors         []string `xml:"Author"`
	Binding         string
	Brand           string
	EAN             string
	Edition         string
	Features        []string `xml:"Features"`
	ISBN            string
	Label           string
	Languages       Languages
	ListPrice       Price
	Manufacturer    string
	ProductGroup    string
	PublicationDate string
	Publisher       string
	ReleaseDate     string
	Studio          string
	Title           string
	UPC             string
}

type Languages struct {
	LanguageList []Language `xml:"Language"`
}

type Language struct {
	Name string
	Type string
}