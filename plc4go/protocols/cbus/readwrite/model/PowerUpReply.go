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

// PowerUpReply is the corresponding interface of PowerUpReply
type PowerUpReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Reply
	// GetPowerUpIndicator returns PowerUpIndicator (property field)
	GetPowerUpIndicator() PowerUp
	// IsPowerUpReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPowerUpReply()
	// CreateBuilder creates a PowerUpReplyBuilder
	CreatePowerUpReplyBuilder() PowerUpReplyBuilder
}

// _PowerUpReply is the data-structure of this message
type _PowerUpReply struct {
	ReplyContract
	PowerUpIndicator PowerUp
}

var _ PowerUpReply = (*_PowerUpReply)(nil)
var _ ReplyRequirements = (*_PowerUpReply)(nil)

// NewPowerUpReply factory function for _PowerUpReply
func NewPowerUpReply(peekedByte byte, powerUpIndicator PowerUp, cBusOptions CBusOptions, requestContext RequestContext) *_PowerUpReply {
	if powerUpIndicator == nil {
		panic("powerUpIndicator of type PowerUp for PowerUpReply must not be nil")
	}
	_result := &_PowerUpReply{
		ReplyContract:    NewReply(peekedByte, cBusOptions, requestContext),
		PowerUpIndicator: powerUpIndicator,
	}
	_result.ReplyContract.(*_Reply)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PowerUpReplyBuilder is a builder for PowerUpReply
type PowerUpReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(powerUpIndicator PowerUp) PowerUpReplyBuilder
	// WithPowerUpIndicator adds PowerUpIndicator (property field)
	WithPowerUpIndicator(PowerUp) PowerUpReplyBuilder
	// WithPowerUpIndicatorBuilder adds PowerUpIndicator (property field) which is build by the builder
	WithPowerUpIndicatorBuilder(func(PowerUpBuilder) PowerUpBuilder) PowerUpReplyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ReplyBuilder
	// Build builds the PowerUpReply or returns an error if something is wrong
	Build() (PowerUpReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PowerUpReply
}

// NewPowerUpReplyBuilder() creates a PowerUpReplyBuilder
func NewPowerUpReplyBuilder() PowerUpReplyBuilder {
	return &_PowerUpReplyBuilder{_PowerUpReply: new(_PowerUpReply)}
}

type _PowerUpReplyBuilder struct {
	*_PowerUpReply

	parentBuilder *_ReplyBuilder

	err *utils.MultiError
}

var _ (PowerUpReplyBuilder) = (*_PowerUpReplyBuilder)(nil)

func (b *_PowerUpReplyBuilder) setParent(contract ReplyContract) {
	b.ReplyContract = contract
	contract.(*_Reply)._SubType = b._PowerUpReply
}

func (b *_PowerUpReplyBuilder) WithMandatoryFields(powerUpIndicator PowerUp) PowerUpReplyBuilder {
	return b.WithPowerUpIndicator(powerUpIndicator)
}

func (b *_PowerUpReplyBuilder) WithPowerUpIndicator(powerUpIndicator PowerUp) PowerUpReplyBuilder {
	b.PowerUpIndicator = powerUpIndicator
	return b
}

func (b *_PowerUpReplyBuilder) WithPowerUpIndicatorBuilder(builderSupplier func(PowerUpBuilder) PowerUpBuilder) PowerUpReplyBuilder {
	builder := builderSupplier(b.PowerUpIndicator.CreatePowerUpBuilder())
	var err error
	b.PowerUpIndicator, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PowerUpBuilder failed"))
	}
	return b
}

func (b *_PowerUpReplyBuilder) Build() (PowerUpReply, error) {
	if b.PowerUpIndicator == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'powerUpIndicator' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PowerUpReply.deepCopy(), nil
}

func (b *_PowerUpReplyBuilder) MustBuild() PowerUpReply {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PowerUpReplyBuilder) Done() ReplyBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewReplyBuilder().(*_ReplyBuilder)
	}
	return b.parentBuilder
}

func (b *_PowerUpReplyBuilder) buildForReply() (Reply, error) {
	return b.Build()
}

func (b *_PowerUpReplyBuilder) DeepCopy() any {
	_copy := b.CreatePowerUpReplyBuilder().(*_PowerUpReplyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePowerUpReplyBuilder creates a PowerUpReplyBuilder
func (b *_PowerUpReply) CreatePowerUpReplyBuilder() PowerUpReplyBuilder {
	if b == nil {
		return NewPowerUpReplyBuilder()
	}
	return &_PowerUpReplyBuilder{_PowerUpReply: b.deepCopy()}
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

func (m *_PowerUpReply) GetParent() ReplyContract {
	return m.ReplyContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PowerUpReply) GetPowerUpIndicator() PowerUp {
	return m.PowerUpIndicator
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPowerUpReply(structType any) PowerUpReply {
	if casted, ok := structType.(PowerUpReply); ok {
		return casted
	}
	if casted, ok := structType.(*PowerUpReply); ok {
		return *casted
	}
	return nil
}

func (m *_PowerUpReply) GetTypeName() string {
	return "PowerUpReply"
}

func (m *_PowerUpReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ReplyContract.(*_Reply).getLengthInBits(ctx))

	// Simple field (powerUpIndicator)
	lengthInBits += m.PowerUpIndicator.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PowerUpReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PowerUpReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Reply, cBusOptions CBusOptions, requestContext RequestContext) (__powerUpReply PowerUpReply, err error) {
	m.ReplyContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PowerUpReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PowerUpReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	powerUpIndicator, err := ReadSimpleField[PowerUp](ctx, "powerUpIndicator", ReadComplex[PowerUp](PowerUpParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'powerUpIndicator' field"))
	}
	m.PowerUpIndicator = powerUpIndicator

	if closeErr := readBuffer.CloseContext("PowerUpReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PowerUpReply")
	}

	return m, nil
}

func (m *_PowerUpReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PowerUpReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PowerUpReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PowerUpReply")
		}

		if err := WriteSimpleField[PowerUp](ctx, "powerUpIndicator", m.GetPowerUpIndicator(), WriteComplex[PowerUp](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'powerUpIndicator' field")
		}

		if popErr := writeBuffer.PopContext("PowerUpReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PowerUpReply")
		}
		return nil
	}
	return m.ReplyContract.(*_Reply).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PowerUpReply) IsPowerUpReply() {}

func (m *_PowerUpReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PowerUpReply) deepCopy() *_PowerUpReply {
	if m == nil {
		return nil
	}
	_PowerUpReplyCopy := &_PowerUpReply{
		m.ReplyContract.(*_Reply).deepCopy(),
		utils.DeepCopy[PowerUp](m.PowerUpIndicator),
	}
	_PowerUpReplyCopy.ReplyContract.(*_Reply)._SubType = m
	return _PowerUpReplyCopy
}

func (m *_PowerUpReply) String() string {
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
