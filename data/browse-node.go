package data

type BrowseNodes struct {
	BrowserNodeList []BrowseNode `xml:"BrowseNode"`
}

type BrowseNode struct {
	BrowseNodeId   int
	Name           string
	Children       []BrowseNode
	Ancestors      []BrowseNode
	IsCategoryRoot bool
}
