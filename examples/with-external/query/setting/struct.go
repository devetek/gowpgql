package setting

import "github.com/devetek/gowpgql/wptype"

// extends model from wptype or you can create your own if you have Advanced Custom Fields
type GeneralSetting struct {
	Data struct {
		GeneralSettings wptype.General `json:"generalSettings"`
	} `json:"data"`
	Extensions wptype.Extensions `json:"extensions"`
}
