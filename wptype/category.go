package wptype

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
	ParentID    string `json:"parentId,omitempty"`
	Parent      struct {
		Node struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"node,omitempty"`
	} `json:"parent,omitempty"`
	Children struct {
		Nodes []Category `json:"nodes"`
	} `json:"children"`
	Posts struct {
		PageInfo struct {
			HasPreviousPage bool `json:"hasPreviousPage"`
			HasNextPage     bool `json:"hasNextPage"`
		} `json:"pageInfo"`
		Nodes []Node `json:"nodes"`
	} `json:"posts,omitempty"`
}
