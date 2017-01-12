package job

func NewHistory(waitQueueLen int) *HistoryJobList{
	h := new(HistoryJobList)
	h.JobList = make([]*Job, 0)
	h.Index = 1
	h.Length = waitQueueLen + 2 // waitQueueLen * waiting + 1 * running + 1 * finished except 1 * empty
	h.Start = 1
	h.JobList = append(h.JobList, &Job{})


	return h
}



func (h *HistoryJobList) forward(step int) {
	for i := 0; i < step; i++ {
		h.JobList = append(h.JobList, &Job{})
	}

	h.JobList = h.JobList[1:]
}

func (h *HistoryJobList) Add(j *Job) {
	j.Id = h.Index
	h.Index ++

	if len(h.JobList) > h.Length - 2 {
		h.forward(1)
	}

	h.JobList = append(h.JobList, j)

}

func (h *HistoryJobList) List() []*Job {
	return h.JobList[1:]
}

