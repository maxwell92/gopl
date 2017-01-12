package agent

import (
	"os"
	"os/exec"
	"strconv"
	"tools/job"

	"github.com/kataras/iris"
)

// waitQueue and History couldn't be too long. 10 is perfect
var waitQueue chan *job.Job
var History *job.HistoryJobList

func syncer(op *OperationAPI) {
	for {
		select {
		case j := <-waitQueue:
			{
				j.Expire, _ = strconv.Atoi(op.cfg.expire)
				j.Status = job.RUNNING
				History.Update(j, job.RUNNING)
				cmd := exec.Command("/bin/bash", "-C", op.cfg.script, j.Image, j.Tag)
				if cmd == nil {
					log.Errorf("Get Command Error")
					os.Exit(1)
				}

				go j.Killer(cmd)

				err := cmd.Run()
				if err != nil {
					j.Status = job.FAILED
					History.Update(j, job.FAILED)
				} else {
					j.Status = job.SUCC
					History.Update(j, job.SUCC)
				}
			}
		}
	}
}

func Init() *OperationAPI {
	op := new(OperationAPI)
	op.cfg = new(OperationConfig)
	op.sc = new(syncController)
	op.lc = new(listController)

	op.cfg.script = os.Getenv("SIAGENT_SCRIPT")
	op.cfg.waitQueueLen = os.Getenv("SIAGENT_WTQUEUE")
	op.cfg.expire = os.Getenv("SIAGENT_EXPIRE")

	f, err := os.Open(op.cfg.script)
	defer f.Close()
	if err != nil {
		log.Infof("Cann't open script file: err=%s", err)
		os.Exit(1)
	}

	if op.cfg.waitQueueLen == "" {
		// op.cfg.waitQueue = "100"
		op.cfg.waitQueueLen = "10"
	}

	if op.cfg.expire == "" {
		// 420 = 24 * 60 / 7
		// op.cfg.expire = "420"
		op.cfg.expire = "10"
	}

	waitQueueLen, _ := strconv.Atoi(op.cfg.waitQueueLen)
	waitQueue = make(chan *job.Job, waitQueueLen)

	History = job.NewHistory(waitQueueLen + 2)

	return op
}

func (op OperationAPI) SetupRouter() {

	iris.API("/api/v1/list", *op.lc)
	iris.API("/api/v1/sync", *op.sc)

	iris.Listen(":8081")
}

func Run() {
	op := Init()
	go syncer(op)
	op.SetupRouter()
}
