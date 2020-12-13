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
	"sync"
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

// loadFiles is a simple generator generates filepath for consumers
func loadFiles(root string) <-chan string {
	fpathCh := make(chan string)

	go func() {
		defer close(fpathCh)

		err := filepath.Walk(
			root,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() {
					return nil
				}

				select {
				case fpathCh <- path:
				}

				return nil
			},
		)
		check(err)
	}()
	return fpathCh
}

func digestor(fpathCh <-chan string, resultCh chan<- []string) {
	for fpath := range fpathCh {
		ips := extractIPsFromFile(fpath)
		resultCh <- ips
	}
}

func workerPool(fpathCh <-chan string) chan []string {
	resultCh := make(chan []string)
	var wg sync.WaitGroup
	const numWorkers = 20

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			digestor(fpathCh, resultCh)
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func collectIPs(resultCh <-chan []string) []string {
	ipSet := make(map[string]struct{}, 0)
	for ips := range resultCh {
		for _, ip := range ips {
			ipSet[ip] = struct{}{}
		}
	}

	ips := make([]string, 0)
	for ip := range ipSet {
		ips = append(ips, ip)
	}
	sort.Strings(ips)

	return ips
}

func analyzingipaddresses(root string) string {
	fpathCh := loadFiles(root)
	ipsCh := workerPool(fpathCh)
	ips := collectIPs(ipsCh)

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
