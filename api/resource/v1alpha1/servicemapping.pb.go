/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: servicemapping.proto

package dubbo_apache_org_v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ServiceNameMappingToClient struct {
	Key                  string              `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Spec                 *ServiceNameMapping `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServiceNameMappingToClient) Reset()         { *m = ServiceNameMappingToClient{} }
func (m *ServiceNameMappingToClient) String() string { return proto.CompactTextString(m) }
func (*ServiceNameMappingToClient) ProtoMessage()    {}
func (*ServiceNameMappingToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0ebb678408b52a, []int{0}
}
func (m *ServiceNameMappingToClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceNameMappingToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceNameMappingToClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceNameMappingToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceNameMappingToClient.Merge(m, src)
}
func (m *ServiceNameMappingToClient) XXX_Size() int {
	return m.Size()
}
func (m *ServiceNameMappingToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceNameMappingToClient.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceNameMappingToClient proto.InternalMessageInfo

func (m *ServiceNameMappingToClient) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ServiceNameMappingToClient) GetSpec() *ServiceNameMapping {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ServiceNameMapping struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	ApplicationNames     []string `protobuf:"bytes,2,rep,name=applicationNames,proto3" json:"applicationNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceNameMapping) Reset()         { *m = ServiceNameMapping{} }
func (m *ServiceNameMapping) String() string { return proto.CompactTextString(m) }
func (*ServiceNameMapping) ProtoMessage()    {}
func (*ServiceNameMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0ebb678408b52a, []int{1}
}
func (m *ServiceNameMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceNameMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceNameMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceNameMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceNameMapping.Merge(m, src)
}
func (m *ServiceNameMapping) XXX_Size() int {
	return m.Size()
}
func (m *ServiceNameMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceNameMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceNameMapping proto.InternalMessageInfo

func (m *ServiceNameMapping) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *ServiceNameMapping) GetApplicationNames() []string {
	if m != nil {
		return m.ApplicationNames
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceNameMappingToClient)(nil), "dubbo.apache.org.v1alpha1.ServiceNameMappingToClient")
	proto.RegisterType((*ServiceNameMapping)(nil), "dubbo.apache.org.v1alpha1.ServiceNameMapping")
}

func init() { proto.RegisterFile("servicemapping.proto", fileDescriptor_4c0ebb678408b52a) }

var fileDescriptor_4c0ebb678408b52a = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xcd, 0x4d, 0x2c, 0x28, 0xc8, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x4c, 0x29, 0x4d, 0x4a, 0xca, 0xd7, 0x4b, 0x2c, 0x48, 0x4c, 0xce, 0x48, 0xd5, 0xcb,
	0x2f, 0x4a, 0xd7, 0x2b, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x54, 0x2a, 0xe4, 0x92, 0x0a,
	0x86, 0x68, 0xf1, 0x4b, 0xcc, 0x4d, 0xf5, 0x85, 0x68, 0x0b, 0xc9, 0x77, 0xce, 0xc9, 0x4c, 0xcd,
	0x2b, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0x31, 0x85, 0x1c, 0xb9, 0x58, 0x8a, 0x0b, 0x52, 0x93, 0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d,
	0x74, 0xf5, 0x70, 0x9a, 0xac, 0x87, 0x69, 0x6c, 0x10, 0x58, 0xab, 0x52, 0x1a, 0x97, 0x10, 0xa6,
	0x9c, 0x90, 0x0a, 0x17, 0x6f, 0x66, 0x5e, 0x49, 0x6a, 0x51, 0x5a, 0x22, 0x44, 0x1c, 0x6a, 0x29,
	0xaa, 0xa0, 0x90, 0x16, 0x97, 0x40, 0x62, 0x41, 0x41, 0x4e, 0x66, 0x72, 0x62, 0x49, 0x66, 0x7e,
	0x1e, 0x48, 0xa8, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x08, 0x43, 0xdc, 0x49, 0xe8, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21,
	0x80, 0x31, 0x89, 0x0d, 0x1c, 0x20, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xa6, 0x0b,
	0x9d, 0x28, 0x01, 0x00, 0x00,
}

func (m *ServiceNameMappingToClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceNameMappingToClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceNameMappingToClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Spec != nil {
		{
			size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintServicemapping(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintServicemapping(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ServiceNameMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceNameMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceNameMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ApplicationNames) > 0 {
		for iNdEx := len(m.ApplicationNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ApplicationNames[iNdEx])
			copy(dAtA[i:], m.ApplicationNames[iNdEx])
			i = encodeVarintServicemapping(dAtA, i, uint64(len(m.ApplicationNames[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.InterfaceName) > 0 {
		i -= len(m.InterfaceName)
		copy(dAtA[i:], m.InterfaceName)
		i = encodeVarintServicemapping(dAtA, i, uint64(len(m.InterfaceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintServicemapping(dAtA []byte, offset int, v uint64) int {
	offset -= sovServicemapping(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceNameMappingToClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovServicemapping(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovServicemapping(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ServiceNameMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterfaceName)
	if l > 0 {
		n += 1 + l + sovServicemapping(uint64(l))
	}
	if len(m.ApplicationNames) > 0 {
		for _, s := range m.ApplicationNames {
			l = len(s)
			n += 1 + l + sovServicemapping(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovServicemapping(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozServicemapping(x uint64) (n int) {
	return sovServicemapping(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceNameMappingToClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServicemapping
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
			return fmt.Errorf("proto: ServiceNameMappingToClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceNameMappingToClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServicemapping
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
				return ErrInvalidLengthServicemapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServicemapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServicemapping
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
				return ErrInvalidLengthServicemapping
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServicemapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &ServiceNameMapping{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServicemapping(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServicemapping
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ServiceNameMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServicemapping
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
			return fmt.Errorf("proto: ServiceNameMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceNameMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServicemapping
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
				return ErrInvalidLengthServicemapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServicemapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServicemapping
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
				return ErrInvalidLengthServicemapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServicemapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationNames = append(m.ApplicationNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServicemapping(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServicemapping
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipServicemapping(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServicemapping
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
					return 0, ErrIntOverflowServicemapping
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
					return 0, ErrIntOverflowServicemapping
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
				return 0, ErrInvalidLengthServicemapping
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupServicemapping
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthServicemapping
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthServicemapping        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServicemapping          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupServicemapping = fmt.Errorf("proto: unexpected end of group")
)
