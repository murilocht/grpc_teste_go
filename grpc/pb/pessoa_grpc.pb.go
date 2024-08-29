// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: pessoa.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PessoaService_CreatePessoa_FullMethodName = "/PessoaService/CreatePessoa"
)

// PessoaServiceClient is the client API for PessoaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PessoaServiceClient interface {
	CreatePessoa(ctx context.Context, in *Pessoa, opts ...grpc.CallOption) (*PessoaResut, error)
}

type pessoaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPessoaServiceClient(cc grpc.ClientConnInterface) PessoaServiceClient {
	return &pessoaServiceClient{cc}
}

func (c *pessoaServiceClient) CreatePessoa(ctx context.Context, in *Pessoa, opts ...grpc.CallOption) (*PessoaResut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PessoaResut)
	err := c.cc.Invoke(ctx, PessoaService_CreatePessoa_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PessoaServiceServer is the server API for PessoaService service.
// All implementations must embed UnimplementedPessoaServiceServer
// for forward compatibility.
type PessoaServiceServer interface {
	CreatePessoa(context.Context, *Pessoa) (*PessoaResut, error)
	mustEmbedUnimplementedPessoaServiceServer()
}

// UnimplementedPessoaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPessoaServiceServer struct{}

func (UnimplementedPessoaServiceServer) CreatePessoa(context.Context, *Pessoa) (*PessoaResut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePessoa not implemented")
}
func (UnimplementedPessoaServiceServer) mustEmbedUnimplementedPessoaServiceServer() {}
func (UnimplementedPessoaServiceServer) testEmbeddedByValue()                       {}

// UnsafePessoaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PessoaServiceServer will
// result in compilation errors.
type UnsafePessoaServiceServer interface {
	mustEmbedUnimplementedPessoaServiceServer()
}

func RegisterPessoaServiceServer(s grpc.ServiceRegistrar, srv PessoaServiceServer) {
	// If the following call pancis, it indicates UnimplementedPessoaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PessoaService_ServiceDesc, srv)
}

func _PessoaService_CreatePessoa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pessoa)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PessoaServiceServer).CreatePessoa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PessoaService_CreatePessoa_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PessoaServiceServer).CreatePessoa(ctx, req.(*Pessoa))
	}
	return interceptor(ctx, in, info, handler)
}

// PessoaService_ServiceDesc is the grpc.ServiceDesc for PessoaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PessoaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PessoaService",
	HandlerType: (*PessoaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePessoa",
			Handler:    _PessoaService_CreatePessoa_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pessoa.proto",
}
