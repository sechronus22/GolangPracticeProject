package container

import (
	"mockproject/internals/config"
	"mockproject/internals/controllers"
	"mockproject/internals/infrastructure/http"
	"mockproject/internals/repositories"
	"mockproject/internals/services"

	"go.uber.org/dig"
)

type Containers struct {
	container *dig.Container
}

func (c *Containers) Configure() error {
	// c.container = dig.New()
	//Config
	if err := c.container.Provide(config.NewConfiguration); err != nil {
		return err
	}
	//infrastructure
	if err := c.container.Provide(http.NewHTTPServer); err != nil {
		return err
	}
	//controllers
	if err := c.container.Provide(controllers.NewServiceController); err != nil {
		return err
	}
	if err := c.container.Provide(controllers.NewPingController); err != nil{
		return err
	}
	if err := c.container.Provide(controllers.NewControllerGateway); err != nil {
		return err
	}
	//services
	if err := c.container.Provide(services.NewStudentService); err != nil {
		return err
	}
	if err := c.container.Provide(services.NewFacultyService); err != nil {
		return err
	}
	//Repositories
	if err := c.container.Provide(repositories.NewStudentRepository); err != nil {
		return err
	}
	if err := c.container.Provide(repositories.NewFacultyRepository); err != nil {
		return err
	}
	return nil
}

func (c *Containers) Start() error {
	if err := c.container.Invoke(func(s *http.Server) {
		s.Start()
	}); err != nil {
		return err
	}
	return nil
}

func NewContainers() (*Containers, error) {
	d := dig.New()
	container := &Containers{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
