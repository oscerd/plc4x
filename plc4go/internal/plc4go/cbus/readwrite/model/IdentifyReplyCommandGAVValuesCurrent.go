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
type IdentifyReplyCommandGAVValuesCurrent struct {
	*IdentifyReplyCommand
	Values []byte
}

// The corresponding interface
type IIdentifyReplyCommandGAVValuesCurrent interface {
	IIdentifyReplyCommand
	// GetValues returns Values
	GetValues() []byte
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
func (m *IdentifyReplyCommandGAVValuesCurrent) Attribute() Attribute {
	return Attribute_GAVValuesCurrent
}

func (m *IdentifyReplyCommandGAVValuesCurrent) GetAttribute() Attribute {
	return Attribute_GAVValuesCurrent
}

func (m *IdentifyReplyCommandGAVValuesCurrent) InitializeParent(parent *IdentifyReplyCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *IdentifyReplyCommandGAVValuesCurrent) GetValues() []byte {
	return m.Values
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandGAVValuesCurrent factory function for IdentifyReplyCommandGAVValuesCurrent
func NewIdentifyReplyCommandGAVValuesCurrent(values []byte) *IdentifyReplyCommand {
	child := &IdentifyReplyCommandGAVValuesCurrent{
		Values:               values,
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandGAVValuesCurrent(structType interface{}) *IdentifyReplyCommandGAVValuesCurrent {
	castFunc := func(typ interface{}) *IdentifyReplyCommandGAVValuesCurrent {
		if casted, ok := typ.(IdentifyReplyCommandGAVValuesCurrent); ok {
			return &casted
		}
		if casted, ok := typ.(*IdentifyReplyCommandGAVValuesCurrent); ok {
			return casted
		}
		if casted, ok := typ.(IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandGAVValuesCurrent(casted.Child)
		}
		if casted, ok := typ.(*IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandGAVValuesCurrent(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *IdentifyReplyCommandGAVValuesCurrent) GetTypeName() string {
	return "IdentifyReplyCommandGAVValuesCurrent"
}

func (m *IdentifyReplyCommandGAVValuesCurrent) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandGAVValuesCurrent) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}

func (m *IdentifyReplyCommandGAVValuesCurrent) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandGAVValuesCurrentParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommand, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVValuesCurrent"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos
	// Byte Array field (values)
	numberOfBytesvalues := int(uint16(16))
	values, _readArrayErr := readBuffer.ReadByteArray("values", numberOfBytesvalues)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'values' field")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVValuesCurrent"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandGAVValuesCurrent{
		Values:               values,
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child.IdentifyReplyCommand, nil
}

func (m *IdentifyReplyCommandGAVValuesCurrent) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVValuesCurrent"); pushErr != nil {
			return pushErr
		}

		// Array Field (values)
		if m.Values != nil {
			// Byte Array field (values)
			_writeArrayErr := writeBuffer.WriteByteArray("values", m.Values)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'values' field")
			}
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVValuesCurrent"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandGAVValuesCurrent) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
