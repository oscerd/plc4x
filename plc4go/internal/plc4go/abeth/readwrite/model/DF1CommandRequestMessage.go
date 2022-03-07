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
type DF1CommandRequestMessage struct {
	*DF1RequestMessage
	Command *DF1RequestCommand
}

// The corresponding interface
type IDF1CommandRequestMessage interface {
	IDF1RequestMessage
	// GetCommand returns Command
	GetCommand() *DF1RequestCommand
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
func (m *DF1CommandRequestMessage) CommandCode() uint8 {
	return 0x0F
}

func (m *DF1CommandRequestMessage) GetCommandCode() uint8 {
	return 0x0F
}

func (m *DF1CommandRequestMessage) InitializeParent(parent *DF1RequestMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) {
	m.DF1RequestMessage.DestinationAddress = destinationAddress
	m.DF1RequestMessage.SourceAddress = sourceAddress
	m.DF1RequestMessage.Status = status
	m.DF1RequestMessage.TransactionCounter = transactionCounter
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *DF1CommandRequestMessage) GetCommand() *DF1RequestCommand {
	return m.Command
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewDF1CommandRequestMessage factory function for DF1CommandRequestMessage
func NewDF1CommandRequestMessage(command *DF1RequestCommand, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) *DF1RequestMessage {
	child := &DF1CommandRequestMessage{
		Command:           command,
		DF1RequestMessage: NewDF1RequestMessage(destinationAddress, sourceAddress, status, transactionCounter),
	}
	child.Child = child
	return child.DF1RequestMessage
}

func CastDF1CommandRequestMessage(structType interface{}) *DF1CommandRequestMessage {
	castFunc := func(typ interface{}) *DF1CommandRequestMessage {
		if casted, ok := typ.(DF1CommandRequestMessage); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1CommandRequestMessage); ok {
			return casted
		}
		if casted, ok := typ.(DF1RequestMessage); ok {
			return CastDF1CommandRequestMessage(casted.Child)
		}
		if casted, ok := typ.(*DF1RequestMessage); ok {
			return CastDF1CommandRequestMessage(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1CommandRequestMessage) GetTypeName() string {
	return "DF1CommandRequestMessage"
}

func (m *DF1CommandRequestMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DF1CommandRequestMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	return lengthInBits
}

func (m *DF1CommandRequestMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1CommandRequestMessageParse(readBuffer utils.ReadBuffer) (*DF1RequestMessage, error) {
	if pullErr := readBuffer.PullContext("DF1CommandRequestMessage"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, pullErr
	}
	_command, _commandErr := DF1RequestCommandParse(readBuffer)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}
	command := CastDF1RequestCommand(_command)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DF1CommandRequestMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DF1CommandRequestMessage{
		Command:           CastDF1RequestCommand(command),
		DF1RequestMessage: &DF1RequestMessage{},
	}
	_child.DF1RequestMessage.Child = _child
	return _child.DF1RequestMessage, nil
}

func (m *DF1CommandRequestMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1CommandRequestMessage"); pushErr != nil {
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

		if popErr := writeBuffer.PopContext("DF1CommandRequestMessage"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DF1CommandRequestMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
