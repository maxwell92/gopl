package main

import (
	"fmt"
	"github.com/kataras/iris"
	"sort"
)

type lessFunc func(d1, d2 *Deployment) bool

type deploymentSorter struct {
	deployments []Deployment
	less        []lessFunc
}

func (ds *deploymentSorter) Sort(deployments []Deployment) {
	ds.deployments = deployments
	sort.Sort(ds)
}

func (ds *deploymentSorter) Len() int {
	return len(ds.deployments)
}

func (ds *deploymentSorter) Swap(i, j int) {
	ds.deployments[i], ds.deployments[j] = ds.deployments[j], ds.deployments[i]
}

func (ds *deploymentSorter) Less(i, j int) bool {
	p, q := &ds.deployments[i], &ds.deployments[j]

	var k int
	// why don't len(ds.less) ?
	for k = 0; k < len(ds.less)-1; k++ {
		less := ds.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}

	return ds.less[k](p, q)
}

type ListDeploymentController struct {
	*deploymentSorter
	*iris.Context
}

func (ldc *ListDeploymentController) OrderedBy(less ...lessFunc) *deploymentSorter {
	return &deploymentSorter{
		less: less,
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

func (lc *ListDeploymentController) GetDeploymentListAndSort() {
	data := new(Deployments)
	data.deployments = []Deployment{
		{"maxwell", 1, "nginx", "20170207"},
		{"maxwell", 2, "busybox", "20170209"},
		{"sheep", 1, "busybox", "20170206"},
		{"magic", 1, "yce", "20170301"},
	}

	userName := func(d1, d2 *Deployment) bool {
		return d1.userName < d2.userName
	}

	deploymentName := func(d1, d2 *Deployment) bool {
		return d1.deploymentName < d2.deploymentName
	}

	dcId := func(d1, d2 *Deployment) bool {
		return d1.dcId < d2.dcId
	}

	updateTime := func(d1, d2 *Deployment) bool {
		return d1.updateTime > d2.updateTime
	}

	//lc := new(ListDeploymentController)
	lc.OrderedBy(userName).Sort(data.deployments)
	fmt.Println("by userName:", data.deployments)

	lc.OrderedBy(dcId).Sort(data.deployments)
	fmt.Println("by dcId:", data.deployments)

	lc.OrderedBy(userName, deploymentName).Sort(data.deployments)
	fmt.Println("by userName, deploymentName:", data.deployments)

	lc.OrderedBy(updateTime, deploymentName).Sort(data.deployments)
	fmt.Println("by updateTime, deploymentName:", data.deployments)

}

func (lc ListDeploymentController) ValidateSession() {
	fmt.Println(lc.RequestHeader("Authorization"))
}

/*
type DeploymentData struct {
	UserId  int32  `json:"userId"`
	OrgId   int32  `json:"orgId"`
	OrgName string `json:"orgName"`
}
*/

func (lc ListDeploymentController) GetNamespace() {
	/*
			data := new(DeploymentData)
				err := lc.ReadJSON(data)
				if err != nil {
					fmt.Println(err)
				}
		fmt.Printf("%v\n", data)
	*/

	fmt.Println(lc.RequestHeader("userId"))
	fmt.Println(lc.RequestHeader("orgId"))
	fmt.Println(lc.RequestHeader("orgName"))
}

func (lc ListDeploymentController) Get() {
	lc.ValidateSession()
	lc.GetNamespace()
	lc.GetDeploymentListAndSort()
}

func main() {
	lc := new(ListDeploymentController)
	iris.API("/api/v2/deployments", *lc)
	iris.Listen(":8081")
}

// test curl:
// curl -XGET -H "Authorization":"abcdefg-1234567" -H "userId":"1" -H "orgId":"2" -H "orgName":"dev" localhost:8081/api/v2/deployments
