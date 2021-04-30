package linkedpackage

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type sourceMap struct {
	Sources []string `json:"sources"`
}

func ParseJSSourcemapFile(path string) ([]Module, error) {
	result := []Module{}
	tmp := make(map[string]Module)

	f, err := os.Open(path)
	if err != nil {
		return result, err
	}
	dec := json.NewDecoder(f)
	var sm sourceMap
	dec.Decode(&sm)

	for _, source := range sm.Sources {
		modulePath := source
		if strings.HasPrefix(source, "webpack:///.") {
			// process "webpack:///./node_modules/@babel/runtime/helpers/wrapNativeSuper/_index.mjs"
			modulePath = strings.TrimPrefix(source, "webpack:///.")
		} else if strings.HasPrefix(source, "../webpack:/") {
			// process "../webpack:/ncc-project/node_modules/trim/index.js"
			// process "../webpack://ncc-project/./node_modules/trim/index.js"
			modulePath = strings.TrimPrefix(source, "../webpack:/")
			modulePath = strings.TrimPrefix(modulePath, "/")
			i := strings.Index(modulePath, "/")
			if i != -1 {
				modulePath = modulePath[i:]
				modulePath = strings.TrimPrefix(modulePath, "/.")
			}
		}
		modules := parseJSModulePaths(modulePath)
		for _, module := range modules {
			tmp[module.Name] = module
		}
	}

	for _, v := range tmp {
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}

func ParseJSWebPack(path string) ([]Module, error) {
	result := []Module{}

	comments := []string{}
	f, err := os.Open(path)
	if err != nil {
		return result, err
	}
	scanner := bufio.NewScanner(f)
	bufLen := bufio.MaxScanTokenSize
	scanner.Buffer(make([]byte, bufLen, 1000*bufLen), 1000*bufLen)

	var insideComment bool
	var lastComments []string
	for scanner.Scan() {
		text := scanner.Text()
		for {
			if !insideComment {
				i := strings.Index(text, "/*")
				if i == -1 {
					break
				}
				text = text[i+2:]
				ei := strings.Index(text, "*/")
				if ei == -1 {
					lastComments = append(lastComments, text)
					insideComment = true
					break
				} else {
					comments = append(comments, text[:ei])
					text = text[:ei+2]
				}
			} else {
				i := strings.Index(text, "*/")
				if i == -1 {
					lastComments = append(lastComments, text)
					break
				}
				lastComments = append(lastComments, text[:i])
				comments = append(comments, strings.Join(lastComments, "\n"))
				lastComments = nil
				insideComment = false
				text = text[:i+2]
			}
		}
	}

	tmp := make(map[string]Module)
	for _, comment := range comments {
		lines := strings.Split(comment, "\n")
		if strings.HasPrefix(lines[0], "!**********") && len(lines) == 3 {
			// pass
		} else {
			continue
		}
		modulePath := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(lines[1]), "!*** "), "***!"))
		modules := parseJSModulePaths(modulePath)
		for _, module := range modules {
			tmp[module.Name] = module
		}
	}

	for _, v := range tmp {
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}

func parseJSModulePaths(origModulePaths string) []Module {
	result := []Module{}
	for _, modulePath := range strings.Split(origModulePaths, "!") {
		i := strings.Index(modulePath, "?")
		if i != -1 {
			modulePath = modulePath[:i]
		}
		i = strings.LastIndex(modulePath, "/node_modules/")
		if i == -1 {
			return result
		}
		parsed := parseJSModulePath(modulePath)
		if parsed != nil {
			result = append(result, *parsed)
		}
	}
	return result
}

func parseJSModulePath(modulePath string) *Module {
	if strings.HasPrefix(modulePath, "./node_modules") {
		modulePath = strings.TrimPrefix(modulePath, ".")
	}
	i := strings.LastIndex(modulePath, "/node_modules/")
	if i == -1 {
		return nil
	}
	moduleNameSrc := modulePath[i:]
	moduleNameSrc = strings.TrimPrefix(moduleNameSrc, "/node_modules/")
	fragments := strings.Split(moduleNameSrc, "/")
	var moduleName string
	if strings.HasPrefix(fragments[0], "@") {
		moduleName = fragments[0] + "/" + fragments[1]
	} else {
		moduleName = fragments[0]
	}
	modulePath = modulePath[:i] + "/node_modules/" + moduleName
	return &Module{
		Lang: "js",
		Name: moduleName,
		Path: modulePath,
	}
}

func projectJSConfigReader(module *Module, root string) error {
	f, err := os.Open(filepath.Join(root, module.Path, "package.json"))
	if err != nil {
		return err
	}
	d := json.NewDecoder(f)
	j := make(map[string]interface{})
	d.Decode(&j)
	author, err := projectJSParseAuthor(j)
	if err != nil {
		return err
	}
	module.Author = author

	l, ok := j["license"]
	if !ok {
		module.LicenseName = "no license"
	} else {
		lname, ok := l.(string)
		if !ok {
			lname = fmt.Sprintf("%v", l)
		}
		module.LicenseName = lname
	}

	err = module.readLicense(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", module.Name, err.Error())
	}

	module.Version = j["version"].(string)

	return nil
}

func projectJSParseAuthor(content map[string]interface{}) (string, error) {
	author, ok := content["author"]
	if ok {
		switch val := author.(type) {
		case string:
			return val, nil
		case map[string]interface{}:
			email, hasEmail := val["email"]
			name, hasName := val["name"]
			if hasName && hasEmail {
				return name.(string) + " <" + email.(string) + ">", nil
			} else if hasName {
				return name.(string), nil
			} else {
				log.Println(val)
				return fmt.Sprintf("%s", val), nil
			}
		}
	}
	name, ok := content["name"]
	if !ok {
		return "",  errors.New("not implemented")
	}
	return name.(string) + " authors", nil
}

func init() {
	RegisterProjectDataReader("js", projectJSConfigReader)
}