package main

import (
	"context"
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/db/dao"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/db/module"
	tags_service "github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service"
)

// TagsServiceImpl implements the last service interface defined in the IDL.
type TagsServiceImpl struct{}

// CreateTags implements the TagsServiceImpl interface.
func (s *TagsServiceImpl) CreateTags(ctx context.Context, req *tags_service.CreateTagsRequest) (resp *tags_service.CreateTagsResponse, err error) {
	// TODO: Your code here...
	dao := dao.GetTagsDAO()
	exist, err := dao.CheckTags(req.Pid, req.Tags)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("重复tags")
	}
	if err := dao.Insert(&module.Tags{
		Pid:  req.Pid,
		Tags: req.Tags,
	}); err != nil {
		return nil, err
	}

	return &tags_service.CreateTagsResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}

// DelTags implements the TagsServiceImpl interface.
func (s *TagsServiceImpl) DelTags(ctx context.Context, req *tags_service.DelTagsRequest) (resp *tags_service.DelTagsResponse, err error) {
	// TODO: Your code here...
	dao := dao.GetTagsDAO()
	if err := dao.DelTags(req.Pid, req.Tags); err != nil {
		return nil, fmt.Errorf("failed to delete tags: %w", err)
	}

	return &tags_service.DelTagsResponse{
		Status: 200,
		Msg:    "success",
	}, nil
}
