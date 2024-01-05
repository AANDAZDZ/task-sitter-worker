package model

type ScheduleTaskReq struct {
	TaskName      string
	TaskType      string
	TaskContext   string
	ScheduleType  int32
	BeginTime     int64
	EndTime       int64
	RetryNum      int32
	RetryInterval int64
	BizName       string
	RateInterval  int64
}

type ScheduleTaskResp struct {
	ErrNo   int
	Message string
}
