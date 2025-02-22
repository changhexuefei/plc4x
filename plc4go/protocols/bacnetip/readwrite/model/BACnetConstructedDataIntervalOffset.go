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

// BACnetConstructedDataIntervalOffset is the corresponding interface of BACnetConstructedDataIntervalOffset
type BACnetConstructedDataIntervalOffset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIntervalOffset returns IntervalOffset (property field)
	GetIntervalOffset() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataIntervalOffset is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntervalOffset()
	// CreateBuilder creates a BACnetConstructedDataIntervalOffsetBuilder
	CreateBACnetConstructedDataIntervalOffsetBuilder() BACnetConstructedDataIntervalOffsetBuilder
}

// _BACnetConstructedDataIntervalOffset is the data-structure of this message
type _BACnetConstructedDataIntervalOffset struct {
	BACnetConstructedDataContract
	IntervalOffset BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataIntervalOffset = (*_BACnetConstructedDataIntervalOffset)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntervalOffset)(nil)

// NewBACnetConstructedDataIntervalOffset factory function for _BACnetConstructedDataIntervalOffset
func NewBACnetConstructedDataIntervalOffset(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, intervalOffset BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntervalOffset {
	if intervalOffset == nil {
		panic("intervalOffset of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataIntervalOffset must not be nil")
	}
	_result := &_BACnetConstructedDataIntervalOffset{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		IntervalOffset:                intervalOffset,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntervalOffsetBuilder is a builder for BACnetConstructedDataIntervalOffset
type BACnetConstructedDataIntervalOffsetBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(intervalOffset BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntervalOffsetBuilder
	// WithIntervalOffset adds IntervalOffset (property field)
	WithIntervalOffset(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntervalOffsetBuilder
	// WithIntervalOffsetBuilder adds IntervalOffset (property field) which is build by the builder
	WithIntervalOffsetBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIntervalOffsetBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIntervalOffset or returns an error if something is wrong
	Build() (BACnetConstructedDataIntervalOffset, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntervalOffset
}

// NewBACnetConstructedDataIntervalOffsetBuilder() creates a BACnetConstructedDataIntervalOffsetBuilder
func NewBACnetConstructedDataIntervalOffsetBuilder() BACnetConstructedDataIntervalOffsetBuilder {
	return &_BACnetConstructedDataIntervalOffsetBuilder{_BACnetConstructedDataIntervalOffset: new(_BACnetConstructedDataIntervalOffset)}
}

type _BACnetConstructedDataIntervalOffsetBuilder struct {
	*_BACnetConstructedDataIntervalOffset

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntervalOffsetBuilder) = (*_BACnetConstructedDataIntervalOffsetBuilder)(nil)

func (b *_BACnetConstructedDataIntervalOffsetBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataIntervalOffset
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) WithMandatoryFields(intervalOffset BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntervalOffsetBuilder {
	return b.WithIntervalOffset(intervalOffset)
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) WithIntervalOffset(intervalOffset BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntervalOffsetBuilder {
	b.IntervalOffset = intervalOffset
	return b
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) WithIntervalOffsetBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIntervalOffsetBuilder {
	builder := builderSupplier(b.IntervalOffset.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.IntervalOffset, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) Build() (BACnetConstructedDataIntervalOffset, error) {
	if b.IntervalOffset == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'intervalOffset' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntervalOffset.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) MustBuild() BACnetConstructedDataIntervalOffset {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntervalOffsetBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntervalOffsetBuilder().(*_BACnetConstructedDataIntervalOffsetBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntervalOffsetBuilder creates a BACnetConstructedDataIntervalOffsetBuilder
func (b *_BACnetConstructedDataIntervalOffset) CreateBACnetConstructedDataIntervalOffsetBuilder() BACnetConstructedDataIntervalOffsetBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntervalOffsetBuilder()
	}
	return &_BACnetConstructedDataIntervalOffsetBuilder{_BACnetConstructedDataIntervalOffset: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIntervalOffset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERVAL_OFFSET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetIntervalOffset() BACnetApplicationTagUnsignedInteger {
	return m.IntervalOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntervalOffset) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIntervalOffset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntervalOffset(structType any) BACnetConstructedDataIntervalOffset {
	if casted, ok := structType.(BACnetConstructedDataIntervalOffset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntervalOffset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntervalOffset) GetTypeName() string {
	return "BACnetConstructedDataIntervalOffset"
}

func (m *_BACnetConstructedDataIntervalOffset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (intervalOffset)
	lengthInBits += m.IntervalOffset.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntervalOffset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntervalOffset) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntervalOffset BACnetConstructedDataIntervalOffset, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntervalOffset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntervalOffset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	intervalOffset, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "intervalOffset", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'intervalOffset' field"))
	}
	m.IntervalOffset = intervalOffset

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), intervalOffset)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntervalOffset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntervalOffset")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntervalOffset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntervalOffset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntervalOffset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntervalOffset")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "intervalOffset", m.GetIntervalOffset(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'intervalOffset' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntervalOffset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntervalOffset")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntervalOffset) IsBACnetConstructedDataIntervalOffset() {}

func (m *_BACnetConstructedDataIntervalOffset) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntervalOffset) deepCopy() *_BACnetConstructedDataIntervalOffset {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntervalOffsetCopy := &_BACnetConstructedDataIntervalOffset{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.IntervalOffset),
	}
	_BACnetConstructedDataIntervalOffsetCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntervalOffsetCopy
}

func (m *_BACnetConstructedDataIntervalOffset) String() string {
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
