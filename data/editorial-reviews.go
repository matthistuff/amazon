package data

// EditorialReviews response group
type EditorialReviews struct {
	EditorialReviewList []EditorialReview `xml:"EditorialReview"`
}

// A EditorialReview in EditorialReviews
type EditorialReview struct {
	Source           string
	Content          string
	IsLinkSuppressed bool
}
