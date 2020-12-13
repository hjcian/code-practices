package analyzingipaddresses

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var possibleIPPattern = regexp.MustCompile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}")

func extractIPsFromFile(fpath string) []string {
	bytes, err := ioutil.ReadFile(fpath)
	check(err)

	content := string(bytes)

	validIPs := make([]string, 0)
	for _, ipStr := range possibleIPPattern.FindAllString(content, -1) {
		if ip := net.ParseIP(ipStr); ip != nil {
			validIPs = append(validIPs, ip.String())
		}
	}
	return validIPs
}

func loadFiles(root string) []string {
	files := make([]string, 0)
	filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			files = append(files, path)
			return err
		},
	)
	return files
}

func extractIPFromFiles(filespaths []string) []string {
	ipSet := make(map[string]struct{}, 0)
	for _, fpath := range filespaths {
		for _, ip := range extractIPsFromFile(fpath) {
			ipSet[ip] = struct{}{}
		}
	}

	ips := make([]string, 0)
	for ip := range ipSet {
		ips = append(ips, ip)
	}

	return ips
}

func analyzingipaddresses(root string) string {
	filespaths := loadFiles(root)
	ips := extractIPFromFiles(filespaths)
	sort.Strings(ips)
	ret := ""
	for _, ip := range ips {
		ret += fmt.Sprintf("%v\n", ip)
		fmt.Println(ip)
	}
	return strings.TrimSpace(ret)
}

func main() {
	analyzingipaddresses("/root/devops/")
}
