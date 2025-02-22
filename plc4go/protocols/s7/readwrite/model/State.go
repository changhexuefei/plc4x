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

// State is the corresponding interface of State
type State interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetSIG_8 returns SIG_8 (property field)
	GetSIG_8() bool
	// GetSIG_7 returns SIG_7 (property field)
	GetSIG_7() bool
	// GetSIG_6 returns SIG_6 (property field)
	GetSIG_6() bool
	// GetSIG_5 returns SIG_5 (property field)
	GetSIG_5() bool
	// GetSIG_4 returns SIG_4 (property field)
	GetSIG_4() bool
	// GetSIG_3 returns SIG_3 (property field)
	GetSIG_3() bool
	// GetSIG_2 returns SIG_2 (property field)
	GetSIG_2() bool
	// GetSIG_1 returns SIG_1 (property field)
	GetSIG_1() bool
	// IsState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsState()
	// CreateBuilder creates a StateBuilder
	CreateStateBuilder() StateBuilder
}

// _State is the data-structure of this message
type _State struct {
	SIG_8 bool
	SIG_7 bool
	SIG_6 bool
	SIG_5 bool
	SIG_4 bool
	SIG_3 bool
	SIG_2 bool
	SIG_1 bool
}

var _ State = (*_State)(nil)

// NewState factory function for _State
func NewState(SIG_8 bool, SIG_7 bool, SIG_6 bool, SIG_5 bool, SIG_4 bool, SIG_3 bool, SIG_2 bool, SIG_1 bool) *_State {
	return &_State{SIG_8: SIG_8, SIG_7: SIG_7, SIG_6: SIG_6, SIG_5: SIG_5, SIG_4: SIG_4, SIG_3: SIG_3, SIG_2: SIG_2, SIG_1: SIG_1}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StateBuilder is a builder for State
type StateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(SIG_8 bool, SIG_7 bool, SIG_6 bool, SIG_5 bool, SIG_4 bool, SIG_3 bool, SIG_2 bool, SIG_1 bool) StateBuilder
	// WithSIG_8 adds SIG_8 (property field)
	WithSIG_8(bool) StateBuilder
	// WithSIG_7 adds SIG_7 (property field)
	WithSIG_7(bool) StateBuilder
	// WithSIG_6 adds SIG_6 (property field)
	WithSIG_6(bool) StateBuilder
	// WithSIG_5 adds SIG_5 (property field)
	WithSIG_5(bool) StateBuilder
	// WithSIG_4 adds SIG_4 (property field)
	WithSIG_4(bool) StateBuilder
	// WithSIG_3 adds SIG_3 (property field)
	WithSIG_3(bool) StateBuilder
	// WithSIG_2 adds SIG_2 (property field)
	WithSIG_2(bool) StateBuilder
	// WithSIG_1 adds SIG_1 (property field)
	WithSIG_1(bool) StateBuilder
	// Build builds the State or returns an error if something is wrong
	Build() (State, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() State
}

// NewStateBuilder() creates a StateBuilder
func NewStateBuilder() StateBuilder {
	return &_StateBuilder{_State: new(_State)}
}

type _StateBuilder struct {
	*_State

	err *utils.MultiError
}

var _ (StateBuilder) = (*_StateBuilder)(nil)

func (b *_StateBuilder) WithMandatoryFields(SIG_8 bool, SIG_7 bool, SIG_6 bool, SIG_5 bool, SIG_4 bool, SIG_3 bool, SIG_2 bool, SIG_1 bool) StateBuilder {
	return b.WithSIG_8(SIG_8).WithSIG_7(SIG_7).WithSIG_6(SIG_6).WithSIG_5(SIG_5).WithSIG_4(SIG_4).WithSIG_3(SIG_3).WithSIG_2(SIG_2).WithSIG_1(SIG_1)
}

func (b *_StateBuilder) WithSIG_8(SIG_8 bool) StateBuilder {
	b.SIG_8 = SIG_8
	return b
}

func (b *_StateBuilder) WithSIG_7(SIG_7 bool) StateBuilder {
	b.SIG_7 = SIG_7
	return b
}

func (b *_StateBuilder) WithSIG_6(SIG_6 bool) StateBuilder {
	b.SIG_6 = SIG_6
	return b
}

func (b *_StateBuilder) WithSIG_5(SIG_5 bool) StateBuilder {
	b.SIG_5 = SIG_5
	return b
}

func (b *_StateBuilder) WithSIG_4(SIG_4 bool) StateBuilder {
	b.SIG_4 = SIG_4
	return b
}

func (b *_StateBuilder) WithSIG_3(SIG_3 bool) StateBuilder {
	b.SIG_3 = SIG_3
	return b
}

func (b *_StateBuilder) WithSIG_2(SIG_2 bool) StateBuilder {
	b.SIG_2 = SIG_2
	return b
}

func (b *_StateBuilder) WithSIG_1(SIG_1 bool) StateBuilder {
	b.SIG_1 = SIG_1
	return b
}

func (b *_StateBuilder) Build() (State, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._State.deepCopy(), nil
}

func (b *_StateBuilder) MustBuild() State {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_StateBuilder) DeepCopy() any {
	_copy := b.CreateStateBuilder().(*_StateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStateBuilder creates a StateBuilder
func (b *_State) CreateStateBuilder() StateBuilder {
	if b == nil {
		return NewStateBuilder()
	}
	return &_StateBuilder{_State: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_State) GetSIG_8() bool {
	return m.SIG_8
}

func (m *_State) GetSIG_7() bool {
	return m.SIG_7
}

func (m *_State) GetSIG_6() bool {
	return m.SIG_6
}

func (m *_State) GetSIG_5() bool {
	return m.SIG_5
}

func (m *_State) GetSIG_4() bool {
	return m.SIG_4
}

func (m *_State) GetSIG_3() bool {
	return m.SIG_3
}

func (m *_State) GetSIG_2() bool {
	return m.SIG_2
}

func (m *_State) GetSIG_1() bool {
	return m.SIG_1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastState(structType any) State {
	if casted, ok := structType.(State); ok {
		return casted
	}
	if casted, ok := structType.(*State); ok {
		return *casted
	}
	return nil
}

func (m *_State) GetTypeName() string {
	return "State"
}

func (m *_State) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (SIG_8)
	lengthInBits += 1

	// Simple field (SIG_7)
	lengthInBits += 1

	// Simple field (SIG_6)
	lengthInBits += 1

	// Simple field (SIG_5)
	lengthInBits += 1

	// Simple field (SIG_4)
	lengthInBits += 1

	// Simple field (SIG_3)
	lengthInBits += 1

	// Simple field (SIG_2)
	lengthInBits += 1

	// Simple field (SIG_1)
	lengthInBits += 1

	return lengthInBits
}

func (m *_State) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StateParse(ctx context.Context, theBytes []byte) (State, error) {
	return StateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StateParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (State, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (State, error) {
		return StateParseWithBuffer(ctx, readBuffer)
	}
}

func StateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (State, error) {
	v, err := (&_State{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_State) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__state State, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("State"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for State")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	SIG_8, err := ReadSimpleField(ctx, "SIG_8", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_8' field"))
	}
	m.SIG_8 = SIG_8

	SIG_7, err := ReadSimpleField(ctx, "SIG_7", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_7' field"))
	}
	m.SIG_7 = SIG_7

	SIG_6, err := ReadSimpleField(ctx, "SIG_6", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_6' field"))
	}
	m.SIG_6 = SIG_6

	SIG_5, err := ReadSimpleField(ctx, "SIG_5", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_5' field"))
	}
	m.SIG_5 = SIG_5

	SIG_4, err := ReadSimpleField(ctx, "SIG_4", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_4' field"))
	}
	m.SIG_4 = SIG_4

	SIG_3, err := ReadSimpleField(ctx, "SIG_3", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_3' field"))
	}
	m.SIG_3 = SIG_3

	SIG_2, err := ReadSimpleField(ctx, "SIG_2", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_2' field"))
	}
	m.SIG_2 = SIG_2

	SIG_1, err := ReadSimpleField(ctx, "SIG_1", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'SIG_1' field"))
	}
	m.SIG_1 = SIG_1

	if closeErr := readBuffer.CloseContext("State"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for State")
	}

	return m, nil
}

func (m *_State) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_State) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("State"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for State")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_8", m.GetSIG_8(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_8' field")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_7", m.GetSIG_7(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_7' field")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_6", m.GetSIG_6(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_6' field")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_5", m.GetSIG_5(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_5' field")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_4", m.GetSIG_4(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_4' field")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_3", m.GetSIG_3(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_3' field")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_2", m.GetSIG_2(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_2' field")
	}

	if err := WriteSimpleField[bool](ctx, "SIG_1", m.GetSIG_1(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'SIG_1' field")
	}

	if popErr := writeBuffer.PopContext("State"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for State")
	}
	return nil
}

func (m *_State) IsState() {}

func (m *_State) DeepCopy() any {
	return m.deepCopy()
}

func (m *_State) deepCopy() *_State {
	if m == nil {
		return nil
	}
	_StateCopy := &_State{
		m.SIG_8,
		m.SIG_7,
		m.SIG_6,
		m.SIG_5,
		m.SIG_4,
		m.SIG_3,
		m.SIG_2,
		m.SIG_1,
	}
	return _StateCopy
}

func (m *_State) String() string {
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
