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

// EUInformation is the corresponding interface of EUInformation
type EUInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNamespaceUri returns NamespaceUri (property field)
	GetNamespaceUri() PascalString
	// GetUnitId returns UnitId (property field)
	GetUnitId() int32
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// IsEUInformation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEUInformation()
	// CreateBuilder creates a EUInformationBuilder
	CreateEUInformationBuilder() EUInformationBuilder
}

// _EUInformation is the data-structure of this message
type _EUInformation struct {
	ExtensionObjectDefinitionContract
	NamespaceUri PascalString
	UnitId       int32
	DisplayName  LocalizedText
	Description  LocalizedText
}

var _ EUInformation = (*_EUInformation)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EUInformation)(nil)

// NewEUInformation factory function for _EUInformation
func NewEUInformation(namespaceUri PascalString, unitId int32, displayName LocalizedText, description LocalizedText) *_EUInformation {
	if namespaceUri == nil {
		panic("namespaceUri of type PascalString for EUInformation must not be nil")
	}
	if displayName == nil {
		panic("displayName of type LocalizedText for EUInformation must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for EUInformation must not be nil")
	}
	_result := &_EUInformation{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NamespaceUri:                      namespaceUri,
		UnitId:                            unitId,
		DisplayName:                       displayName,
		Description:                       description,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EUInformationBuilder is a builder for EUInformation
type EUInformationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(namespaceUri PascalString, unitId int32, displayName LocalizedText, description LocalizedText) EUInformationBuilder
	// WithNamespaceUri adds NamespaceUri (property field)
	WithNamespaceUri(PascalString) EUInformationBuilder
	// WithNamespaceUriBuilder adds NamespaceUri (property field) which is build by the builder
	WithNamespaceUriBuilder(func(PascalStringBuilder) PascalStringBuilder) EUInformationBuilder
	// WithUnitId adds UnitId (property field)
	WithUnitId(int32) EUInformationBuilder
	// WithDisplayName adds DisplayName (property field)
	WithDisplayName(LocalizedText) EUInformationBuilder
	// WithDisplayNameBuilder adds DisplayName (property field) which is build by the builder
	WithDisplayNameBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) EUInformationBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) EUInformationBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) EUInformationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the EUInformation or returns an error if something is wrong
	Build() (EUInformation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EUInformation
}

// NewEUInformationBuilder() creates a EUInformationBuilder
func NewEUInformationBuilder() EUInformationBuilder {
	return &_EUInformationBuilder{_EUInformation: new(_EUInformation)}
}

type _EUInformationBuilder struct {
	*_EUInformation

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (EUInformationBuilder) = (*_EUInformationBuilder)(nil)

func (b *_EUInformationBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._EUInformation
}

func (b *_EUInformationBuilder) WithMandatoryFields(namespaceUri PascalString, unitId int32, displayName LocalizedText, description LocalizedText) EUInformationBuilder {
	return b.WithNamespaceUri(namespaceUri).WithUnitId(unitId).WithDisplayName(displayName).WithDescription(description)
}

func (b *_EUInformationBuilder) WithNamespaceUri(namespaceUri PascalString) EUInformationBuilder {
	b.NamespaceUri = namespaceUri
	return b
}

func (b *_EUInformationBuilder) WithNamespaceUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) EUInformationBuilder {
	builder := builderSupplier(b.NamespaceUri.CreatePascalStringBuilder())
	var err error
	b.NamespaceUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_EUInformationBuilder) WithUnitId(unitId int32) EUInformationBuilder {
	b.UnitId = unitId
	return b
}

func (b *_EUInformationBuilder) WithDisplayName(displayName LocalizedText) EUInformationBuilder {
	b.DisplayName = displayName
	return b
}

func (b *_EUInformationBuilder) WithDisplayNameBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) EUInformationBuilder {
	builder := builderSupplier(b.DisplayName.CreateLocalizedTextBuilder())
	var err error
	b.DisplayName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_EUInformationBuilder) WithDescription(description LocalizedText) EUInformationBuilder {
	b.Description = description
	return b
}

func (b *_EUInformationBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) EUInformationBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_EUInformationBuilder) Build() (EUInformation, error) {
	if b.NamespaceUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'namespaceUri' not set"))
	}
	if b.DisplayName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'displayName' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EUInformation.deepCopy(), nil
}

func (b *_EUInformationBuilder) MustBuild() EUInformation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EUInformationBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_EUInformationBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_EUInformationBuilder) DeepCopy() any {
	_copy := b.CreateEUInformationBuilder().(*_EUInformationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEUInformationBuilder creates a EUInformationBuilder
func (b *_EUInformation) CreateEUInformationBuilder() EUInformationBuilder {
	if b == nil {
		return NewEUInformationBuilder()
	}
	return &_EUInformationBuilder{_EUInformation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EUInformation) GetExtensionId() int32 {
	return int32(889)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EUInformation) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EUInformation) GetNamespaceUri() PascalString {
	return m.NamespaceUri
}

func (m *_EUInformation) GetUnitId() int32 {
	return m.UnitId
}

func (m *_EUInformation) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_EUInformation) GetDescription() LocalizedText {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEUInformation(structType any) EUInformation {
	if casted, ok := structType.(EUInformation); ok {
		return casted
	}
	if casted, ok := structType.(*EUInformation); ok {
		return *casted
	}
	return nil
}

func (m *_EUInformation) GetTypeName() string {
	return "EUInformation"
}

func (m *_EUInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (namespaceUri)
	lengthInBits += m.NamespaceUri.GetLengthInBits(ctx)

	// Simple field (unitId)
	lengthInBits += 32

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_EUInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EUInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__eUInformation EUInformation, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EUInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EUInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceUri, err := ReadSimpleField[PascalString](ctx, "namespaceUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceUri' field"))
	}
	m.NamespaceUri = namespaceUri

	unitId, err := ReadSimpleField(ctx, "unitId", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitId' field"))
	}
	m.UnitId = unitId

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	if closeErr := readBuffer.CloseContext("EUInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EUInformation")
	}

	return m, nil
}

func (m *_EUInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EUInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EUInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EUInformation")
		}

		if err := WriteSimpleField[PascalString](ctx, "namespaceUri", m.GetNamespaceUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceUri' field")
		}

		if err := WriteSimpleField[int32](ctx, "unitId", m.GetUnitId(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitId' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if popErr := writeBuffer.PopContext("EUInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EUInformation")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EUInformation) IsEUInformation() {}

func (m *_EUInformation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EUInformation) deepCopy() *_EUInformation {
	if m == nil {
		return nil
	}
	_EUInformationCopy := &_EUInformation{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.NamespaceUri),
		m.UnitId,
		utils.DeepCopy[LocalizedText](m.DisplayName),
		utils.DeepCopy[LocalizedText](m.Description),
	}
	_EUInformationCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _EUInformationCopy
}

func (m *_EUInformation) String() string {
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
