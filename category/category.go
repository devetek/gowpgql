package category

import (
	"context"
	"log"

	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
	"github.com/devetek/gowpgql/wptype"
)

type Category struct {
	Data struct {
		Category wptype.Category `json:"category"`
	} `json:"data"`
	Extensions wptype.Extensions `json:"extensions"`
}

func GetCategory(
	client *gqlclient.Client,
	model *gqlmodel.Graphql,
	id string,
	idType string,
	postFirst int,
	postLast int,
	postBefore string,
	postAfter string,
	postWhere map[string]interface{},
) *Category {
	ctx := context.Background()
	// response
	var category = &Category{}

	getcategory, err := model.Query("category/query/category.graphql")
	if err != nil {
		log.Println(err)
	}

	reqGql := gqlclient.NewRequest(getcategory)

	// if variables not set
	if id == "" || idType == "" {
		return category
	}

	// set variables
	reqGql.Var("id", id)
	reqGql.Var("idType", idType)

	if postFirst != 0 {
		reqGql.Var("postFirst", postFirst)
	}

	if postLast != 0 {
		reqGql.Var("postLast", postLast)
	}

	if postBefore != "" {
		reqGql.Var("postBefore", postBefore)
	}

	if postAfter != "" {
		reqGql.Var("postAfter", postAfter)
	}

	if postWhere != nil {
		reqGql.Var("postWhere", postWhere)
	}

	// run and capture the response
	if err := client.Run(ctx, reqGql, &category); err != nil {
		log.Println("[gowpgql/category/category.go] in line 68, error: ", err)
	}

	return category
}
