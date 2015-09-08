package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/pics/server"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "server",
		Short: "Run the pics server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			server.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
