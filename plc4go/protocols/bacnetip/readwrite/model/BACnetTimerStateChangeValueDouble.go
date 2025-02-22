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

// BACnetTimerStateChangeValueDouble is the corresponding interface of BACnetTimerStateChangeValueDouble
type BACnetTimerStateChangeValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// IsBACnetTimerStateChangeValueDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueDouble()
	// CreateBuilder creates a BACnetTimerStateChangeValueDoubleBuilder
	CreateBACnetTimerStateChangeValueDoubleBuilder() BACnetTimerStateChangeValueDoubleBuilder
}

// _BACnetTimerStateChangeValueDouble is the data-structure of this message
type _BACnetTimerStateChangeValueDouble struct {
	BACnetTimerStateChangeValueContract
	DoubleValue BACnetApplicationTagDouble
}

var _ BACnetTimerStateChangeValueDouble = (*_BACnetTimerStateChangeValueDouble)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueDouble)(nil)

// NewBACnetTimerStateChangeValueDouble factory function for _BACnetTimerStateChangeValueDouble
func NewBACnetTimerStateChangeValueDouble(peekedTagHeader BACnetTagHeader, doubleValue BACnetApplicationTagDouble, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueDouble {
	if doubleValue == nil {
		panic("doubleValue of type BACnetApplicationTagDouble for BACnetTimerStateChangeValueDouble must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueDouble{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		DoubleValue:                         doubleValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueDoubleBuilder is a builder for BACnetTimerStateChangeValueDouble
type BACnetTimerStateChangeValueDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetTimerStateChangeValueDoubleBuilder
	// WithDoubleValue adds DoubleValue (property field)
	WithDoubleValue(BACnetApplicationTagDouble) BACnetTimerStateChangeValueDoubleBuilder
	// WithDoubleValueBuilder adds DoubleValue (property field) which is build by the builder
	WithDoubleValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetTimerStateChangeValueDoubleBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetTimerStateChangeValueBuilder
	// Build builds the BACnetTimerStateChangeValueDouble or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueDouble
}

// NewBACnetTimerStateChangeValueDoubleBuilder() creates a BACnetTimerStateChangeValueDoubleBuilder
func NewBACnetTimerStateChangeValueDoubleBuilder() BACnetTimerStateChangeValueDoubleBuilder {
	return &_BACnetTimerStateChangeValueDoubleBuilder{_BACnetTimerStateChangeValueDouble: new(_BACnetTimerStateChangeValueDouble)}
}

type _BACnetTimerStateChangeValueDoubleBuilder struct {
	*_BACnetTimerStateChangeValueDouble

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueDoubleBuilder) = (*_BACnetTimerStateChangeValueDoubleBuilder)(nil)

func (b *_BACnetTimerStateChangeValueDoubleBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
	contract.(*_BACnetTimerStateChangeValue)._SubType = b._BACnetTimerStateChangeValueDouble
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetTimerStateChangeValueDoubleBuilder {
	return b.WithDoubleValue(doubleValue)
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) WithDoubleValue(doubleValue BACnetApplicationTagDouble) BACnetTimerStateChangeValueDoubleBuilder {
	b.DoubleValue = doubleValue
	return b
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) WithDoubleValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetTimerStateChangeValueDoubleBuilder {
	builder := builderSupplier(b.DoubleValue.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.DoubleValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) Build() (BACnetTimerStateChangeValueDouble, error) {
	if b.DoubleValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doubleValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueDouble.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) MustBuild() BACnetTimerStateChangeValueDouble {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) Done() BACnetTimerStateChangeValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetTimerStateChangeValueBuilder().(*_BACnetTimerStateChangeValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueDoubleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueDoubleBuilder().(*_BACnetTimerStateChangeValueDoubleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueDoubleBuilder creates a BACnetTimerStateChangeValueDoubleBuilder
func (b *_BACnetTimerStateChangeValueDouble) CreateBACnetTimerStateChangeValueDoubleBuilder() BACnetTimerStateChangeValueDoubleBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueDoubleBuilder()
	}
	return &_BACnetTimerStateChangeValueDoubleBuilder{_BACnetTimerStateChangeValueDouble: b.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueDouble) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueDouble(structType any) BACnetTimerStateChangeValueDouble {
	if casted, ok := structType.(BACnetTimerStateChangeValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueDouble) GetTypeName() string {
	return "BACnetTimerStateChangeValueDouble"
}

func (m *_BACnetTimerStateChangeValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueDouble BACnetTimerStateChangeValueDouble, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doubleValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	m.DoubleValue = doubleValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueDouble")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueDouble")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", m.GetDoubleValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueDouble")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueDouble) IsBACnetTimerStateChangeValueDouble() {}

func (m *_BACnetTimerStateChangeValueDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueDouble) deepCopy() *_BACnetTimerStateChangeValueDouble {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueDoubleCopy := &_BACnetTimerStateChangeValueDouble{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.DoubleValue),
	}
	_BACnetTimerStateChangeValueDoubleCopy.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueDoubleCopy
}

func (m *_BACnetTimerStateChangeValueDouble) String() string {
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
