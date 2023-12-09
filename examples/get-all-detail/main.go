package main

import (
	"fmt"

	"github.com/devetek/gowpgql"
	"github.com/devetek/gowpgql/examples/internal/constants"
)

func main() {
	gql := gowpgql.New(constants.GQL_ENDPOINT)

	post := gql.GetPost("hello-world", "SLUG", false)
	category := gql.GetCategory("uncategorized", "SLUG", 0, 0, "", "", nil)

	fmt.Println("Post Title: ", post.Data.Post.Title)
	fmt.Println("Category Name: ", category.Data.Category.Name)
}
