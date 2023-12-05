package wptype

type Menu struct {
	ID        string   `json:"id"`
	Slug      string   `json:"slug"`
	Name      string   `json:"name"`
	Locations []string `json:"locations"`
	MenuItems struct {
		Nodes []struct {
			ID    string `json:"id"`
			Label string `json:"label"`
			URI   string `json:"uri"`
		} `json:"nodes"`
	} `json:"menuItems"`
}
