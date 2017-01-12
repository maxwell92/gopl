package agent

import (
	"encoding/json"
	"tools/job"
)

func (lc listController) Get() {
	jobs := new(job.JobList)
	jobs.List = History.List()

	listJSON, err := json.Marshal(jobs.List)
	if err != nil {
		log.Errorf("Marshal Json Error: %s", err)
	}

	lc.WriteOk(string(listJSON))
	return
}

func (sc syncController) Post() {
	j := new(job.Job)
	err := sc.ReadJSON(j)
	if err != nil {
		log.Errorf("ReadJSON Error: err=%s", err)
		//TODO: temporarily exit
		return
	}
	waitQueue <- j
	j.Status = job.WAITING
	History.Add(j)
	sc.WriteOk("")
	return
}
