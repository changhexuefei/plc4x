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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Orientation is the corresponding interface of Orientation
type Orientation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsOrientation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOrientation()
	// CreateBuilder creates a OrientationBuilder
	CreateOrientationBuilder() OrientationBuilder
}

// _Orientation is the data-structure of this message
type _Orientation struct {
	ExtensionObjectDefinitionContract
}

var _ Orientation = (*_Orientation)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_Orientation)(nil)

// NewOrientation factory function for _Orientation
func NewOrientation() *_Orientation {
	_result := &_Orientation{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OrientationBuilder is a builder for Orientation
type OrientationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() OrientationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the Orientation or returns an error if something is wrong
	Build() (Orientation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Orientation
}

// NewOrientationBuilder() creates a OrientationBuilder
func NewOrientationBuilder() OrientationBuilder {
	return &_OrientationBuilder{_Orientation: new(_Orientation)}
}

type _OrientationBuilder struct {
	*_Orientation

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (OrientationBuilder) = (*_OrientationBuilder)(nil)

func (b *_OrientationBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._Orientation
}

func (b *_OrientationBuilder) WithMandatoryFields() OrientationBuilder {
	return b
}

func (b *_OrientationBuilder) Build() (Orientation, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Orientation.deepCopy(), nil
}

func (b *_OrientationBuilder) MustBuild() Orientation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_OrientationBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_OrientationBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_OrientationBuilder) DeepCopy() any {
	_copy := b.CreateOrientationBuilder().(*_OrientationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOrientationBuilder creates a OrientationBuilder
func (b *_Orientation) CreateOrientationBuilder() OrientationBuilder {
	if b == nil {
		return NewOrientationBuilder()
	}
	return &_OrientationBuilder{_Orientation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Orientation) GetExtensionId() int32 {
	return int32(18813)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Orientation) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastOrientation(structType any) Orientation {
	if casted, ok := structType.(Orientation); ok {
		return casted
	}
	if casted, ok := structType.(*Orientation); ok {
		return *casted
	}
	return nil
}

func (m *_Orientation) GetTypeName() string {
	return "Orientation"
}

func (m *_Orientation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_Orientation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_Orientation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__orientation Orientation, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Orientation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Orientation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Orientation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Orientation")
	}

	return m, nil
}

func (m *_Orientation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Orientation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Orientation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Orientation")
		}

		if popErr := writeBuffer.PopContext("Orientation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Orientation")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Orientation) IsOrientation() {}

func (m *_Orientation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Orientation) deepCopy() *_Orientation {
	if m == nil {
		return nil
	}
	_OrientationCopy := &_Orientation{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	_OrientationCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _OrientationCopy
}

func (m *_Orientation) String() string {
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
