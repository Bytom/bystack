package main

import (
	"runtime"

	cmd "github.com/bytom/bystack/cmd/bystackcli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
