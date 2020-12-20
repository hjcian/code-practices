package analyzinglogs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
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

type Record struct {
	timestamp time.Time
	logtype   string
	filename  string
	line      uint64
	function  string
	errmsg    string
}

func (r *Record) IsErrorType() bool {
	if r.logtype == "ERROR" {
		return true
	}
	return false
}

func parseRecord(record string) *Record {
	match := LogFormatPattern.FindStringSubmatch(record)

	ts, err := strconv.ParseInt(match[1], 10, 64)
	check(err)

	line, err := strconv.ParseUint(match[4], 10, 64)
	check(err)

	return &Record{
		timestamp: time.Unix(ts, 0),
		logtype:   match[2],
		filename:  match[3],
		line:      line,
		function:  match[5],
		errmsg:    match[6],
	}
}

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

func readLog2Lines(logpath string) []string {
	file, err := os.Open(logpath)
	check(err)
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if text := strings.TrimSpace(scanner.Text()); len(text) > 0 {
			lines = append(lines, text)
		}
	}

	check(scanner.Err())
	return lines
}

type Result struct {
	filename string
	count    int
}

func convert2Results(errorCounts map[string]int) []Result {
	fname2Count := make([]Result, 0, len(errorCounts))

	for filename := range errorCounts {
		fname2Count = append(fname2Count, Result{
			filename: filename,
			count:    errorCounts[filename],
		})
	}

	sort.Slice(fname2Count, func(i, j int) bool {
		if fname2Count[i].count > fname2Count[j].count {
			return true
		} else if fname2Count[i].count == fname2Count[j].count &&
			fname2Count[i].filename < fname2Count[j].filename {
			return true
		}
		return false
	})

	return fname2Count
}

func analyzingLogs(root string) string {
	errorCounts := make(map[string]int, 0)

	logs := listLogs(root)
	for _, log := range logs {
		for _, line := range readLog2Lines(log) {
			if record := parseRecord(line); record.IsErrorType() {
				errorCounts[record.filename]++
			}
		}
	}

	results := convert2Results(errorCounts)
	ret := ""
	for _, result := range results {
		ret += fmt.Sprintf("%s %d\n", result.filename, result.count)
	}
	return ret
}

func main() {
	analyzingLogs("/root/devops/")
}
