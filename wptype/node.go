package wptype

type Node struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Excerpt    string `json:"excerpt"`
	Content    string `json:"content"`
	URI        string `json:"uri"`
	Slug       string `json:"slug"`
	Author     Author `json:"author"`
	Date       string `json:"date"`
	Modified   string `json:"modified"`
	Link       string `json:"link"`
	Categories struct {
		Nodes []Category `json:"nodes"`
	} `json:"categories"`
	FeaturedImage struct {
		Node struct {
			MediaItemURL string `json:"mediaItemUrl"`
			Caption      string `json:"caption,omitempty"`
			Sizes        string `json:"sizes"`
			MediaDetails struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"mediaDetails"`
		} `json:"node"`
	} `json:"featuredImage"`
	// this is example how to add custom field in your graphql response post using Advanced Custom Fields
	// PageCusom struct {
	// 	Excerpt string `json:"excerpt,omitempty"`
	// } `json:"pageCusom,omitempty"`
}
