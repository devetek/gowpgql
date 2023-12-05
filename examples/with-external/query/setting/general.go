package setting

import (
	"context"
	"fmt"

	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
)

func GetGeneralSetting(client *gqlclient.Client, model *gqlmodel.Graphql) GeneralSetting {
	ctx := context.Background()
	var generalSetting GeneralSetting

	// request gql
	getgeneralsetting, err := model.Query("query/setting/general.graphql")
	if err != nil {
		fmt.Println(err)
	}

	reqGql := gqlclient.NewRequest(getgeneralsetting)

	// run it and capture the response
	if err := client.Run(ctx, reqGql, &generalSetting); err != nil {
		fmt.Println(err)
	}

	return generalSetting
}
