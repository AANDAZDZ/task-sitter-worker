package model

type ScheduleTaskReq struct {
	TaskName      string
	TaskType      string
	TaskContext   string
	BeginTime     int64
	EndTime       int64
	RetryNum      int32
	RetryInterval int64
	BizName       string
}

type ScheduleTaskResp struct {
	Message string
}
