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

// TelephonyDataInternetConnectionRequestMade is the corresponding interface of TelephonyDataInternetConnectionRequestMade
type TelephonyDataInternetConnectionRequestMade interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TelephonyData
	// IsTelephonyDataInternetConnectionRequestMade is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataInternetConnectionRequestMade()
	// CreateBuilder creates a TelephonyDataInternetConnectionRequestMadeBuilder
	CreateTelephonyDataInternetConnectionRequestMadeBuilder() TelephonyDataInternetConnectionRequestMadeBuilder
}

// _TelephonyDataInternetConnectionRequestMade is the data-structure of this message
type _TelephonyDataInternetConnectionRequestMade struct {
	TelephonyDataContract
}

var _ TelephonyDataInternetConnectionRequestMade = (*_TelephonyDataInternetConnectionRequestMade)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataInternetConnectionRequestMade)(nil)

// NewTelephonyDataInternetConnectionRequestMade factory function for _TelephonyDataInternetConnectionRequestMade
func NewTelephonyDataInternetConnectionRequestMade(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataInternetConnectionRequestMade {
	_result := &_TelephonyDataInternetConnectionRequestMade{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TelephonyDataInternetConnectionRequestMadeBuilder is a builder for TelephonyDataInternetConnectionRequestMade
type TelephonyDataInternetConnectionRequestMadeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() TelephonyDataInternetConnectionRequestMadeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() TelephonyDataBuilder
	// Build builds the TelephonyDataInternetConnectionRequestMade or returns an error if something is wrong
	Build() (TelephonyDataInternetConnectionRequestMade, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TelephonyDataInternetConnectionRequestMade
}

// NewTelephonyDataInternetConnectionRequestMadeBuilder() creates a TelephonyDataInternetConnectionRequestMadeBuilder
func NewTelephonyDataInternetConnectionRequestMadeBuilder() TelephonyDataInternetConnectionRequestMadeBuilder {
	return &_TelephonyDataInternetConnectionRequestMadeBuilder{_TelephonyDataInternetConnectionRequestMade: new(_TelephonyDataInternetConnectionRequestMade)}
}

type _TelephonyDataInternetConnectionRequestMadeBuilder struct {
	*_TelephonyDataInternetConnectionRequestMade

	parentBuilder *_TelephonyDataBuilder

	err *utils.MultiError
}

var _ (TelephonyDataInternetConnectionRequestMadeBuilder) = (*_TelephonyDataInternetConnectionRequestMadeBuilder)(nil)

func (b *_TelephonyDataInternetConnectionRequestMadeBuilder) setParent(contract TelephonyDataContract) {
	b.TelephonyDataContract = contract
	contract.(*_TelephonyData)._SubType = b._TelephonyDataInternetConnectionRequestMade
}

func (b *_TelephonyDataInternetConnectionRequestMadeBuilder) WithMandatoryFields() TelephonyDataInternetConnectionRequestMadeBuilder {
	return b
}

func (b *_TelephonyDataInternetConnectionRequestMadeBuilder) Build() (TelephonyDataInternetConnectionRequestMade, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TelephonyDataInternetConnectionRequestMade.deepCopy(), nil
}

func (b *_TelephonyDataInternetConnectionRequestMadeBuilder) MustBuild() TelephonyDataInternetConnectionRequestMade {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TelephonyDataInternetConnectionRequestMadeBuilder) Done() TelephonyDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewTelephonyDataBuilder().(*_TelephonyDataBuilder)
	}
	return b.parentBuilder
}

func (b *_TelephonyDataInternetConnectionRequestMadeBuilder) buildForTelephonyData() (TelephonyData, error) {
	return b.Build()
}

func (b *_TelephonyDataInternetConnectionRequestMadeBuilder) DeepCopy() any {
	_copy := b.CreateTelephonyDataInternetConnectionRequestMadeBuilder().(*_TelephonyDataInternetConnectionRequestMadeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTelephonyDataInternetConnectionRequestMadeBuilder creates a TelephonyDataInternetConnectionRequestMadeBuilder
func (b *_TelephonyDataInternetConnectionRequestMade) CreateTelephonyDataInternetConnectionRequestMadeBuilder() TelephonyDataInternetConnectionRequestMadeBuilder {
	if b == nil {
		return NewTelephonyDataInternetConnectionRequestMadeBuilder()
	}
	return &_TelephonyDataInternetConnectionRequestMadeBuilder{_TelephonyDataInternetConnectionRequestMade: b.deepCopy()}
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

func (m *_TelephonyDataInternetConnectionRequestMade) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataInternetConnectionRequestMade(structType any) TelephonyDataInternetConnectionRequestMade {
	if casted, ok := structType.(TelephonyDataInternetConnectionRequestMade); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataInternetConnectionRequestMade); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataInternetConnectionRequestMade) GetTypeName() string {
	return "TelephonyDataInternetConnectionRequestMade"
}

func (m *_TelephonyDataInternetConnectionRequestMade) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_TelephonyDataInternetConnectionRequestMade) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataInternetConnectionRequestMade) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataInternetConnectionRequestMade TelephonyDataInternetConnectionRequestMade, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataInternetConnectionRequestMade"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataInternetConnectionRequestMade")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TelephonyDataInternetConnectionRequestMade"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataInternetConnectionRequestMade")
	}

	return m, nil
}

func (m *_TelephonyDataInternetConnectionRequestMade) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataInternetConnectionRequestMade) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataInternetConnectionRequestMade"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataInternetConnectionRequestMade")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataInternetConnectionRequestMade"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataInternetConnectionRequestMade")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataInternetConnectionRequestMade) IsTelephonyDataInternetConnectionRequestMade() {
}

func (m *_TelephonyDataInternetConnectionRequestMade) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TelephonyDataInternetConnectionRequestMade) deepCopy() *_TelephonyDataInternetConnectionRequestMade {
	if m == nil {
		return nil
	}
	_TelephonyDataInternetConnectionRequestMadeCopy := &_TelephonyDataInternetConnectionRequestMade{
		m.TelephonyDataContract.(*_TelephonyData).deepCopy(),
	}
	_TelephonyDataInternetConnectionRequestMadeCopy.TelephonyDataContract.(*_TelephonyData)._SubType = m
	return _TelephonyDataInternetConnectionRequestMadeCopy
}

func (m *_TelephonyDataInternetConnectionRequestMade) String() string {
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
