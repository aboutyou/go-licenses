package export

import (
	"encoding/json"
	"io"

	"github.com/aboutyou/go-licenses/modules"
)

type JSONWriter struct {
	json *json.Encoder
}

func NewJSONWriter(output io.Writer) *JSONWriter {
	return &JSONWriter{
		json: json.NewEncoder(output),
	}
}

func (writer *JSONWriter) WriteModules(mods []modules.Module) error {
	jsonMods := make([]map[string]string, 0, len(mods))
	for _, mod := range mods {
		jsonMods = append(jsonMods, map[string]string{
			"path":    mod.Path,
			"version": mod.Path,
			"license": mod.License,
			"url":     mod.PackageURL().String(),
		})
	}

	return writer.json.Encode(jsonMods)
}
