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

// NLMDisconnectConnectionToNetwork is the corresponding interface of NLMDisconnectConnectionToNetwork
type NLMDisconnectConnectionToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() uint16
	// IsNLMDisconnectConnectionToNetwork is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMDisconnectConnectionToNetwork()
	// CreateBuilder creates a NLMDisconnectConnectionToNetworkBuilder
	CreateNLMDisconnectConnectionToNetworkBuilder() NLMDisconnectConnectionToNetworkBuilder
}

// _NLMDisconnectConnectionToNetwork is the data-structure of this message
type _NLMDisconnectConnectionToNetwork struct {
	NLMContract
	DestinationNetworkAddress uint16
}

var _ NLMDisconnectConnectionToNetwork = (*_NLMDisconnectConnectionToNetwork)(nil)
var _ NLMRequirements = (*_NLMDisconnectConnectionToNetwork)(nil)

// NewNLMDisconnectConnectionToNetwork factory function for _NLMDisconnectConnectionToNetwork
func NewNLMDisconnectConnectionToNetwork(destinationNetworkAddress uint16, apduLength uint16) *_NLMDisconnectConnectionToNetwork {
	_result := &_NLMDisconnectConnectionToNetwork{
		NLMContract:               NewNLM(apduLength),
		DestinationNetworkAddress: destinationNetworkAddress,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMDisconnectConnectionToNetworkBuilder is a builder for NLMDisconnectConnectionToNetwork
type NLMDisconnectConnectionToNetworkBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationNetworkAddress uint16) NLMDisconnectConnectionToNetworkBuilder
	// WithDestinationNetworkAddress adds DestinationNetworkAddress (property field)
	WithDestinationNetworkAddress(uint16) NLMDisconnectConnectionToNetworkBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() NLMBuilder
	// Build builds the NLMDisconnectConnectionToNetwork or returns an error if something is wrong
	Build() (NLMDisconnectConnectionToNetwork, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMDisconnectConnectionToNetwork
}

// NewNLMDisconnectConnectionToNetworkBuilder() creates a NLMDisconnectConnectionToNetworkBuilder
func NewNLMDisconnectConnectionToNetworkBuilder() NLMDisconnectConnectionToNetworkBuilder {
	return &_NLMDisconnectConnectionToNetworkBuilder{_NLMDisconnectConnectionToNetwork: new(_NLMDisconnectConnectionToNetwork)}
}

type _NLMDisconnectConnectionToNetworkBuilder struct {
	*_NLMDisconnectConnectionToNetwork

	parentBuilder *_NLMBuilder

	err *utils.MultiError
}

var _ (NLMDisconnectConnectionToNetworkBuilder) = (*_NLMDisconnectConnectionToNetworkBuilder)(nil)

func (b *_NLMDisconnectConnectionToNetworkBuilder) setParent(contract NLMContract) {
	b.NLMContract = contract
	contract.(*_NLM)._SubType = b._NLMDisconnectConnectionToNetwork
}

func (b *_NLMDisconnectConnectionToNetworkBuilder) WithMandatoryFields(destinationNetworkAddress uint16) NLMDisconnectConnectionToNetworkBuilder {
	return b.WithDestinationNetworkAddress(destinationNetworkAddress)
}

func (b *_NLMDisconnectConnectionToNetworkBuilder) WithDestinationNetworkAddress(destinationNetworkAddress uint16) NLMDisconnectConnectionToNetworkBuilder {
	b.DestinationNetworkAddress = destinationNetworkAddress
	return b
}

func (b *_NLMDisconnectConnectionToNetworkBuilder) Build() (NLMDisconnectConnectionToNetwork, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NLMDisconnectConnectionToNetwork.deepCopy(), nil
}

func (b *_NLMDisconnectConnectionToNetworkBuilder) MustBuild() NLMDisconnectConnectionToNetwork {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NLMDisconnectConnectionToNetworkBuilder) Done() NLMBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewNLMBuilder().(*_NLMBuilder)
	}
	return b.parentBuilder
}

func (b *_NLMDisconnectConnectionToNetworkBuilder) buildForNLM() (NLM, error) {
	return b.Build()
}

func (b *_NLMDisconnectConnectionToNetworkBuilder) DeepCopy() any {
	_copy := b.CreateNLMDisconnectConnectionToNetworkBuilder().(*_NLMDisconnectConnectionToNetworkBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNLMDisconnectConnectionToNetworkBuilder creates a NLMDisconnectConnectionToNetworkBuilder
func (b *_NLMDisconnectConnectionToNetwork) CreateNLMDisconnectConnectionToNetworkBuilder() NLMDisconnectConnectionToNetworkBuilder {
	if b == nil {
		return NewNLMDisconnectConnectionToNetworkBuilder()
	}
	return &_NLMDisconnectConnectionToNetworkBuilder{_NLMDisconnectConnectionToNetwork: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMDisconnectConnectionToNetwork) GetMessageType() uint8 {
	return 0x09
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMDisconnectConnectionToNetwork) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMDisconnectConnectionToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMDisconnectConnectionToNetwork(structType any) NLMDisconnectConnectionToNetwork {
	if casted, ok := structType.(NLMDisconnectConnectionToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMDisconnectConnectionToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMDisconnectConnectionToNetwork) GetTypeName() string {
	return "NLMDisconnectConnectionToNetwork"
}

func (m *_NLMDisconnectConnectionToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_NLMDisconnectConnectionToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMDisconnectConnectionToNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMDisconnectConnectionToNetwork NLMDisconnectConnectionToNetwork, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMDisconnectConnectionToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMDisconnectConnectionToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationNetworkAddress, err := ReadSimpleField(ctx, "destinationNetworkAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationNetworkAddress' field"))
	}
	m.DestinationNetworkAddress = destinationNetworkAddress

	if closeErr := readBuffer.CloseContext("NLMDisconnectConnectionToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMDisconnectConnectionToNetwork")
	}

	return m, nil
}

func (m *_NLMDisconnectConnectionToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMDisconnectConnectionToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMDisconnectConnectionToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMDisconnectConnectionToNetwork")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationNetworkAddress", m.GetDestinationNetworkAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationNetworkAddress' field")
		}

		if popErr := writeBuffer.PopContext("NLMDisconnectConnectionToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMDisconnectConnectionToNetwork")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMDisconnectConnectionToNetwork) IsNLMDisconnectConnectionToNetwork() {}

func (m *_NLMDisconnectConnectionToNetwork) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMDisconnectConnectionToNetwork) deepCopy() *_NLMDisconnectConnectionToNetwork {
	if m == nil {
		return nil
	}
	_NLMDisconnectConnectionToNetworkCopy := &_NLMDisconnectConnectionToNetwork{
		m.NLMContract.(*_NLM).deepCopy(),
		m.DestinationNetworkAddress,
	}
	_NLMDisconnectConnectionToNetworkCopy.NLMContract.(*_NLM)._SubType = m
	return _NLMDisconnectConnectionToNetworkCopy
}

func (m *_NLMDisconnectConnectionToNetwork) String() string {
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
