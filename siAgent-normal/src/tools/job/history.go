package job

import (
	"fmt"
	"tools/queue"
)

func NewHistory(waitQueueLen int) *HistoryJobList {
	return &HistoryJobList{
		JobList: queue.New(waitQueueLen),
	}
}

func (h *HistoryJobList) Add(j *Job) {
	j.Id = int32(h.JobList.Bottom % h.JobList.Maxlen)
	h.JobList.EnQueue(j)
}

func (h *HistoryJobList) List() []interface{} {
	return h.JobList.Elems
}

func (h *HistoryJobList) Update(j *Job, status string) {
	jo := h.JobList.Elems[j.Id].(*Job)
	jo.Status = status

	// for test
	h.Watch()
}

func (h *HistoryJobList) Watch() {
	var start, end int
	if h.JobList.Full {
		start, end = h.JobList.Top, h.JobList.Maxlen
	} else {
		start, end = h.JobList.Top, h.JobList.Bottom
	}


	fmt.Printf("----------- %d ----------\n", len(h.JobList.Elems))
	for i := start; i < end; i++ {
		fmt.Println(i)
		jo := h.JobList.Elems[i].(*Job)

		fmt.Printf("[%d] %s, %s\n", jo.Id, jo.Status, jo.Image+":"+jo.Tag)

	}
	fmt.Println("----------------------")
}
