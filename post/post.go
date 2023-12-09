package post

import (
	"context"
	"log"

	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
	"github.com/devetek/gowpgql/wptype"
)

type Post struct {
	Data struct {
		Post wptype.Node `json:"post"`
	} `json:"data"`
	Extensions wptype.Extensions `json:"extensions"`
}

func GetPost(client *gqlclient.Client, model *gqlmodel.Graphql, id string, idType string, isPreview bool) *Post {
	ctx := context.Background()
	// response
	var post = &Post{}

	getpost, err := model.Query("post/query/post.graphql")
	if err != nil {
		log.Println(err)
	}

	reqGql := gqlclient.NewRequest(getpost)

	// if variables not set
	if id == "" || idType == "" {
		return post
	}

	// set variables
	reqGql.Var("id", id)
	reqGql.Var("idType", idType)
	reqGql.Var("isPreview", isPreview)

	// run and capture the response
	if err := client.Run(ctx, reqGql, &post); err != nil {
		log.Println("[gowpgql/post/post.go] in line 45, error: ", err)
	}

	return post
}
