// Code generated by Kitex v0.12.1. DO NOT EDIT.

package tagsservice

import (
	"context"
	"errors"
	tags_service "github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PidTidCreate": kitex.NewMethodInfo(
		pidTidCreateHandler,
		newPidTidCreateArgs,
		newPidTidCreateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetTag": kitex.NewMethodInfo(
		getTagHandler,
		newGetTagArgs,
		newGetTagResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetTagIDsWithPid": kitex.NewMethodInfo(
		getTagIDsWithPidHandler,
		newGetTagIDsWithPidArgs,
		newGetTagIDsWithPidResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	tagsServiceServiceInfo                = NewServiceInfo()
	tagsServiceServiceInfoForClient       = NewServiceInfoForClient()
	tagsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return tagsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return tagsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return tagsServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "TagsService"
	handlerType := (*tags_service.TagsService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "tags",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func pidTidCreateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(tags_service.PidTidCreateRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(tags_service.TagsService).PidTidCreate(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PidTidCreateArgs:
		success, err := handler.(tags_service.TagsService).PidTidCreate(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PidTidCreateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPidTidCreateArgs() interface{} {
	return &PidTidCreateArgs{}
}

func newPidTidCreateResult() interface{} {
	return &PidTidCreateResult{}
}

type PidTidCreateArgs struct {
	Req *tags_service.PidTidCreateRequest
}

func (p *PidTidCreateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(tags_service.PidTidCreateRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PidTidCreateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PidTidCreateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PidTidCreateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PidTidCreateArgs) Unmarshal(in []byte) error {
	msg := new(tags_service.PidTidCreateRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PidTidCreateArgs_Req_DEFAULT *tags_service.PidTidCreateRequest

func (p *PidTidCreateArgs) GetReq() *tags_service.PidTidCreateRequest {
	if !p.IsSetReq() {
		return PidTidCreateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PidTidCreateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PidTidCreateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PidTidCreateResult struct {
	Success *tags_service.PidTidCreateResponse
}

var PidTidCreateResult_Success_DEFAULT *tags_service.PidTidCreateResponse

func (p *PidTidCreateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(tags_service.PidTidCreateResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PidTidCreateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PidTidCreateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PidTidCreateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PidTidCreateResult) Unmarshal(in []byte) error {
	msg := new(tags_service.PidTidCreateResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PidTidCreateResult) GetSuccess() *tags_service.PidTidCreateResponse {
	if !p.IsSetSuccess() {
		return PidTidCreateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PidTidCreateResult) SetSuccess(x interface{}) {
	p.Success = x.(*tags_service.PidTidCreateResponse)
}

func (p *PidTidCreateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PidTidCreateResult) GetResult() interface{} {
	return p.Success
}

func getTagHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(tags_service.GetTagRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(tags_service.TagsService).GetTag(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetTagArgs:
		success, err := handler.(tags_service.TagsService).GetTag(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetTagResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetTagArgs() interface{} {
	return &GetTagArgs{}
}

func newGetTagResult() interface{} {
	return &GetTagResult{}
}

type GetTagArgs struct {
	Req *tags_service.GetTagRequest
}

func (p *GetTagArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(tags_service.GetTagRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetTagArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetTagArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetTagArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetTagArgs) Unmarshal(in []byte) error {
	msg := new(tags_service.GetTagRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetTagArgs_Req_DEFAULT *tags_service.GetTagRequest

func (p *GetTagArgs) GetReq() *tags_service.GetTagRequest {
	if !p.IsSetReq() {
		return GetTagArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetTagArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetTagArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetTagResult struct {
	Success *tags_service.GetTagResponse
}

var GetTagResult_Success_DEFAULT *tags_service.GetTagResponse

func (p *GetTagResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(tags_service.GetTagResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetTagResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetTagResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetTagResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetTagResult) Unmarshal(in []byte) error {
	msg := new(tags_service.GetTagResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetTagResult) GetSuccess() *tags_service.GetTagResponse {
	if !p.IsSetSuccess() {
		return GetTagResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetTagResult) SetSuccess(x interface{}) {
	p.Success = x.(*tags_service.GetTagResponse)
}

func (p *GetTagResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetTagResult) GetResult() interface{} {
	return p.Success
}

func getTagIDsWithPidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(tags_service.GetTagIDRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(tags_service.TagsService).GetTagIDsWithPid(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetTagIDsWithPidArgs:
		success, err := handler.(tags_service.TagsService).GetTagIDsWithPid(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetTagIDsWithPidResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetTagIDsWithPidArgs() interface{} {
	return &GetTagIDsWithPidArgs{}
}

func newGetTagIDsWithPidResult() interface{} {
	return &GetTagIDsWithPidResult{}
}

type GetTagIDsWithPidArgs struct {
	Req *tags_service.GetTagIDRequest
}

func (p *GetTagIDsWithPidArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(tags_service.GetTagIDRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetTagIDsWithPidArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetTagIDsWithPidArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetTagIDsWithPidArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetTagIDsWithPidArgs) Unmarshal(in []byte) error {
	msg := new(tags_service.GetTagIDRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetTagIDsWithPidArgs_Req_DEFAULT *tags_service.GetTagIDRequest

func (p *GetTagIDsWithPidArgs) GetReq() *tags_service.GetTagIDRequest {
	if !p.IsSetReq() {
		return GetTagIDsWithPidArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetTagIDsWithPidArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetTagIDsWithPidArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetTagIDsWithPidResult struct {
	Success *tags_service.GetTagIDResponse
}

var GetTagIDsWithPidResult_Success_DEFAULT *tags_service.GetTagIDResponse

func (p *GetTagIDsWithPidResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(tags_service.GetTagIDResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetTagIDsWithPidResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetTagIDsWithPidResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetTagIDsWithPidResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetTagIDsWithPidResult) Unmarshal(in []byte) error {
	msg := new(tags_service.GetTagIDResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetTagIDsWithPidResult) GetSuccess() *tags_service.GetTagIDResponse {
	if !p.IsSetSuccess() {
		return GetTagIDsWithPidResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetTagIDsWithPidResult) SetSuccess(x interface{}) {
	p.Success = x.(*tags_service.GetTagIDResponse)
}

func (p *GetTagIDsWithPidResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetTagIDsWithPidResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PidTidCreate(ctx context.Context, Req *tags_service.PidTidCreateRequest) (r *tags_service.PidTidCreateResponse, err error) {
	var _args PidTidCreateArgs
	_args.Req = Req
	var _result PidTidCreateResult
	if err = p.c.Call(ctx, "PidTidCreate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTag(ctx context.Context, Req *tags_service.GetTagRequest) (r *tags_service.GetTagResponse, err error) {
	var _args GetTagArgs
	_args.Req = Req
	var _result GetTagResult
	if err = p.c.Call(ctx, "GetTag", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTagIDsWithPid(ctx context.Context, Req *tags_service.GetTagIDRequest) (r *tags_service.GetTagIDResponse, err error) {
	var _args GetTagIDsWithPidArgs
	_args.Req = Req
	var _result GetTagIDsWithPidResult
	if err = p.c.Call(ctx, "GetTagIDsWithPid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
