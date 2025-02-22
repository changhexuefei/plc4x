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

// BACnetConfirmedServiceRequestRequestKey is the corresponding interface of BACnetConfirmedServiceRequestRequestKey
type BACnetConfirmedServiceRequestRequestKey interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
	// IsBACnetConfirmedServiceRequestRequestKey is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestRequestKey()
	// CreateBuilder creates a BACnetConfirmedServiceRequestRequestKeyBuilder
	CreateBACnetConfirmedServiceRequestRequestKeyBuilder() BACnetConfirmedServiceRequestRequestKeyBuilder
}

// _BACnetConfirmedServiceRequestRequestKey is the data-structure of this message
type _BACnetConfirmedServiceRequestRequestKey struct {
	BACnetConfirmedServiceRequestContract
	BytesOfRemovedService []byte

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestRequestKey = (*_BACnetConfirmedServiceRequestRequestKey)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestRequestKey)(nil)

// NewBACnetConfirmedServiceRequestRequestKey factory function for _BACnetConfirmedServiceRequestRequestKey
func NewBACnetConfirmedServiceRequestRequestKey(bytesOfRemovedService []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestRequestKey {
	_result := &_BACnetConfirmedServiceRequestRequestKey{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		BytesOfRemovedService:                 bytesOfRemovedService,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestRequestKeyBuilder is a builder for BACnetConfirmedServiceRequestRequestKey
type BACnetConfirmedServiceRequestRequestKeyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bytesOfRemovedService []byte) BACnetConfirmedServiceRequestRequestKeyBuilder
	// WithBytesOfRemovedService adds BytesOfRemovedService (property field)
	WithBytesOfRemovedService(...byte) BACnetConfirmedServiceRequestRequestKeyBuilder
	// WithArgServiceRequestPayloadLength sets a parser argument
	WithArgServiceRequestPayloadLength(uint32) BACnetConfirmedServiceRequestRequestKeyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestBuilder
	// Build builds the BACnetConfirmedServiceRequestRequestKey or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestRequestKey, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestRequestKey
}

// NewBACnetConfirmedServiceRequestRequestKeyBuilder() creates a BACnetConfirmedServiceRequestRequestKeyBuilder
func NewBACnetConfirmedServiceRequestRequestKeyBuilder() BACnetConfirmedServiceRequestRequestKeyBuilder {
	return &_BACnetConfirmedServiceRequestRequestKeyBuilder{_BACnetConfirmedServiceRequestRequestKey: new(_BACnetConfirmedServiceRequestRequestKey)}
}

type _BACnetConfirmedServiceRequestRequestKeyBuilder struct {
	*_BACnetConfirmedServiceRequestRequestKey

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestRequestKeyBuilder) = (*_BACnetConfirmedServiceRequestRequestKeyBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
	contract.(*_BACnetConfirmedServiceRequest)._SubType = b._BACnetConfirmedServiceRequestRequestKey
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) WithMandatoryFields(bytesOfRemovedService []byte) BACnetConfirmedServiceRequestRequestKeyBuilder {
	return b.WithBytesOfRemovedService(bytesOfRemovedService...)
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) WithBytesOfRemovedService(bytesOfRemovedService ...byte) BACnetConfirmedServiceRequestRequestKeyBuilder {
	b.BytesOfRemovedService = bytesOfRemovedService
	return b
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) WithArgServiceRequestPayloadLength(serviceRequestPayloadLength uint32) BACnetConfirmedServiceRequestRequestKeyBuilder {
	b.ServiceRequestPayloadLength = serviceRequestPayloadLength
	return b
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) Build() (BACnetConfirmedServiceRequestRequestKey, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestRequestKey.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) MustBuild() BACnetConfirmedServiceRequestRequestKey {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestBuilder().(*_BACnetConfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestRequestKeyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestRequestKeyBuilder().(*_BACnetConfirmedServiceRequestRequestKeyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestRequestKeyBuilder creates a BACnetConfirmedServiceRequestRequestKeyBuilder
func (b *_BACnetConfirmedServiceRequestRequestKey) CreateBACnetConfirmedServiceRequestRequestKeyBuilder() BACnetConfirmedServiceRequestRequestKeyBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestRequestKeyBuilder()
	}
	return &_BACnetConfirmedServiceRequestRequestKeyBuilder{_BACnetConfirmedServiceRequestRequestKey: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REQUEST_KEY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestRequestKey(structType any) BACnetConfirmedServiceRequestRequestKey {
	if casted, ok := structType.(BACnetConfirmedServiceRequestRequestKey); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestRequestKey); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetTypeName() string {
	return "BACnetConfirmedServiceRequestRequestKey"
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestRequestKey) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestRequestKey BACnetConfirmedServiceRequestRequestKey, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestRequestKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestRequestKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bytesOfRemovedService, err := readBuffer.ReadByteArray("bytesOfRemovedService", int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bytesOfRemovedService' field"))
	}
	m.BytesOfRemovedService = bytesOfRemovedService

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestRequestKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestRequestKey")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestRequestKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestRequestKey")
		}

		if err := WriteByteArrayField(ctx, "bytesOfRemovedService", m.GetBytesOfRemovedService(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestRequestKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestRequestKey")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestRequestKey) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestRequestKey) IsBACnetConfirmedServiceRequestRequestKey() {}

func (m *_BACnetConfirmedServiceRequestRequestKey) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestRequestKey) deepCopy() *_BACnetConfirmedServiceRequestRequestKey {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestRequestKeyCopy := &_BACnetConfirmedServiceRequestRequestKey{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.BytesOfRemovedService),
		m.ServiceRequestPayloadLength,
	}
	_BACnetConfirmedServiceRequestRequestKeyCopy.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestRequestKeyCopy
}

func (m *_BACnetConfirmedServiceRequestRequestKey) String() string {
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
