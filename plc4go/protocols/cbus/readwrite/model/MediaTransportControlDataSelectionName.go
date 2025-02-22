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

// MediaTransportControlDataSelectionName is the corresponding interface of MediaTransportControlDataSelectionName
type MediaTransportControlDataSelectionName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// GetSelectionName returns SelectionName (property field)
	GetSelectionName() string
	// IsMediaTransportControlDataSelectionName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataSelectionName()
	// CreateBuilder creates a MediaTransportControlDataSelectionNameBuilder
	CreateMediaTransportControlDataSelectionNameBuilder() MediaTransportControlDataSelectionNameBuilder
}

// _MediaTransportControlDataSelectionName is the data-structure of this message
type _MediaTransportControlDataSelectionName struct {
	MediaTransportControlDataContract
	SelectionName string
}

var _ MediaTransportControlDataSelectionName = (*_MediaTransportControlDataSelectionName)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataSelectionName)(nil)

// NewMediaTransportControlDataSelectionName factory function for _MediaTransportControlDataSelectionName
func NewMediaTransportControlDataSelectionName(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, selectionName string) *_MediaTransportControlDataSelectionName {
	_result := &_MediaTransportControlDataSelectionName{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		SelectionName:                     selectionName,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataSelectionNameBuilder is a builder for MediaTransportControlDataSelectionName
type MediaTransportControlDataSelectionNameBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(selectionName string) MediaTransportControlDataSelectionNameBuilder
	// WithSelectionName adds SelectionName (property field)
	WithSelectionName(string) MediaTransportControlDataSelectionNameBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MediaTransportControlDataBuilder
	// Build builds the MediaTransportControlDataSelectionName or returns an error if something is wrong
	Build() (MediaTransportControlDataSelectionName, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataSelectionName
}

// NewMediaTransportControlDataSelectionNameBuilder() creates a MediaTransportControlDataSelectionNameBuilder
func NewMediaTransportControlDataSelectionNameBuilder() MediaTransportControlDataSelectionNameBuilder {
	return &_MediaTransportControlDataSelectionNameBuilder{_MediaTransportControlDataSelectionName: new(_MediaTransportControlDataSelectionName)}
}

type _MediaTransportControlDataSelectionNameBuilder struct {
	*_MediaTransportControlDataSelectionName

	parentBuilder *_MediaTransportControlDataBuilder

	err *utils.MultiError
}

var _ (MediaTransportControlDataSelectionNameBuilder) = (*_MediaTransportControlDataSelectionNameBuilder)(nil)

func (b *_MediaTransportControlDataSelectionNameBuilder) setParent(contract MediaTransportControlDataContract) {
	b.MediaTransportControlDataContract = contract
	contract.(*_MediaTransportControlData)._SubType = b._MediaTransportControlDataSelectionName
}

func (b *_MediaTransportControlDataSelectionNameBuilder) WithMandatoryFields(selectionName string) MediaTransportControlDataSelectionNameBuilder {
	return b.WithSelectionName(selectionName)
}

func (b *_MediaTransportControlDataSelectionNameBuilder) WithSelectionName(selectionName string) MediaTransportControlDataSelectionNameBuilder {
	b.SelectionName = selectionName
	return b
}

func (b *_MediaTransportControlDataSelectionNameBuilder) Build() (MediaTransportControlDataSelectionName, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MediaTransportControlDataSelectionName.deepCopy(), nil
}

func (b *_MediaTransportControlDataSelectionNameBuilder) MustBuild() MediaTransportControlDataSelectionName {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataSelectionNameBuilder) Done() MediaTransportControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMediaTransportControlDataBuilder().(*_MediaTransportControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MediaTransportControlDataSelectionNameBuilder) buildForMediaTransportControlData() (MediaTransportControlData, error) {
	return b.Build()
}

func (b *_MediaTransportControlDataSelectionNameBuilder) DeepCopy() any {
	_copy := b.CreateMediaTransportControlDataSelectionNameBuilder().(*_MediaTransportControlDataSelectionNameBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMediaTransportControlDataSelectionNameBuilder creates a MediaTransportControlDataSelectionNameBuilder
func (b *_MediaTransportControlDataSelectionName) CreateMediaTransportControlDataSelectionNameBuilder() MediaTransportControlDataSelectionNameBuilder {
	if b == nil {
		return NewMediaTransportControlDataSelectionNameBuilder()
	}
	return &_MediaTransportControlDataSelectionNameBuilder{_MediaTransportControlDataSelectionName: b.deepCopy()}
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

func (m *_MediaTransportControlDataSelectionName) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSelectionName) GetSelectionName() string {
	return m.SelectionName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataSelectionName(structType any) MediaTransportControlDataSelectionName {
	if casted, ok := structType.(MediaTransportControlDataSelectionName); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataSelectionName); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataSelectionName) GetTypeName() string {
	return "MediaTransportControlDataSelectionName"
}

func (m *_MediaTransportControlDataSelectionName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (selectionName)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_MediaTransportControlDataSelectionName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataSelectionName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer) (__mediaTransportControlDataSelectionName MediaTransportControlDataSelectionName, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSelectionName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSelectionName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	selectionName, err := ReadSimpleField(ctx, "selectionName", ReadString(readBuffer, uint32(int32((int32(commandTypeContainer.NumBytes())-int32(int32(1))))*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'selectionName' field"))
	}
	m.SelectionName = selectionName

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSelectionName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSelectionName")
	}

	return m, nil
}

func (m *_MediaTransportControlDataSelectionName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataSelectionName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataSelectionName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataSelectionName")
		}

		if err := WriteSimpleField[string](ctx, "selectionName", m.GetSelectionName(), WriteString(writeBuffer, int32(int32((int32(m.GetCommandTypeContainer().NumBytes())-int32(int32(1))))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'selectionName' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSelectionName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSelectionName")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataSelectionName) IsMediaTransportControlDataSelectionName() {}

func (m *_MediaTransportControlDataSelectionName) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataSelectionName) deepCopy() *_MediaTransportControlDataSelectionName {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataSelectionNameCopy := &_MediaTransportControlDataSelectionName{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
		m.SelectionName,
	}
	_MediaTransportControlDataSelectionNameCopy.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataSelectionNameCopy
}

func (m *_MediaTransportControlDataSelectionName) String() string {
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
