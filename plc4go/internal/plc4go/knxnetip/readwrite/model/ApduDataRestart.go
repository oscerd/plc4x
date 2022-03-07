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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataRestart struct {
	*ApduData

	// Arguments.
	DataLength uint8
}

// The corresponding interface
type IApduDataRestart interface {
	IApduData
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
func (m *ApduDataRestart) ApciType() uint8 {
	return 0xE
}

func (m *ApduDataRestart) GetApciType() uint8 {
	return 0xE
}

func (m *ApduDataRestart) InitializeParent(parent *ApduData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataRestart factory function for ApduDataRestart
func NewApduDataRestart(dataLength uint8) *ApduData {
	child := &ApduDataRestart{
		ApduData: NewApduData(dataLength),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataRestart(structType interface{}) *ApduDataRestart {
	castFunc := func(typ interface{}) *ApduDataRestart {
		if casted, ok := typ.(ApduDataRestart); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataRestart); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataRestart(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataRestart(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataRestart) GetTypeName() string {
	return "ApduDataRestart"
}

func (m *ApduDataRestart) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataRestart) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataRestart) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataRestartParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataRestart"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataRestart"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataRestart{
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataRestart) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataRestart"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataRestart"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataRestart) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
