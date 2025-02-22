/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CancelResponse is the corresponding interface of CancelResponse
type CancelResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetCancelCount returns CancelCount (property field)
	GetCancelCount() uint32
	// IsCancelResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCancelResponse()
	// CreateBuilder creates a CancelResponseBuilder
	CreateCancelResponseBuilder() CancelResponseBuilder
}

// _CancelResponse is the data-structure of this message
type _CancelResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ResponseHeader
	CancelCount    uint32
}

var _ CancelResponse = (*_CancelResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CancelResponse)(nil)

// NewCancelResponse factory function for _CancelResponse
func NewCancelResponse(responseHeader ResponseHeader, cancelCount uint32) *_CancelResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for CancelResponse must not be nil")
	}
	_result := &_CancelResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		CancelCount:                       cancelCount,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CancelResponseBuilder is a builder for CancelResponse
type CancelResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, cancelCount uint32) CancelResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) CancelResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) CancelResponseBuilder
	// WithCancelCount adds CancelCount (property field)
	WithCancelCount(uint32) CancelResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the CancelResponse or returns an error if something is wrong
	Build() (CancelResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CancelResponse
}

// NewCancelResponseBuilder() creates a CancelResponseBuilder
func NewCancelResponseBuilder() CancelResponseBuilder {
	return &_CancelResponseBuilder{_CancelResponse: new(_CancelResponse)}
}

type _CancelResponseBuilder struct {
	*_CancelResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CancelResponseBuilder) = (*_CancelResponseBuilder)(nil)

func (b *_CancelResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._CancelResponse
}

func (b *_CancelResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, cancelCount uint32) CancelResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithCancelCount(cancelCount)
}

func (b *_CancelResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) CancelResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_CancelResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) CancelResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateResponseHeaderBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseHeaderBuilder failed"))
	}
	return b
}

func (b *_CancelResponseBuilder) WithCancelCount(cancelCount uint32) CancelResponseBuilder {
	b.CancelCount = cancelCount
	return b
}

func (b *_CancelResponseBuilder) Build() (CancelResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CancelResponse.deepCopy(), nil
}

func (b *_CancelResponseBuilder) MustBuild() CancelResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CancelResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_CancelResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CancelResponseBuilder) DeepCopy() any {
	_copy := b.CreateCancelResponseBuilder().(*_CancelResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCancelResponseBuilder creates a CancelResponseBuilder
func (b *_CancelResponse) CreateCancelResponseBuilder() CancelResponseBuilder {
	if b == nil {
		return NewCancelResponseBuilder()
	}
	return &_CancelResponseBuilder{_CancelResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CancelResponse) GetExtensionId() int32 {
	return int32(482)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CancelResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CancelResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_CancelResponse) GetCancelCount() uint32 {
	return m.CancelCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCancelResponse(structType any) CancelResponse {
	if casted, ok := structType.(CancelResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CancelResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CancelResponse) GetTypeName() string {
	return "CancelResponse"
}

func (m *_CancelResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (cancelCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CancelResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CancelResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__cancelResponse CancelResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CancelResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CancelResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	cancelCount, err := ReadSimpleField(ctx, "cancelCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cancelCount' field"))
	}
	m.CancelCount = cancelCount

	if closeErr := readBuffer.CloseContext("CancelResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CancelResponse")
	}

	return m, nil
}

func (m *_CancelResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CancelResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CancelResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CancelResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "cancelCount", m.GetCancelCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'cancelCount' field")
		}

		if popErr := writeBuffer.PopContext("CancelResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CancelResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CancelResponse) IsCancelResponse() {}

func (m *_CancelResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CancelResponse) deepCopy() *_CancelResponse {
	if m == nil {
		return nil
	}
	_CancelResponseCopy := &_CancelResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ResponseHeader](m.ResponseHeader),
		m.CancelCount,
	}
	_CancelResponseCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CancelResponseCopy
}

func (m *_CancelResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
