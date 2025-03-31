// Code generated by Kitex v0.12.1. DO NOT EDIT.

package tagsservice

import (
	"context"
	tags_service "github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PidTidCreate(ctx context.Context, Req *tags_service.PidTidCreateRequest, callOptions ...callopt.Option) (r *tags_service.PidTidCreateResponse, err error)
	GetTag(ctx context.Context, Req *tags_service.GetTagRequest, callOptions ...callopt.Option) (r *tags_service.GetTagResponse, err error)
	GetTagIDsWithPid(ctx context.Context, Req *tags_service.GetTagIDRequest, callOptions ...callopt.Option) (r *tags_service.GetTagIDResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTagsServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTagsServiceClient struct {
	*kClient
}

func (p *kTagsServiceClient) PidTidCreate(ctx context.Context, Req *tags_service.PidTidCreateRequest, callOptions ...callopt.Option) (r *tags_service.PidTidCreateResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PidTidCreate(ctx, Req)
}

func (p *kTagsServiceClient) GetTag(ctx context.Context, Req *tags_service.GetTagRequest, callOptions ...callopt.Option) (r *tags_service.GetTagResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTag(ctx, Req)
}

func (p *kTagsServiceClient) GetTagIDsWithPid(ctx context.Context, Req *tags_service.GetTagIDRequest, callOptions ...callopt.Option) (r *tags_service.GetTagIDResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTagIDsWithPid(ctx, Req)
}
