package gowpgql

import (
	"github.com/devetek/go-core/gqlclient"
	"github.com/devetek/go-core/gqlmodel"
	"github.com/devetek/gowpgql/category"
	"github.com/devetek/gowpgql/post"
)

// public interface
type GoWPGQL interface {
	Client() *gqlclient.Client
	// get list of posts
	GetPosts(
		first int,
		last int,
		before string,
		after string,
		where map[string]interface{},
	) *post.Posts
	// get post detail
	GetPost(id string, idType string, isPreview bool) *post.Post
	// get category list
	GetCategories(
		first int,
		last int,
		before string,
		after string,
		where map[string]interface{},
	) *category.Categories
	GetCategory(
		id string,
		idType string,
		postFirst int,
		postLast int,
		postBefore string,
		postAfter string,
		postWhere map[string]interface{},
	) *category.Category
	// expose external model
	ExtQuery() *gqlmodel.Graphql
}

func New(endpoint string, options ...Option) GoWPGQL {
	m := goWPGQL{
		client: gqlclient.NewClient(endpoint),
		model:  gqlmodel.New(mdlPath),
	}

	for _, option := range options {
		option(&m)
	}

	return &m
}

// return gql client library
func (gwp *goWPGQL) Client() *gqlclient.Client {
	return gwp.client
}

func (gwp *goWPGQL) ExtQuery() *gqlmodel.Graphql {
	return gwp.ext
}
