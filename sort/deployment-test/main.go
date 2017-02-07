package main

import (
	"fmt"
	"sorter"
)

func main() {
	data := new(Deployments)
	data.deployments = []Deployment{
		{"maxwell", 1, "nginx", "20170207"},
		{"maxwell", 2, "busybox", "20170209"},
		{"sheep", 1, "busybox", "20170206"},
		{"magic", 1, "yce", "20170301"},
	}

	lc := &ListDeploymentController{
		UserName: func(d1, d2 sorter.IResource) bool {
			return d1.(Deployment).userName < d2.(Deployment).userName
		},
		DeploymentName: func(d1, d2 sorter.IResource) bool {
			return d1.(Deployment).deploymentName < d2.(Deployment).deploymentName
		},
		DcId: func(d1, d2 sorter.IResource) bool {
			return d1.(Deployment).dcId < d2.(Deployment).dcId
		},
		UpdateTime: func(d1, d2 sorter.IResource) bool {
			return d1.(Deployment).updateTime > d2.(Deployment).updateTime
		},
	}

	var sr []sorter.IResource = make([]sorter.IResource, len(data.deployments))

	for i, d := range data.deployments {
		sr[i] = d
	}

	lc.OrderedBy(lc.UserName).Sort(sr)
	fmt.Println("by userName:", sr)

	lc.OrderedBy(lc.DcId).Sort(sr)
	fmt.Println("by dcId:", sr)

	lc.OrderedBy(lc.UserName, lc.DeploymentName).Sort(sr)
	fmt.Println("by userName, deploymentName:", sr)

	lc.OrderedBy(lc.UpdateTime, lc.DeploymentName).Sort(sr)
	fmt.Println("by updateTime, deploymentName:", sr)

}
