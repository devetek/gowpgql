package wptype

type General struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// type GeneralSetting struct {
// 	Data struct {
// 		GeneralSettings General `json:"generalSettings"`
// 	} `json:"data"`
// 	Extensions Extensions `json:"extensions"`
// }
