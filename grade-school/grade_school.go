// Package school implements a solution of the exercise titled `Grade School'.
package school

import "sort"

// Grade is an entry of the roll per grade.
type Grade struct {
	grade    int
	students []string
}

// School is a student directory.
type School []Grade

// New is the ctor of School
func New() *School {
	return &School{}

}

// Add enrolls a student to the grade.
func (s *School) Add(student string, grade int) {
	for i := range *s {
		if (*s)[i].grade == grade {
			(*s)[i].students = append((*s)[i].students, student)
			return
		}
	}
	*s = append(*s, Grade{grade, []string{student}})
}

// Grade returns names of the grade enrolled.
func (s *School) Grade(grade int) []string {
	for i := range *s {
		if (*s)[i].grade == grade {
			return (*s)[i].students
		}
	}
	return nil
}

// Enrollment returns the student directory.
func (s *School) Enrollment() []Grade {
	for i := range *s {
		sort.Slice((*s)[i].students, func(j, k int) bool { return (*s)[i].students[j] < (*s)[i].students[k] })
	}
	sort.Slice(*s, func(i, j int) bool { return (*s)[i].grade < (*s)[j].grade })
	return *s
}

/*
BenchmarkAddStudents-4   	   41923	     25217 ns/op	   19080 B/op	      67 allocs/op
BenchmarkEnrollment-4    	   51295	     22903 ns/op	     688 B/op	      21 allocs/op
*/
