package CLI

import (
	"example.com/mod/apiCall"
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
)

var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "web-server is about a Library archive",
	Long: `This is a demo of Library archive api-server, where you can use the basic methods of
			PUT, POST, GET, DELETE from postman. This is the server hit-point
			This is the CLI(Cmd) portion. 
			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server is called successfully")
		apiCall.ServerStartPoint(port)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Default Port For the HTTP Server")

	if err := serverCmd.MarkFlagRequired("port"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(serverCmd)
}
