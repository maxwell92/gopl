package agent

import (
	"tools/controller"
	"tools/mylog"
)

var log = mylog.Log

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
}
