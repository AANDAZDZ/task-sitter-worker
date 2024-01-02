package rpc

import (
	"context"
	proto2 "task-sitter-worker/sdk/rpc/proto"
	"task-sitter-worker/sdk/worker"
)

type ScheduleServer struct {
	proto2.UnimplementedScheduleServer
}

func (s *ScheduleServer) ScheduleTask(context.Context, *proto2.ScheduleTaskReq) (*proto2.ScheduleTaskResp, error) {
	worker.ScheduleTask()
	return nil, nil
}
