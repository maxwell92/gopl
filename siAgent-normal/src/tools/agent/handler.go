package agent

import (
	"encoding/json"
	//myerror "tools/error"
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

	/*
		if !History.Empty() {
			if History.Exist(j) {
				sc.Ye = myerror.NewYceError(myerror.ESYNC_EXIST, "")
				sc.WriteError()
			} else {
				waitQueue <- j
				History.Add(j)
				sc.WriteOk("")
			}
		} else {
			waitQueue <- j
			History.Add(j)
			sc.WriteOk("")
		}

	*/
	waitQueue <- j
	History.Add(j)
	sc.WriteOk("")

	return
}
