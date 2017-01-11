package agent

import (
	"tools/mylog"
	"tools/controller"
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

type OperationConfig struct {
	script    string
	waitQueue string
	expire    string
}

type syncController struct {
	controller.Controller
}

type listController struct {
	controller.Controller
}

type OperationAPI struct {
	sc *syncController
	lc *listController
	cfg *OperationConfig
	Job *Job `json:"job"`
}

type clientPage struct {
	Title string
	Host  string
}


