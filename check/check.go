package check

import (
	"fmt"
	"strings"

	"github.com/aboutyou/go-licenses/modules"
	"github.com/google/licenseclassifier"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "check",
	Short: "Prints all licenses that apply to a Go package and its dependencies",
	RunE:  runCommand,
}

// TODO(HenriBeck): Add support for more licenses
var allowedLicenses = []string{
	licenseclassifier.Apache20,
	licenseclassifier.BSD2Clause,
	licenseclassifier.BSD3Clause,
	licenseclassifier.MIT,
	licenseclassifier.Facebook3Clause,
	licenseclassifier.Ruby,
	licenseclassifier.PHP301,
	licenseclassifier.Python20,
}

func runCommand(_ *cobra.Command, args []string) error {
	mods, err := modules.LoadModules(modules.LoadModuleOptions{})
	if err != nil {
		return err
	}

	forbiddenMods := CheckLicenses(mods)

	if len(forbiddenMods) > 0 {
		packageNames := make([]string, len(forbiddenMods))
		for index, mod := range forbiddenMods {
			packageNames[index] = mod.Path
		}

		return fmt.Errorf(
			"invalid licenses for packages: %s",
			strings.Join(packageNames, ", "),
		)
	}

	fmt.Println("All Licenses are valid")

	return nil
}

func CheckLicenses(mods []modules.Module) []modules.Module {
	forbiddenMods := make([]modules.Module, 0, len(mods))

	for _, mod := range mods {

		isAllowed := false
		for _, allowedLicense := range allowedLicenses {
			if mod.License == allowedLicense {
				isAllowed = true
			}
		}

		if !isAllowed {
			forbiddenMods = append(forbiddenMods, mod)
		}
	}

	return forbiddenMods
}
