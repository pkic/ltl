package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkic/ltl/ltl"

	"github.com/go-playground/validator"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	var result []ltl.PKI

	validate := validator.New()

	err := filepath.WalkDir("./data", func(s string, f fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if f.IsDir() {
			return nil
		}

		// Skip hidden files and directories such as .git
		if s[0:1] == "." || strings.HasPrefix(s, "ltl.") {
			return nil
		}

		switch filepath.Ext(f.Name()) {
		case ".yaml":
			file, err := os.Open(s)
			if err != nil {
				return err
			}

			var pki ltl.PKI

			d := yaml.NewDecoder(file)
			d.KnownFields(true)
			err = d.Decode(&pki)
			if err != nil {
				return err
			}

			err = validate.Struct(pki)
			if err != nil {
				return err
			}

			if pki.ID == "" {
				pki.ID = strings.ToLower(strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
			}

			result = append(result, pki)

		case ".go", ".mod", ".sum", ".md", "":
			return nil
		default:
			return fmt.Errorf("unexpected file type: %s", s)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	yamlData, err := yaml.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("ltl.yaml", yamlData, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("ltl.json", jsonData, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
