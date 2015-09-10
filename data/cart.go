package data

// CartResponse returned from a cart request
type CartResponse struct {
	Cart Cart
}

// The Cart of the CartResponse
type Cart struct {
	Request        Request
	CartID         string `xml:"CartId"`
	HMAC           string
	URLEncodedHMAC string
	PurchaseURL    string
	SubTotal       Price
	CartItems      CartItems
}

// CartItems of a Cart
type CartItems struct {
	SubTotal     Price
	CartItemList []CartItem `xml:"CartItem"`
}

// CartItem in CartItems
type CartItem struct {
	CartItemID   string `xml:"CartItemId"`
	ASIN         string
	Quantity     int
	Title        string
	ProductGroup string
	Price        Price
	ItemTotal    Price
}
