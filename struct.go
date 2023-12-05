package gowpgql

import (
	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
)

type goWPgql struct {
	client *gqlclient.Client
	model  *gqlmodel.Graphql
	ext    *gqlmodel.Graphql
}
