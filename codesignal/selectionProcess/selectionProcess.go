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
			return err
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

	students := make([]*Student, 0)

	scanner := bufio.NewScanner(file)

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

// Job contains parsed subject and students' information
type Job struct {
	subject  string
	students []*Student
}

func asyncLoadTranscript(filename string) <-chan Job {
	ch := make(chan Job)
	go func() {
		subject := extractSubject(filename)
		students := loadStudents(filename)
		ch <- Job{subject, students}
		close(ch)
	}()
	return ch
}

func mergeJobs(jobs ...<-chan Job) <-chan Job {
	// a Fan-in pattern
	var wg sync.WaitGroup
	out := make(chan Job)

	output := func(j <-chan Job) {
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

// ImportStudents import students information from filepaths concurrently
func (cs *CourseSystem) ImportStudents(files []string) {
	jobChans := make([]<-chan Job, 0, len(files))

	for _, filename := range files {
		jobChans = append(jobChans, asyncLoadTranscript(filename))
	}

	for job := range mergeJobs(jobChans...) {
		sm := cs.GetSubjectManager(job.subject)
		sm.AppendStudents(job.students)
	}
}

// GetItem is a generator returns CourseSystemItem (Subject-SubjectManager) in ascending order
func (cs *CourseSystem) GetItem() <-chan CourseSystemItem {
	subjects := make([]string, 0, len(*cs))
	for subject := range *cs {
		subjects = append(subjects, subject)
	}
	sort.Strings(subjects)

	ch := make(chan CourseSystemItem, len(subjects)) // buffered channel is OK
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
