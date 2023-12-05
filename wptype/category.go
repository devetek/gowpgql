package wptype

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Link        string `json:"link"`
	Description string `json:"description,omitempty"`
	Posts       struct {
		PageInfo struct {
			HasPreviousPage bool `json:"hasPreviousPage"`
			HasNextPage     bool `json:"hasNextPage"`
		} `json:"pageInfo"`
	} `json:"posts"`
}
