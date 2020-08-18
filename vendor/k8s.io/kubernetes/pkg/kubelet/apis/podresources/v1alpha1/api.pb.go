/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ListPodResourcesRequest is the request made to the PodResourcesLister service
type ListPodResourcesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPodResourcesRequest) Reset()      { *m = ListPodResourcesRequest{} }
func (*ListPodResourcesRequest) ProtoMessage() {}
func (*ListPodResourcesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *ListPodResourcesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListPodResourcesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListPodResourcesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListPodResourcesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodResourcesRequest.Merge(m, src)
}
func (m *ListPodResourcesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListPodResourcesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodResourcesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodResourcesRequest proto.InternalMessageInfo

// ListPodResourcesResponse is the response returned by List function
type ListPodResourcesResponse struct {
	PodResources         []*PodResources `protobuf:"bytes,1,rep,name=pod_resources,json=podResources,proto3" json:"pod_resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListPodResourcesResponse) Reset()      { *m = ListPodResourcesResponse{} }
func (*ListPodResourcesResponse) ProtoMessage() {}
func (*ListPodResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *ListPodResourcesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListPodResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListPodResourcesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListPodResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodResourcesResponse.Merge(m, src)
}
func (m *ListPodResourcesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListPodResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodResourcesResponse proto.InternalMessageInfo

func (m *ListPodResourcesResponse) GetPodResources() []*PodResources {
	if m != nil {
		return m.PodResources
	}
	return nil
}

// PodResources contains information about the node resources assigned to a pod
type PodResources struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string                `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Containers           []*ContainerResources `protobuf:"bytes,3,rep,name=containers,proto3" json:"containers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PodResources) Reset()      { *m = PodResources{} }
func (*PodResources) ProtoMessage() {}
func (*PodResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}
func (m *PodResources) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PodResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PodResources.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PodResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodResources.Merge(m, src)
}
func (m *PodResources) XXX_Size() int {
	return m.Size()
}
func (m *PodResources) XXX_DiscardUnknown() {
	xxx_messageInfo_PodResources.DiscardUnknown(m)
}

var xxx_messageInfo_PodResources proto.InternalMessageInfo

func (m *PodResources) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PodResources) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PodResources) GetContainers() []*ContainerResources {
	if m != nil {
		return m.Containers
	}
	return nil
}

// ContainerResources contains information about the resources assigned to a container
type ContainerResources struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Devices              []*ContainerDevices `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
	CpuIds               []uint32            `protobuf:"varint,3,rep,packed,name=cpu_ids,json=cpuIds,proto3" json:"cpu_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ContainerResources) Reset()      { *m = ContainerResources{} }
func (*ContainerResources) ProtoMessage() {}
func (*ContainerResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}
func (m *ContainerResources) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContainerResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContainerResources.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContainerResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerResources.Merge(m, src)
}
func (m *ContainerResources) XXX_Size() int {
	return m.Size()
}
func (m *ContainerResources) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerResources.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerResources proto.InternalMessageInfo

func (m *ContainerResources) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerResources) GetDevices() []*ContainerDevices {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *ContainerResources) GetCpuIds() []uint32 {
	if m != nil {
		return m.CpuIds
	}
	return nil
}

// ContainerDevices contains information about the devices assigned to a container
type ContainerDevices struct {
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	DeviceIds            []string `protobuf:"bytes,2,rep,name=device_ids,json=deviceIds,proto3" json:"device_ids,omitempty"`
	Numaid               int64    `protobuf:"varint,3,opt,name=numaid,proto3" json:"numaid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerDevices) Reset()      { *m = ContainerDevices{} }
func (*ContainerDevices) ProtoMessage() {}
func (*ContainerDevices) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}
func (m *ContainerDevices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContainerDevices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContainerDevices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContainerDevices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerDevices.Merge(m, src)
}
func (m *ContainerDevices) XXX_Size() int {
	return m.Size()
}
func (m *ContainerDevices) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerDevices.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerDevices proto.InternalMessageInfo

func (m *ContainerDevices) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ContainerDevices) GetDeviceIds() []string {
	if m != nil {
		return m.DeviceIds
	}
	return nil
}

func (m *ContainerDevices) GetNumaid() int64 {
	if m != nil {
		return m.Numaid
	}
	return 0
}

func init() {
	proto.RegisterType((*ListPodResourcesRequest)(nil), "v1alpha1.ListPodResourcesRequest")
	proto.RegisterType((*ListPodResourcesResponse)(nil), "v1alpha1.ListPodResourcesResponse")
	proto.RegisterType((*PodResources)(nil), "v1alpha1.PodResources")
	proto.RegisterType((*ContainerResources)(nil), "v1alpha1.ContainerResources")
	proto.RegisterType((*ContainerDevices)(nil), "v1alpha1.ContainerDevices")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x65, 0x28, 0x81, 0xd7, 0xfb, 0x20, 0x79, 0x99, 0x05, 0xf4, 0x11, 0x5e, 0xc3, 0xab, 0x1b,
	0x36, 0x96, 0x80, 0xee, 0x74, 0xa5, 0x6e, 0x48, 0x8c, 0x9a, 0x6e, 0x5c, 0x92, 0xd2, 0x19, 0x61,
	0x12, 0xe9, 0x8c, 0x9d, 0x0e, 0xee, 0x8c, 0x9f, 0xe0, 0x67, 0xb1, 0x74, 0xe9, 0x52, 0xea, 0x8f,
	0x98, 0x4e, 0x6d, 0x5a, 0x05, 0x5d, 0xf5, 0x9e, 0x7b, 0xce, 0x9d, 0x73, 0x66, 0x7a, 0xc1, 0xf4,
	0x05, 0x73, 0x45, 0xc4, 0x63, 0x8e, 0x7f, 0xad, 0x46, 0xfe, 0xad, 0x58, 0xf8, 0xa3, 0xee, 0xfe,
	0x9c, 0xc5, 0x0b, 0x35, 0x73, 0x03, 0xbe, 0x1c, 0xce, 0xf9, 0x9c, 0x0f, 0xb5, 0x60, 0xa6, 0x6e,
	0x34, 0xd2, 0x40, 0x57, 0xd9, 0xa0, 0xf3, 0x17, 0x3a, 0xe7, 0x4c, 0xc6, 0x57, 0x9c, 0x78, 0x54,
	0x72, 0x15, 0x05, 0x54, 0x7a, 0xf4, 0x4e, 0x51, 0x19, 0x3b, 0xd7, 0x60, 0x6d, 0x53, 0x52, 0xf0,
	0x50, 0x52, 0x7c, 0x04, 0x2d, 0xc1, 0xc9, 0x34, 0xca, 0x09, 0x0b, 0xf5, 0x8d, 0xc1, 0xef, 0x71,
	0xdb, 0xcd, 0x73, 0xb8, 0x9f, 0xc6, 0x9a, 0xa2, 0x84, 0x9c, 0x07, 0x68, 0x96, 0x59, 0x8c, 0xa1,
	0x16, 0xfa, 0x4b, 0x6a, 0xa1, 0x3e, 0x1a, 0x98, 0x9e, 0xae, 0x71, 0x0f, 0xcc, 0xf4, 0x2b, 0x85,
	0x1f, 0x50, 0xab, 0xaa, 0x89, 0xa2, 0x81, 0x8f, 0x01, 0x02, 0x1e, 0xc6, 0x3e, 0x0b, 0x69, 0x24,
	0x2d, 0x43, 0x7b, 0xf7, 0x0a, 0xef, 0xd3, 0x9c, 0x2b, 0x12, 0x94, 0xf4, 0xce, 0x3d, 0xe0, 0x6d,
	0xc5, 0xce, 0x14, 0x87, 0xd0, 0x20, 0x74, 0xc5, 0xd2, 0x0b, 0x56, 0xb5, 0x49, 0x77, 0x87, 0xc9,
	0x59, 0xa6, 0xf0, 0x72, 0x29, 0xee, 0x40, 0x23, 0x10, 0x6a, 0xca, 0x48, 0x16, 0xad, 0xe5, 0xd5,
	0x03, 0xa1, 0x26, 0x44, 0x3a, 0x21, 0xfc, 0xf9, 0x3a, 0x85, 0xf7, 0xa0, 0x95, 0xbf, 0xe2, 0xb4,
	0xe4, 0xdf, 0xcc, 0x9b, 0x17, 0x69, 0x8e, 0x7f, 0x00, 0xd9, 0xe1, 0xfa, 0xd0, 0x34, 0x8a, 0xe9,
	0x99, 0x59, 0x67, 0x42, 0x24, 0x6e, 0x43, 0x3d, 0x54, 0x4b, 0x9f, 0x11, 0xcb, 0xe8, 0xa3, 0x81,
	0xe1, 0x7d, 0xa0, 0x31, 0x05, 0x5c, 0x7e, 0xe8, 0xf4, 0x6f, 0xd2, 0x08, 0x5f, 0x42, 0x2d, 0xad,
	0xf0, 0xff, 0xe2, 0x2e, 0xdf, 0xac, 0x40, 0xd7, 0xf9, 0x49, 0x92, 0xad, 0x82, 0x53, 0x39, 0xe9,
	0xad, 0x37, 0x36, 0x7a, 0xd9, 0xd8, 0x95, 0xc7, 0xc4, 0x46, 0xeb, 0xc4, 0x46, 0xcf, 0x89, 0x8d,
	0x5e, 0x13, 0x1b, 0x3d, 0xbd, 0xd9, 0x95, 0x59, 0x5d, 0x2f, 0xda, 0xc1, 0x7b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0xa0, 0xb9, 0x67, 0xae, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PodResourcesListerClient is the client API for PodResourcesLister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PodResourcesListerClient interface {
	List(ctx context.Context, in *ListPodResourcesRequest, opts ...grpc.CallOption) (*ListPodResourcesResponse, error)
}

type podResourcesListerClient struct {
	cc *grpc.ClientConn
}

func NewPodResourcesListerClient(cc *grpc.ClientConn) PodResourcesListerClient {
	return &podResourcesListerClient{cc}
}

func (c *podResourcesListerClient) List(ctx context.Context, in *ListPodResourcesRequest, opts ...grpc.CallOption) (*ListPodResourcesResponse, error) {
	out := new(ListPodResourcesResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.PodResourcesLister/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodResourcesListerServer is the server API for PodResourcesLister service.
type PodResourcesListerServer interface {
	List(context.Context, *ListPodResourcesRequest) (*ListPodResourcesResponse, error)
}

// UnimplementedPodResourcesListerServer can be embedded to have forward compatible implementations.
type UnimplementedPodResourcesListerServer struct {
}

func (*UnimplementedPodResourcesListerServer) List(ctx context.Context, req *ListPodResourcesRequest) (*ListPodResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterPodResourcesListerServer(s *grpc.Server, srv PodResourcesListerServer) {
	s.RegisterService(&_PodResourcesLister_serviceDesc, srv)
}

func _PodResourcesLister_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodResourcesListerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.PodResourcesLister/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodResourcesListerServer).List(ctx, req.(*ListPodResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PodResourcesLister_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.PodResourcesLister",
	HandlerType: (*PodResourcesListerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PodResourcesLister_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *ListPodResourcesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListPodResourcesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListPodResourcesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListPodResourcesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListPodResourcesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListPodResourcesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PodResources) > 0 {
		for iNdEx := len(m.PodResources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PodResources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PodResources) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PodResources) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PodResources) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Containers) > 0 {
		for iNdEx := len(m.Containers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Containers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContainerResources) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerResources) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContainerResources) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CpuIds) > 0 {
		dAtA2 := make([]byte, len(m.CpuIds)*10)
		var j1 int
		for _, num := range m.CpuIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintApi(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Devices) > 0 {
		for iNdEx := len(m.Devices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Devices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContainerDevices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerDevices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContainerDevices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Numaid != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Numaid))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DeviceIds) > 0 {
		for iNdEx := len(m.DeviceIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DeviceIds[iNdEx])
			copy(dAtA[i:], m.DeviceIds[iNdEx])
			i = encodeVarintApi(dAtA, i, uint64(len(m.DeviceIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ResourceName) > 0 {
		i -= len(m.ResourceName)
		copy(dAtA[i:], m.ResourceName)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ResourceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListPodResourcesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListPodResourcesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PodResources) > 0 {
		for _, e := range m.PodResources {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func (m *PodResources) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if len(m.Containers) > 0 {
		for _, e := range m.Containers {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func (m *ContainerResources) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if len(m.CpuIds) > 0 {
		l = 0
		for _, e := range m.CpuIds {
			l += sovApi(uint64(e))
		}
		n += 1 + sovApi(uint64(l)) + l
	}
	return n
}

func (m *ContainerDevices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ResourceName)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if len(m.DeviceIds) > 0 {
		for _, s := range m.DeviceIds {
			l = len(s)
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if m.Numaid != 0 {
		n += 1 + sovApi(uint64(m.Numaid))
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ListPodResourcesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListPodResourcesRequest{`,
		`}`,
	}, "")
	return s
}
func (this *ListPodResourcesResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForPodResources := "[]*PodResources{"
	for _, f := range this.PodResources {
		repeatedStringForPodResources += strings.Replace(f.String(), "PodResources", "PodResources", 1) + ","
	}
	repeatedStringForPodResources += "}"
	s := strings.Join([]string{`&ListPodResourcesResponse{`,
		`PodResources:` + repeatedStringForPodResources + `,`,
		`}`,
	}, "")
	return s
}
func (this *PodResources) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForContainers := "[]*ContainerResources{"
	for _, f := range this.Containers {
		repeatedStringForContainers += strings.Replace(f.String(), "ContainerResources", "ContainerResources", 1) + ","
	}
	repeatedStringForContainers += "}"
	s := strings.Join([]string{`&PodResources{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Containers:` + repeatedStringForContainers + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerResources) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForDevices := "[]*ContainerDevices{"
	for _, f := range this.Devices {
		repeatedStringForDevices += strings.Replace(f.String(), "ContainerDevices", "ContainerDevices", 1) + ","
	}
	repeatedStringForDevices += "}"
	s := strings.Join([]string{`&ContainerResources{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Devices:` + repeatedStringForDevices + `,`,
		`CpuIds:` + fmt.Sprintf("%v", this.CpuIds) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerDevices) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerDevices{`,
		`ResourceName:` + fmt.Sprintf("%v", this.ResourceName) + `,`,
		`DeviceIds:` + fmt.Sprintf("%v", this.DeviceIds) + `,`,
		`Numaid:` + fmt.Sprintf("%v", this.Numaid) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListPodResourcesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListPodResourcesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListPodResourcesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListPodResourcesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListPodResourcesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListPodResourcesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodResources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodResources = append(m.PodResources, &PodResources{})
			if err := m.PodResources[len(m.PodResources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PodResources) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PodResources: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PodResources: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Containers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Containers = append(m.Containers, &ContainerResources{})
			if err := m.Containers[len(m.Containers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerResources) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContainerResources: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerResources: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Devices = append(m.Devices, &ContainerDevices{})
			if err := m.Devices[len(m.Devices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowApi
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CpuIds = append(m.CpuIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowApi
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthApi
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthApi
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CpuIds) == 0 {
					m.CpuIds = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowApi
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CpuIds = append(m.CpuIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CpuIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerDevices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContainerDevices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerDevices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceIds = append(m.DeviceIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Numaid", wireType)
			}
			m.Numaid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Numaid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
