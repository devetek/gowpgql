package gowpgql

import "github.com/devetek/gowpgql/post"

// get list of posts
// you can use this interface or directly access functionality inside the post independently
// by accessing directly to post, you need to import some packages separately
//
//	import "github.com/devetek/go-core/gqlclient"
//
// import "github.com/devetek/go-core/gqlmodel"
func (gwp *goWPgql) Posts(first int, last int, before string, after string, where map[string]interface{}) *post.Posts {

	return post.GetPosts(gwp.client, gwp.model, first, last, before, after, where)
}
