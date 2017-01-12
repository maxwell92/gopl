package agent

import (
	"tools/controller"
	"tools/mylog"
)

var log = mylog.Log

type Job struct {
	Id        int    `json:"id,omitempty"`
	Image     string `json:"image"`
	Tag       string `json:"tag"`
	Status    string `json:"status,omitempty"`
	Expire    int    `json:"expire"`
	Op        string `json:"op"`
	CreatedAt string `json:"createdAt"`
}

type JobList struct {
	jobList []*Job `json:"jobList"`
}

type HistoryJobList struct {
	jobList []*Job
	index int
	finished int
}

type OperationConfig struct {
	script       string
	waitQueueLen string
	expire       string
}

type syncController struct {
	controller.Controller
}

type listController struct {
	controller.Controller
}

type OperationAPI struct {
	sc  *syncController
	lc  *listController
	cfg *OperationConfig
	Job *Job `json:"job"`
	//History *HistoryJobList
}

type clientPage struct {
	Title string
	Host  string
}
