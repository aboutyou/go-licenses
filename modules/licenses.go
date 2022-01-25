package modules

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"regexp"

	"github.com/google/licenseclassifier"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

var (
	licenseRegexp = regexp.MustCompile(`^(?i)(LICENSE|LICENCE|COPYING|README|NOTICE)`)
)

func findLicenseFile(dep *modfile.Require) ([]byte, error) {
	fsPath, err := module.EscapePath(dep.Mod.Path)
	if err != nil {
		return nil, fmt.Errorf("escaping the module path failed: %w", err)
	}

	version, err := module.EscapeVersion(dep.Mod.Version)
	if err != nil {
		return nil, fmt.Errorf("escaping the module version failed: %w", err)
	}

	modBasePath := path.Join(build.Default.GOPATH, "pkg/mod")
	if os.Getenv("GOMODCACHE") != "" {
		modBasePath = os.Getenv("GOMODCACHE")
	}
	modFolder := path.Join(
		modBasePath,
		fmt.Sprintf("%s@%s", fsPath, version),
	)

	dir, err := os.ReadDir(modFolder)
	if err != nil {
		return nil, fmt.Errorf("reading the modules directory failed: %w", err)
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if !licenseRegexp.MatchString(file.Name()) {
			continue
		}

		license, err := os.ReadFile(path.Join(modFolder, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("reading the license file failed: %w", err)
		}

		return license, nil

	}

	return nil, fmt.Errorf("no license found for module %s", dep.Mod.String())
}

func resolveLicenseType(
	classifier *licenseclassifier.License,
	dep *modfile.Require,
) string {
	licenseFile, err := findLicenseFile(dep)
	if err != nil {
		return "unknown"
	}

	match := classifier.NearestMatch(string(licenseFile))
	if match == nil {
		return "unknown"
	}

	return match.Name
}
