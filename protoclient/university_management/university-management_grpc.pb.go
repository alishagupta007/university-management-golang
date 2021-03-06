// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package university_management

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

// UniversityManagementServiceClient is the client API for UniversityManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UniversityManagementServiceClient interface {
	GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*GetDepartmentResponse, error)
	GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error)
	GetStaffForStudent(ctx context.Context, in *GetStaffForStudentRequest, opts ...grpc.CallOption) (*GetStaffForStudentResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Notify(ctx context.Context, in *GetNotifyRequest, opts ...grpc.CallOption) (*GetNotifyResponse, error)
}

type universityManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUniversityManagementServiceClient(cc grpc.ClientConnInterface) UniversityManagementServiceClient {
	return &universityManagementServiceClient{cc}
}

func (c *universityManagementServiceClient) GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*GetDepartmentResponse, error) {
	out := new(GetDepartmentResponse)
	err := c.cc.Invoke(ctx, "/university_management.UniversityManagementService/GetDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityManagementServiceClient) GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error) {
	out := new(GetStudentsResponse)
	err := c.cc.Invoke(ctx, "/university_management.UniversityManagementService/GetStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityManagementServiceClient) GetStaffForStudent(ctx context.Context, in *GetStaffForStudentRequest, opts ...grpc.CallOption) (*GetStaffForStudentResponse, error) {
	out := new(GetStaffForStudentResponse)
	err := c.cc.Invoke(ctx, "/university_management.UniversityManagementService/GetStaffForStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityManagementServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/university_management.UniversityManagementService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityManagementServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/university_management.UniversityManagementService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *universityManagementServiceClient) Notify(ctx context.Context, in *GetNotifyRequest, opts ...grpc.CallOption) (*GetNotifyResponse, error) {
	out := new(GetNotifyResponse)
	err := c.cc.Invoke(ctx, "/university_management.UniversityManagementService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UniversityManagementServiceServer is the server API for UniversityManagementService service.
// All implementations must embed UnimplementedUniversityManagementServiceServer
// for forward compatibility
type UniversityManagementServiceServer interface {
	GetDepartment(context.Context, *GetDepartmentRequest) (*GetDepartmentResponse, error)
	GetStudents(context.Context, *GetStudentsRequest) (*GetStudentsResponse, error)
	GetStaffForStudent(context.Context, *GetStaffForStudentRequest) (*GetStaffForStudentResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Notify(context.Context, *GetNotifyRequest) (*GetNotifyResponse, error)
	mustEmbedUnimplementedUniversityManagementServiceServer()
}

// UnimplementedUniversityManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUniversityManagementServiceServer struct {
}

func (UnimplementedUniversityManagementServiceServer) GetDepartment(context.Context, *GetDepartmentRequest) (*GetDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartment not implemented")
}
func (UnimplementedUniversityManagementServiceServer) GetStudents(context.Context, *GetStudentsRequest) (*GetStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}
func (UnimplementedUniversityManagementServiceServer) GetStaffForStudent(context.Context, *GetStaffForStudentRequest) (*GetStaffForStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaffForStudent not implemented")
}
func (UnimplementedUniversityManagementServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUniversityManagementServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUniversityManagementServiceServer) Notify(context.Context, *GetNotifyRequest) (*GetNotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedUniversityManagementServiceServer) mustEmbedUnimplementedUniversityManagementServiceServer() {
}

// UnsafeUniversityManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UniversityManagementServiceServer will
// result in compilation errors.
type UnsafeUniversityManagementServiceServer interface {
	mustEmbedUnimplementedUniversityManagementServiceServer()
}

func RegisterUniversityManagementServiceServer(s grpc.ServiceRegistrar, srv UniversityManagementServiceServer) {
	s.RegisterService(&UniversityManagementService_ServiceDesc, srv)
}

func _UniversityManagementService_GetDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityManagementServiceServer).GetDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/university_management.UniversityManagementService/GetDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityManagementServiceServer).GetDepartment(ctx, req.(*GetDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityManagementService_GetStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityManagementServiceServer).GetStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/university_management.UniversityManagementService/GetStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityManagementServiceServer).GetStudents(ctx, req.(*GetStudentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityManagementService_GetStaffForStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaffForStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityManagementServiceServer).GetStaffForStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/university_management.UniversityManagementService/GetStaffForStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityManagementServiceServer).GetStaffForStudent(ctx, req.(*GetStaffForStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityManagementService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityManagementServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/university_management.UniversityManagementService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityManagementServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityManagementService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityManagementServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/university_management.UniversityManagementService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityManagementServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UniversityManagementService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UniversityManagementServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/university_management.UniversityManagementService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UniversityManagementServiceServer).Notify(ctx, req.(*GetNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UniversityManagementService_ServiceDesc is the grpc.ServiceDesc for UniversityManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UniversityManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "university_management.UniversityManagementService",
	HandlerType: (*UniversityManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDepartment",
			Handler:    _UniversityManagementService_GetDepartment_Handler,
		},
		{
			MethodName: "GetStudents",
			Handler:    _UniversityManagementService_GetStudents_Handler,
		},
		{
			MethodName: "GetStaffForStudent",
			Handler:    _UniversityManagementService_GetStaffForStudent_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UniversityManagementService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UniversityManagementService_Logout_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _UniversityManagementService_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "university-management.proto",
}
