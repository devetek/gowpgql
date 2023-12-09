package post

import (
	"context"
	"log"

	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
	"github.com/devetek/gowpgql/wptype"
)

type Posts struct {
	Data struct {
		Posts struct {
			Nodes []wptype.Node `json:"nodes"`
		} `json:"posts"`
	} `json:"data"`
	Extensions wptype.Extensions `json:"extensions"`
}

func GetPosts(client *gqlclient.Client, model *gqlmodel.Graphql, first int, last int, before string, after string, where map[string]interface{}) *Posts {
	ctx := context.Background()
	// response
	var posts = &Posts{}

	getposts, err := model.Query("post/query/posts.graphql")
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
	if err := client.Run(ctx, reqGql, &posts); err != nil {
		log.Println("[gowpgql/post/posts.go] in line 56, error: ", err)
	}

	return posts
}
