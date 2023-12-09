## Description

GoWPGQL is an interface to communicate with WordPress through GraphQL by [WPGraphQL](https://www.wpgraphql.com/). Integrate with repository [Wordpress Playground](https://github.com/devetek/wpgraphql-playground) to play around with Wordpress.

### Prerequisite

Wordpress engine with WPGraphQL enabled, use [Wordpress Playground](https://github.com/devetek/wpgraphql-playground) to test GoWPGQL examples.

### Usage

1. Get Posts

```sh
package main

import (
	"fmt"

	"github.com/devetek/gowpgql"
)

func main() {
    var wpGQLURL = "https://graphql.terpusat.com/"

	gql := gowpgql.New(wpGQLURL)

    // get post without variables
	posts := gql.Posts(1, 0, "", "", nil)

	fmt.Println(posts)
}
```

### Examples

1. Use common available interface from GoWPGQL

```sh
make run-example-all-list
```

1. Define your custom model, if it not supported to your wordpress data structure

```sh
make run-example-with-external
```

## Todo

- [ ] Migrate wordpress common data to completing golang wordpress interface
