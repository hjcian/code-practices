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

func executor(fpath string) <-chan []string {
	ch := make(chan []string)
	go func() {
		ips := extractIPsFromFile(fpath)
		ch <- ips
		close(ch)
	}()
	return ch
}

func fanOut(fpaths []string) []chan []string {
	workers := make([]chan []string, len(fpaths))
	for i := range workers {
		workers[i] = make(chan []string)
	}

	for i := range fpaths {
		go func(i int) {
			select {
			case out := <-executor(fpaths[i]):
				workers[i] <- out
				close(workers[i])
			}
		}(i)
	}

	return workers
}

func fanIn(jobs ...chan []string) chan []string {
	var wg sync.WaitGroup
	out := make(chan []string)

	output := func(j <-chan []string) {
		defer wg.Done()
		for n := range j {
			out <- n
		}
	}

	wg.Add(len(jobs))

	for _, job := range jobs {
		go output(job)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func concurrentExtractIPs(fpaths []string) []string {
	workers := fanOut(fpaths)

	ips := make([]string, 0)
	for results := range fanIn(workers...) {
		ips = append(ips, results...)
	}
	return ips
}

func extractIPFromFiles(fpaths []string) []string {
	ipSet := make(map[string]struct{}, 0)
	for _, ip := range concurrentExtractIPs(fpaths) {
		ipSet[ip] = struct{}{}
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
