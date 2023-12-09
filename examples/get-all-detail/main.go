package main

import (
	"fmt"

	"github.com/devetek/gowpgql"
	"github.com/devetek/gowpgql/examples/internal/constants"
)

func main() {
	gql := gowpgql.New(constants.GQL_ENDPOINT)

	post := gql.Post("hello-world", "SLUG", false)

	fmt.Println(post.Data.Post.Title)
}
