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
// source: k8s.io/kubernetes/vendor/k8s.io/api/coordination/v1beta1/generated.proto

package v1beta1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/apis/meta/v1"

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

func (m *Lease) Reset()      { *m = Lease{} }
func (*Lease) ProtoMessage() {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_daca6bcd2ff63a80, []int{0}
}
func (m *Lease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(m, src)
}
func (m *Lease) XXX_Size() int {
	return m.Size()
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

func (m *LeaseList) Reset()      { *m = LeaseList{} }
func (*LeaseList) ProtoMessage() {}
func (*LeaseList) Descriptor() ([]byte, []int) {
	return fileDescriptor_daca6bcd2ff63a80, []int{1}
}
func (m *LeaseList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LeaseList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseList.Merge(m, src)
}
func (m *LeaseList) XXX_Size() int {
	return m.Size()
}
func (m *LeaseList) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseList.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseList proto.InternalMessageInfo

func (m *LeaseSpec) Reset()      { *m = LeaseSpec{} }
func (*LeaseSpec) ProtoMessage() {}
func (*LeaseSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_daca6bcd2ff63a80, []int{2}
}
func (m *LeaseSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaseSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LeaseSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaseSpec.Merge(m, src)
}
func (m *LeaseSpec) XXX_Size() int {
	return m.Size()
}
func (m *LeaseSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaseSpec.DiscardUnknown(m)
}

var xxx_messageInfo_LeaseSpec proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Lease)(nil), "k8s.io.api.coordination.v1beta1.Lease")
	proto.RegisterType((*LeaseList)(nil), "k8s.io.api.coordination.v1beta1.LeaseList")
	proto.RegisterType((*LeaseSpec)(nil), "k8s.io.api.coordination.v1beta1.LeaseSpec")
}

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/api/coordination/v1beta1/generated.proto", fileDescriptor_daca6bcd2ff63a80)
}

var fileDescriptor_daca6bcd2ff63a80 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe3, 0xb6, 0x11, 0xcd, 0x86, 0x96, 0xc8, 0xca, 0xc1, 0xca, 0xc1, 0xae, 0x72, 0x40,
	0x15, 0x52, 0x77, 0x49, 0x85, 0x10, 0xe2, 0x04, 0x16, 0x87, 0x56, 0xb8, 0x42, 0x72, 0x7b, 0x42,
	0x3d, 0xb0, 0xb6, 0x07, 0x67, 0x49, 0xed, 0x35, 0xbb, 0xeb, 0xa0, 0xde, 0x78, 0x04, 0xae, 0xbc,
	0x08, 0xbc, 0x42, 0x8e, 0x3d, 0xf6, 0x64, 0x11, 0xf3, 0x22, 0xc8, 0x1b, 0xb7, 0x09, 0x49, 0x51,
	0x23, 0x6e, 0xde, 0x99, 0xf9, 0xbf, 0xf9, 0xe7, 0x37, 0x3a, 0x1a, 0xbd, 0x90, 0x98, 0x71, 0x32,
	0xca, 0x03, 0x10, 0x29, 0x28, 0x90, 0x64, 0x0c, 0x69, 0xc4, 0x05, 0xa9, 0x1b, 0x34, 0x63, 0x24,
	0xe4, 0x5c, 0x44, 0x2c, 0xa5, 0x8a, 0xf1, 0x94, 0x8c, 0x07, 0x01, 0x28, 0x3a, 0x20, 0x31, 0xa4,
	0x20, 0xa8, 0x82, 0x08, 0x67, 0x82, 0x2b, 0x6e, 0x3a, 0x33, 0x01, 0xa6, 0x19, 0xc3, 0x8b, 0x02,
	0x5c, 0x0b, 0x7a, 0x07, 0x31, 0x53, 0xc3, 0x3c, 0xc0, 0x21, 0x4f, 0x48, 0xcc, 0x63, 0x4e, 0xb4,
	0x2e, 0xc8, 0x3f, 0xea, 0x97, 0x7e, 0xe8, 0xaf, 0x19, 0xaf, 0xf7, 0x6c, 0x6e, 0x20, 0xa1, 0xe1,
	0x90, 0xa5, 0x20, 0x2e, 0x49, 0x36, 0x8a, 0xab, 0x82, 0x24, 0x09, 0x28, 0x4a, 0xc6, 0x2b, 0x2e,
	0x7a, 0xe4, 0x5f, 0x2a, 0x91, 0xa7, 0x8a, 0x25, 0xb0, 0x22, 0x78, 0x7e, 0x9f, 0x40, 0x86, 0x43,
	0x48, 0xe8, 0xb2, 0xae, 0xff, 0xd3, 0x40, 0x4d, 0x0f, 0xa8, 0x04, 0xf3, 0x03, 0xda, 0xae, 0xdc,
	0x44, 0x54, 0x51, 0xcb, 0xd8, 0x33, 0xf6, 0xdb, 0x87, 0x4f, 0xf1, 0x3c, 0x8b, 0x5b, 0x28, 0xce,
	0x46, 0x71, 0x55, 0x90, 0xb8, 0x9a, 0xc6, 0xe3, 0x01, 0x7e, 0x17, 0x7c, 0x82, 0x50, 0x9d, 0x80,
	0xa2, 0xae, 0x39, 0x29, 0x9c, 0x46, 0x59, 0x38, 0x68, 0x5e, 0xf3, 0x6f, 0xa9, 0xa6, 0x87, 0xb6,
	0x64, 0x06, 0xa1, 0xb5, 0xa1, 0xe9, 0x4f, 0xf0, 0x3d, 0x49, 0x63, 0xed, 0xeb, 0x34, 0x83, 0xd0,
	0x7d, 0x58, 0x73, 0xb7, 0xaa, 0x97, 0xaf, 0x29, 0xfd, 0x1f, 0x06, 0x6a, 0xe9, 0x09, 0x8f, 0x49,
	0x65, 0x9e, 0xaf, 0xb8, 0xc7, 0xeb, 0xb9, 0xaf, 0xd4, 0xda, 0x7b, 0xa7, 0xde, 0xb1, 0x7d, 0x53,
	0x59, 0x70, 0xfe, 0x16, 0x35, 0x99, 0x82, 0x44, 0x5a, 0x1b, 0x7b, 0x9b, 0xfb, 0xed, 0xc3, 0xc7,
	0xeb, 0x59, 0x77, 0x77, 0x6a, 0x64, 0xf3, 0xb8, 0x12, 0xfb, 0x33, 0x46, 0xff, 0xfb, 0x66, 0x6d,
	0xbc, 0x3a, 0xc6, 0x7c, 0x89, 0x76, 0x87, 0xfc, 0x22, 0x02, 0x71, 0x1c, 0x41, 0xaa, 0x98, 0xba,
	0xd4, 0xf6, 0x5b, 0xae, 0x59, 0x16, 0xce, 0xee, 0xd1, 0x5f, 0x1d, 0x7f, 0x69, 0xd2, 0xf4, 0x50,
	0xf7, 0xa2, 0x02, 0xbd, 0xc9, 0x85, 0x5e, 0x7f, 0x0a, 0x21, 0x4f, 0x23, 0xa9, 0x03, 0x6e, 0xba,
	0x56, 0x59, 0x38, 0x5d, 0xef, 0x8e, 0xbe, 0x7f, 0xa7, 0xca, 0x0c, 0x50, 0x9b, 0x86, 0x9f, 0x73,
	0x26, 0xe0, 0x8c, 0x25, 0x60, 0x6d, 0xea, 0x14, 0xc9, 0x7a, 0x29, 0x9e, 0xb0, 0x50, 0xf0, 0x4a,
	0xe6, 0x3e, 0x2a, 0x0b, 0xa7, 0xfd, 0x7a, 0xce, 0xf1, 0x17, 0xa1, 0xe6, 0x39, 0x6a, 0x09, 0x48,
	0xe1, 0x8b, 0xde, 0xb0, 0xf5, 0x7f, 0x1b, 0x76, 0xca, 0xc2, 0x69, 0xf9, 0x37, 0x14, 0x7f, 0x0e,
	0x34, 0x5f, 0xa1, 0x8e, 0xbe, 0xec, 0x4c, 0xd0, 0x54, 0xb2, 0xea, 0x36, 0x69, 0x35, 0x75, 0x16,
	0xdd, 0xb2, 0x70, 0x3a, 0xde, 0x52, 0xcf, 0x5f, 0x99, 0x76, 0x0f, 0x26, 0x53, 0xbb, 0x71, 0x35,
	0xb5, 0x1b, 0xd7, 0x53, 0xbb, 0xf1, 0xb5, 0xb4, 0x8d, 0x49, 0x69, 0x1b, 0x57, 0xa5, 0x6d, 0x5c,
	0x97, 0xb6, 0xf1, 0xab, 0xb4, 0x8d, 0x6f, 0xbf, 0xed, 0xc6, 0xfb, 0x07, 0xf5, 0x6f, 0xfe, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x51, 0x34, 0x6a, 0x0f, 0x77, 0x04, 0x00, 0x00,
}

func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LeaseList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LeaseSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaseSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaseSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LeaseTransitions != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.LeaseTransitions))
		i--
		dAtA[i] = 0x28
	}
	if m.RenewTime != nil {
		{
			size, err := m.RenewTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.AcquireTime != nil {
		{
			size, err := m.AcquireTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LeaseDurationSeconds != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.LeaseDurationSeconds))
		i--
		dAtA[i] = 0x10
	}
	if m.HolderIdentity != nil {
		i -= len(*m.HolderIdentity)
		copy(dAtA[i:], *m.HolderIdentity)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.HolderIdentity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *LeaseList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *LeaseSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HolderIdentity != nil {
		l = len(*m.HolderIdentity)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LeaseDurationSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.LeaseDurationSeconds))
	}
	if m.AcquireTime != nil {
		l = m.AcquireTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.RenewTime != nil {
		l = m.RenewTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LeaseTransitions != nil {
		n += 1 + sovGenerated(uint64(*m.LeaseTransitions))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Lease) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Lease{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "LeaseSpec", "LeaseSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LeaseList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Lease{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Lease", "Lease", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&LeaseList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *LeaseSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LeaseSpec{`,
		`HolderIdentity:` + valueToStringGenerated(this.HolderIdentity) + `,`,
		`LeaseDurationSeconds:` + valueToStringGenerated(this.LeaseDurationSeconds) + `,`,
		`AcquireTime:` + strings.Replace(fmt.Sprintf("%v", this.AcquireTime), "MicroTime", "v1.MicroTime", 1) + `,`,
		`RenewTime:` + strings.Replace(fmt.Sprintf("%v", this.RenewTime), "MicroTime", "v1.MicroTime", 1) + `,`,
		`LeaseTransitions:` + valueToStringGenerated(this.LeaseTransitions) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *LeaseList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: LeaseList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Lease{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *LeaseSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: LeaseSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HolderIdentity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.HolderIdentity = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseDurationSeconds", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeaseDurationSeconds = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcquireTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AcquireTime == nil {
				m.AcquireTime = &v1.MicroTime{}
			}
			if err := m.AcquireTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RenewTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RenewTime == nil {
				m.RenewTime = &v1.MicroTime{}
			}
			if err := m.RenewTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseTransitions", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeaseTransitions = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
