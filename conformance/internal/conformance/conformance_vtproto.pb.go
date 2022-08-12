// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: conformance/conformance.proto

package conformance

import (
	fmt "fmt"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (this *FailureSet) EqualVT(that *FailureSet) bool {
	if this == nil {
		return that == nil || that.String() == ""
	} else if that == nil {
		return this.String() == ""
	}
	if len(this.Failure) != len(that.Failure) {
		return false
	}
	for i := range this.Failure {
		if this.Failure[i] != that.Failure[i] {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConformanceRequest) EqualVT(that *ConformanceRequest) bool {
	if this == nil {
		return that == nil || that.String() == ""
	} else if that == nil {
		return this.String() == ""
	}
	if this.Payload == nil && that.Payload != nil {
		return false
	} else if this.Payload != nil {
		if that.Payload == nil {
			return false
		}
		if string(this.GetProtobufPayload()) != string(that.GetProtobufPayload()) {
			return false
		}
		if this.GetJsonPayload() != that.GetJsonPayload() {
			return false
		}
		if this.GetJspbPayload() != that.GetJspbPayload() {
			return false
		}
		if this.GetTextPayload() != that.GetTextPayload() {
			return false
		}
	}
	if this.RequestedOutputFormat != that.RequestedOutputFormat {
		return false
	}
	if this.MessageType != that.MessageType {
		return false
	}
	if this.TestCategory != that.TestCategory {
		return false
	}
	if !this.JspbEncodingOptions.EqualVT(that.JspbEncodingOptions) {
		return false
	}
	if this.PrintUnknownFields != that.PrintUnknownFields {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *ConformanceResponse) EqualVT(that *ConformanceResponse) bool {
	if this == nil {
		return that == nil || that.String() == ""
	} else if that == nil {
		return this.String() == ""
	}
	if this.Result == nil && that.Result != nil {
		return false
	} else if this.Result != nil {
		if that.Result == nil {
			return false
		}
		if this.GetParseError() != that.GetParseError() {
			return false
		}
		if this.GetRuntimeError() != that.GetRuntimeError() {
			return false
		}
		if string(this.GetProtobufPayload()) != string(that.GetProtobufPayload()) {
			return false
		}
		if this.GetJsonPayload() != that.GetJsonPayload() {
			return false
		}
		if this.GetSkipped() != that.GetSkipped() {
			return false
		}
		if this.GetSerializeError() != that.GetSerializeError() {
			return false
		}
		if this.GetJspbPayload() != that.GetJspbPayload() {
			return false
		}
		if this.GetTextPayload() != that.GetTextPayload() {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *JspbEncodingConfig) EqualVT(that *JspbEncodingConfig) bool {
	if this == nil {
		return that == nil || that.String() == ""
	} else if that == nil {
		return this.String() == ""
	}
	if this.UseJspbArrayAnyFormat != that.UseJspbArrayAnyFormat {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (m *FailureSet) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FailureSet) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *FailureSet) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Failure) > 0 {
		for iNdEx := len(m.Failure) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Failure[iNdEx])
			copy(dAtA[i:], m.Failure[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.Failure[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ConformanceRequest) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConformanceRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Payload.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.PrintUnknownFields {
		i--
		if m.PrintUnknownFields {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.JspbEncodingOptions != nil {
		size, err := m.JspbEncodingOptions.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x32
	}
	if m.TestCategory != 0 {
		i = encodeVarint(dAtA, i, uint64(m.TestCategory))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MessageType) > 0 {
		i -= len(m.MessageType)
		copy(dAtA[i:], m.MessageType)
		i = encodeVarint(dAtA, i, uint64(len(m.MessageType)))
		i--
		dAtA[i] = 0x22
	}
	if m.RequestedOutputFormat != 0 {
		i = encodeVarint(dAtA, i, uint64(m.RequestedOutputFormat))
		i--
		dAtA[i] = 0x18
	}
	return len(dAtA) - i, nil
}

func (m *ConformanceRequest_ProtobufPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceRequest_ProtobufPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ProtobufPayload)
	copy(dAtA[i:], m.ProtobufPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.ProtobufPayload)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *ConformanceRequest_JsonPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceRequest_JsonPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.JsonPayload)
	copy(dAtA[i:], m.JsonPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.JsonPayload)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *ConformanceRequest_JspbPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceRequest_JspbPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.JspbPayload)
	copy(dAtA[i:], m.JspbPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.JspbPayload)))
	i--
	dAtA[i] = 0x3a
	return len(dAtA) - i, nil
}
func (m *ConformanceRequest_TextPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceRequest_TextPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.TextPayload)
	copy(dAtA[i:], m.TextPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.TextPayload)))
	i--
	dAtA[i] = 0x42
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConformanceResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Result.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *ConformanceResponse_ParseError) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_ParseError) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ParseError)
	copy(dAtA[i:], m.ParseError)
	i = encodeVarint(dAtA, i, uint64(len(m.ParseError)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse_RuntimeError) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_RuntimeError) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.RuntimeError)
	copy(dAtA[i:], m.RuntimeError)
	i = encodeVarint(dAtA, i, uint64(len(m.RuntimeError)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse_ProtobufPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_ProtobufPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ProtobufPayload)
	copy(dAtA[i:], m.ProtobufPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.ProtobufPayload)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse_JsonPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_JsonPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.JsonPayload)
	copy(dAtA[i:], m.JsonPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.JsonPayload)))
	i--
	dAtA[i] = 0x22
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse_Skipped) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_Skipped) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Skipped)
	copy(dAtA[i:], m.Skipped)
	i = encodeVarint(dAtA, i, uint64(len(m.Skipped)))
	i--
	dAtA[i] = 0x2a
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse_SerializeError) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_SerializeError) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.SerializeError)
	copy(dAtA[i:], m.SerializeError)
	i = encodeVarint(dAtA, i, uint64(len(m.SerializeError)))
	i--
	dAtA[i] = 0x32
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse_JspbPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_JspbPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.JspbPayload)
	copy(dAtA[i:], m.JspbPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.JspbPayload)))
	i--
	dAtA[i] = 0x3a
	return len(dAtA) - i, nil
}
func (m *ConformanceResponse_TextPayload) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ConformanceResponse_TextPayload) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.TextPayload)
	copy(dAtA[i:], m.TextPayload)
	i = encodeVarint(dAtA, i, uint64(len(m.TextPayload)))
	i--
	dAtA[i] = 0x42
	return len(dAtA) - i, nil
}
func (m *JspbEncodingConfig) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JspbEncodingConfig) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *JspbEncodingConfig) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.UseJspbArrayAnyFormat {
		i--
		if m.UseJspbArrayAnyFormat {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FailureSet) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Failure) > 0 {
		for _, s := range m.Failure {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *ConformanceRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Payload.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.RequestedOutputFormat != 0 {
		n += 1 + sov(uint64(m.RequestedOutputFormat))
	}
	l = len(m.MessageType)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.TestCategory != 0 {
		n += 1 + sov(uint64(m.TestCategory))
	}
	if m.JspbEncodingOptions != nil {
		l = m.JspbEncodingOptions.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.PrintUnknownFields {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *ConformanceRequest_ProtobufPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtobufPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceRequest_JsonPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JsonPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceRequest_JspbPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JspbPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceRequest_TextPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TextPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Result.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *ConformanceResponse_ParseError) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ParseError)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse_RuntimeError) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RuntimeError)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse_ProtobufPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtobufPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse_JsonPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JsonPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse_Skipped) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Skipped)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse_SerializeError) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SerializeError)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse_JspbPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JspbPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *ConformanceResponse_TextPayload) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TextPayload)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *JspbEncodingConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UseJspbArrayAnyFormat {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *FailureSet) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: FailureSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FailureSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failure", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Failure = append(m.Failure, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConformanceRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: ConformanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConformanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtobufPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Payload = &ConformanceRequest_ProtobufPayload{ProtobufPayload: v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = &ConformanceRequest_JsonPayload{JsonPayload: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedOutputFormat", wireType)
			}
			m.RequestedOutputFormat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestedOutputFormat |= WireFormat(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestCategory", wireType)
			}
			m.TestCategory = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TestCategory |= TestCategory(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JspbEncodingOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.JspbEncodingOptions == nil {
				m.JspbEncodingOptions = &JspbEncodingConfig{}
			}
			if err := m.JspbEncodingOptions.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JspbPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = &ConformanceRequest_JspbPayload{JspbPayload: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = &ConformanceRequest_TextPayload{TextPayload: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrintUnknownFields", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrintUnknownFields = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConformanceResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: ConformanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConformanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParseError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &ConformanceResponse_ParseError{ParseError: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &ConformanceResponse_RuntimeError{RuntimeError: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtobufPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Result = &ConformanceResponse_ProtobufPayload{ProtobufPayload: v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &ConformanceResponse_JsonPayload{JsonPayload: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skipped", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &ConformanceResponse_Skipped{Skipped: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerializeError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &ConformanceResponse_SerializeError{SerializeError: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JspbPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &ConformanceResponse_JspbPayload{JspbPayload: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
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
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &ConformanceResponse_TextPayload{TextPayload: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JspbEncodingConfig) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
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
			return fmt.Errorf("proto: JspbEncodingConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JspbEncodingConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UseJspbArrayAnyFormat", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UseJspbArrayAnyFormat = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
