package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/pics/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "pics"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
