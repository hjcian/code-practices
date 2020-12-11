package analyzingipaddresses

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var possibleIPPattern = regexp.MustCompile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}")

func extractIPsFromFile(fpath string) {
	bytes, err := ioutil.ReadFile(fpath)
	check(err)
	content := string(bytes)

	fmt.Println(content)
	fmt.Println(possibleIPPattern.FindAllString(content, -1))
}

func loadFiles(root string) []string {
	files := make([]string, 0)
	filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			fmt.Println(path)
			extractIPsFromFile(path)
			files = append(files, path)
			return err
		},
	)
	return files
}

func analyzingipaddresses(root string) string {
	loadFiles(root)
	return ""
}
