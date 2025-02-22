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

// ModbusPDUReadCoilsRequest is the corresponding interface of ModbusPDUReadCoilsRequest
type ModbusPDUReadCoilsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// IsModbusPDUReadCoilsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUReadCoilsRequest()
	// CreateBuilder creates a ModbusPDUReadCoilsRequestBuilder
	CreateModbusPDUReadCoilsRequestBuilder() ModbusPDUReadCoilsRequestBuilder
}

// _ModbusPDUReadCoilsRequest is the data-structure of this message
type _ModbusPDUReadCoilsRequest struct {
	ModbusPDUContract
	StartingAddress uint16
	Quantity        uint16
}

var _ ModbusPDUReadCoilsRequest = (*_ModbusPDUReadCoilsRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadCoilsRequest)(nil)

// NewModbusPDUReadCoilsRequest factory function for _ModbusPDUReadCoilsRequest
func NewModbusPDUReadCoilsRequest(startingAddress uint16, quantity uint16) *_ModbusPDUReadCoilsRequest {
	_result := &_ModbusPDUReadCoilsRequest{
		ModbusPDUContract: NewModbusPDU(),
		StartingAddress:   startingAddress,
		Quantity:          quantity,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUReadCoilsRequestBuilder is a builder for ModbusPDUReadCoilsRequest
type ModbusPDUReadCoilsRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUReadCoilsRequestBuilder
	// WithStartingAddress adds StartingAddress (property field)
	WithStartingAddress(uint16) ModbusPDUReadCoilsRequestBuilder
	// WithQuantity adds Quantity (property field)
	WithQuantity(uint16) ModbusPDUReadCoilsRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ModbusPDUBuilder
	// Build builds the ModbusPDUReadCoilsRequest or returns an error if something is wrong
	Build() (ModbusPDUReadCoilsRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUReadCoilsRequest
}

// NewModbusPDUReadCoilsRequestBuilder() creates a ModbusPDUReadCoilsRequestBuilder
func NewModbusPDUReadCoilsRequestBuilder() ModbusPDUReadCoilsRequestBuilder {
	return &_ModbusPDUReadCoilsRequestBuilder{_ModbusPDUReadCoilsRequest: new(_ModbusPDUReadCoilsRequest)}
}

type _ModbusPDUReadCoilsRequestBuilder struct {
	*_ModbusPDUReadCoilsRequest

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUReadCoilsRequestBuilder) = (*_ModbusPDUReadCoilsRequestBuilder)(nil)

func (b *_ModbusPDUReadCoilsRequestBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
	contract.(*_ModbusPDU)._SubType = b._ModbusPDUReadCoilsRequest
}

func (b *_ModbusPDUReadCoilsRequestBuilder) WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUReadCoilsRequestBuilder {
	return b.WithStartingAddress(startingAddress).WithQuantity(quantity)
}

func (b *_ModbusPDUReadCoilsRequestBuilder) WithStartingAddress(startingAddress uint16) ModbusPDUReadCoilsRequestBuilder {
	b.StartingAddress = startingAddress
	return b
}

func (b *_ModbusPDUReadCoilsRequestBuilder) WithQuantity(quantity uint16) ModbusPDUReadCoilsRequestBuilder {
	b.Quantity = quantity
	return b
}

func (b *_ModbusPDUReadCoilsRequestBuilder) Build() (ModbusPDUReadCoilsRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUReadCoilsRequest.deepCopy(), nil
}

func (b *_ModbusPDUReadCoilsRequestBuilder) MustBuild() ModbusPDUReadCoilsRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ModbusPDUReadCoilsRequestBuilder) Done() ModbusPDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewModbusPDUBuilder().(*_ModbusPDUBuilder)
	}
	return b.parentBuilder
}

func (b *_ModbusPDUReadCoilsRequestBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUReadCoilsRequestBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUReadCoilsRequestBuilder().(*_ModbusPDUReadCoilsRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUReadCoilsRequestBuilder creates a ModbusPDUReadCoilsRequestBuilder
func (b *_ModbusPDUReadCoilsRequest) CreateModbusPDUReadCoilsRequestBuilder() ModbusPDUReadCoilsRequestBuilder {
	if b == nil {
		return NewModbusPDUReadCoilsRequestBuilder()
	}
	return &_ModbusPDUReadCoilsRequestBuilder{_ModbusPDUReadCoilsRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadCoilsRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadCoilsRequest) GetFunctionFlag() uint8 {
	return 0x01
}

func (m *_ModbusPDUReadCoilsRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadCoilsRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadCoilsRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUReadCoilsRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUReadCoilsRequest(structType any) ModbusPDUReadCoilsRequest {
	if casted, ok := structType.(ModbusPDUReadCoilsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadCoilsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadCoilsRequest) GetTypeName() string {
	return "ModbusPDUReadCoilsRequest"
}

func (m *_ModbusPDUReadCoilsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadCoilsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadCoilsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadCoilsRequest ModbusPDUReadCoilsRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadCoilsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadCoilsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startingAddress, err := ReadSimpleField(ctx, "startingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingAddress' field"))
	}
	m.StartingAddress = startingAddress

	quantity, err := ReadSimpleField(ctx, "quantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'quantity' field"))
	}
	m.Quantity = quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadCoilsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadCoilsRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadCoilsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadCoilsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadCoilsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadCoilsRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "startingAddress", m.GetStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "quantity", m.GetQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadCoilsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadCoilsRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadCoilsRequest) IsModbusPDUReadCoilsRequest() {}

func (m *_ModbusPDUReadCoilsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUReadCoilsRequest) deepCopy() *_ModbusPDUReadCoilsRequest {
	if m == nil {
		return nil
	}
	_ModbusPDUReadCoilsRequestCopy := &_ModbusPDUReadCoilsRequest{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.StartingAddress,
		m.Quantity,
	}
	_ModbusPDUReadCoilsRequestCopy.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUReadCoilsRequestCopy
}

func (m *_ModbusPDUReadCoilsRequest) String() string {
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
