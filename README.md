## Description

Golang interface to communicate with wordpress graphql by ['wpgraphql'](https://www.wpgraphql.com/). By using this you can easily to get, update or delete data from golang to wordpress engine.

### Usage

1. Get Posts

```sh
package main

import (
	"fmt"

	"github.com/devetek/gowpgql"
)

func main() {
    var gqlUrl = "https://graphql.terpusat.com/"

	gql := gowpgql.New(gqlUrl)

    // get post without variables
	posts := gql.Posts(1, 0, "", "", nil)

	fmt.Println(posts)
}
```

### Examples

1. Use common available interface from this package

```sh
make run-example-getposts
```

1. Define your own model, if it not supported to your wordpress data structure

```sh
make run-example-with-external
```

## Todo

- [ ] Migrate wordpress common data to completing golang wordpress interface
