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

// BACnetChannelValueOctetString is the corresponding interface of BACnetChannelValueOctetString
type BACnetChannelValueOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
	// IsBACnetChannelValueOctetString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueOctetString()
	// CreateBuilder creates a BACnetChannelValueOctetStringBuilder
	CreateBACnetChannelValueOctetStringBuilder() BACnetChannelValueOctetStringBuilder
}

// _BACnetChannelValueOctetString is the data-structure of this message
type _BACnetChannelValueOctetString struct {
	BACnetChannelValueContract
	OctetStringValue BACnetApplicationTagOctetString
}

var _ BACnetChannelValueOctetString = (*_BACnetChannelValueOctetString)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueOctetString)(nil)

// NewBACnetChannelValueOctetString factory function for _BACnetChannelValueOctetString
func NewBACnetChannelValueOctetString(peekedTagHeader BACnetTagHeader, octetStringValue BACnetApplicationTagOctetString) *_BACnetChannelValueOctetString {
	if octetStringValue == nil {
		panic("octetStringValue of type BACnetApplicationTagOctetString for BACnetChannelValueOctetString must not be nil")
	}
	_result := &_BACnetChannelValueOctetString{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		OctetStringValue:           octetStringValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueOctetStringBuilder is a builder for BACnetChannelValueOctetString
type BACnetChannelValueOctetStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(octetStringValue BACnetApplicationTagOctetString) BACnetChannelValueOctetStringBuilder
	// WithOctetStringValue adds OctetStringValue (property field)
	WithOctetStringValue(BACnetApplicationTagOctetString) BACnetChannelValueOctetStringBuilder
	// WithOctetStringValueBuilder adds OctetStringValue (property field) which is build by the builder
	WithOctetStringValueBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetChannelValueOctetStringBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetChannelValueBuilder
	// Build builds the BACnetChannelValueOctetString or returns an error if something is wrong
	Build() (BACnetChannelValueOctetString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueOctetString
}

// NewBACnetChannelValueOctetStringBuilder() creates a BACnetChannelValueOctetStringBuilder
func NewBACnetChannelValueOctetStringBuilder() BACnetChannelValueOctetStringBuilder {
	return &_BACnetChannelValueOctetStringBuilder{_BACnetChannelValueOctetString: new(_BACnetChannelValueOctetString)}
}

type _BACnetChannelValueOctetStringBuilder struct {
	*_BACnetChannelValueOctetString

	parentBuilder *_BACnetChannelValueBuilder

	err *utils.MultiError
}

var _ (BACnetChannelValueOctetStringBuilder) = (*_BACnetChannelValueOctetStringBuilder)(nil)

func (b *_BACnetChannelValueOctetStringBuilder) setParent(contract BACnetChannelValueContract) {
	b.BACnetChannelValueContract = contract
	contract.(*_BACnetChannelValue)._SubType = b._BACnetChannelValueOctetString
}

func (b *_BACnetChannelValueOctetStringBuilder) WithMandatoryFields(octetStringValue BACnetApplicationTagOctetString) BACnetChannelValueOctetStringBuilder {
	return b.WithOctetStringValue(octetStringValue)
}

func (b *_BACnetChannelValueOctetStringBuilder) WithOctetStringValue(octetStringValue BACnetApplicationTagOctetString) BACnetChannelValueOctetStringBuilder {
	b.OctetStringValue = octetStringValue
	return b
}

func (b *_BACnetChannelValueOctetStringBuilder) WithOctetStringValueBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetChannelValueOctetStringBuilder {
	builder := builderSupplier(b.OctetStringValue.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.OctetStringValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetChannelValueOctetStringBuilder) Build() (BACnetChannelValueOctetString, error) {
	if b.OctetStringValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'octetStringValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetChannelValueOctetString.deepCopy(), nil
}

func (b *_BACnetChannelValueOctetStringBuilder) MustBuild() BACnetChannelValueOctetString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetChannelValueOctetStringBuilder) Done() BACnetChannelValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetChannelValueBuilder().(*_BACnetChannelValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetChannelValueOctetStringBuilder) buildForBACnetChannelValue() (BACnetChannelValue, error) {
	return b.Build()
}

func (b *_BACnetChannelValueOctetStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetChannelValueOctetStringBuilder().(*_BACnetChannelValueOctetStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetChannelValueOctetStringBuilder creates a BACnetChannelValueOctetStringBuilder
func (b *_BACnetChannelValueOctetString) CreateBACnetChannelValueOctetStringBuilder() BACnetChannelValueOctetStringBuilder {
	if b == nil {
		return NewBACnetChannelValueOctetStringBuilder()
	}
	return &_BACnetChannelValueOctetStringBuilder{_BACnetChannelValueOctetString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueOctetString) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueOctetString) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueOctetString(structType any) BACnetChannelValueOctetString {
	if casted, ok := structType.(BACnetChannelValueOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueOctetString) GetTypeName() string {
	return "BACnetChannelValueOctetString"
}

func (m *_BACnetChannelValueOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueOctetString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueOctetString BACnetChannelValueOctetString, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	octetStringValue, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "octetStringValue", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octetStringValue' field"))
	}
	m.OctetStringValue = octetStringValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueOctetString")
	}

	return m, nil
}

func (m *_BACnetChannelValueOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueOctetString")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "octetStringValue", m.GetOctetStringValue(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueOctetString")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueOctetString) IsBACnetChannelValueOctetString() {}

func (m *_BACnetChannelValueOctetString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueOctetString) deepCopy() *_BACnetChannelValueOctetString {
	if m == nil {
		return nil
	}
	_BACnetChannelValueOctetStringCopy := &_BACnetChannelValueOctetString{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.OctetStringValue),
	}
	_BACnetChannelValueOctetStringCopy.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueOctetStringCopy
}

func (m *_BACnetChannelValueOctetString) String() string {
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
