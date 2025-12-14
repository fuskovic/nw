package main

import "github.com/fuskovic/nw/v4/cmd"

func main() {
	cmd.Root.CompletionOptions.DisableDefaultCmd = true
	cmd.Root.Execute()
}
