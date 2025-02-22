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

// EncodedReplyCALReply is the corresponding interface of EncodedReplyCALReply
type EncodedReplyCALReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EncodedReply
	// GetCalReply returns CalReply (property field)
	GetCalReply() CALReply
	// IsEncodedReplyCALReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEncodedReplyCALReply()
	// CreateBuilder creates a EncodedReplyCALReplyBuilder
	CreateEncodedReplyCALReplyBuilder() EncodedReplyCALReplyBuilder
}

// _EncodedReplyCALReply is the data-structure of this message
type _EncodedReplyCALReply struct {
	EncodedReplyContract
	CalReply CALReply
}

var _ EncodedReplyCALReply = (*_EncodedReplyCALReply)(nil)
var _ EncodedReplyRequirements = (*_EncodedReplyCALReply)(nil)

// NewEncodedReplyCALReply factory function for _EncodedReplyCALReply
func NewEncodedReplyCALReply(peekedByte byte, calReply CALReply, cBusOptions CBusOptions, requestContext RequestContext) *_EncodedReplyCALReply {
	if calReply == nil {
		panic("calReply of type CALReply for EncodedReplyCALReply must not be nil")
	}
	_result := &_EncodedReplyCALReply{
		EncodedReplyContract: NewEncodedReply(peekedByte, cBusOptions, requestContext),
		CalReply:             calReply,
	}
	_result.EncodedReplyContract.(*_EncodedReply)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EncodedReplyCALReplyBuilder is a builder for EncodedReplyCALReply
type EncodedReplyCALReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(calReply CALReply) EncodedReplyCALReplyBuilder
	// WithCalReply adds CalReply (property field)
	WithCalReply(CALReply) EncodedReplyCALReplyBuilder
	// WithCalReplyBuilder adds CalReply (property field) which is build by the builder
	WithCalReplyBuilder(func(CALReplyBuilder) CALReplyBuilder) EncodedReplyCALReplyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() EncodedReplyBuilder
	// Build builds the EncodedReplyCALReply or returns an error if something is wrong
	Build() (EncodedReplyCALReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EncodedReplyCALReply
}

// NewEncodedReplyCALReplyBuilder() creates a EncodedReplyCALReplyBuilder
func NewEncodedReplyCALReplyBuilder() EncodedReplyCALReplyBuilder {
	return &_EncodedReplyCALReplyBuilder{_EncodedReplyCALReply: new(_EncodedReplyCALReply)}
}

type _EncodedReplyCALReplyBuilder struct {
	*_EncodedReplyCALReply

	parentBuilder *_EncodedReplyBuilder

	err *utils.MultiError
}

var _ (EncodedReplyCALReplyBuilder) = (*_EncodedReplyCALReplyBuilder)(nil)

func (b *_EncodedReplyCALReplyBuilder) setParent(contract EncodedReplyContract) {
	b.EncodedReplyContract = contract
	contract.(*_EncodedReply)._SubType = b._EncodedReplyCALReply
}

func (b *_EncodedReplyCALReplyBuilder) WithMandatoryFields(calReply CALReply) EncodedReplyCALReplyBuilder {
	return b.WithCalReply(calReply)
}

func (b *_EncodedReplyCALReplyBuilder) WithCalReply(calReply CALReply) EncodedReplyCALReplyBuilder {
	b.CalReply = calReply
	return b
}

func (b *_EncodedReplyCALReplyBuilder) WithCalReplyBuilder(builderSupplier func(CALReplyBuilder) CALReplyBuilder) EncodedReplyCALReplyBuilder {
	builder := builderSupplier(b.CalReply.CreateCALReplyBuilder())
	var err error
	b.CalReply, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CALReplyBuilder failed"))
	}
	return b
}

func (b *_EncodedReplyCALReplyBuilder) Build() (EncodedReplyCALReply, error) {
	if b.CalReply == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'calReply' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EncodedReplyCALReply.deepCopy(), nil
}

func (b *_EncodedReplyCALReplyBuilder) MustBuild() EncodedReplyCALReply {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EncodedReplyCALReplyBuilder) Done() EncodedReplyBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewEncodedReplyBuilder().(*_EncodedReplyBuilder)
	}
	return b.parentBuilder
}

func (b *_EncodedReplyCALReplyBuilder) buildForEncodedReply() (EncodedReply, error) {
	return b.Build()
}

func (b *_EncodedReplyCALReplyBuilder) DeepCopy() any {
	_copy := b.CreateEncodedReplyCALReplyBuilder().(*_EncodedReplyCALReplyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEncodedReplyCALReplyBuilder creates a EncodedReplyCALReplyBuilder
func (b *_EncodedReplyCALReply) CreateEncodedReplyCALReplyBuilder() EncodedReplyCALReplyBuilder {
	if b == nil {
		return NewEncodedReplyCALReplyBuilder()
	}
	return &_EncodedReplyCALReplyBuilder{_EncodedReplyCALReply: b.deepCopy()}
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

func (m *_EncodedReplyCALReply) GetParent() EncodedReplyContract {
	return m.EncodedReplyContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EncodedReplyCALReply) GetCalReply() CALReply {
	return m.CalReply
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEncodedReplyCALReply(structType any) EncodedReplyCALReply {
	if casted, ok := structType.(EncodedReplyCALReply); ok {
		return casted
	}
	if casted, ok := structType.(*EncodedReplyCALReply); ok {
		return *casted
	}
	return nil
}

func (m *_EncodedReplyCALReply) GetTypeName() string {
	return "EncodedReplyCALReply"
}

func (m *_EncodedReplyCALReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EncodedReplyContract.(*_EncodedReply).getLengthInBits(ctx))

	// Simple field (calReply)
	lengthInBits += m.CalReply.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_EncodedReplyCALReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EncodedReplyCALReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EncodedReply, cBusOptions CBusOptions, requestContext RequestContext) (__encodedReplyCALReply EncodedReplyCALReply, err error) {
	m.EncodedReplyContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EncodedReplyCALReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EncodedReplyCALReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	calReply, err := ReadSimpleField[CALReply](ctx, "calReply", ReadComplex[CALReply](CALReplyParseWithBufferProducer[CALReply]((CBusOptions)(cBusOptions), (RequestContext)(requestContext)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calReply' field"))
	}
	m.CalReply = calReply

	if closeErr := readBuffer.CloseContext("EncodedReplyCALReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EncodedReplyCALReply")
	}

	return m, nil
}

func (m *_EncodedReplyCALReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EncodedReplyCALReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EncodedReplyCALReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EncodedReplyCALReply")
		}

		if err := WriteSimpleField[CALReply](ctx, "calReply", m.GetCalReply(), WriteComplex[CALReply](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'calReply' field")
		}

		if popErr := writeBuffer.PopContext("EncodedReplyCALReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EncodedReplyCALReply")
		}
		return nil
	}
	return m.EncodedReplyContract.(*_EncodedReply).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EncodedReplyCALReply) IsEncodedReplyCALReply() {}

func (m *_EncodedReplyCALReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EncodedReplyCALReply) deepCopy() *_EncodedReplyCALReply {
	if m == nil {
		return nil
	}
	_EncodedReplyCALReplyCopy := &_EncodedReplyCALReply{
		m.EncodedReplyContract.(*_EncodedReply).deepCopy(),
		utils.DeepCopy[CALReply](m.CalReply),
	}
	_EncodedReplyCALReplyCopy.EncodedReplyContract.(*_EncodedReply)._SubType = m
	return _EncodedReplyCALReplyCopy
}

func (m *_EncodedReplyCALReply) String() string {
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
