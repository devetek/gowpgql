package wptype

type Extensions struct {
	Debug []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"debug"`
}
