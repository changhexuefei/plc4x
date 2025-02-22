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

// BVLCBroadcastDistributionTableEntry is the corresponding interface of BVLCBroadcastDistributionTableEntry
type BVLCBroadcastDistributionTableEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetBroadcastDistributionMap returns BroadcastDistributionMap (property field)
	GetBroadcastDistributionMap() []uint8
	// IsBVLCBroadcastDistributionTableEntry is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCBroadcastDistributionTableEntry()
	// CreateBuilder creates a BVLCBroadcastDistributionTableEntryBuilder
	CreateBVLCBroadcastDistributionTableEntryBuilder() BVLCBroadcastDistributionTableEntryBuilder
}

// _BVLCBroadcastDistributionTableEntry is the data-structure of this message
type _BVLCBroadcastDistributionTableEntry struct {
	Ip                       []uint8
	Port                     uint16
	BroadcastDistributionMap []uint8
}

var _ BVLCBroadcastDistributionTableEntry = (*_BVLCBroadcastDistributionTableEntry)(nil)

// NewBVLCBroadcastDistributionTableEntry factory function for _BVLCBroadcastDistributionTableEntry
func NewBVLCBroadcastDistributionTableEntry(ip []uint8, port uint16, broadcastDistributionMap []uint8) *_BVLCBroadcastDistributionTableEntry {
	return &_BVLCBroadcastDistributionTableEntry{Ip: ip, Port: port, BroadcastDistributionMap: broadcastDistributionMap}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCBroadcastDistributionTableEntryBuilder is a builder for BVLCBroadcastDistributionTableEntry
type BVLCBroadcastDistributionTableEntryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ip []uint8, port uint16, broadcastDistributionMap []uint8) BVLCBroadcastDistributionTableEntryBuilder
	// WithIp adds Ip (property field)
	WithIp(...uint8) BVLCBroadcastDistributionTableEntryBuilder
	// WithPort adds Port (property field)
	WithPort(uint16) BVLCBroadcastDistributionTableEntryBuilder
	// WithBroadcastDistributionMap adds BroadcastDistributionMap (property field)
	WithBroadcastDistributionMap(...uint8) BVLCBroadcastDistributionTableEntryBuilder
	// Build builds the BVLCBroadcastDistributionTableEntry or returns an error if something is wrong
	Build() (BVLCBroadcastDistributionTableEntry, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCBroadcastDistributionTableEntry
}

// NewBVLCBroadcastDistributionTableEntryBuilder() creates a BVLCBroadcastDistributionTableEntryBuilder
func NewBVLCBroadcastDistributionTableEntryBuilder() BVLCBroadcastDistributionTableEntryBuilder {
	return &_BVLCBroadcastDistributionTableEntryBuilder{_BVLCBroadcastDistributionTableEntry: new(_BVLCBroadcastDistributionTableEntry)}
}

type _BVLCBroadcastDistributionTableEntryBuilder struct {
	*_BVLCBroadcastDistributionTableEntry

	err *utils.MultiError
}

var _ (BVLCBroadcastDistributionTableEntryBuilder) = (*_BVLCBroadcastDistributionTableEntryBuilder)(nil)

func (b *_BVLCBroadcastDistributionTableEntryBuilder) WithMandatoryFields(ip []uint8, port uint16, broadcastDistributionMap []uint8) BVLCBroadcastDistributionTableEntryBuilder {
	return b.WithIp(ip...).WithPort(port).WithBroadcastDistributionMap(broadcastDistributionMap...)
}

func (b *_BVLCBroadcastDistributionTableEntryBuilder) WithIp(ip ...uint8) BVLCBroadcastDistributionTableEntryBuilder {
	b.Ip = ip
	return b
}

func (b *_BVLCBroadcastDistributionTableEntryBuilder) WithPort(port uint16) BVLCBroadcastDistributionTableEntryBuilder {
	b.Port = port
	return b
}

func (b *_BVLCBroadcastDistributionTableEntryBuilder) WithBroadcastDistributionMap(broadcastDistributionMap ...uint8) BVLCBroadcastDistributionTableEntryBuilder {
	b.BroadcastDistributionMap = broadcastDistributionMap
	return b
}

func (b *_BVLCBroadcastDistributionTableEntryBuilder) Build() (BVLCBroadcastDistributionTableEntry, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCBroadcastDistributionTableEntry.deepCopy(), nil
}

func (b *_BVLCBroadcastDistributionTableEntryBuilder) MustBuild() BVLCBroadcastDistributionTableEntry {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BVLCBroadcastDistributionTableEntryBuilder) DeepCopy() any {
	_copy := b.CreateBVLCBroadcastDistributionTableEntryBuilder().(*_BVLCBroadcastDistributionTableEntryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCBroadcastDistributionTableEntryBuilder creates a BVLCBroadcastDistributionTableEntryBuilder
func (b *_BVLCBroadcastDistributionTableEntry) CreateBVLCBroadcastDistributionTableEntryBuilder() BVLCBroadcastDistributionTableEntryBuilder {
	if b == nil {
		return NewBVLCBroadcastDistributionTableEntryBuilder()
	}
	return &_BVLCBroadcastDistributionTableEntryBuilder{_BVLCBroadcastDistributionTableEntry: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCBroadcastDistributionTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCBroadcastDistributionTableEntry) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCBroadcastDistributionTableEntry) GetBroadcastDistributionMap() []uint8 {
	return m.BroadcastDistributionMap
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBVLCBroadcastDistributionTableEntry(structType any) BVLCBroadcastDistributionTableEntry {
	if casted, ok := structType.(BVLCBroadcastDistributionTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCBroadcastDistributionTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCBroadcastDistributionTableEntry) GetTypeName() string {
	return "BVLCBroadcastDistributionTableEntry"
}

func (m *_BVLCBroadcastDistributionTableEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Array field
	if len(m.BroadcastDistributionMap) > 0 {
		lengthInBits += 8 * uint16(len(m.BroadcastDistributionMap))
	}

	return lengthInBits
}

func (m *_BVLCBroadcastDistributionTableEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCBroadcastDistributionTableEntryParse(ctx context.Context, theBytes []byte) (BVLCBroadcastDistributionTableEntry, error) {
	return BVLCBroadcastDistributionTableEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BVLCBroadcastDistributionTableEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCBroadcastDistributionTableEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCBroadcastDistributionTableEntry, error) {
		return BVLCBroadcastDistributionTableEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BVLCBroadcastDistributionTableEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCBroadcastDistributionTableEntry, error) {
	v, err := (&_BVLCBroadcastDistributionTableEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BVLCBroadcastDistributionTableEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bVLCBroadcastDistributionTableEntry BVLCBroadcastDistributionTableEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCBroadcastDistributionTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCBroadcastDistributionTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ip, err := ReadCountArrayField[uint8](ctx, "ip", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ip' field"))
	}
	m.Ip = ip

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	broadcastDistributionMap, err := ReadCountArrayField[uint8](ctx, "broadcastDistributionMap", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'broadcastDistributionMap' field"))
	}
	m.BroadcastDistributionMap = broadcastDistributionMap

	if closeErr := readBuffer.CloseContext("BVLCBroadcastDistributionTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCBroadcastDistributionTableEntry")
	}

	return m, nil
}

func (m *_BVLCBroadcastDistributionTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCBroadcastDistributionTableEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BVLCBroadcastDistributionTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLCBroadcastDistributionTableEntry")
	}

	if err := WriteSimpleTypeArrayField(ctx, "ip", m.GetIp(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'ip' field")
	}

	if err := WriteSimpleField[uint16](ctx, "port", m.GetPort(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'port' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "broadcastDistributionMap", m.GetBroadcastDistributionMap(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'broadcastDistributionMap' field")
	}

	if popErr := writeBuffer.PopContext("BVLCBroadcastDistributionTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLCBroadcastDistributionTableEntry")
	}
	return nil
}

func (m *_BVLCBroadcastDistributionTableEntry) IsBVLCBroadcastDistributionTableEntry() {}

func (m *_BVLCBroadcastDistributionTableEntry) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCBroadcastDistributionTableEntry) deepCopy() *_BVLCBroadcastDistributionTableEntry {
	if m == nil {
		return nil
	}
	_BVLCBroadcastDistributionTableEntryCopy := &_BVLCBroadcastDistributionTableEntry{
		utils.DeepCopySlice[uint8, uint8](m.Ip),
		m.Port,
		utils.DeepCopySlice[uint8, uint8](m.BroadcastDistributionMap),
	}
	return _BVLCBroadcastDistributionTableEntryCopy
}

func (m *_BVLCBroadcastDistributionTableEntry) String() string {
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
