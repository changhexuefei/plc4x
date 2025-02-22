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

// OpcuaNodeIdServicesVariablePort is an enum
type OpcuaNodeIdServicesVariablePort int32

type IOpcuaNodeIdServicesVariablePort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariablePort_PortIdSubtype_EnumValues OpcuaNodeIdServicesVariablePort = 18950
)

var OpcuaNodeIdServicesVariablePortValues []OpcuaNodeIdServicesVariablePort

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariablePortValues = []OpcuaNodeIdServicesVariablePort{
		OpcuaNodeIdServicesVariablePort_PortIdSubtype_EnumValues,
	}
}

func OpcuaNodeIdServicesVariablePortByValue(value int32) (enum OpcuaNodeIdServicesVariablePort, ok bool) {
	switch value {
	case 18950:
		return OpcuaNodeIdServicesVariablePort_PortIdSubtype_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariablePortByName(value string) (enum OpcuaNodeIdServicesVariablePort, ok bool) {
	switch value {
	case "PortIdSubtype_EnumValues":
		return OpcuaNodeIdServicesVariablePort_PortIdSubtype_EnumValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariablePortKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariablePortValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariablePort(structType any) OpcuaNodeIdServicesVariablePort {
	castFunc := func(typ any) OpcuaNodeIdServicesVariablePort {
		if sOpcuaNodeIdServicesVariablePort, ok := typ.(OpcuaNodeIdServicesVariablePort); ok {
			return sOpcuaNodeIdServicesVariablePort
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariablePort) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariablePort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariablePortParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariablePort, error) {
	return OpcuaNodeIdServicesVariablePortParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariablePortParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariablePort, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariablePort", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariablePort")
	}
	if enum, ok := OpcuaNodeIdServicesVariablePortByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariablePort")
		return OpcuaNodeIdServicesVariablePort(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariablePort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariablePort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariablePort", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariablePort) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariablePort) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariablePort_PortIdSubtype_EnumValues:
		return "PortIdSubtype_EnumValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariablePort) String() string {
	return e.PLC4XEnumName()
}
