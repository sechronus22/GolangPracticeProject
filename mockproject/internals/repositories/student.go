package repositories

import (
	"fmt"
	"mockproject/internals/models"
)

type StudentRepository struct {
	StudentDB []*models.Student
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		StudentDB: []*models.Student{},
	}
}

func (sr *StudentRepository) AddStudent(student *models.Student) error {
	sr.StudentDB = append(sr.StudentDB, student)
	return nil
}

func (sr *StudentRepository) DeleteStudent(sid int) error {
	n := len(sr.StudentDB)
	var pos int
	for i := 0; i < n; i++ {
		if sr.StudentDB[i].Id == sid {
			pos = i
			sr.StudentDB = append(sr.StudentDB[:pos], sr.StudentDB[pos+1:]...)
			break
		}
		if i == n-1 {
			fmt.Println("There is no student with this ID")
		}
	}

	return nil
}

func (sr *StudentRepository) GetStudentList() ([]*models.Student, error) {
	return sr.StudentDB, nil
}
