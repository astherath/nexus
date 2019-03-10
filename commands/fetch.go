package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fetchCommand = &cobra.Command{

	// name of the cmd
	// TODO redo all this after testing
	Use:   "fetch [input]", // TODO make input the source of the json
	Short: "Fetch from JSON file",
	Long: `Given the pathname of a JSON file holding the
    	info of pro League of Legends matches, stores, and
	parses the data from it`,
	Run: fetchRun,
}

// fetch func TODO call exec.Command
func fetchRun(cmd *cobra.Command, args []string) {
	fmt.Println("Fetched ", args)
}

// when root command is called, add this command as well in init
func init() {
	RootCommand.AddCommand(fetchCommand)
}
