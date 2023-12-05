package main

import (
	"embed"

	"github.com/devetek/go-core/mdfs"
)

// create your model folder structure
var (
	//go:embed query/**/*.graphql
	mdl embed.FS

	extmdl = mdfs.New(mdl, "", "development")
)
