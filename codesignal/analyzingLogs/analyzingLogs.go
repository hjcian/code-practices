package analyzinglogs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// LogType is enumeration of log type
var LogType = `ERROR|WARNING|TRACE|INFO`

// LogFormatPattern is a named matches of Regexp
var LogFormatPattern = regexp.MustCompile(fmt.Sprintf(
	`(?P<timestamp>\d+)\|`+
		`(?P<logtype>%[1]v)\|`+
		`(?P<filename>.+)\|`+
		`(?P<line>\d+)\|`+
		`(?P<func>.+)\|`+
		`(?P<errmsg>.+)`,
	LogType,
))

func listLogs(root string) []string {
	logs := []string{}

	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if strings.HasSuffix(path, ".log") {
				logs = append(logs, path)
			}

			return nil
		},
	)

	check(err)

	return logs
}

func extractErrorFile(line string) string {
	match := LogFormatPattern.FindStringSubmatch(line)
	fmt.Println(match)
	// for i, name := range LogFormatPattern.SubexpNames() {
	// 	fmt.Println(i, name)
	// 	if i != 0 && name == "ERROR" {
	// 		result[name] = match[i]
	// 	}
	// }
	return ""
}

func analyzingLogs(root string) string {
	logs := listLogs(root)
	fmt.Println(logs)

	ioutil.ReadFile()
	s := "1477123675|ERROR|handler.cpp|127|findHandlers|Division by zero"

	// result := make(map[string]string)

	// match := LogFormatPattern.FindStringSubmatch(s)
	// for i, name := range LogFormatPattern.SubexpNames() {
	// 	fmt.Println(i, name)
	// 	if i != 0 && name != "" {
	// 		result[name] = match[i]
	// 	}
	// }

	fmt.Println(result)
	return ""
}

func main() {
	analyzingLogs("/root/devops/")
}
