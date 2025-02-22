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

// ActionState is an enum
type ActionState uint32

type IActionState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ActionState_actionStateIdle      ActionState = 0
	ActionState_actionStateExecuting ActionState = 1
	ActionState_actionStateDone      ActionState = 2
)

var ActionStateValues []ActionState

func init() {
	_ = errors.New
	ActionStateValues = []ActionState{
		ActionState_actionStateIdle,
		ActionState_actionStateExecuting,
		ActionState_actionStateDone,
	}
}

func ActionStateByValue(value uint32) (enum ActionState, ok bool) {
	switch value {
	case 0:
		return ActionState_actionStateIdle, true
	case 1:
		return ActionState_actionStateExecuting, true
	case 2:
		return ActionState_actionStateDone, true
	}
	return 0, false
}

func ActionStateByName(value string) (enum ActionState, ok bool) {
	switch value {
	case "actionStateIdle":
		return ActionState_actionStateIdle, true
	case "actionStateExecuting":
		return ActionState_actionStateExecuting, true
	case "actionStateDone":
		return ActionState_actionStateDone, true
	}
	return 0, false
}

func ActionStateKnows(value uint32) bool {
	for _, typeValue := range ActionStateValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastActionState(structType any) ActionState {
	castFunc := func(typ any) ActionState {
		if sActionState, ok := typ.(ActionState); ok {
			return sActionState
		}
		return 0
	}
	return castFunc(structType)
}

func (m ActionState) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m ActionState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ActionStateParse(ctx context.Context, theBytes []byte) (ActionState, error) {
	return ActionStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ActionStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ActionState, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("ActionState", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ActionState")
	}
	if enum, ok := ActionStateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ActionState")
		return ActionState(val), nil
	} else {
		return enum, nil
	}
}

func (e ActionState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ActionState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("ActionState", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e ActionState) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ActionState) PLC4XEnumName() string {
	switch e {
	case ActionState_actionStateIdle:
		return "actionStateIdle"
	case ActionState_actionStateExecuting:
		return "actionStateExecuting"
	case ActionState_actionStateDone:
		return "actionStateDone"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e ActionState) String() string {
	return e.PLC4XEnumName()
}
