package selectionprocess

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractSubject(filename string) string {
	return strings.Split(filepath.Base(filename), ".")[0]
}

func findTxtFiles(root string) []string {
	files := make([]string, 0)

	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".txt") {
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

// Text2Student makes a Student by given string input
func Text2Student(line string) *Student {
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
		students = append(students, Text2Student(scanner.Text()))
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
		// use '>' for descending order
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

// CourseSystem manges all SubjectManager
type CourseSystem map[string]*SubjectManager

// CourseSystemItem is key-value pair of Subject and SubjectManager
type CourseSystemItem struct {
	subject        string
	subjectManager *SubjectManager
}

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

// ImportStudents import students information from filepaths concurrently
func (cs *CourseSystem) ImportStudents(files []string) {
	type Job struct {
		subject string
		result  []*Student
	}
	jobs := make(chan Job, len(files))
	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, filename := range files {
		go func(filename string) {
			defer wg.Done()
			subject := extractSubject(filename)
			result := loadStudents(filename)
			jobs <- Job{subject, result}
		}(filename)
	}

	wg.Wait()
	close(jobs)

	for job := range jobs {
		sm := cs.GetSubjectManager(job.subject)
		sm.AppendStudents(job.result)
	}
}

// GetItem is a generator returns CourseSystemItem (Subject-SubjectManager) in ascending order
func (cs *CourseSystem) GetItem() <-chan CourseSystemItem {
	var subjects []string
	for subject := range *cs {
		subjects = append(subjects, subject)
	}
	sort.Strings(subjects)

	ch := make(chan CourseSystemItem)
	go func() {
		for _, subject := range subjects {
			ch <- CourseSystemItem{
				subject,
				(*cs)[subject],
			}
		}
		close(ch)
	}()

	return ch
}

func showBest4BySubject(subject string, sm *SubjectManager) string {
	ret := ""

	fmt.Printf("%s:\n", subject) // print this line for codesignal tests
	ret += fmt.Sprintf("%s:\n", subject)

	for _, student := range sm.GetBest4() {
		fmt.Println(student.FullName()) // print this line for codesignal tests
		ret += fmt.Sprintf("%s\n", student.FullName())
	}

	fmt.Println() // print this line for codesignal tests
	ret += fmt.Sprintf("\n")

	return ret
}

func selectionProcess(root string) string {
	files := findTxtFiles(root)
	cs := make(CourseSystem)

	cs.ImportStudents(files)

	ret := ""
	for item := range cs.GetItem() {
		ret += showBest4BySubject(
			item.subject,
			item.subjectManager,
		)
	}

	return ret[:len(ret)-1]
}

func main() {
	selectionProcess("/root/devops/")
}
