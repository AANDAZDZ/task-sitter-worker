package rpc

import (
	"context"
	"task-sitter-worker/sdk/model"
	proto2 "task-sitter-worker/sdk/rpc/proto"
	"task-sitter-worker/sdk/worker"
)

type ScheduleServer struct {
	proto2.UnimplementedScheduleServer
}

func (s *ScheduleServer) ScheduleTask(ctx context.Context, req *proto2.ScheduleTaskReq) (resp *proto2.ScheduleTaskResp, err error) {
	handleReq := &model.ScheduleTaskReq{
		TaskName:      req.TaskName,
		TaskContext:   req.TaskContext,
		BeginTime:     req.BeginTime,
		EndTime:       req.EndTime,
		RetryNum:      req.RetryNum,
		RetryInterval: req.RetryInterval,
		BizName:       req.BizName,
	}
	handleResp := worker.ScheduleTask(handleReq)
	resp.Message = handleResp.Message
	return resp, nil
}
