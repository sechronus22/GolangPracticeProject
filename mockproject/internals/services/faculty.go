package services

import (
	"mockproject/internals/models"
	"mockproject/internals/repositories"
)

type FacultyService struct {
	FacultyRepository *repositories.FacultyRepository
}

func NewFacultyService() *FacultyService {
	return &FacultyService{
		FacultyRepository: repositories.NewFacultyRepository(),
	}
}

func (fs *FacultyService) ListFaculty()([]*models.Faculty,error){
	facultyList,err := fs.FacultyRepository.GetFacultyList()
	if err != nil {
		return []*models.Faculty{}, err
	}
	return facultyList,nil
}

func (fs *FacultyService) AddFaculty(f *models.Faculty) error{
	err := fs.FacultyRepository.AddFaculty(f)
	if err != nil{
		return err
	}
	return nil
}

func (fs *FacultyService) DeleteFaculty(fid int) error{
	err := fs.FacultyRepository.DeleteFaculty(fid)
	if err != nil{
		return err
	}
	return nil

}

func (fs *FacultyService) AddDepartment(fid int,dep *models.Department) error  {
	err := fs.FacultyRepository.AddDepartment(fid, dep)
	if err != nil{
		return err
	}
	return nil
}

func (fs *FacultyService) DeleteDepartment(fid int,abb string) error{
	err := fs.FacultyRepository.DeleteDepartment(fid,abb)
	if err != nil{
		return err
	}
	return nil
}

func (fs *FacultyService) GetInformation(fName string) (string,error){
	info,err := fs.FacultyRepository.GetInformation(fName)
	if err != nil{
		return "",err
	}
	return info,nil
}