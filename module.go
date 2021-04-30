package linkedpackage

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Module struct {
	Lang           string
	Name           string
	Path           string
	Author         string
	LicenseName    string
	LicenseContent string
	Version        string
}

func (m *Module) readLicense(root string) error {
	// Find LICENSE*
	entries, err := ioutil.ReadDir(filepath.Join(root, m.Path))
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(strings.ToUpper(entry.Name()), "LICENSE") {
			licensePath := filepath.Join(root, m.Path, entry.Name())
			licenseContent, err := ioutil.ReadFile(licensePath)
			if err == nil {
				m.LicenseContent = strings.TrimSpace(string(licenseContent))
				return nil
			}
		}
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(strings.ToUpper(entry.Name()), "README") {
			readmePath := filepath.Join(root, m.Path, entry.Name())
			f, err := os.Open(readmePath)
			if err != nil {
				continue
			}
			s := bufio.NewScanner(f)
			var lines []string
			insideLicense := false
			var heading int
			for s.Scan() {
				text := s.Text()
				if !insideLicense {
					if strings.HasPrefix(text, "#") && strings.Contains(strings.ToUpper(text), "LICENSE") {
						left := strings.TrimLeft(text, "#")
						heading = len(text) - len(left)
						insideLicense = true
					}
 				} else {
					if strings.HasPrefix(text, "#") {
						left := strings.TrimLeft(text, "#")
						if len(text) - len(left) <= heading {
							m.LicenseContent = strings.Join(lines, "\n")
							return nil
						}
					}
					lines = append(lines, text)
				}
			}
		}
	}
	return errors.New("license file missing")
}

func UniqueModules(modules []Module) []Module {
	used := make(map[string]bool)
	result := []Module{}
	for _, module := range modules {
		key := module.Lang + "----" + module.Path
		if used[key] {
			continue
		}
		result = append(result, module)
		used[key] = true
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

type GroupedModule struct {
	Author string
	License string
	Modules []Module
}

func GroupingModulesByLicense(modules []Module) []GroupedModule {
	result := []GroupedModule{}
	indexes := make(map[string]int)
	for _, module := range modules {
		key := module.Author + "------" + module.LicenseName
		index, ok := indexes[key]
		if ok {
			result[index].Modules = append(result[index].Modules, module)
		} else {
			indexes[key] = len(result)
			result = append(result, GroupedModule{
				Author:  module.Author,
				License: module.LicenseName,
				Modules: []Module{
					module,
				},
			})
		}
	}

	return result
}

var projectDataReaders = map[string]func(*Module, string) error{}

func RegisterProjectDataReader(language string, reader func(*Module, string) error) {
	projectDataReaders[language] = reader
}

func ReadProjectData(module *Module, root string) error {
	reader, ok := projectDataReaders[module.Lang]
	if !ok {
		return fmt.Errorf("lang %s is not supported", module.Lang)
	}
	return reader(module, root)
}