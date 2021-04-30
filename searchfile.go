package linkedpackage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Search(dir string, extensions ...string) []string {
	result := []string{}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			for _, extension := range extensions {
				if strings.HasSuffix(info.Name(), extension) {
					result = append(result, path)
				}
			}
		}
		return nil
	})
	return result
}