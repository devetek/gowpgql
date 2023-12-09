package main

import (
	"fmt"

	"github.com/devetek/gowpgql"
	"github.com/devetek/gowpgql/examples/internal/constants"
)

func main() {
	gql := gowpgql.New(constants.GQL_ENDPOINT)

	// get data lists
	respPosts := gql.GetPosts(1, 0, "", "", nil)
	respCategories := gql.GetCategories(1, 0, "", "", nil)

	// print posts
	posts := respPosts.Data.Posts.Nodes
	if len(posts) > 0 {
		for _, post := range posts {
			fmt.Println("Post Title: ", post.Title)
		}
	}

	// print categories
	categories := respCategories.Data.Categories.Nodes
	if len(categories) > 0 {
		for _, category := range categories {
			fmt.Println("Category Name: ", category.Name)
		}
	}

}
