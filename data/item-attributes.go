package data

// ItemAttributes response group
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

// Languages for the item
type Languages struct {
	LanguageList []Language `xml:"Language"`
}

// Languages returns item languages as an easy to work with map
func (l Languages) Languages() []string {
	var langs []string
	unique := map[string]bool{}

	for _, lang := range l.LanguageList {
		if exists, _ := unique[lang.Name]; !exists {
			langs = append(langs, lang.Name)
			unique[lang.Name] = true
		}

	}

	return langs
}

// Language of an item
type Language struct {
	Name string
	Type string
}
