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

// BACnetTagPayloadOctetString is the corresponding interface of BACnetTagPayloadOctetString
type BACnetTagPayloadOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOctets returns Octets (property field)
	GetOctets() []byte
	// IsBACnetTagPayloadOctetString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadOctetString()
	// CreateBuilder creates a BACnetTagPayloadOctetStringBuilder
	CreateBACnetTagPayloadOctetStringBuilder() BACnetTagPayloadOctetStringBuilder
}

// _BACnetTagPayloadOctetString is the data-structure of this message
type _BACnetTagPayloadOctetString struct {
	Octets []byte

	// Arguments.
	ActualLength uint32
}

var _ BACnetTagPayloadOctetString = (*_BACnetTagPayloadOctetString)(nil)

// NewBACnetTagPayloadOctetString factory function for _BACnetTagPayloadOctetString
func NewBACnetTagPayloadOctetString(octets []byte, actualLength uint32) *_BACnetTagPayloadOctetString {
	return &_BACnetTagPayloadOctetString{Octets: octets, ActualLength: actualLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTagPayloadOctetStringBuilder is a builder for BACnetTagPayloadOctetString
type BACnetTagPayloadOctetStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(octets []byte) BACnetTagPayloadOctetStringBuilder
	// WithOctets adds Octets (property field)
	WithOctets(...byte) BACnetTagPayloadOctetStringBuilder
	// WithArgActualLength sets a parser argument
	WithArgActualLength(uint32) BACnetTagPayloadOctetStringBuilder
	// Build builds the BACnetTagPayloadOctetString or returns an error if something is wrong
	Build() (BACnetTagPayloadOctetString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTagPayloadOctetString
}

// NewBACnetTagPayloadOctetStringBuilder() creates a BACnetTagPayloadOctetStringBuilder
func NewBACnetTagPayloadOctetStringBuilder() BACnetTagPayloadOctetStringBuilder {
	return &_BACnetTagPayloadOctetStringBuilder{_BACnetTagPayloadOctetString: new(_BACnetTagPayloadOctetString)}
}

type _BACnetTagPayloadOctetStringBuilder struct {
	*_BACnetTagPayloadOctetString

	err *utils.MultiError
}

var _ (BACnetTagPayloadOctetStringBuilder) = (*_BACnetTagPayloadOctetStringBuilder)(nil)

func (b *_BACnetTagPayloadOctetStringBuilder) WithMandatoryFields(octets []byte) BACnetTagPayloadOctetStringBuilder {
	return b.WithOctets(octets...)
}

func (b *_BACnetTagPayloadOctetStringBuilder) WithOctets(octets ...byte) BACnetTagPayloadOctetStringBuilder {
	b.Octets = octets
	return b
}

func (b *_BACnetTagPayloadOctetStringBuilder) WithArgActualLength(actualLength uint32) BACnetTagPayloadOctetStringBuilder {
	b.ActualLength = actualLength
	return b
}

func (b *_BACnetTagPayloadOctetStringBuilder) Build() (BACnetTagPayloadOctetString, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTagPayloadOctetString.deepCopy(), nil
}

func (b *_BACnetTagPayloadOctetStringBuilder) MustBuild() BACnetTagPayloadOctetString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTagPayloadOctetStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTagPayloadOctetStringBuilder().(*_BACnetTagPayloadOctetStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTagPayloadOctetStringBuilder creates a BACnetTagPayloadOctetStringBuilder
func (b *_BACnetTagPayloadOctetString) CreateBACnetTagPayloadOctetStringBuilder() BACnetTagPayloadOctetStringBuilder {
	if b == nil {
		return NewBACnetTagPayloadOctetStringBuilder()
	}
	return &_BACnetTagPayloadOctetStringBuilder{_BACnetTagPayloadOctetString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadOctetString) GetOctets() []byte {
	return m.Octets
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadOctetString(structType any) BACnetTagPayloadOctetString {
	if casted, ok := structType.(BACnetTagPayloadOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadOctetString) GetTypeName() string {
	return "BACnetTagPayloadOctetString"
}

func (m *_BACnetTagPayloadOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Octets) > 0 {
		lengthInBits += 8 * uint16(len(m.Octets))
	}

	return lengthInBits
}

func (m *_BACnetTagPayloadOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadOctetStringParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadOctetString, error) {
	return BACnetTagPayloadOctetStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadOctetStringParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadOctetString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadOctetString, error) {
		return BACnetTagPayloadOctetStringParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetTagPayloadOctetStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadOctetString, error) {
	v, err := (&_BACnetTagPayloadOctetString{ActualLength: actualLength}).parse(ctx, readBuffer, actualLength)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadOctetString) parse(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (__bACnetTagPayloadOctetString BACnetTagPayloadOctetString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	octets, err := readBuffer.ReadByteArray("octets", int(actualLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octets' field"))
	}
	m.Octets = octets

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadOctetString")
	}

	return m, nil
}

func (m *_BACnetTagPayloadOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadOctetString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadOctetString")
	}

	if err := WriteByteArrayField(ctx, "octets", m.GetOctets(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octets' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadOctetString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadOctetString")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadOctetString) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadOctetString) IsBACnetTagPayloadOctetString() {}

func (m *_BACnetTagPayloadOctetString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagPayloadOctetString) deepCopy() *_BACnetTagPayloadOctetString {
	if m == nil {
		return nil
	}
	_BACnetTagPayloadOctetStringCopy := &_BACnetTagPayloadOctetString{
		utils.DeepCopySlice[byte, byte](m.Octets),
		m.ActualLength,
	}
	return _BACnetTagPayloadOctetStringCopy
}

func (m *_BACnetTagPayloadOctetString) String() string {
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
