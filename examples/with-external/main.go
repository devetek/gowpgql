package main

import (
	"fmt"

	"github.com/devetek/gowpgql"
	"github.com/devetek/gowpgql/examples/with-external/query/setting"
)

func main() {
	gql := gowpgql.New("https://internal.jurnalistika.id/graphql", gowpgql.WithExternalModel(extmdl))

	fmt.Println(setting.GetGeneralSetting(gql.Client(), gql.ExtQuery()))
}
