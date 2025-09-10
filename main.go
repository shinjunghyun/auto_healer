package main

import (
	"auto_healer/cmd"
)

var (
	GitCommit string
	BuildTime string
)

func main() {
	cmd.AutoHealerStart(GitCommit, BuildTime)
}
