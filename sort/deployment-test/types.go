package main

import (
	"sorter"
)

type Controller struct {
	*sorter.Sorter
}

type ListDeploymentController struct {
	Controller
	UserName       func(sorter.IResource, sorter.IResource) bool
	DeploymentName func(sorter.IResource, sorter.IResource) bool
	DcId           func(sorter.IResource, sorter.IResource) bool
	UpdateTime     func(sorter.IResource, sorter.IResource) bool
}

/*
func (ldc *ListDeploymentController) OrderedBy(less ...sorter.LessFunc) *sorter.Sorter {
	return &sorter.Sorter{
		Lessf: less,
	}
}
*/

func (c *Controller) OrderedBy(less ...sorter.LessFunc) *sorter.Sorter {
	return &sorter.Sorter{
		Lessf: less,
	}
}

type Deployment struct {
	userName       string
	dcId           int
	deploymentName string
	updateTime     string
}

type Deployments struct {
	deployments []Deployment
}

func (d Deployment) Print() {

}
