package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getExePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(ex)
	return dir, nil
}

func getPWD() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func getLeetCodePath() (string, error) {
	exePath, err := getExePath()
	if err != nil {
		return "", err
	}

	pwd, err := getPWD()
	if err != nil {
		return "", err
	}

	if strings.HasSuffix(exePath, "leetcode") {
		return exePath, nil
	} else if strings.HasSuffix(pwd, "leetcode") {
		return pwd, nil
	}

	if info, err := os.Stat(pwd + "/leetcode"); err == nil && info.IsDir() {
		log.Println(pwd + "/leetcode")
		return pwd + "/leetcode", nil
	}

	return "", errors.New("where are you?")
}

func genName(title string) (string, error) {
	tokens := strings.Split(title, ".")
	if len(tokens) != 2 {
		return "", errors.New("expect title is formatted as: `<N>. <Title>`")
	}
	questionNumber, err := strconv.Atoi(tokens[0])
	if err != nil {
		return "", fmt.Errorf("expect title has number, N: `<N>. <Title>`: %w", err)
	}
	formattedNum := fmt.Sprintf("%04d", questionNumber)
	formattedText := strings.ReplaceAll(tokens[1], " ", "")
	return formattedNum + "." + formattedText, nil
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
	rootPath, err := getLeetCodePath()
	logOrFatal(err, "checked rootPath:", rootPath)

	name, err := genName(title)
	logOrFatal(err, "formatted title:", name)

	folderPath := filepath.Join(rootPath, name)
	logfOrFatal(nil, "%s will be created", folderPath)
	os.Mkdir(folderPath, 0755)

	logfOrFatal(writeReadmeFile(folderPath), "create readme (%s)", folderPath)
	logfOrFatal(writeMainFile(folderPath, name), "create main (%s/%s)", folderPath, name)
	logfOrFatal(writeTestFile(folderPath, name), "creat test (%s/%s)", folderPath, name)
}

func logOrFatal(err error, v ...any) {
	log.Println(v...)
	if err != nil {
		log.Fatal("[ERROR]", err.Error())
	}
}

func logfOrFatal(err error, format string, v ...any) {
	log.Printf(format, v...)
	if err != nil {
		log.Fatal("[ERROR]", err.Error())
	}
}
