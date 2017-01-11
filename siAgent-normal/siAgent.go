package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"tools/mylog"
	"tools/controller"

	"github.com/kataras/iris"
)

var log = mylog.Log

const (
	SUCC    = "SUCCEED"
	FAILED  = "FAILED"
	EXPIRED = "EXPIRED"
	RUNNING = "RUNNING"
	WAITING = "WAITING"
)

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

type OperationAPI struct {
	*controller.Controller
	cfg *OperationConfig
	Job *Job `json:"job"`
}

type clientPage struct {
	Title string
	Host  string
}

var usage = `Usage: siAgent <filepath>
`

var waitJobQueue chan *Job
var writeJobQueue chan *Job
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

func syncer(script, expire string) {
	for {
		select {
		case job := <-waitJobQueue:
			{
				job.Expire, _ = strconv.Atoi(expire)
				job.Status = RUNNING
				log.Infof("%s: Waiting --> Running", job.Image+":"+job.Tag)
				cmd := exec.Command("/bin/bash", "-C", script, job.Image, job.Tag)
				if cmd == nil {
					log.Infof("Get Command Error")
					os.Exit(1)
				}

				go job.killer(cmd)

				writeJobQueue <- job
				runFlag = true
				err := cmd.Run()
				runFlag = false
				if err != nil {
					job.Status = FAILED
					log.Infof("%s: Running --> Killed", job.Image+":"+job.Tag)
					writeJobQueue <- job
				} else {
					job.Status = SUCC
					log.Infof("%s: Running --> Finish", job.Image+":"+job.Tag)
					writeJobQueue <- job
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

func UsageAndExit() {
	if len(os.Args) != 1 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

}

func Init() *OperationAPI {
	op := new(OperationAPI)
	op.cfg = new(OperationConfig)

	op.cfg.script = os.Getenv("SIAGENT_SCRIPT")
	op.cfg.waitQueue = os.Getenv("SIAGENT_WTQUEUE")
	op.cfg.expire = os.Getenv("SIAGENT_EXPIRE")

	f, err := os.Open(op.cfg.script)
	defer f.Close()
	if err != nil {
		log.Infof("Cann't open script file: err=%s", err)
		os.Exit(1)
	}

	if op.cfg.waitQueue == "" {
		// op.cfg.waitQueue = "100"
		op.cfg.waitQueue = "10"
	}

	if op.cfg.expire == "" {
		// op.cfg.expire = "420"
		op.cfg.expire = "10"
	}

	waitJobLen, _ := strconv.Atoi(op.cfg.waitQueue)
	waitJobQueue = make(chan *Job, waitJobLen)
	writeJobLen := waitJobLen
	writeJobQueue = make(chan *Job, writeJobLen)

	return op
}

func (op OperationAPI) Get() {
	jobList := make([]Job, 0)

	for idx, job := range waitJobQueue {
		jobList := append(jobList, job)
	}

	json, err := json.Marshal(jobList)
	if err != nil {
		log.Errorf("Get")
	}

	op.Write()
}

func (op OperationAPI) Post() {
	job := new(Job)
		err := op.ReadJSON(job)
		if err != nil {
			log.Errorf("ReadJSON Error: err=%s", err)
			//TODO: temporarily exit
			os.Exit(1)
		}
		waitJobQueue <- job
		op.WriteOk("操作成功")
	}

func (op OperationAPI) SetupRouter() {

	iris.Static("js", "./static/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("client.html", clientPage{"Client Page", ctx.HostString()})
	})

	iris.Get("/api/v1/list", op.Get)
	iris.Post("/api/v1/sync", op.Post)

	iris.Listen(":8081")
}

func main() {
	UsageAndExit()
	op := Init()
	go syncer(op.cfg.script, op.cfg.expire)
	op.SetupRouter()
}
