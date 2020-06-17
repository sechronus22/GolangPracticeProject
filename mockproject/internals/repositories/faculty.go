package repositories

import (
	"errors"
	"mockproject/internals/models"
	"fmt"
)

type FacultyRepository struct {
	FacultyDB []*models.Faculty
}

func NewFacultyRepository() *FacultyRepository {
	return &FacultyRepository{
		FacultyDB: []*models.Faculty{},
	}
}

func (fr *FacultyRepository) AddFaculty(f *models.Faculty) error{
	fr.FacultyDB = append(fr.FacultyDB,f)
	return nil
}

func (fr *FacultyRepository) AddDepartment(fid int,dep *models.Department) error{
	for i, e := range fr.FacultyDB{
		if e.Code == fid{
			fr.FacultyDB[i].Department = append(fr.FacultyDB[i].Department, dep)
			return nil
		}
	}
	return errors.New("code not found")
}

func (fr *FacultyRepository) GetFacultyList() ([]*models.Faculty,error){
	return fr.FacultyDB,nil
}

func (fr *FacultyRepository) DeleteFaculty(fid int) error{
	n := len(fr.FacultyDB)
	var pos int
	for i := 0; i < n; i++ {
		if fr.FacultyDB[i].Code == fid  {
			pos = i
			fr.FacultyDB = append(fr.FacultyDB[:pos], fr.FacultyDB[pos+1:]...)
			break
		}
		if i == n-1 {
			fmt.Println("There is no faculty with this Code")
		}
	}

	return nil
}
func (fr *FacultyRepository) DeleteDepartment(fid int,abb string) error{
	for i,e := range fr.FacultyDB{
		if e.Code == fid{
			for j,d := range e.Department{
				if d.Abbreviation == abb{
					fr.FacultyDB[i].Department = append(fr.FacultyDB[i].Department[:j],fr.FacultyDB[i].Department[:j+1]...)
				}
			}
		}
	}
	return nil
}