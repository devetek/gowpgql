package gowpgql

import (
	"github.com/devetek/gowpgql/category"
)

// get list of categories
func (gwp *goWPgql) GetCategories(first int, last int, before string, after string, where map[string]interface{}) *category.Categories {
	return category.GetCategories(gwp.client, gwp.model, first, last, before, after, where)
}
