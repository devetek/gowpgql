package gowpgql

import (
	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
	"github.com/devetek/gowpgql/post"
)

// available interface
type GoWPGql interface {
	Client() *gqlclient.Client
	// get list of posts
	Posts(first int, last int, before string, after string, where map[string]interface{}) *post.Posts
	// expose external model
	ExtQuery() *gqlmodel.Graphql
}

func New(endpoint string, options ...Option) GoWPGql {
	m := goWPgql{
		client: gqlclient.NewClient(endpoint),
		model:  gqlmodel.New(mdlPath),
	}

	for _, option := range options {
		option(&m)
	}

	return &m
}

// return gql client library
func (gwp *goWPgql) Client() *gqlclient.Client {
	return gwp.client
}

func (gwp *goWPgql) ExtQuery() *gqlmodel.Graphql {
	return gwp.ext
}
