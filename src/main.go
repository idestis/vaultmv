package main

import "github.com/idestis/vaultmv/src/cmd"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Execute(cmd.Version{
		Commit:  commit,
		Version: version,
		Date:    date,
	})
}