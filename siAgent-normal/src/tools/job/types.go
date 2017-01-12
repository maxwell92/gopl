package job

import (
	"tools/mylog"
	"tools/queue"
)

var log = mylog.Log

type Job struct {
	Id        int32  `json:"id,omitempty"`
	Image     string `json:"image"`
	Tag       string `json:"tag"`
	Status    string `json:"status,omitempty"`
	Expire    int    `json:"expire"`
	Op        string `json:"op"`
	CreatedAt string `json:"createdAt"`
}

type JobList struct {
	List []*Job `json:"jobList"`
}

type HistoryJobList struct {
	JobList *queue.Queue
}
