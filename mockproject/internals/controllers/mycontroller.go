package controllers

import (
	"mockproject/internals/models"
	"mockproject/internals/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceController struct {
	StudentService *services.StudentService
	FacultyService *services.FacultyService
}

func NewServiceController(
	s *services.StudentService,
	f *services.FacultyService) *ServiceController {
	return &ServiceController{
		StudentService: s,
		FacultyService: f,
	}
}

//Student
func (sc *ServiceController) GetStudents(c *gin.Context) {
	students, err := sc.StudentService.ListStudent()
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, students)
}
func (sc *ServiceController) AddStudent(c *gin.Context) {
	student := &models.Student{}

	if err := c.BindJSON(student); err != nil {
		c.AbortWithError(500, err)
		return
	}
	err := sc.StudentService.AddStudent(student)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, student)
}
func (sc ServiceController) DeleteStudent(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))

	if  err != nil {
		c.AbortWithError(500, err)
		return
	}
	if err := sc.StudentService.DeleteStudent(id);err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, id)
}

//Faculty
func (sc ServiceController) GetFaculty(c *gin.Context){
	faculty,err := sc.FacultyService.ListFaculty()
	if err != nil{
		c.AbortWithError(500, err)
	}
	c.JSON(200,faculty)
}
func (sc ServiceController) GetInformation(c *gin.Context){
	fName := c.Param("name")
	if info,err := sc.FacultyService.GetInformation(fName); err != nil{
		c.AbortWithError(500,err)
	}
	c.JSON(200,info)
}
func (sc ServiceController) AddFaculty(c *gin.Context){
	faculty := &models.Faculty{}
	if err :=c.BindJSON(faculty); err != nil{
		c.AbortWithError(500,err)
	}
	if err := sc.FacultyService.AddFaculty(faculty); err != nil{
		c.AbortWithError(500,err)
	}
	c.JSON(200, faculty)
}
func (sc ServiceController) DeleteFaculty(c *gin.Context){
	code,err := strconv.Atoi(c.Param("code"))
	if err != nil{
		c.AbortWithError(500,err)
	}
	if err := sc.FacultyService.DeleteFaculty(code);err != nil{
		c.AbortWithError(500,err)
	}
	c.JSON(200,code)
}
func (sc ServiceController) AddDepartment(c *gin.Context){
	code,err := strconv.Atoi(c.Param("code"))
	department := &models.Department{}
	if err != nil{
		c.AbortWithError(500,err)
	}
	if err := c.BindJSON(department);err != nil{
		c.AbortWithError(500,err)
	}
	if err := sc.FacultyService.AddDepartment(code,department);err != nil{
		c.AbortWithError(500,err)
	}
	c.JSON(200,code)
}
func (sc ServiceController) DeleteDepartment(c *gin.Context){
	code,err := strconv.Atoi(c.Param("code"))
	if err != nil{
		c.AbortWithError(500,err)
	}
	abb := c.Param("abb")

	if err := sc.FacultyService.DeleteDepartment(code ,abb);err != nil{
		c.AbortWithError(500,err)
	}
	c.JSON(200,code)
}


//Ping
type PingController struct{}

func (PingController) GetPing(c *gin.Context) {
	c.String(200, "pong")
}

func NewPingController() *PingController {
	return &PingController{}
}