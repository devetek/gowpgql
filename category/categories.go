package category

import (
	"context"
	"log"

	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
	"github.com/devetek/gowpgql/wptype"
)

type Categories struct {
	Data struct {
		Categories struct {
			Nodes []wptype.Category `json:"nodes"`
		} `json:"categories"`
	} `json:"data"`
	Extensions wptype.Extensions `json:"extensions"`
}

func GetCategories(client *gqlclient.Client, model *gqlmodel.Graphql, first int, last int, before string, after string, where map[string]interface{}) *Categories {
	ctx := context.Background()
	// response
	var categories = &Categories{}

	getposts, err := model.Query("category/query/categories.graphql")
	if err != nil {
		log.Println(err)
	}

	reqGql := gqlclient.NewRequest(getposts)

	// set variables
	if first != 0 {
		reqGql.Var("first", first)
	}

	if last != 0 {
		reqGql.Var("last", last)
	}

	if before != "" {
		reqGql.Var("before", before)
	}

	if after != "" {
		reqGql.Var("after", after)
	}

	if where != nil {
		reqGql.Var("where", where)
	}

	// run and capture the response
	if err := client.Run(ctx, reqGql, &categories); err != nil {
		log.Println("[gowpgql/category/categories.go] in line 56, error: ", err)
	}

	return categories
}
