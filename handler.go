package main

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/db/dao"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/db/model"
	tags_service "github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service"
)

//todo 用户输入传入的是pid和标签string。然后string到标签库中查询。如果存在，返回对应的tid，不存在，则创建，然后返回tid。然后在pidtid表中插入对应的pid和tid，在传出来tid

// TagsServiceImpl implements the last service interface defined in the IDL.
type TagsServiceImpl struct{}

// PidTidCreate implements the TagsServiceImpl interface.
func (s *TagsServiceImpl) PidTidCreate(ctx context.Context, req *tags_service.PidTidCreateRequest) (resp *tags_service.PidTidCreateResponse, err error) {
	// TODO: Your code here...
	PidTidDAO := dao.GetTidPidDao()
	TagDAO := dao.GetTagDao()
	//在tag表中检查tag是否存在
	judgment, err := TagDAO.GetTidByTag(req.Tag)
	if err != nil {
		return nil, err
	} else if judgment == -1 {
		err = TagDAO.InsertByTag(&model.Tag{})
		if err != nil {
			return nil, err
		}
	}
	err = PidTidDAO.InsertByPidTid(&model.PidTid{})
	if err != nil {
		return nil, err
	}
	return &tags_service.PidTidCreateResponse{
		Status: 200,
		Msg:    "success",
		Pid:    req.Pid,
		Tid:    req.Pid,
	}, nil
}

// GetTag implements the TagsServiceImpl interface.
func (s *TagsServiceImpl) GetTag(ctx context.Context, req *tags_service.GetTagRequest) (resp *tags_service.GetTagResponse, err error) {
	// TODO: Your code here...

	TagDAO := dao.GetTagDao()
	tag, err := TagDAO.GetTagByTid(req.Tid)
	if err != nil {
		return nil, err
	}
	return &tags_service.GetTagResponse{
		Status: 200,
		Msg:    "success",
		Tag:    tag,
	}, nil

}

// GetTagIDsWithPid implements the TagsServiceImpl interface.
func (s *TagsServiceImpl) GetTagIDsWithPid(ctx context.Context, req *tags_service.GetTagIDRequest) (resp *tags_service.GetTagIDResponse, err error) {
	// TODO: Your code here...
	PidTidDAO := dao.GetTidPidDao()
	getTid, err := PidTidDAO.GetTidByPid(req.Pid)
	if err != nil {
		return nil, err
	}
	return &tags_service.GetTagIDResponse{
		Status: 200,
		Msg:    "success",
		Tids:   getTid,
	}, nil
}
