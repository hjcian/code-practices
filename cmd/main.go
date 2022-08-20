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
		log.Printf("executable path: %s", exePath)
		return exePath
	} else if strings.HasSuffix(pwd, "leetcode") {
		log.Printf("pwd: %s", pwd)
		return pwd
	}

	if info, err := os.Stat(pwd + "/leetcode"); err == nil && info.IsDir() {
		log.Println(pwd + "/leetcode")
		return pwd + "/leetcode"
	}
	log.Fatal("where are you?")
	return ""
}

func genName(title string) string {
	tokens := strings.Split(title, ".")
	if len(tokens) != 2 {
		log.Fatal("expect title is formatted as: `<N>. <Title>`")
	}
	questionNumber, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatal(err, "expect title has number, N: `<N>. <Title>`")
	}
	formattedNum := fmt.Sprintf("%04d", questionNumber)
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

func writeReadmeFile(root string) error {
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

func writeMainFile(root, name string) error {
	fpath := filepath.Join(root, name+".go")
	if err := checkFileExists(fpath); err != nil {
		return err
	}

	packageName := strings.ToLower(strings.Split(name, ".")[1])
	content := fmt.Sprintf(`
package %s
`, packageName)
	return ioutil.WriteFile(fpath, []byte(strings.TrimSpace(content)), 0644)
}

func writeTestFile(root, name string) error {
	fpath := filepath.Join(root, name+"_test.go")
	if err := checkFileExists(fpath); err != nil {
		return err
	}

	packageName := strings.ToLower(strings.Split(name, ".")[1])

	testFuncName := strings.ReplaceAll(name, ".", "_")
	content := fmt.Sprintf(`
package %s

import "testing"

func Test_%s(t *testing.T) {

}

`, packageName, testFuncName)
	return ioutil.WriteFile(fpath, []byte(strings.TrimSpace(content)), 0644)
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
	log.Printf("%s will be created", folderPath)

	os.Mkdir(folderPath, 0755)

	if err := writeReadmeFile(folderPath); err != nil {
		log.Printf("%s: %s", "create readme error", err)
	}

	if err := writeMainFile(folderPath, name); err != nil {
		log.Printf("%s: %s", "create main error", err)
	}

	if err := writeTestFile(folderPath, name); err != nil {
		log.Printf("%s: %s", "creat test error", err)
	}
}
