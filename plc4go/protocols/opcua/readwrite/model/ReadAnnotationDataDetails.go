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

// ReadAnnotationDataDetails is the corresponding interface of ReadAnnotationDataDetails
type ReadAnnotationDataDetails interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetReqTimes returns ReqTimes (property field)
	GetReqTimes() []int64
	// IsReadAnnotationDataDetails is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReadAnnotationDataDetails()
	// CreateBuilder creates a ReadAnnotationDataDetailsBuilder
	CreateReadAnnotationDataDetailsBuilder() ReadAnnotationDataDetailsBuilder
}

// _ReadAnnotationDataDetails is the data-structure of this message
type _ReadAnnotationDataDetails struct {
	ExtensionObjectDefinitionContract
	ReqTimes []int64
}

var _ ReadAnnotationDataDetails = (*_ReadAnnotationDataDetails)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReadAnnotationDataDetails)(nil)

// NewReadAnnotationDataDetails factory function for _ReadAnnotationDataDetails
func NewReadAnnotationDataDetails(reqTimes []int64) *_ReadAnnotationDataDetails {
	_result := &_ReadAnnotationDataDetails{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ReqTimes:                          reqTimes,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReadAnnotationDataDetailsBuilder is a builder for ReadAnnotationDataDetails
type ReadAnnotationDataDetailsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reqTimes []int64) ReadAnnotationDataDetailsBuilder
	// WithReqTimes adds ReqTimes (property field)
	WithReqTimes(...int64) ReadAnnotationDataDetailsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ReadAnnotationDataDetails or returns an error if something is wrong
	Build() (ReadAnnotationDataDetails, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReadAnnotationDataDetails
}

// NewReadAnnotationDataDetailsBuilder() creates a ReadAnnotationDataDetailsBuilder
func NewReadAnnotationDataDetailsBuilder() ReadAnnotationDataDetailsBuilder {
	return &_ReadAnnotationDataDetailsBuilder{_ReadAnnotationDataDetails: new(_ReadAnnotationDataDetails)}
}

type _ReadAnnotationDataDetailsBuilder struct {
	*_ReadAnnotationDataDetails

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ReadAnnotationDataDetailsBuilder) = (*_ReadAnnotationDataDetailsBuilder)(nil)

func (b *_ReadAnnotationDataDetailsBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ReadAnnotationDataDetails
}

func (b *_ReadAnnotationDataDetailsBuilder) WithMandatoryFields(reqTimes []int64) ReadAnnotationDataDetailsBuilder {
	return b.WithReqTimes(reqTimes...)
}

func (b *_ReadAnnotationDataDetailsBuilder) WithReqTimes(reqTimes ...int64) ReadAnnotationDataDetailsBuilder {
	b.ReqTimes = reqTimes
	return b
}

func (b *_ReadAnnotationDataDetailsBuilder) Build() (ReadAnnotationDataDetails, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReadAnnotationDataDetails.deepCopy(), nil
}

func (b *_ReadAnnotationDataDetailsBuilder) MustBuild() ReadAnnotationDataDetails {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ReadAnnotationDataDetailsBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ReadAnnotationDataDetailsBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ReadAnnotationDataDetailsBuilder) DeepCopy() any {
	_copy := b.CreateReadAnnotationDataDetailsBuilder().(*_ReadAnnotationDataDetailsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReadAnnotationDataDetailsBuilder creates a ReadAnnotationDataDetailsBuilder
func (b *_ReadAnnotationDataDetails) CreateReadAnnotationDataDetailsBuilder() ReadAnnotationDataDetailsBuilder {
	if b == nil {
		return NewReadAnnotationDataDetailsBuilder()
	}
	return &_ReadAnnotationDataDetailsBuilder{_ReadAnnotationDataDetails: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReadAnnotationDataDetails) GetExtensionId() int32 {
	return int32(23499)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReadAnnotationDataDetails) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReadAnnotationDataDetails) GetReqTimes() []int64 {
	return m.ReqTimes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReadAnnotationDataDetails(structType any) ReadAnnotationDataDetails {
	if casted, ok := structType.(ReadAnnotationDataDetails); ok {
		return casted
	}
	if casted, ok := structType.(*ReadAnnotationDataDetails); ok {
		return *casted
	}
	return nil
}

func (m *_ReadAnnotationDataDetails) GetTypeName() string {
	return "ReadAnnotationDataDetails"
}

func (m *_ReadAnnotationDataDetails) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfReqTimes)
	lengthInBits += 32

	// Array field
	if len(m.ReqTimes) > 0 {
		lengthInBits += 64 * uint16(len(m.ReqTimes))
	}

	return lengthInBits
}

func (m *_ReadAnnotationDataDetails) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReadAnnotationDataDetails) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__readAnnotationDataDetails ReadAnnotationDataDetails, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReadAnnotationDataDetails"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReadAnnotationDataDetails")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfReqTimes, err := ReadImplicitField[int32](ctx, "noOfReqTimes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReqTimes' field"))
	}
	_ = noOfReqTimes

	reqTimes, err := ReadCountArrayField[int64](ctx, "reqTimes", ReadSignedLong(readBuffer, uint8(64)), uint64(noOfReqTimes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reqTimes' field"))
	}
	m.ReqTimes = reqTimes

	if closeErr := readBuffer.CloseContext("ReadAnnotationDataDetails"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReadAnnotationDataDetails")
	}

	return m, nil
}

func (m *_ReadAnnotationDataDetails) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReadAnnotationDataDetails) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReadAnnotationDataDetails"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReadAnnotationDataDetails")
		}
		noOfReqTimes := int32(utils.InlineIf(bool((m.GetReqTimes()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetReqTimes()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfReqTimes", noOfReqTimes, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReqTimes' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "reqTimes", m.GetReqTimes(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'reqTimes' field")
		}

		if popErr := writeBuffer.PopContext("ReadAnnotationDataDetails"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReadAnnotationDataDetails")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReadAnnotationDataDetails) IsReadAnnotationDataDetails() {}

func (m *_ReadAnnotationDataDetails) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReadAnnotationDataDetails) deepCopy() *_ReadAnnotationDataDetails {
	if m == nil {
		return nil
	}
	_ReadAnnotationDataDetailsCopy := &_ReadAnnotationDataDetails{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[int64, int64](m.ReqTimes),
	}
	_ReadAnnotationDataDetailsCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReadAnnotationDataDetailsCopy
}

func (m *_ReadAnnotationDataDetails) String() string {
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
