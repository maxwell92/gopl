package job

import (
	"encoding/json"
	"os"
	"os/exec"
	"time"
)

func (j Job) Encode() string {
	jobBYTE, err := json.Marshal(j)
	if err != nil {
		log.Infof("Job Encoding Error: err=%s", err)
		os.Exit(1)
	}
	jobSTR := string(jobBYTE)
	return jobSTR
}

func (j Job) Killer(c *exec.Cmd) {
	k := time.NewTimer(time.Duration(j.Expire) * time.Second)
	<-k.C
	c.Process.Kill()
}
