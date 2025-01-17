// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/services/content/v1/content.proto
package content

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TTRPCContentService interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	List(context.Context, *ListContentRequest, TTRPCContent_ListServer) error
	Delete(context.Context, *DeleteContentRequest) (*emptypb.Empty, error)
	Read(context.Context, *ReadContentRequest, TTRPCContent_ReadServer) error
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	ListStatuses(context.Context, *ListStatusesRequest) (*ListStatusesResponse, error)
	Write(context.Context, TTRPCContent_WriteServer) error
	Abort(context.Context, *AbortRequest) (*emptypb.Empty, error)
}

type TTRPCContent_ListServer interface {
	Send(*ListContentResponse) error
	ttrpc.StreamServer
}

type ttrpccontentListServer struct {
	ttrpc.StreamServer
}

func (x *ttrpccontentListServer) Send(m *ListContentResponse) error {
	return x.StreamServer.SendMsg(m)
}

type TTRPCContent_ReadServer interface {
	Send(*ReadContentResponse) error
	ttrpc.StreamServer
}

type ttrpccontentReadServer struct {
	ttrpc.StreamServer
}

func (x *ttrpccontentReadServer) Send(m *ReadContentResponse) error {
	return x.StreamServer.SendMsg(m)
}

type TTRPCContent_WriteServer interface {
	Send(*WriteContentResponse) error
	Recv() (*WriteContentRequest, error)
	ttrpc.StreamServer
}

type ttrpccontentWriteServer struct {
	ttrpc.StreamServer
}

func (x *ttrpccontentWriteServer) Send(m *WriteContentResponse) error {
	return x.StreamServer.SendMsg(m)
}

func (x *ttrpccontentWriteServer) Recv() (*WriteContentRequest, error) {
	m := new(WriteContentRequest)
	if err := x.StreamServer.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func RegisterTTRPCContentService(srv *ttrpc.Server, svc TTRPCContentService) {
	srv.RegisterService("containerd.services.content.v1.Content", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"Info": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req InfoRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Info(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req UpdateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req DeleteContentRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Delete(ctx, &req)
			},
			"Status": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StatusRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Status(ctx, &req)
			},
			"ListStatuses": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ListStatusesRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ListStatuses(ctx, &req)
			},
			"Abort": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req AbortRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Abort(ctx, &req)
			},
		},
		Streams: map[string]ttrpc.Stream{
			"List": {
				Handler: func(ctx context.Context, stream ttrpc.StreamServer) (interface{}, error) {
					m := new(ListContentRequest)
					if err := stream.RecvMsg(m); err != nil {
						return nil, err
					}
					return nil, svc.List(ctx, m, &ttrpccontentListServer{stream})
				},
				StreamingClient: false,
				StreamingServer: true,
			},
			"Read": {
				Handler: func(ctx context.Context, stream ttrpc.StreamServer) (interface{}, error) {
					m := new(ReadContentRequest)
					if err := stream.RecvMsg(m); err != nil {
						return nil, err
					}
					return nil, svc.Read(ctx, m, &ttrpccontentReadServer{stream})
				},
				StreamingClient: false,
				StreamingServer: true,
			},
			"Write": {
				Handler: func(ctx context.Context, stream ttrpc.StreamServer) (interface{}, error) {
					return nil, svc.Write(ctx, &ttrpccontentWriteServer{stream})
				},
				StreamingClient: true,
				StreamingServer: true,
			},
		},
	})
}

type TTRPCContentClient interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	List(context.Context, *ListContentRequest) (TTRPCContent_ListClient, error)
	Delete(context.Context, *DeleteContentRequest) (*emptypb.Empty, error)
	Read(context.Context, *ReadContentRequest) (TTRPCContent_ReadClient, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	ListStatuses(context.Context, *ListStatusesRequest) (*ListStatusesResponse, error)
	Write(context.Context) (TTRPCContent_WriteClient, error)
	Abort(context.Context, *AbortRequest) (*emptypb.Empty, error)
}

type ttrpccontentClient struct {
	client *ttrpc.Client
}

func NewTTRPCContentClient(client *ttrpc.Client) TTRPCContentClient {
	return &ttrpccontentClient{
		client: client,
	}
}

func (c *ttrpccontentClient) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, error) {
	var resp InfoResponse
	if err := c.client.Call(ctx, "containerd.services.content.v1.Content", "Info", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontentClient) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	var resp UpdateResponse
	if err := c.client.Call(ctx, "containerd.services.content.v1.Content", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontentClient) List(ctx context.Context, req *ListContentRequest) (TTRPCContent_ListClient, error) {
	stream, err := c.client.NewStream(ctx, &ttrpc.StreamDesc{
		StreamingClient: false,
		StreamingServer: true,
	}, "containerd.services.content.v1.Content", "List", req)
	if err != nil {
		return nil, err
	}
	x := &ttrpccontentListClient{stream}
	return x, nil
}

type TTRPCContent_ListClient interface {
	Recv() (*ListContentResponse, error)
	ttrpc.ClientStream
}

type ttrpccontentListClient struct {
	ttrpc.ClientStream
}

func (x *ttrpccontentListClient) Recv() (*ListContentResponse, error) {
	m := new(ListContentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ttrpccontentClient) Delete(ctx context.Context, req *DeleteContentRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.content.v1.Content", "Delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontentClient) Read(ctx context.Context, req *ReadContentRequest) (TTRPCContent_ReadClient, error) {
	stream, err := c.client.NewStream(ctx, &ttrpc.StreamDesc{
		StreamingClient: false,
		StreamingServer: true,
	}, "containerd.services.content.v1.Content", "Read", req)
	if err != nil {
		return nil, err
	}
	x := &ttrpccontentReadClient{stream}
	return x, nil
}

type TTRPCContent_ReadClient interface {
	Recv() (*ReadContentResponse, error)
	ttrpc.ClientStream
}

type ttrpccontentReadClient struct {
	ttrpc.ClientStream
}

func (x *ttrpccontentReadClient) Recv() (*ReadContentResponse, error) {
	m := new(ReadContentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ttrpccontentClient) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	var resp StatusResponse
	if err := c.client.Call(ctx, "containerd.services.content.v1.Content", "Status", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontentClient) ListStatuses(ctx context.Context, req *ListStatusesRequest) (*ListStatusesResponse, error) {
	var resp ListStatusesResponse
	if err := c.client.Call(ctx, "containerd.services.content.v1.Content", "ListStatuses", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpccontentClient) Write(ctx context.Context) (TTRPCContent_WriteClient, error) {
	stream, err := c.client.NewStream(ctx, &ttrpc.StreamDesc{
		StreamingClient: true,
		StreamingServer: true,
	}, "containerd.services.content.v1.Content", "Write", nil)
	if err != nil {
		return nil, err
	}
	x := &ttrpccontentWriteClient{stream}
	return x, nil
}

type TTRPCContent_WriteClient interface {
	Send(*WriteContentRequest) error
	Recv() (*WriteContentResponse, error)
	ttrpc.ClientStream
}

type ttrpccontentWriteClient struct {
	ttrpc.ClientStream
}

func (x *ttrpccontentWriteClient) Send(m *WriteContentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ttrpccontentWriteClient) Recv() (*WriteContentResponse, error) {
	m := new(WriteContentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ttrpccontentClient) Abort(ctx context.Context, req *AbortRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.content.v1.Content", "Abort", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
