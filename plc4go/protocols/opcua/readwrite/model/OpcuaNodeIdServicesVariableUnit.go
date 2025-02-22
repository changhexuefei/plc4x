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

// OpcuaNodeIdServicesVariableUnit is an enum
type OpcuaNodeIdServicesVariableUnit int32

type IOpcuaNodeIdServicesVariableUnit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableUnit_UnitType_Symbol     OpcuaNodeIdServicesVariableUnit = 32443
	OpcuaNodeIdServicesVariableUnit_UnitType_UnitSystem OpcuaNodeIdServicesVariableUnit = 32445
	OpcuaNodeIdServicesVariableUnit_UnitType_Discipline OpcuaNodeIdServicesVariableUnit = 32446
)

var OpcuaNodeIdServicesVariableUnitValues []OpcuaNodeIdServicesVariableUnit

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableUnitValues = []OpcuaNodeIdServicesVariableUnit{
		OpcuaNodeIdServicesVariableUnit_UnitType_Symbol,
		OpcuaNodeIdServicesVariableUnit_UnitType_UnitSystem,
		OpcuaNodeIdServicesVariableUnit_UnitType_Discipline,
	}
}

func OpcuaNodeIdServicesVariableUnitByValue(value int32) (enum OpcuaNodeIdServicesVariableUnit, ok bool) {
	switch value {
	case 32443:
		return OpcuaNodeIdServicesVariableUnit_UnitType_Symbol, true
	case 32445:
		return OpcuaNodeIdServicesVariableUnit_UnitType_UnitSystem, true
	case 32446:
		return OpcuaNodeIdServicesVariableUnit_UnitType_Discipline, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUnitByName(value string) (enum OpcuaNodeIdServicesVariableUnit, ok bool) {
	switch value {
	case "UnitType_Symbol":
		return OpcuaNodeIdServicesVariableUnit_UnitType_Symbol, true
	case "UnitType_UnitSystem":
		return OpcuaNodeIdServicesVariableUnit_UnitType_UnitSystem, true
	case "UnitType_Discipline":
		return OpcuaNodeIdServicesVariableUnit_UnitType_Discipline, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUnitKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableUnitValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableUnit(structType any) OpcuaNodeIdServicesVariableUnit {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableUnit {
		if sOpcuaNodeIdServicesVariableUnit, ok := typ.(OpcuaNodeIdServicesVariableUnit); ok {
			return sOpcuaNodeIdServicesVariableUnit
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableUnit) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableUnit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableUnitParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableUnit, error) {
	return OpcuaNodeIdServicesVariableUnitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableUnitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableUnit, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableUnit", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableUnit")
	}
	if enum, ok := OpcuaNodeIdServicesVariableUnitByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableUnit")
		return OpcuaNodeIdServicesVariableUnit(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableUnit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableUnit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableUnit", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableUnit) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableUnit) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableUnit_UnitType_Symbol:
		return "UnitType_Symbol"
	case OpcuaNodeIdServicesVariableUnit_UnitType_UnitSystem:
		return "UnitType_UnitSystem"
	case OpcuaNodeIdServicesVariableUnit_UnitType_Discipline:
		return "UnitType_Discipline"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableUnit) String() string {
	return e.PLC4XEnumName()
}
