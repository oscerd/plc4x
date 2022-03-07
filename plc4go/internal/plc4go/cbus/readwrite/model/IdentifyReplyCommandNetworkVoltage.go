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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const IdentifyReplyCommandNetworkVoltage_DOT byte = 0x2C
const IdentifyReplyCommandNetworkVoltage_V byte = 0x56

// The data-structure of this message
type IdentifyReplyCommandNetworkVoltage struct {
	*IdentifyReplyCommand
	Volts             string
	VoltsDecimalPlace string
}

// The corresponding interface
type IIdentifyReplyCommandNetworkVoltage interface {
	IIdentifyReplyCommand
	// GetVolts returns Volts
	GetVolts() string
	// GetVoltsDecimalPlace returns VoltsDecimalPlace
	GetVoltsDecimalPlace() string
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
func (m *IdentifyReplyCommandNetworkVoltage) Attribute() Attribute {
	return Attribute_NetworkVoltage
}

func (m *IdentifyReplyCommandNetworkVoltage) GetAttribute() Attribute {
	return Attribute_NetworkVoltage
}

func (m *IdentifyReplyCommandNetworkVoltage) InitializeParent(parent *IdentifyReplyCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *IdentifyReplyCommandNetworkVoltage) GetVolts() string {
	return m.Volts
}

func (m *IdentifyReplyCommandNetworkVoltage) GetVoltsDecimalPlace() string {
	return m.VoltsDecimalPlace
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandNetworkVoltage factory function for IdentifyReplyCommandNetworkVoltage
func NewIdentifyReplyCommandNetworkVoltage(volts string, voltsDecimalPlace string) *IdentifyReplyCommand {
	child := &IdentifyReplyCommandNetworkVoltage{
		Volts:                volts,
		VoltsDecimalPlace:    voltsDecimalPlace,
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandNetworkVoltage(structType interface{}) *IdentifyReplyCommandNetworkVoltage {
	castFunc := func(typ interface{}) *IdentifyReplyCommandNetworkVoltage {
		if casted, ok := typ.(IdentifyReplyCommandNetworkVoltage); ok {
			return &casted
		}
		if casted, ok := typ.(*IdentifyReplyCommandNetworkVoltage); ok {
			return casted
		}
		if casted, ok := typ.(IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandNetworkVoltage(casted.Child)
		}
		if casted, ok := typ.(*IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandNetworkVoltage(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *IdentifyReplyCommandNetworkVoltage) GetTypeName() string {
	return "IdentifyReplyCommandNetworkVoltage"
}

func (m *IdentifyReplyCommandNetworkVoltage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandNetworkVoltage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (volts)
	lengthInBits += 2

	// Const Field (dot)
	lengthInBits += 8

	// Simple field (voltsDecimalPlace)
	lengthInBits += 2

	// Const Field (v)
	lengthInBits += 8

	return lengthInBits
}

func (m *IdentifyReplyCommandNetworkVoltage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandNetworkVoltageParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommand, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandNetworkVoltage"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (volts)
	_volts, _voltsErr := readBuffer.ReadString("volts", uint32(2))
	if _voltsErr != nil {
		return nil, errors.Wrap(_voltsErr, "Error parsing 'volts' field")
	}
	volts := _volts

	// Const Field (dot)
	dot, _dotErr := readBuffer.ReadByte("dot")
	if _dotErr != nil {
		return nil, errors.Wrap(_dotErr, "Error parsing 'dot' field")
	}
	if dot != IdentifyReplyCommandNetworkVoltage_DOT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", IdentifyReplyCommandNetworkVoltage_DOT) + " but got " + fmt.Sprintf("%d", dot))
	}

	// Simple Field (voltsDecimalPlace)
	_voltsDecimalPlace, _voltsDecimalPlaceErr := readBuffer.ReadString("voltsDecimalPlace", uint32(2))
	if _voltsDecimalPlaceErr != nil {
		return nil, errors.Wrap(_voltsDecimalPlaceErr, "Error parsing 'voltsDecimalPlace' field")
	}
	voltsDecimalPlace := _voltsDecimalPlace

	// Const Field (v)
	v, _vErr := readBuffer.ReadByte("v")
	if _vErr != nil {
		return nil, errors.Wrap(_vErr, "Error parsing 'v' field")
	}
	if v != IdentifyReplyCommandNetworkVoltage_V {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", IdentifyReplyCommandNetworkVoltage_V) + " but got " + fmt.Sprintf("%d", v))
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandNetworkVoltage"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandNetworkVoltage{
		Volts:                volts,
		VoltsDecimalPlace:    voltsDecimalPlace,
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child.IdentifyReplyCommand, nil
}

func (m *IdentifyReplyCommandNetworkVoltage) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandNetworkVoltage"); pushErr != nil {
			return pushErr
		}

		// Simple Field (volts)
		volts := string(m.Volts)
		_voltsErr := writeBuffer.WriteString("volts", uint32(2), "UTF-8", (volts))
		if _voltsErr != nil {
			return errors.Wrap(_voltsErr, "Error serializing 'volts' field")
		}

		// Const Field (dot)
		_dotErr := writeBuffer.WriteByte("dot", 0x2C)
		if _dotErr != nil {
			return errors.Wrap(_dotErr, "Error serializing 'dot' field")
		}

		// Simple Field (voltsDecimalPlace)
		voltsDecimalPlace := string(m.VoltsDecimalPlace)
		_voltsDecimalPlaceErr := writeBuffer.WriteString("voltsDecimalPlace", uint32(2), "UTF-8", (voltsDecimalPlace))
		if _voltsDecimalPlaceErr != nil {
			return errors.Wrap(_voltsDecimalPlaceErr, "Error serializing 'voltsDecimalPlace' field")
		}

		// Const Field (v)
		_vErr := writeBuffer.WriteByte("v", 0x56)
		if _vErr != nil {
			return errors.Wrap(_vErr, "Error serializing 'v' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandNetworkVoltage"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandNetworkVoltage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
