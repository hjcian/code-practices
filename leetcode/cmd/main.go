package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getExePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(ex)
	return dir
}

func getPWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return pwd
}

func getLeetCodePath() string {
	exePath := getExePath()
	pwd := getPWD()
	if strings.HasSuffix(exePath, "leetcode") {
		return exePath
	} else if strings.HasSuffix(pwd, "leetcode") {
		return pwd
	}
	log.Fatal("where are you?")
	return ""
}

func genName(title string) string {
	tokens := strings.Split(title, ".")
	if len(tokens) != 2 {
		log.Fatal("expect title is formatted as: `<N>. <Camel> <Case>`")
	}
	num, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatal(err, "expect title has number, N: `<N>. <Camel> <Case>`")
	}
	formattedNum := fmt.Sprintf("%04d", num)
	formattedText := strings.ReplaceAll(tokens[1], " ", "")
	return formattedNum + "." + formattedText
}

func checkFileExists(path string) error {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	return fmt.Errorf("`%s` exists (%d bytes)", stat.Name(), stat.Size())
}

func writeREADME(root string) error {
	readmeFile := filepath.Join(root, "README.md")
	if err := checkFileExists(readmeFile); err != nil {
		return err
	}

	content := `
# Question

# Thoughts
`
	return ioutil.WriteFile(readmeFile, []byte(strings.TrimSpace(content)), 0644)
}

func writeMain(root, name string) error {
	readmeFile := filepath.Join(root, name+".go")
	if err := checkFileExists(readmeFile); err != nil {
		return err
	}

	packageName := strings.ToLower(strings.Split(name, ".")[1])
	content := fmt.Sprintf(`
package %s
`, packageName)
	return ioutil.WriteFile(readmeFile, []byte(strings.TrimSpace(content)), 0644)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatal("forgot to give me question title?")
	}

	title := args[0]
	rootPath := getLeetCodePath()
	log.Println("checked rootPath:", rootPath)

	name := genName(title)
	log.Println("formatted title:", name)

	folderPath := filepath.Join(rootPath, name)
	os.Mkdir(folderPath, 0755)

	if err := writeREADME(folderPath); err != nil {
		log.Printf("%s: %s", "writeREADME error", err)
	}

	if err := writeMain(folderPath, name); err != nil {
		log.Printf("%s: %s", "writeMain error", err)
	}
}
