package main

import (
	"fmt"

	"github.com/devetek/gowpgql"
)

func main() {
	gql := gowpgql.New("https://www.okcarlomboktransport.com/graphql")

	posts := gql.Posts(1, 0, "", "", nil)

	fmt.Println(posts)
}
