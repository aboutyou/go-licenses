package export

import (
	"encoding/csv"
	"io"

	"github.com/aboutyou/go-licenses/modules"
)

type CSVWriter struct {
	csv *csv.Writer
}

func NewCSVWriter(output io.Writer) *CSVWriter {
	return &CSVWriter{
		csv: csv.NewWriter(output),
	}
}

func (writer *CSVWriter) WriteModules(mods []modules.Module) error {
	defer writer.csv.Flush()

	err := writer.csv.Write([]string{
		"Name",
		"Version",
		"License",
		"Link",
	})
	if err != nil {
		return err
	}

	for _, mod := range mods {
		err := writer.csv.Write([]string{
			mod.Path,
			mod.Version,
			mod.License,
			mod.PackageURL().String(),
		})

		if err != nil {
			return err
		}
	}

	return writer.csv.Error()
}
