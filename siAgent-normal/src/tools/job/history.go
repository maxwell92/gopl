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
	var queueLen int
	if h.JobList.Full {
		queueLen = h.JobList.Maxlen
	} else {
		queueLen = h.JobList.Bottom - h.JobList.Top
	}

	log.Infof("HistoryJobList: len(jobList)=%d", queueLen)

	return h.JobList.Elems
}

func (h *HistoryJobList) Update(j *Job, status string) {
	jo := h.JobList.Elems[j.Id].(*Job)
	oldStatus := jo.Status
	jo.Status = status

	// for test
	// h.Watch()

	log.Infof("Job %s:%s change from %s to %s", j.Image, j.Tag, oldStatus, j.Status)
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

/*
func (h *HistoryJobList) Exist(j *Job) bool {
	for i := 0; i < h.JobList.Maxlen; i++ {
		job := h.JobList.Elems[i].(*Job)
		if j.Image == job.Image && j.Tag == job.Tag {
			return true
		}
	}

	return false
}
*/

func (h *HistoryJobList) Exist(j *Job) bool {
	if h.JobList.Full {
		for i := h.JobList.Top; i < h.JobList.Maxlen; i++ {
			job := h.JobList.Elems[i].(*Job)
			if j.Image == job.Image && j.Tag == job.Tag {
				return true
			}
		}
		return false
	} else {
		for i := h.JobList.Top; i < h.JobList.Bottom; i++ {
			job := h.JobList.Elems[i].(*Job)
			if j.Image == job.Image && j.Tag == job.Tag {
				return true
			}
		}
		return false
	}
}

func (h *HistoryJobList) Empty() bool {
	if h.JobList.Top == h.JobList.Bottom && !h.JobList.Full {
		return true
	} else {
		return false
	}
}
