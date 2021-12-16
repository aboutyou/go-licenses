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
	licenseRegexp = regexp.MustCompile(`^(?i)(LICENSE|LICENCE|COPYING|README|NOTICE).*$`)
)

func findLicenseFile(dep *modfile.Require) ([]byte, error) {
	fsPath, err := module.EscapePath(dep.Mod.Path)
	if err != nil {
		return nil, err
	}

	version, err := module.EscapeVersion(dep.Mod.Version)
	if err != nil {
		return nil, err
	}

	modFolder := path.Join(
		build.Default.GOPATH,
		"pkg/mod",
		fmt.Sprintf("%s@%s", fsPath, version),
	)

	dir, err := os.ReadDir(modFolder)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if licenseRegexp.MatchString(file.Name()) {
			return os.ReadFile(path.Join(modFolder, file.Name()))
		}
	}

	return nil, fmt.Errorf("no license found")
}

func getLicenseType(
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
