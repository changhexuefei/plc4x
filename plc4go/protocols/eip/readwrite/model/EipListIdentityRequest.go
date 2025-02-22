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

// EipListIdentityRequest is the corresponding interface of EipListIdentityRequest
type EipListIdentityRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EipPacket
	// IsEipListIdentityRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEipListIdentityRequest()
	// CreateBuilder creates a EipListIdentityRequestBuilder
	CreateEipListIdentityRequestBuilder() EipListIdentityRequestBuilder
}

// _EipListIdentityRequest is the data-structure of this message
type _EipListIdentityRequest struct {
	EipPacketContract
}

var _ EipListIdentityRequest = (*_EipListIdentityRequest)(nil)
var _ EipPacketRequirements = (*_EipListIdentityRequest)(nil)

// NewEipListIdentityRequest factory function for _EipListIdentityRequest
func NewEipListIdentityRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipListIdentityRequest {
	_result := &_EipListIdentityRequest{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EipListIdentityRequestBuilder is a builder for EipListIdentityRequest
type EipListIdentityRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() EipListIdentityRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() EipPacketBuilder
	// Build builds the EipListIdentityRequest or returns an error if something is wrong
	Build() (EipListIdentityRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EipListIdentityRequest
}

// NewEipListIdentityRequestBuilder() creates a EipListIdentityRequestBuilder
func NewEipListIdentityRequestBuilder() EipListIdentityRequestBuilder {
	return &_EipListIdentityRequestBuilder{_EipListIdentityRequest: new(_EipListIdentityRequest)}
}

type _EipListIdentityRequestBuilder struct {
	*_EipListIdentityRequest

	parentBuilder *_EipPacketBuilder

	err *utils.MultiError
}

var _ (EipListIdentityRequestBuilder) = (*_EipListIdentityRequestBuilder)(nil)

func (b *_EipListIdentityRequestBuilder) setParent(contract EipPacketContract) {
	b.EipPacketContract = contract
	contract.(*_EipPacket)._SubType = b._EipListIdentityRequest
}

func (b *_EipListIdentityRequestBuilder) WithMandatoryFields() EipListIdentityRequestBuilder {
	return b
}

func (b *_EipListIdentityRequestBuilder) Build() (EipListIdentityRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EipListIdentityRequest.deepCopy(), nil
}

func (b *_EipListIdentityRequestBuilder) MustBuild() EipListIdentityRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EipListIdentityRequestBuilder) Done() EipPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewEipPacketBuilder().(*_EipPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_EipListIdentityRequestBuilder) buildForEipPacket() (EipPacket, error) {
	return b.Build()
}

func (b *_EipListIdentityRequestBuilder) DeepCopy() any {
	_copy := b.CreateEipListIdentityRequestBuilder().(*_EipListIdentityRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEipListIdentityRequestBuilder creates a EipListIdentityRequestBuilder
func (b *_EipListIdentityRequest) CreateEipListIdentityRequestBuilder() EipListIdentityRequestBuilder {
	if b == nil {
		return NewEipListIdentityRequestBuilder()
	}
	return &_EipListIdentityRequestBuilder{_EipListIdentityRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipListIdentityRequest) GetCommand() uint16 {
	return 0x0063
}

func (m *_EipListIdentityRequest) GetResponse() bool {
	return bool(false)
}

func (m *_EipListIdentityRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipListIdentityRequest) GetParent() EipPacketContract {
	return m.EipPacketContract
}

// Deprecated: use the interface for direct cast
func CastEipListIdentityRequest(structType any) EipListIdentityRequest {
	if casted, ok := structType.(EipListIdentityRequest); ok {
		return casted
	}
	if casted, ok := structType.(*EipListIdentityRequest); ok {
		return *casted
	}
	return nil
}

func (m *_EipListIdentityRequest) GetTypeName() string {
	return "EipListIdentityRequest"
}

func (m *_EipListIdentityRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_EipListIdentityRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EipListIdentityRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__eipListIdentityRequest EipListIdentityRequest, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipListIdentityRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipListIdentityRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("EipListIdentityRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipListIdentityRequest")
	}

	return m, nil
}

func (m *_EipListIdentityRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipListIdentityRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipListIdentityRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipListIdentityRequest")
		}

		if popErr := writeBuffer.PopContext("EipListIdentityRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipListIdentityRequest")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipListIdentityRequest) IsEipListIdentityRequest() {}

func (m *_EipListIdentityRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EipListIdentityRequest) deepCopy() *_EipListIdentityRequest {
	if m == nil {
		return nil
	}
	_EipListIdentityRequestCopy := &_EipListIdentityRequest{
		m.EipPacketContract.(*_EipPacket).deepCopy(),
	}
	_EipListIdentityRequestCopy.EipPacketContract.(*_EipPacket)._SubType = m
	return _EipListIdentityRequestCopy
}

func (m *_EipListIdentityRequest) String() string {
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
