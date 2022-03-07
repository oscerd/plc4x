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
type BACnetApplicationTagBitString struct {
	*BACnetApplicationTag
	Payload *BACnetTagPayloadBitString
}

// The corresponding interface
type IBACnetApplicationTagBitString interface {
	IBACnetApplicationTag
	// GetPayload returns Payload
	GetPayload() *BACnetTagPayloadBitString
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
func (m *BACnetApplicationTagBitString) ActualTagNumber() uint8 {
	return 0x8
}

func (m *BACnetApplicationTagBitString) GetActualTagNumber() uint8 {
	return 0x8
}

func (m *BACnetApplicationTagBitString) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader) {
	m.BACnetApplicationTag.Header = header
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagBitString) GetPayload() *BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagBitString factory function for BACnetApplicationTagBitString
func NewBACnetApplicationTagBitString(payload *BACnetTagPayloadBitString, header *BACnetTagHeader) *BACnetApplicationTag {
	child := &BACnetApplicationTagBitString{
		Payload:              payload,
		BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagBitString(structType interface{}) *BACnetApplicationTagBitString {
	castFunc := func(typ interface{}) *BACnetApplicationTagBitString {
		if casted, ok := typ.(BACnetApplicationTagBitString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagBitString); ok {
			return casted
		}
		if casted, ok := typ.(BACnetApplicationTag); ok {
			return CastBACnetApplicationTagBitString(casted.Child)
		}
		if casted, ok := typ.(*BACnetApplicationTag); ok {
			return CastBACnetApplicationTagBitString(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagBitString) GetTypeName() string {
	return "BACnetApplicationTagBitString"
}

func (m *BACnetApplicationTagBitString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetApplicationTagBitString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetApplicationTagBitString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagBitStringParse(readBuffer utils.ReadBuffer, header *BACnetTagHeader) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagBitString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadBitString(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagBitString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagBitString{
		Payload:              CastBACnetTagPayloadBitString(payload),
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagBitString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagBitString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagBitString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
