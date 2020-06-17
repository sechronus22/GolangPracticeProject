package services

import (
	//"mockproject/internals/models"
	"mockproject/internals/models"
	"mockproject/internals/repositories"
)

type StudentService struct {
	StudentRepository *repositories.StudentRepository
}

func NewStudentService() *StudentService {
	return &StudentService{
		StudentRepository: repositories.NewStudentRepository(),
	}
}

func (ss *StudentService) ListStudent() ([]*models.Student, error) {
	studentList, err := ss.StudentRepository.GetStudentList()
	if err != nil {
		return []*models.Student{}, err
	}
	return studentList, nil
}

func (ss *StudentService) AddStudent(s *models.Student) error {
	err := ss.StudentRepository.AddStudent(s)
	if err != nil {
		return err
	}
	return nil
}

func (ss *StudentService) DeleteStudent(sid int) error {
	err := ss.StudentRepository.DeleteStudent(sid)
	if err != nil {
		return err
	}
	return nil
}
