package http

import (
	"mockproject/internals/config"
	"mockproject/internals/controllers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	route       *gin.Engine
	gatewayCtrl controllers.ControllerGateway
	env         config.Configuration
}

func (h *Server) Configure() {
	r := h.route
	r.GET("/ping", h.gatewayCtrl.PingController.GetPing)

	//Student
	s := r.Group("/student")
	{
		s.GET("/getStudent",h.gatewayCtrl.ServiceController.GetStudents)
		s.POST("/addStudent", h.gatewayCtrl.ServiceController.AddStudent)
		s.DELETE("/deleteStudent/:id", h.gatewayCtrl.ServiceController.DeleteStudent)
	}

	//Faculty
	f := r.Group("/faculty")
	{
		f.GET("/getFaculty",h.gatewayCtrl.ServiceController.GetFaculty)
		f.GET("/getInformation/:name",h.gatewayCtrl.ServiceController.GetInformation)
		f.POST("/addFaculty",h.gatewayCtrl.ServiceController.AddFaculty)
		f.DELETE("/deleteFaculty/:code", h.gatewayCtrl.ServiceController.DeleteFaculty)
		f.POST("/addDepartment/:code",h.gatewayCtrl.ServiceController.AddDepartment)
		f.DELETE("/deleteDepartment/:code/:abb",h.gatewayCtrl.ServiceController.DeleteDepartment)
	}

}

func (h *Server) Start() {
	h.Configure()
	if err := h.route.Run(":3000"); err != nil {
		panic(err)
	}
}

func NewHTTPServer(s controllers.ControllerGateway, env config.Configuration) *Server {
	return &Server{
		route:       gin.Default(),
		gatewayCtrl: s,
		env:         env,
	}
}
