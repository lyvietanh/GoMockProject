// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: booking.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookingGrpcClient is the client API for BookingGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingGrpcClient interface {
	BookFlight(ctx context.Context, in *GrpcBookFlightRequest, opts ...grpc.CallOption) (*GrpcBookFlightResponse, error)
	ViewBooking(ctx context.Context, in *GrpcViewBookingRequest, opts ...grpc.CallOption) (*GrpcBookFlightResponse, error)
	CancelBooking(ctx context.Context, in *GrpcCancelBookingRequest, opts ...grpc.CallOption) (*GrpcCancelBookingResponse, error)
}

type bookingGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingGrpcClient(cc grpc.ClientConnInterface) BookingGrpcClient {
	return &bookingGrpcClient{cc}
}

func (c *bookingGrpcClient) BookFlight(ctx context.Context, in *GrpcBookFlightRequest, opts ...grpc.CallOption) (*GrpcBookFlightResponse, error) {
	out := new(GrpcBookFlightResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingGrpc/BookFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingGrpcClient) ViewBooking(ctx context.Context, in *GrpcViewBookingRequest, opts ...grpc.CallOption) (*GrpcBookFlightResponse, error) {
	out := new(GrpcBookFlightResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingGrpc/ViewBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingGrpcClient) CancelBooking(ctx context.Context, in *GrpcCancelBookingRequest, opts ...grpc.CallOption) (*GrpcCancelBookingResponse, error) {
	out := new(GrpcCancelBookingResponse)
	err := c.cc.Invoke(ctx, "/pb.BookingGrpc/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingGrpcServer is the server API for BookingGrpc service.
// All implementations must embed UnimplementedBookingGrpcServer
// for forward compatibility
type BookingGrpcServer interface {
	BookFlight(context.Context, *GrpcBookFlightRequest) (*GrpcBookFlightResponse, error)
	ViewBooking(context.Context, *GrpcViewBookingRequest) (*GrpcBookFlightResponse, error)
	CancelBooking(context.Context, *GrpcCancelBookingRequest) (*GrpcCancelBookingResponse, error)
	mustEmbedUnimplementedBookingGrpcServer()
}

// UnimplementedBookingGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedBookingGrpcServer struct {
}

func (UnimplementedBookingGrpcServer) BookFlight(context.Context, *GrpcBookFlightRequest) (*GrpcBookFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookFlight not implemented")
}
func (UnimplementedBookingGrpcServer) ViewBooking(context.Context, *GrpcViewBookingRequest) (*GrpcBookFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewBooking not implemented")
}
func (UnimplementedBookingGrpcServer) CancelBooking(context.Context, *GrpcCancelBookingRequest) (*GrpcCancelBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedBookingGrpcServer) mustEmbedUnimplementedBookingGrpcServer() {}

// UnsafeBookingGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingGrpcServer will
// result in compilation errors.
type UnsafeBookingGrpcServer interface {
	mustEmbedUnimplementedBookingGrpcServer()
}

func RegisterBookingGrpcServer(s grpc.ServiceRegistrar, srv BookingGrpcServer) {
	s.RegisterService(&BookingGrpc_ServiceDesc, srv)
}

func _BookingGrpc_BookFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcBookFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGrpcServer).BookFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingGrpc/BookFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGrpcServer).BookFlight(ctx, req.(*GrpcBookFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingGrpc_ViewBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcViewBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGrpcServer).ViewBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingGrpc/ViewBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGrpcServer).ViewBooking(ctx, req.(*GrpcViewBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingGrpc_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcCancelBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGrpcServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BookingGrpc/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGrpcServer).CancelBooking(ctx, req.(*GrpcCancelBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingGrpc_ServiceDesc is the grpc.ServiceDesc for BookingGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BookingGrpc",
	HandlerType: (*BookingGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookFlight",
			Handler:    _BookingGrpc_BookFlight_Handler,
		},
		{
			MethodName: "ViewBooking",
			Handler:    _BookingGrpc_ViewBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _BookingGrpc_CancelBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}