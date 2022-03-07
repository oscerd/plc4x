/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type FirmataMessageCommand struct {
	*FirmataMessage
	Command *FirmataCommand

	// Arguments.
	Response bool
}

// The corresponding interface
type IFirmataMessageCommand interface {
	IFirmataMessage
	// GetCommand returns Command
	GetCommand() *FirmataCommand
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataMessageCommand) MessageType() uint8 {
	return 0xF
}

func (m *FirmataMessageCommand) GetMessageType() uint8 {
	return 0xF
}

func (m *FirmataMessageCommand) InitializeParent(parent *FirmataMessage) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *FirmataMessageCommand) GetCommand() *FirmataCommand {
	return m.Command
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewFirmataMessageCommand factory function for FirmataMessageCommand
func NewFirmataMessageCommand(command *FirmataCommand, response bool) *FirmataMessage {
	child := &FirmataMessageCommand{
		Command:        command,
		FirmataMessage: NewFirmataMessage(response),
	}
	child.Child = child
	return child.FirmataMessage
}

func CastFirmataMessageCommand(structType interface{}) *FirmataMessageCommand {
	castFunc := func(typ interface{}) *FirmataMessageCommand {
		if casted, ok := typ.(FirmataMessageCommand); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataMessageCommand); ok {
			return casted
		}
		if casted, ok := typ.(FirmataMessage); ok {
			return CastFirmataMessageCommand(casted.Child)
		}
		if casted, ok := typ.(*FirmataMessage); ok {
			return CastFirmataMessageCommand(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataMessageCommand) GetTypeName() string {
	return "FirmataMessageCommand"
}

func (m *FirmataMessageCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataMessageCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	return lengthInBits
}

func (m *FirmataMessageCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataMessageCommandParse(readBuffer utils.ReadBuffer, response bool) (*FirmataMessage, error) {
	if pullErr := readBuffer.PullContext("FirmataMessageCommand"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, pullErr
	}
	_command, _commandErr := FirmataCommandParse(readBuffer, bool(response))
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}
	command := CastFirmataCommand(_command)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("FirmataMessageCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataMessageCommand{
		Command:        CastFirmataCommand(command),
		FirmataMessage: &FirmataMessage{},
	}
	_child.FirmataMessage.Child = _child
	return _child.FirmataMessage, nil
}

func (m *FirmataMessageCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageCommand"); pushErr != nil {
			return pushErr
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return pushErr
		}
		_commandErr := m.Command.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return popErr
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageCommand"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataMessageCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
