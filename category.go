package gowpgql

import (
	"github.com/devetek/gowpgql/category"
)

// get list of categories
func (gwp *goWPGQL) GetCategories(first int, last int, before string, after string, where map[string]interface{}) *category.Categories {
	return category.GetCategories(gwp.client, gwp.model, first, last, before, after, where)
}

// get post detail, idType can be [SLUG, NAME, DATABASE_ID, ID, URI]
func (gwp *goWPGQL) GetCategory(
	id string,
	idType string,
	postFirst int,
	postLast int,
	postBefore string,
	postAfter string,
	postWhere map[string]interface{}) *category.Category {
	return category.GetCategory(
		gwp.client,
		gwp.model,
		id,
		idType,
		postFirst,
		postLast,
		postBefore,
		postAfter,
		postWhere,
	)
}
