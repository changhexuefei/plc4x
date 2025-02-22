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

// BACnetEventParameterBufferReady is the corresponding interface of BACnetEventParameterBufferReady
type BACnetEventParameterBufferReady interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetNotificationThreshold returns NotificationThreshold (property field)
	GetNotificationThreshold() BACnetContextTagUnsignedInteger
	// GetPreviousNotificationCount returns PreviousNotificationCount (property field)
	GetPreviousNotificationCount() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterBufferReady is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterBufferReady()
	// CreateBuilder creates a BACnetEventParameterBufferReadyBuilder
	CreateBACnetEventParameterBufferReadyBuilder() BACnetEventParameterBufferReadyBuilder
}

// _BACnetEventParameterBufferReady is the data-structure of this message
type _BACnetEventParameterBufferReady struct {
	BACnetEventParameterContract
	OpeningTag                BACnetOpeningTag
	NotificationThreshold     BACnetContextTagUnsignedInteger
	PreviousNotificationCount BACnetContextTagUnsignedInteger
	ClosingTag                BACnetClosingTag
}

var _ BACnetEventParameterBufferReady = (*_BACnetEventParameterBufferReady)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterBufferReady)(nil)

// NewBACnetEventParameterBufferReady factory function for _BACnetEventParameterBufferReady
func NewBACnetEventParameterBufferReady(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, notificationThreshold BACnetContextTagUnsignedInteger, previousNotificationCount BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) *_BACnetEventParameterBufferReady {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterBufferReady must not be nil")
	}
	if notificationThreshold == nil {
		panic("notificationThreshold of type BACnetContextTagUnsignedInteger for BACnetEventParameterBufferReady must not be nil")
	}
	if previousNotificationCount == nil {
		panic("previousNotificationCount of type BACnetContextTagUnsignedInteger for BACnetEventParameterBufferReady must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterBufferReady must not be nil")
	}
	_result := &_BACnetEventParameterBufferReady{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		NotificationThreshold:        notificationThreshold,
		PreviousNotificationCount:    previousNotificationCount,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterBufferReadyBuilder is a builder for BACnetEventParameterBufferReady
type BACnetEventParameterBufferReadyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, notificationThreshold BACnetContextTagUnsignedInteger, previousNotificationCount BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterBufferReadyBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterBufferReadyBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterBufferReadyBuilder
	// WithNotificationThreshold adds NotificationThreshold (property field)
	WithNotificationThreshold(BACnetContextTagUnsignedInteger) BACnetEventParameterBufferReadyBuilder
	// WithNotificationThresholdBuilder adds NotificationThreshold (property field) which is build by the builder
	WithNotificationThresholdBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterBufferReadyBuilder
	// WithPreviousNotificationCount adds PreviousNotificationCount (property field)
	WithPreviousNotificationCount(BACnetContextTagUnsignedInteger) BACnetEventParameterBufferReadyBuilder
	// WithPreviousNotificationCountBuilder adds PreviousNotificationCount (property field) which is build by the builder
	WithPreviousNotificationCountBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterBufferReadyBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterBufferReadyBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterBufferReadyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetEventParameterBuilder
	// Build builds the BACnetEventParameterBufferReady or returns an error if something is wrong
	Build() (BACnetEventParameterBufferReady, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterBufferReady
}

// NewBACnetEventParameterBufferReadyBuilder() creates a BACnetEventParameterBufferReadyBuilder
func NewBACnetEventParameterBufferReadyBuilder() BACnetEventParameterBufferReadyBuilder {
	return &_BACnetEventParameterBufferReadyBuilder{_BACnetEventParameterBufferReady: new(_BACnetEventParameterBufferReady)}
}

type _BACnetEventParameterBufferReadyBuilder struct {
	*_BACnetEventParameterBufferReady

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterBufferReadyBuilder) = (*_BACnetEventParameterBufferReadyBuilder)(nil)

func (b *_BACnetEventParameterBufferReadyBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
	contract.(*_BACnetEventParameter)._SubType = b._BACnetEventParameterBufferReady
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, notificationThreshold BACnetContextTagUnsignedInteger, previousNotificationCount BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterBufferReadyBuilder {
	return b.WithOpeningTag(openingTag).WithNotificationThreshold(notificationThreshold).WithPreviousNotificationCount(previousNotificationCount).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterBufferReadyBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterBufferReadyBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithNotificationThreshold(notificationThreshold BACnetContextTagUnsignedInteger) BACnetEventParameterBufferReadyBuilder {
	b.NotificationThreshold = notificationThreshold
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithNotificationThresholdBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterBufferReadyBuilder {
	builder := builderSupplier(b.NotificationThreshold.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.NotificationThreshold, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithPreviousNotificationCount(previousNotificationCount BACnetContextTagUnsignedInteger) BACnetEventParameterBufferReadyBuilder {
	b.PreviousNotificationCount = previousNotificationCount
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithPreviousNotificationCountBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterBufferReadyBuilder {
	builder := builderSupplier(b.PreviousNotificationCount.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.PreviousNotificationCount, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterBufferReadyBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterBufferReadyBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterBufferReadyBuilder) Build() (BACnetEventParameterBufferReady, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.NotificationThreshold == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'notificationThreshold' not set"))
	}
	if b.PreviousNotificationCount == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'previousNotificationCount' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventParameterBufferReady.deepCopy(), nil
}

func (b *_BACnetEventParameterBufferReadyBuilder) MustBuild() BACnetEventParameterBufferReady {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventParameterBufferReadyBuilder) Done() BACnetEventParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetEventParameterBuilder().(*_BACnetEventParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetEventParameterBufferReadyBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterBufferReadyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterBufferReadyBuilder().(*_BACnetEventParameterBufferReadyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterBufferReadyBuilder creates a BACnetEventParameterBufferReadyBuilder
func (b *_BACnetEventParameterBufferReady) CreateBACnetEventParameterBufferReadyBuilder() BACnetEventParameterBufferReadyBuilder {
	if b == nil {
		return NewBACnetEventParameterBufferReadyBuilder()
	}
	return &_BACnetEventParameterBufferReadyBuilder{_BACnetEventParameterBufferReady: b.deepCopy()}
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

func (m *_BACnetEventParameterBufferReady) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterBufferReady) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterBufferReady) GetNotificationThreshold() BACnetContextTagUnsignedInteger {
	return m.NotificationThreshold
}

func (m *_BACnetEventParameterBufferReady) GetPreviousNotificationCount() BACnetContextTagUnsignedInteger {
	return m.PreviousNotificationCount
}

func (m *_BACnetEventParameterBufferReady) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterBufferReady(structType any) BACnetEventParameterBufferReady {
	if casted, ok := structType.(BACnetEventParameterBufferReady); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterBufferReady); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterBufferReady) GetTypeName() string {
	return "BACnetEventParameterBufferReady"
}

func (m *_BACnetEventParameterBufferReady) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (notificationThreshold)
	lengthInBits += m.NotificationThreshold.GetLengthInBits(ctx)

	// Simple field (previousNotificationCount)
	lengthInBits += m.PreviousNotificationCount.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterBufferReady) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterBufferReady) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterBufferReady BACnetEventParameterBufferReady, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterBufferReady"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterBufferReady")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(10))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	notificationThreshold, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "notificationThreshold", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationThreshold' field"))
	}
	m.NotificationThreshold = notificationThreshold

	previousNotificationCount, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "previousNotificationCount", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'previousNotificationCount' field"))
	}
	m.PreviousNotificationCount = previousNotificationCount

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(10))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterBufferReady"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterBufferReady")
	}

	return m, nil
}

func (m *_BACnetEventParameterBufferReady) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterBufferReady) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterBufferReady"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterBufferReady")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "notificationThreshold", m.GetNotificationThreshold(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationThreshold' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "previousNotificationCount", m.GetPreviousNotificationCount(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'previousNotificationCount' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterBufferReady"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterBufferReady")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterBufferReady) IsBACnetEventParameterBufferReady() {}

func (m *_BACnetEventParameterBufferReady) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterBufferReady) deepCopy() *_BACnetEventParameterBufferReady {
	if m == nil {
		return nil
	}
	_BACnetEventParameterBufferReadyCopy := &_BACnetEventParameterBufferReady{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.NotificationThreshold),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.PreviousNotificationCount),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	_BACnetEventParameterBufferReadyCopy.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterBufferReadyCopy
}

func (m *_BACnetEventParameterBufferReady) String() string {
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
