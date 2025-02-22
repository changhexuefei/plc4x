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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageSubscribeAnalogPinValue is the corresponding interface of FirmataMessageSubscribeAnalogPinValue
type FirmataMessageSubscribeAnalogPinValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	FirmataMessage
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetEnable returns Enable (property field)
	GetEnable() bool
	// IsFirmataMessageSubscribeAnalogPinValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFirmataMessageSubscribeAnalogPinValue()
	// CreateBuilder creates a FirmataMessageSubscribeAnalogPinValueBuilder
	CreateFirmataMessageSubscribeAnalogPinValueBuilder() FirmataMessageSubscribeAnalogPinValueBuilder
}

// _FirmataMessageSubscribeAnalogPinValue is the data-structure of this message
type _FirmataMessageSubscribeAnalogPinValue struct {
	FirmataMessageContract
	Pin    uint8
	Enable bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ FirmataMessageSubscribeAnalogPinValue = (*_FirmataMessageSubscribeAnalogPinValue)(nil)
var _ FirmataMessageRequirements = (*_FirmataMessageSubscribeAnalogPinValue)(nil)

// NewFirmataMessageSubscribeAnalogPinValue factory function for _FirmataMessageSubscribeAnalogPinValue
func NewFirmataMessageSubscribeAnalogPinValue(pin uint8, enable bool, response bool) *_FirmataMessageSubscribeAnalogPinValue {
	_result := &_FirmataMessageSubscribeAnalogPinValue{
		FirmataMessageContract: NewFirmataMessage(response),
		Pin:                    pin,
		Enable:                 enable,
	}
	_result.FirmataMessageContract.(*_FirmataMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FirmataMessageSubscribeAnalogPinValueBuilder is a builder for FirmataMessageSubscribeAnalogPinValue
type FirmataMessageSubscribeAnalogPinValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(pin uint8, enable bool) FirmataMessageSubscribeAnalogPinValueBuilder
	// WithPin adds Pin (property field)
	WithPin(uint8) FirmataMessageSubscribeAnalogPinValueBuilder
	// WithEnable adds Enable (property field)
	WithEnable(bool) FirmataMessageSubscribeAnalogPinValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() FirmataMessageBuilder
	// Build builds the FirmataMessageSubscribeAnalogPinValue or returns an error if something is wrong
	Build() (FirmataMessageSubscribeAnalogPinValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FirmataMessageSubscribeAnalogPinValue
}

// NewFirmataMessageSubscribeAnalogPinValueBuilder() creates a FirmataMessageSubscribeAnalogPinValueBuilder
func NewFirmataMessageSubscribeAnalogPinValueBuilder() FirmataMessageSubscribeAnalogPinValueBuilder {
	return &_FirmataMessageSubscribeAnalogPinValueBuilder{_FirmataMessageSubscribeAnalogPinValue: new(_FirmataMessageSubscribeAnalogPinValue)}
}

type _FirmataMessageSubscribeAnalogPinValueBuilder struct {
	*_FirmataMessageSubscribeAnalogPinValue

	parentBuilder *_FirmataMessageBuilder

	err *utils.MultiError
}

var _ (FirmataMessageSubscribeAnalogPinValueBuilder) = (*_FirmataMessageSubscribeAnalogPinValueBuilder)(nil)

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) setParent(contract FirmataMessageContract) {
	b.FirmataMessageContract = contract
	contract.(*_FirmataMessage)._SubType = b._FirmataMessageSubscribeAnalogPinValue
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) WithMandatoryFields(pin uint8, enable bool) FirmataMessageSubscribeAnalogPinValueBuilder {
	return b.WithPin(pin).WithEnable(enable)
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) WithPin(pin uint8) FirmataMessageSubscribeAnalogPinValueBuilder {
	b.Pin = pin
	return b
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) WithEnable(enable bool) FirmataMessageSubscribeAnalogPinValueBuilder {
	b.Enable = enable
	return b
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) Build() (FirmataMessageSubscribeAnalogPinValue, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FirmataMessageSubscribeAnalogPinValue.deepCopy(), nil
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) MustBuild() FirmataMessageSubscribeAnalogPinValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) Done() FirmataMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewFirmataMessageBuilder().(*_FirmataMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) buildForFirmataMessage() (FirmataMessage, error) {
	return b.Build()
}

func (b *_FirmataMessageSubscribeAnalogPinValueBuilder) DeepCopy() any {
	_copy := b.CreateFirmataMessageSubscribeAnalogPinValueBuilder().(*_FirmataMessageSubscribeAnalogPinValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFirmataMessageSubscribeAnalogPinValueBuilder creates a FirmataMessageSubscribeAnalogPinValueBuilder
func (b *_FirmataMessageSubscribeAnalogPinValue) CreateFirmataMessageSubscribeAnalogPinValueBuilder() FirmataMessageSubscribeAnalogPinValueBuilder {
	if b == nil {
		return NewFirmataMessageSubscribeAnalogPinValueBuilder()
	}
	return &_FirmataMessageSubscribeAnalogPinValueBuilder{_FirmataMessageSubscribeAnalogPinValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageSubscribeAnalogPinValue) GetMessageType() uint8 {
	return 0xC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageSubscribeAnalogPinValue) GetParent() FirmataMessageContract {
	return m.FirmataMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageSubscribeAnalogPinValue) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetEnable() bool {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFirmataMessageSubscribeAnalogPinValue(structType any) FirmataMessageSubscribeAnalogPinValue {
	if casted, ok := structType.(FirmataMessageSubscribeAnalogPinValue); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageSubscribeAnalogPinValue); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetTypeName() string {
	return "FirmataMessageSubscribeAnalogPinValue"
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.FirmataMessageContract.(*_FirmataMessage).getLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 4

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enable)
	lengthInBits += 1

	return lengthInBits
}

func (m *_FirmataMessageSubscribeAnalogPinValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FirmataMessageSubscribeAnalogPinValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_FirmataMessage, response bool) (__firmataMessageSubscribeAnalogPinValue FirmataMessageSubscribeAnalogPinValue, err error) {
	m.FirmataMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataMessageSubscribeAnalogPinValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageSubscribeAnalogPinValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	pin, err := ReadSimpleField(ctx, "pin", ReadUnsignedByte(readBuffer, uint8(4)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pin' field"))
	}
	m.Pin = pin

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enable, err := ReadSimpleField(ctx, "enable", ReadBoolean(readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enable' field"))
	}
	m.Enable = enable

	if closeErr := readBuffer.CloseContext("FirmataMessageSubscribeAnalogPinValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageSubscribeAnalogPinValue")
	}

	return m, nil
}

func (m *_FirmataMessageSubscribeAnalogPinValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageSubscribeAnalogPinValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageSubscribeAnalogPinValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageSubscribeAnalogPinValue")
		}

		if err := WriteSimpleField[uint8](ctx, "pin", m.GetPin(), WriteUnsignedByte(writeBuffer, 4), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pin' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enable", m.GetEnable(), WriteBoolean(writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'enable' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageSubscribeAnalogPinValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageSubscribeAnalogPinValue")
		}
		return nil
	}
	return m.FirmataMessageContract.(*_FirmataMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageSubscribeAnalogPinValue) IsFirmataMessageSubscribeAnalogPinValue() {}

func (m *_FirmataMessageSubscribeAnalogPinValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FirmataMessageSubscribeAnalogPinValue) deepCopy() *_FirmataMessageSubscribeAnalogPinValue {
	if m == nil {
		return nil
	}
	_FirmataMessageSubscribeAnalogPinValueCopy := &_FirmataMessageSubscribeAnalogPinValue{
		m.FirmataMessageContract.(*_FirmataMessage).deepCopy(),
		m.Pin,
		m.Enable,
		m.reservedField0,
	}
	_FirmataMessageSubscribeAnalogPinValueCopy.FirmataMessageContract.(*_FirmataMessage)._SubType = m
	return _FirmataMessageSubscribeAnalogPinValueCopy
}

func (m *_FirmataMessageSubscribeAnalogPinValue) String() string {
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
