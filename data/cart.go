package data

type CartGetResponse struct {
	Cart Cart
}

type CartCreateResponse struct {
	Cart Cart
}

type CartAddResponse struct {
	Cart Cart
}

type Cart struct {
	Request        Request
	CartId         string
	HMAC           string
	URLEncodedHMAC string
	PurchaseURL    string
	SubTotal       Price
	CartItems      CartItems
}

type CartItems struct {
	SubTotal     Price
	CartItemList []CartItem `xml:"CartItem"`
}

type CartItem struct {
	CartItemId   string
	ASIN         string
	Quantity     int
	Title        string
	ProductGroup string
	Price        Price
	ItemTotal    Price
}