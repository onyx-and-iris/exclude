package main

import (
	"context"
	"os"
	"runtime/debug"
	"strings"

	"github.com/charmbracelet/fang"
	"github.com/onyx-and-iris/exclude/cmd"
)

var version string // Version of the CLI, set during build time

func versionFromBuild() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(unable to read version)"
	}
	return strings.Split(info.Main.Version, "-")[0]
}

func main() {
	if err := fang.Execute(context.Background(), cmd.RootCmd, fang.WithVersion(versionFromBuild())); err != nil {
		os.Exit(1)
	}
}
