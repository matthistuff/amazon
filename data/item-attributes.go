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

func (l Languages) Languages() []string {
	langs := make([]string, 0)
	unique := map[string]bool{}

	for _, lang := range l.LanguageList {
		if exists, _ := unique[lang.Name]; !exists {
			langs = append(langs, lang.Name)
			unique[lang.Name] = true
		}

	}

	return langs
}

type Language struct {
	Name string
	Type string
}
