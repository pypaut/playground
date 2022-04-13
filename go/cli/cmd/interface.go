package cmd

import (
	"io/ioutil"
	"list-switches/constants"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface [switch_id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		l_out := log.New(os.Stdout, "", 0)
		l_err := log.New(os.Stderr, "", 1)

		if len(args) != 1 {
			l_err.Printf("Usage : nss interface [switch_id]")
			return
		}

		switch_id := args[0]

		url := constants.ApiRoot + "/switch/" + switch_id + "/interface"
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			l_err.Printf("Could not build request - %v", err)
			return
		}

		l_err.Printf("Fetching switch %v interface...", switch_id)
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			l_err.Printf("Could not request interface - %v", err)
			return
		}

		if response.StatusCode != 200 {
			l_err.Printf("Error, status code : %v", response.StatusCode)
			l_err.Printf("URL was : %v", url)
			return
		}

		responseBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			l_err.Printf("Could not read response body - %v", err)
			return
		}

		l_out.Printf(string(responseBytes) + "\n")
	},
}

func init() {
	rootCmd.AddCommand(interfaceCmd)
}
