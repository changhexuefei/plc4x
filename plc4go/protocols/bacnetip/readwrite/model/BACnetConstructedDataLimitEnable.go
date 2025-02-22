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

// BACnetConstructedDataLimitEnable is the corresponding interface of BACnetConstructedDataLimitEnable
type BACnetConstructedDataLimitEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLimitEnable returns LimitEnable (property field)
	GetLimitEnable() BACnetLimitEnableTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLimitEnableTagged
	// IsBACnetConstructedDataLimitEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLimitEnable()
	// CreateBuilder creates a BACnetConstructedDataLimitEnableBuilder
	CreateBACnetConstructedDataLimitEnableBuilder() BACnetConstructedDataLimitEnableBuilder
}

// _BACnetConstructedDataLimitEnable is the data-structure of this message
type _BACnetConstructedDataLimitEnable struct {
	BACnetConstructedDataContract
	LimitEnable BACnetLimitEnableTagged
}

var _ BACnetConstructedDataLimitEnable = (*_BACnetConstructedDataLimitEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLimitEnable)(nil)

// NewBACnetConstructedDataLimitEnable factory function for _BACnetConstructedDataLimitEnable
func NewBACnetConstructedDataLimitEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, limitEnable BACnetLimitEnableTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLimitEnable {
	if limitEnable == nil {
		panic("limitEnable of type BACnetLimitEnableTagged for BACnetConstructedDataLimitEnable must not be nil")
	}
	_result := &_BACnetConstructedDataLimitEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LimitEnable:                   limitEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLimitEnableBuilder is a builder for BACnetConstructedDataLimitEnable
type BACnetConstructedDataLimitEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(limitEnable BACnetLimitEnableTagged) BACnetConstructedDataLimitEnableBuilder
	// WithLimitEnable adds LimitEnable (property field)
	WithLimitEnable(BACnetLimitEnableTagged) BACnetConstructedDataLimitEnableBuilder
	// WithLimitEnableBuilder adds LimitEnable (property field) which is build by the builder
	WithLimitEnableBuilder(func(BACnetLimitEnableTaggedBuilder) BACnetLimitEnableTaggedBuilder) BACnetConstructedDataLimitEnableBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLimitEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataLimitEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLimitEnable
}

// NewBACnetConstructedDataLimitEnableBuilder() creates a BACnetConstructedDataLimitEnableBuilder
func NewBACnetConstructedDataLimitEnableBuilder() BACnetConstructedDataLimitEnableBuilder {
	return &_BACnetConstructedDataLimitEnableBuilder{_BACnetConstructedDataLimitEnable: new(_BACnetConstructedDataLimitEnable)}
}

type _BACnetConstructedDataLimitEnableBuilder struct {
	*_BACnetConstructedDataLimitEnable

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLimitEnableBuilder) = (*_BACnetConstructedDataLimitEnableBuilder)(nil)

func (b *_BACnetConstructedDataLimitEnableBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLimitEnable
}

func (b *_BACnetConstructedDataLimitEnableBuilder) WithMandatoryFields(limitEnable BACnetLimitEnableTagged) BACnetConstructedDataLimitEnableBuilder {
	return b.WithLimitEnable(limitEnable)
}

func (b *_BACnetConstructedDataLimitEnableBuilder) WithLimitEnable(limitEnable BACnetLimitEnableTagged) BACnetConstructedDataLimitEnableBuilder {
	b.LimitEnable = limitEnable
	return b
}

func (b *_BACnetConstructedDataLimitEnableBuilder) WithLimitEnableBuilder(builderSupplier func(BACnetLimitEnableTaggedBuilder) BACnetLimitEnableTaggedBuilder) BACnetConstructedDataLimitEnableBuilder {
	builder := builderSupplier(b.LimitEnable.CreateBACnetLimitEnableTaggedBuilder())
	var err error
	b.LimitEnable, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLimitEnableTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLimitEnableBuilder) Build() (BACnetConstructedDataLimitEnable, error) {
	if b.LimitEnable == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'limitEnable' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLimitEnable.deepCopy(), nil
}

func (b *_BACnetConstructedDataLimitEnableBuilder) MustBuild() BACnetConstructedDataLimitEnable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLimitEnableBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLimitEnableBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLimitEnableBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLimitEnableBuilder().(*_BACnetConstructedDataLimitEnableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLimitEnableBuilder creates a BACnetConstructedDataLimitEnableBuilder
func (b *_BACnetConstructedDataLimitEnable) CreateBACnetConstructedDataLimitEnableBuilder() BACnetConstructedDataLimitEnableBuilder {
	if b == nil {
		return NewBACnetConstructedDataLimitEnableBuilder()
	}
	return &_BACnetConstructedDataLimitEnableBuilder{_BACnetConstructedDataLimitEnable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLimitEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIMIT_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetLimitEnable() BACnetLimitEnableTagged {
	return m.LimitEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetActualValue() BACnetLimitEnableTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLimitEnableTagged(m.GetLimitEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLimitEnable(structType any) BACnetConstructedDataLimitEnable {
	if casted, ok := structType.(BACnetConstructedDataLimitEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLimitEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLimitEnable) GetTypeName() string {
	return "BACnetConstructedDataLimitEnable"
}

func (m *_BACnetConstructedDataLimitEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (limitEnable)
	lengthInBits += m.LimitEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLimitEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLimitEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLimitEnable BACnetConstructedDataLimitEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLimitEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLimitEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	limitEnable, err := ReadSimpleField[BACnetLimitEnableTagged](ctx, "limitEnable", ReadComplex[BACnetLimitEnableTagged](BACnetLimitEnableTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limitEnable' field"))
	}
	m.LimitEnable = limitEnable

	actualValue, err := ReadVirtualField[BACnetLimitEnableTagged](ctx, "actualValue", (*BACnetLimitEnableTagged)(nil), limitEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLimitEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLimitEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLimitEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLimitEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLimitEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLimitEnable")
		}

		if err := WriteSimpleField[BACnetLimitEnableTagged](ctx, "limitEnable", m.GetLimitEnable(), WriteComplex[BACnetLimitEnableTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limitEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLimitEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLimitEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLimitEnable) IsBACnetConstructedDataLimitEnable() {}

func (m *_BACnetConstructedDataLimitEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLimitEnable) deepCopy() *_BACnetConstructedDataLimitEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLimitEnableCopy := &_BACnetConstructedDataLimitEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetLimitEnableTagged](m.LimitEnable),
	}
	_BACnetConstructedDataLimitEnableCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLimitEnableCopy
}

func (m *_BACnetConstructedDataLimitEnable) String() string {
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
