package data

// BrowseNodes response group
type BrowseNodes struct {
	BrowserNodeList []BrowseNode `xml:"BrowseNode"`
}

// A BrowseNode is an entry in BrowseNodes
type BrowseNode struct {
	BrowseNodeID   int `xml:"BrowseNodeId"`
	Name           string
	Children       []BrowseNode
	Ancestors      []BrowseNode
	IsCategoryRoot bool
}
