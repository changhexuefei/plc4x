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

// OpcuaNodeIdServicesVariableNode is an enum
type OpcuaNodeIdServicesVariableNode int32

type IOpcuaNodeIdServicesVariableNode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableNode_NodeVersion                   OpcuaNodeIdServicesVariableNode = 3068
	OpcuaNodeIdServicesVariableNode_NodeClass_EnumValues          OpcuaNodeIdServicesVariableNode = 11878
	OpcuaNodeIdServicesVariableNode_NodeAttributesMask_EnumValues OpcuaNodeIdServicesVariableNode = 11881
)

var OpcuaNodeIdServicesVariableNodeValues []OpcuaNodeIdServicesVariableNode

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableNodeValues = []OpcuaNodeIdServicesVariableNode{
		OpcuaNodeIdServicesVariableNode_NodeVersion,
		OpcuaNodeIdServicesVariableNode_NodeClass_EnumValues,
		OpcuaNodeIdServicesVariableNode_NodeAttributesMask_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableNodeByValue(value int32) (enum OpcuaNodeIdServicesVariableNode, ok bool) {
	switch value {
	case 11878:
		return OpcuaNodeIdServicesVariableNode_NodeClass_EnumValues, true
	case 11881:
		return OpcuaNodeIdServicesVariableNode_NodeAttributesMask_EnumValues, true
	case 3068:
		return OpcuaNodeIdServicesVariableNode_NodeVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNodeByName(value string) (enum OpcuaNodeIdServicesVariableNode, ok bool) {
	switch value {
	case "NodeClass_EnumValues":
		return OpcuaNodeIdServicesVariableNode_NodeClass_EnumValues, true
	case "NodeAttributesMask_EnumValues":
		return OpcuaNodeIdServicesVariableNode_NodeAttributesMask_EnumValues, true
	case "NodeVersion":
		return OpcuaNodeIdServicesVariableNode_NodeVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNodeKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableNodeValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableNode(structType any) OpcuaNodeIdServicesVariableNode {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableNode {
		if sOpcuaNodeIdServicesVariableNode, ok := typ.(OpcuaNodeIdServicesVariableNode); ok {
			return sOpcuaNodeIdServicesVariableNode
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableNode) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableNode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableNodeParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableNode, error) {
	return OpcuaNodeIdServicesVariableNodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableNodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableNode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableNode", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableNode")
	}
	if enum, ok := OpcuaNodeIdServicesVariableNodeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableNode")
		return OpcuaNodeIdServicesVariableNode(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableNode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableNode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableNode", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableNode) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableNode) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableNode_NodeClass_EnumValues:
		return "NodeClass_EnumValues"
	case OpcuaNodeIdServicesVariableNode_NodeAttributesMask_EnumValues:
		return "NodeAttributesMask_EnumValues"
	case OpcuaNodeIdServicesVariableNode_NodeVersion:
		return "NodeVersion"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableNode) String() string {
	return e.PLC4XEnumName()
}
