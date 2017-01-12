package agent

import (
	"encoding/json"
	"tools/job"
)


func (lc listController) Get() {
	jobs := History.List()

	list := jobs

	listJSON, err := json.Marshal(list)
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

	j.Status = job.WAITING
	waitQueue <- j
	History.Add(j)
	sc.WriteOk("")
	return
}
