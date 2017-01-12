package job

import (
	"fmt"
	"tools/queue"
)
/*
func NewHistory(waitQueueLen int) *HistoryJobList {
	h := new(HistoryJobList)
	h.JobList = make([]*Job, 0)
	h.Index = 1
	h.Length = waitQueueLen + 2 // waitQueueLen * waiting + 1 * running + 1 * finished except 1 * empty
	h.Start = 1
	h.JobList = append(h.JobList, &Job{})
	h.Full = false

	return h
}
*/


/*
func (h *HistoryJobList) Update(j *Job, status string) {

	if h.Full {
		// h.JobList[h.Start].Status = status
		h.Add(j)
	} else {
		h.JobList[j.Id].Status = status
	}

	h.Watch()
}

func (h *HistoryJobList) forward(step int, j *Job) {
	for i := 0; i < step; i++ {
		h.JobList = append(h.JobList, j)
	}

	h.JobList = h.JobList[1:]
}

func (h *HistoryJobList) Add(j *Job) {
	j.Id = h.Index
	h.Index++

	// fmt.Printf("j.Id: %d, image: %s\n", j.Id, j.Image+":"+j.Tag)

	// fmt.Printf("history len: %d, length: %d\n", len(h.JobList), h.Length)
	if len(h.JobList) > h.Length-2 {
		h.forward(1, j)
		h.Full = true
	} else {
		h.JobList = append(h.JobList, j)
	}

	// h.Watch()
}

func (h *HistoryJobList) List() []*Job {
	return h.JobList[1:]
}

func (h *HistoryJobList) Watch() {
	fmt.Printf("----------- %d ----------\n", len(h.JobList) - 1)

	for _, v := range h.JobList[1:] {
		fmt.Printf("[%d] %s, %s\n", v.Id, v.Status, v.Image+":"+v.Tag)
	}

	fmt.Println("----------------------")
}
*/



func NewHistory(waitQueueLen int) *HistoryJobList {
	return &HistoryJobList{
		JobList: queue.New(waitQueueLen),
	}
}

func (h *HistoryJobList) Add(j *Job) {
	//fmt.Printf("Add: %d, %s\n", h.JobList.Bottom, j.Image+":"+j.Tag)
	j.Id = int32(h.JobList.Bottom % h.JobList.Maxlen)
	h.JobList.EnQueue(j)
}

func (h *HistoryJobList) List() []interface{} {
	return h.JobList.Elems
}

func (h *HistoryJobList) Update(j *Job, status string) {
	//fmt.Printf("Update: %d, %s\n", j.Id, j.Image+":"+j.Tag)
	jo := h.JobList.Elems[j.Id].(*Job)
	jo.Status = status

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
