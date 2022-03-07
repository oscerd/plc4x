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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type FirmataCommandSysex struct {
	*FirmataCommand
	Command *SysexCommand

	// Arguments.
	Response bool
}

// The corresponding interface
type IFirmataCommandSysex interface {
	IFirmataCommand
	// GetCommand returns Command
	GetCommand() *SysexCommand
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
func (m *FirmataCommandSysex) CommandCode() uint8 {
	return 0x0
}

func (m *FirmataCommandSysex) GetCommandCode() uint8 {
	return 0x0
}

func (m *FirmataCommandSysex) InitializeParent(parent *FirmataCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *FirmataCommandSysex) GetCommand() *SysexCommand {
	return m.Command
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewFirmataCommandSysex factory function for FirmataCommandSysex
func NewFirmataCommandSysex(command *SysexCommand, response bool) *FirmataCommand {
	child := &FirmataCommandSysex{
		Command:        command,
		FirmataCommand: NewFirmataCommand(response),
	}
	child.Child = child
	return child.FirmataCommand
}

func CastFirmataCommandSysex(structType interface{}) *FirmataCommandSysex {
	castFunc := func(typ interface{}) *FirmataCommandSysex {
		if casted, ok := typ.(FirmataCommandSysex); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandSysex); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandSysex(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandSysex(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandSysex) GetTypeName() string {
	return "FirmataCommandSysex"
}

func (m *FirmataCommandSysex) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataCommandSysex) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *FirmataCommandSysex) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandSysexParse(readBuffer utils.ReadBuffer, response bool) (*FirmataCommand, error) {
	if pullErr := readBuffer.PullContext("FirmataCommandSysex"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, pullErr
	}
	_command, _commandErr := SysexCommandParse(readBuffer, bool(response))
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}
	command := CastSysexCommand(_command)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0xF7) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0xF7),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("FirmataCommandSysex"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandSysex{
		Command:        CastSysexCommand(command),
		FirmataCommand: &FirmataCommand{},
	}
	_child.FirmataCommand.Child = _child
	return _child.FirmataCommand, nil
}

func (m *FirmataCommandSysex) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSysex"); pushErr != nil {
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

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0xF7))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSysex"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataCommandSysex) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
