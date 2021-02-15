package model

type Task struct {
	StartTime      int64 `json:"start_time"`
	StartTimeStr   string `json:"start_time_str"`
	EndTime        int64  `json:"end_time"`
	EndTimeStr     string `json:"end_time_str"`
	Desc           string `json:"desc"`
	Tags            []string `json:"tags"`
	TaskCompletion int `json:"task_completion"` // [0, 10]
}
type Journal struct {
	DateUnixTime int64  `json:"date_unix_time"`
	DateStr      string  `json:"date_str"`
	Weather      string `json:"weather"`
	Tasks        []*Task `json:"tasks"`
	Reflections  string `json:"reflections"`
}
