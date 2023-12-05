package wptype

type Author struct {
	Node struct {
		ID        string `json:"id"`
		Slug      string `json:"slug"`
		Name      string `json:"name"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Avatar    struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"avatar"`
	} `json:"node"`
}
