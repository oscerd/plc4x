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
type KnxGroupAddressFreeLevel struct {
	*KnxGroupAddress
	SubGroup uint16
}

// The corresponding interface
type IKnxGroupAddressFreeLevel interface {
	IKnxGroupAddress
	// GetSubGroup returns SubGroup
	GetSubGroup() uint16
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
func (m *KnxGroupAddressFreeLevel) NumLevels() uint8 {
	return uint8(1)
}

func (m *KnxGroupAddressFreeLevel) GetNumLevels() uint8 {
	return uint8(1)
}

func (m *KnxGroupAddressFreeLevel) InitializeParent(parent *KnxGroupAddress) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *KnxGroupAddressFreeLevel) GetSubGroup() uint16 {
	return m.SubGroup
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewKnxGroupAddressFreeLevel factory function for KnxGroupAddressFreeLevel
func NewKnxGroupAddressFreeLevel(subGroup uint16) *KnxGroupAddress {
	child := &KnxGroupAddressFreeLevel{
		SubGroup:        subGroup,
		KnxGroupAddress: NewKnxGroupAddress(),
	}
	child.Child = child
	return child.KnxGroupAddress
}

func CastKnxGroupAddressFreeLevel(structType interface{}) *KnxGroupAddressFreeLevel {
	castFunc := func(typ interface{}) *KnxGroupAddressFreeLevel {
		if casted, ok := typ.(KnxGroupAddressFreeLevel); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxGroupAddressFreeLevel); ok {
			return casted
		}
		if casted, ok := typ.(KnxGroupAddress); ok {
			return CastKnxGroupAddressFreeLevel(casted.Child)
		}
		if casted, ok := typ.(*KnxGroupAddress); ok {
			return CastKnxGroupAddressFreeLevel(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxGroupAddressFreeLevel) GetTypeName() string {
	return "KnxGroupAddressFreeLevel"
}

func (m *KnxGroupAddressFreeLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxGroupAddressFreeLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subGroup)
	lengthInBits += 16

	return lengthInBits
}

func (m *KnxGroupAddressFreeLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxGroupAddressFreeLevelParse(readBuffer utils.ReadBuffer, numLevels uint8) (*KnxGroupAddress, error) {
	if pullErr := readBuffer.PullContext("KnxGroupAddressFreeLevel"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (subGroup)
	_subGroup, _subGroupErr := readBuffer.ReadUint16("subGroup", 16)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field")
	}
	subGroup := _subGroup

	if closeErr := readBuffer.CloseContext("KnxGroupAddressFreeLevel"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxGroupAddressFreeLevel{
		SubGroup:        subGroup,
		KnxGroupAddress: &KnxGroupAddress{},
	}
	_child.KnxGroupAddress.Child = _child
	return _child.KnxGroupAddress, nil
}

func (m *KnxGroupAddressFreeLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxGroupAddressFreeLevel"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subGroup)
		subGroup := uint16(m.SubGroup)
		_subGroupErr := writeBuffer.WriteUint16("subGroup", 16, (subGroup))
		if _subGroupErr != nil {
			return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
		}

		if popErr := writeBuffer.PopContext("KnxGroupAddressFreeLevel"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxGroupAddressFreeLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
