package export

import (
	"fmt"
	"os"

	"github.com/aboutyou/go-licenses/modules"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "export {csv | json}",
	Short:   "Prints all modules and their licenses in either CSV or JSON format for license compliance usecases.",
	Example: "go-licenses export csv > licenses.csv",
	Args:    cobra.ExactArgs(1),
	RunE:    runCommand,
}

func runCommand(_ *cobra.Command, args []string) error {
	mods, err := modules.LoadModules(modules.LoadModuleOptions{})
	if err != nil {
		return err
	}

	switch args[0] {
	case "csv":
		csvWriter := NewCSVWriter(os.Stdout)
		return csvWriter.WriteModules(mods)

	case "json":
		jsonWriter := NewJSONWriter(os.Stdout)
		return jsonWriter.WriteModules(mods)
	}

	return fmt.Errorf("unknown formatter: %s", args[0])
}
