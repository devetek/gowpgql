package gowpgql

import (
	"github.com/devetek/go-core/gqlmodel"
	"github.com/devetek/go-core/mdfs"
)

type Option func(*goWPGQL)

// WithExternalModel sets other folder location used as model
// ext default is nil, no external model included
func WithExternalModel(extmodel mdfs.FS) Option {
	return func(g *goWPGQL) {
		var extm = gqlmodel.New(extmodel)

		g.ext = extm
	}
}
