package controllers

import (
	"go.uber.org/dig"
)

type Controller struct {
	ControllerGateway *ControllerGateway
}

type ControllerGateway struct {
	dig.In
	ServiceController *ServiceController
	PingController	*PingController
}

func NewControllerGateway(cg ControllerGateway) *Controller {
	return &Controller{
		ControllerGateway: &cg,
	}
}
