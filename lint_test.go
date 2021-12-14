package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/pkic/ltl/ltl"
	yaml "gopkg.in/yaml.v3"
)

func TestDecode(t *testing.T) {
	validate := validator.New()

	err := filepath.WalkDir("./data", func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			return nil
		}

		// Skip hidden files and directories such as .git
		if s[0:1] == "." || strings.HasPrefix(s, "ltl.") {
			return nil
		}

		switch filepath.Ext(d.Name()) {
		case ".yaml":
			t.Run(s, func(t *testing.T) {
				file, err := os.Open(s)
				if err != nil {
					t.Errorf("failed to open file: %v", err)
					return
				}

				var pki ltl.PKI

				d := yaml.NewDecoder(file)
				d.KnownFields(true)
				err = d.Decode(&pki)
				if err != nil {
					t.Error(err)
					return
				}

				err = validate.Struct(pki)
				if err != nil {
					t.Error(err)
					return
				}

			})
		case ".go", ".mod", ".sum", ".md", "":
			return nil
		default:
			t.Errorf("unexpected file type: %s", s)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
