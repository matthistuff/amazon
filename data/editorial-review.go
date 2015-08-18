package data

type EditorialReviews struct {
	EditorialReviewList []EditorialReview `xml:"EditorialReview"`
}

type EditorialReview struct {
	Source           string
	Content          string
	IsLinkSuppressed bool
}
