package modules

import (
	"fmt"
	"net/url"
	"os"

	"github.com/google/licenseclassifier"
	"golang.org/x/mod/modfile"
)

type Module struct {
	Path    string
	Version string
	License string
}

func (module *Module) GetPackageURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   "pkg.go.dev",
		Path:   fmt.Sprintf("%s@%s", module.Path, module.Version),
	}
}

type LoadModuleOptions struct {
	IncludeIndirectModules bool
}

func LoadModules(options LoadModuleOptions) ([]Module, error) {
	content, err := os.ReadFile("./go.mod")
	if err != nil {
		return nil, err
	}

	file, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return nil, err
	}

	classifier, err := licenseclassifier.New(0.8)
	if err != nil {
		return nil, err
	}

	modules := make([]Module, 0, len(file.Require))

	for _, dep := range file.Require {
		if dep.Indirect && !options.IncludeIndirectModules {
			continue
		}

		modules = append(modules, Module{
			Path:    dep.Mod.Path,
			Version: dep.Mod.Version,
			License: getLicenseType(classifier, dep),
		})
	}

	return modules, nil
}
