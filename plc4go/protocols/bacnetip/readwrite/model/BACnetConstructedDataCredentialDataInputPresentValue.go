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

// BACnetConstructedDataCredentialDataInputPresentValue is the corresponding interface of BACnetConstructedDataCredentialDataInputPresentValue
type BACnetConstructedDataCredentialDataInputPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetAuthenticationFactor
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAuthenticationFactor
	// IsBACnetConstructedDataCredentialDataInputPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCredentialDataInputPresentValue()
	// CreateBuilder creates a BACnetConstructedDataCredentialDataInputPresentValueBuilder
	CreateBACnetConstructedDataCredentialDataInputPresentValueBuilder() BACnetConstructedDataCredentialDataInputPresentValueBuilder
}

// _BACnetConstructedDataCredentialDataInputPresentValue is the data-structure of this message
type _BACnetConstructedDataCredentialDataInputPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetAuthenticationFactor
}

var _ BACnetConstructedDataCredentialDataInputPresentValue = (*_BACnetConstructedDataCredentialDataInputPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCredentialDataInputPresentValue)(nil)

// NewBACnetConstructedDataCredentialDataInputPresentValue factory function for _BACnetConstructedDataCredentialDataInputPresentValue
func NewBACnetConstructedDataCredentialDataInputPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetAuthenticationFactor, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialDataInputPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetAuthenticationFactor for BACnetConstructedDataCredentialDataInputPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataCredentialDataInputPresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCredentialDataInputPresentValueBuilder is a builder for BACnetConstructedDataCredentialDataInputPresentValue
type BACnetConstructedDataCredentialDataInputPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetAuthenticationFactor) BACnetConstructedDataCredentialDataInputPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetAuthenticationFactor) BACnetConstructedDataCredentialDataInputPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorBuilder) BACnetConstructedDataCredentialDataInputPresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCredentialDataInputPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataCredentialDataInputPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCredentialDataInputPresentValue
}

// NewBACnetConstructedDataCredentialDataInputPresentValueBuilder() creates a BACnetConstructedDataCredentialDataInputPresentValueBuilder
func NewBACnetConstructedDataCredentialDataInputPresentValueBuilder() BACnetConstructedDataCredentialDataInputPresentValueBuilder {
	return &_BACnetConstructedDataCredentialDataInputPresentValueBuilder{_BACnetConstructedDataCredentialDataInputPresentValue: new(_BACnetConstructedDataCredentialDataInputPresentValue)}
}

type _BACnetConstructedDataCredentialDataInputPresentValueBuilder struct {
	*_BACnetConstructedDataCredentialDataInputPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCredentialDataInputPresentValueBuilder) = (*_BACnetConstructedDataCredentialDataInputPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCredentialDataInputPresentValue
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) WithMandatoryFields(presentValue BACnetAuthenticationFactor) BACnetConstructedDataCredentialDataInputPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) WithPresentValue(presentValue BACnetAuthenticationFactor) BACnetConstructedDataCredentialDataInputPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetAuthenticationFactorBuilder) BACnetAuthenticationFactorBuilder) BACnetConstructedDataCredentialDataInputPresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetAuthenticationFactorBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAuthenticationFactorBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) Build() (BACnetConstructedDataCredentialDataInputPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCredentialDataInputPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) MustBuild() BACnetConstructedDataCredentialDataInputPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCredentialDataInputPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCredentialDataInputPresentValueBuilder().(*_BACnetConstructedDataCredentialDataInputPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCredentialDataInputPresentValueBuilder creates a BACnetConstructedDataCredentialDataInputPresentValueBuilder
func (b *_BACnetConstructedDataCredentialDataInputPresentValue) CreateBACnetConstructedDataCredentialDataInputPresentValueBuilder() BACnetConstructedDataCredentialDataInputPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataCredentialDataInputPresentValueBuilder()
	}
	return &_BACnetConstructedDataCredentialDataInputPresentValueBuilder{_BACnetConstructedDataCredentialDataInputPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CREDENTIAL_DATA_INPUT
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetPresentValue() BACnetAuthenticationFactor {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetActualValue() BACnetAuthenticationFactor {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAuthenticationFactor(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialDataInputPresentValue(structType any) BACnetConstructedDataCredentialDataInputPresentValue {
	if casted, ok := structType.(BACnetConstructedDataCredentialDataInputPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialDataInputPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetTypeName() string {
	return "BACnetConstructedDataCredentialDataInputPresentValue"
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCredentialDataInputPresentValue BACnetConstructedDataCredentialDataInputPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialDataInputPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialDataInputPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetAuthenticationFactor](ctx, "presentValue", ReadComplex[BACnetAuthenticationFactor](BACnetAuthenticationFactorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetAuthenticationFactor](ctx, "actualValue", (*BACnetAuthenticationFactor)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialDataInputPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialDataInputPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialDataInputPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialDataInputPresentValue")
		}

		if err := WriteSimpleField[BACnetAuthenticationFactor](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetAuthenticationFactor](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialDataInputPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialDataInputPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) IsBACnetConstructedDataCredentialDataInputPresentValue() {
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) deepCopy() *_BACnetConstructedDataCredentialDataInputPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCredentialDataInputPresentValueCopy := &_BACnetConstructedDataCredentialDataInputPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetAuthenticationFactor](m.PresentValue),
	}
	_BACnetConstructedDataCredentialDataInputPresentValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCredentialDataInputPresentValueCopy
}

func (m *_BACnetConstructedDataCredentialDataInputPresentValue) String() string {
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
