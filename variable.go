package gowpgql

import (
	"embed"
	"os"

	"github.com/devetek/go-core/mdfs"
)

var (
	//go:embed **/query/*.graphql
	mdl embed.FS

	mdlPath = mdfs.New(mdl, "", os.Getenv("ENV"))
)
