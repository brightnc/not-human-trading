package cmds

import (
	"github.com/brightnc/not-human-trading/protocol"

	"github.com/spf13/cobra"
)

/*
	|--------------------------------------------------------------------------
	| Command to serve REST protocol
	|--------------------------------------------------------------------------
	|
	| Here is the command that makes your application serve the REST protocol
	| for client.
	|
*/

var serveRESTCmd = &cobra.Command{
	Use:   "serve-rest",
	Short: "start a http server",
	RunE:  serveREST,
}

// @title github.com/brightnc/not-human-trading
// @version 1.0
// @host localhost:8080
// @schemes http https
func serveREST(cmd *cobra.Command, args []string) error {
	return protocol.ServeREST()
}

func init() {
	rootCmd.AddCommand(serveRESTCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveRESTCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveRESTCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
