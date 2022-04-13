package cmd

import (
	"io/ioutil"
	"list-switches/constants"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "List all switches",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		l_out := log.New(os.Stdout, "", 0)
		l_err := log.New(os.Stderr, "", 1)

		url := constants.ApiRoot + "/switch"
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			l_err.Printf("Could not build request - %v", err)
			return
		}

		l_err.Printf("Fetching switches...")
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			l_err.Printf("Could not request switches - %v", err)
			return
		}

		responseBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			l_err.Printf("Could not read response body - %v", err)
			return
		}

		l_out.Printf(string(responseBytes))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
