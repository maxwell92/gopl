package agent

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

var waitQueue chan *Job
var History *HistoryJobList

var runFlag bool

func (j Job) Encode() string {
	jobBYTE, err := json.Marshal(j)
	if err != nil {
		log.Infof("Job Encoding Error: err=%s", err)
		os.Exit(1)
	}
	jobSTR := string(jobBYTE)
	return jobSTR
}

func (h *HistoryJobList) Clean() {
	/*
	var split int
	if judge := len(h.jobList) % 2; judge == 0 {
		split = h.index / 2
	} else {
		split = (h.index - 1) / 2
	}

	for i := 0; i < split; i++ {
		h.jobList = append(h.jobList, &Job{})
	}

	h.jobList = h.jobList[split:]
	h.index = len(h.jobList) - split
	*/

	split := h.finished - 2
	for i := 1; i < split; i++ {
		h.jobList = append(h.jobList, &Job{})
	}

	h.jobList = h.jobList[split:]
	h.index = len(h.jobList) - split


}

func gc(history *HistoryJobList) {
	t := time.NewTimer(time.Duration(30) * time.Second)
	for {
		select {
		case <-t.C:
			{
				if len(history.jobList) == history.index {
					history.Clean()
					log.Infof("finQueue gc has removed half of the queue")
				}
			}
		}
	}

}

func syncer(op *OperationAPI) {
	for {
		select {
		case job := <-waitQueue:
			{
				job.Expire, _ = strconv.Atoi(op.cfg.expire)
				History.jobList[job.Id].Status = RUNNING
				job.Status = RUNNING
				log.Infof("%s: Waiting --> Running", job.Image+":"+job.Tag)
				cmd := exec.Command("/bin/bash", "-C", op.cfg.script, job.Image, job.Tag)
				if cmd == nil {
					log.Infof("Get Command Error")
					os.Exit(1)
				}

				go job.killer(cmd)

				runFlag = true
				err := cmd.Run()
				runFlag = false
				if err != nil {
					job.Status = FAILED
					History.jobList[job.Id].Status = FAILED
					History.finished = job.Id
					log.Infof("%s: Running --> Killed", job.Image+":"+job.Tag)
				} else {
					job.Status = SUCC
					History.jobList[job.Id].Status = SUCC
					History.finished = job.Id
					log.Infof("%s: Running --> Finish", job.Image+":"+job.Tag)
				}
			}
		}
	}
}

func (j Job) killer(c *exec.Cmd) {
	k := time.NewTimer(time.Duration(j.Expire) * time.Second)
	<-k.C
	c.Process.Kill()
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
		// op.cfg.expire = "420"
		op.cfg.expire = "10"
	}

	waitQueueLen, _ := strconv.Atoi(op.cfg.waitQueueLen)
	waitQueue = make(chan *Job, waitQueueLen)

	History = new(HistoryJobList)
	History.jobList = make([]*Job, waitQueueLen)
	History.index = 1
	History.finished = 1
	go gc(History)

	return op
}

func (lc listController) Get() {
	jobs := new(JobList)
	jobs.jobList = History.jobList[1:History.finished + 1]

	listJSON, err := json.Marshal(jobs.jobList)
	if err != nil {
		log.Errorf("Get")
	}

	lc.WriteOk(string(listJSON))
	return
}

func (sc syncController) Post() {
	job := new(Job)
	err := sc.ReadJSON(job)
	if err != nil {
		log.Errorf("ReadJSON Error: err=%s", err)
		//TODO: temporarily exit
		return
	}
	waitQueue <- job
	job.Id = History.index
	job.Status = WAITING
	History.jobList[History.index] = job
	History.finished = History.index
	History.index++
	sc.WriteOk("")
	return
}

func (op OperationAPI) SetupRouter() {
	iris.Static("js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	iris.API("/api/v1/list", *op.lc)
	iris.API("/api/v1/sync", *op.sc)

	iris.Listen(":8081")
}

func Run() {
	op := Init()
	go syncer(op)
	op.SetupRouter()
}
