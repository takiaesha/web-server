package CLI

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "web-server",
	Short: "web-server is about a Library archive",
	Long: `This is a demo of Library archive api-server, where you can use the basic methods of
			PUT, POST, GET, DELETE from postman. This is the web-server hit-point, This is the CLI(Cmd) portion. 
			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1) // if get error, terminate the OS, why error because this rootCmd returns an error
	}
}

func init() {
	rootCmd.Flags().BoolP("detailsToggle", "t", false, "Effective help command for toggle-bar")

}
