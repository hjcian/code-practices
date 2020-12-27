package footballfanids

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(fpath string) []string {
	file, err := os.Open(fpath)
	check(err)

	scanner := bufio.NewScanner(file)

	ret := []string{}
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

// compareList compares the given path to a list by given matcher
// matcher is expected either a strings.HasSuffix or strings.Contains
func compareList(path string, list []string, matcher func(a, b string) bool) bool {
	dotPath := strings.ReplaceAll(
		filepath.Dir(path),
		"/", ".")

	for _, pattern := range list {
		if matcher(dotPath, pattern) {
			return true
		}
	}
	return false
}

func listInvitedFans(root string, invited []string) []string {
	ret := []string{}

	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() &&
				info.Name() == "fan.info" &&
				compareList(path, invited, strings.Contains) {
				ret = append(ret, path)
			}
			return nil
		},
	)
	check(err)

	return ret
}

func filterBannedFans(invitedFans, banned []string) []string {
	ret := []string{}

	for _, path := range invitedFans {
		if !compareList(path, banned, strings.HasSuffix) {
			ret = append(ret, path)
		}
	}
	return ret
}

func loadFanNOs(fpaths []string) []string {
	allFans := []string{}

	for _, path := range fpaths {
		fans := readLines(path)
		allFans = append(allFans, fans...)
	}

	return allFans
}

func footballFanIds(root string) string {
	invitedList := readLines(fmt.Sprintf("%s/invite.info", root))
	bannedList := readLines(fmt.Sprintf("%s/ban.info", root))

	invitedFans := listInvitedFans(root, invitedList)
	filteredInvitedFans := filterBannedFans(invitedFans, bannedList)

	nos := loadFanNOs(filteredInvitedFans)
	sort.Strings(nos) // sorted lexicographically

	ret := strings.Join(nos, "\n")
	fmt.Printf(ret)
	return ret
}

func main() {
	footballFanIds("/root/devops")
}
