package selectionprocess

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	chemistry = "chemistry.txt"
	maths     = "maths.txt"
	physics   = "physics.txt"
)

func findTxtFiles(root string) []string {
	files := make([]string, 0)
	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, ".txt") {
				files = append(files, path)
			}
			return nil
		})

	check(err)
	return files
}

// Student contains student's basic information
type Student struct {
	Name    string
	Surname string
	Score   int
}

// FullName return composition of "Name Surname"
func (s *Student) FullName() string {
	return s.Name + " " + s.Surname
}

// MakeStudent makes a Student by given string input
func MakeStudent(line string) *Student {
	info := strings.Split(line, " ")
	scroe, err := strconv.Atoi(info[2])
	check(err)

	return &Student{
		Name:    info[0],
		Surname: info[1],
		Score:   scroe,
	}
}

func loadStudents(fpath string) []*Student {
	file, err := os.Open(fpath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	students := make([]*Student, 0)
	for scanner.Scan() {
		students = append(students, MakeStudent(scanner.Text()))
	}

	return students
}

// SubjectManager manage the student's scores
type SubjectManager struct {
	Students []*Student
}

// GetBest4 sort the student by their score and return best 4
func (sm *SubjectManager) GetBest4() []*Student {
	sort.Slice(sm.Students, func(i, j int) bool {
		// use > as descending
		return sm.Students[i].Score > sm.Students[j].Score
	})
	end := 4
	if len(sm.Students) < end {
		end = len(sm.Students)
	}
	return sm.Students[:end]
}

// AppendStudents appends the students in its students
func (sm *SubjectManager) AppendStudents(students []*Student) {
	sm.Students = append(sm.Students, students...)
}

func extractSubject(filename string) string {
	return strings.Split(filepath.Base(filename), ".")[0]
}

// CourseSystem manges all SubjectManager
type CourseSystem map[string]*SubjectManager

// GetSubjectManager return a *SubjectManager if it exists,
// or create a new one before return
func (cs *CourseSystem) GetSubjectManager(subject string) *SubjectManager {
	sm, ok := (*cs)[subject]
	if !ok {
		(*cs)[subject] = &SubjectManager{make([]*Student, 0)}
		sm = (*cs)[subject]
	}
	return sm
}

func selectionProcess(root string) string {
	cs := make(CourseSystem)

	files := findTxtFiles(root)

	for _, filename := range files {
		subject := extractSubject(filename)
		sm := cs.GetSubjectManager(subject)
		sm.AppendStudents(loadStudents(filename))
	}

	var subjects []string
	for subject := range cs {
		subjects = append(subjects, subject)
	}
	sort.Strings(subjects)

	ret := ""
	for _, subject := range subjects {
		sm := cs.GetSubjectManager(subject)

		fmt.Printf("%s:\n", subject)
		ret += fmt.Sprintf("%s:\n", subject)
		for _, student := range sm.GetBest4() {
			fmt.Println(student.FullName())
			ret += fmt.Sprintf("%s\n", student.FullName())
		}
		fmt.Println()
		ret += fmt.Sprintf("\n")
	}

	return ret[:len(ret)-1]
}

func main() {
	selectionProcess("/root/devops/")
}
