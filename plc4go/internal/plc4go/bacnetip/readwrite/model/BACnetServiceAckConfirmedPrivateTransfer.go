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
type BACnetServiceAckConfirmedPrivateTransfer struct {
	*BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckConfirmedPrivateTransfer interface {
	IBACnetServiceAck
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
func (m *BACnetServiceAckConfirmedPrivateTransfer) ServiceChoice() uint8 {
	return 0x12
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetServiceChoice() uint8 {
	return 0x12
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) InitializeParent(parent *BACnetServiceAck) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetServiceAckConfirmedPrivateTransfer factory function for BACnetServiceAckConfirmedPrivateTransfer
func NewBACnetServiceAckConfirmedPrivateTransfer() *BACnetServiceAck {
	child := &BACnetServiceAckConfirmedPrivateTransfer{
		BACnetServiceAck: NewBACnetServiceAck(),
	}
	child.Child = child
	return child.BACnetServiceAck
}

func CastBACnetServiceAckConfirmedPrivateTransfer(structType interface{}) *BACnetServiceAckConfirmedPrivateTransfer {
	castFunc := func(typ interface{}) *BACnetServiceAckConfirmedPrivateTransfer {
		if casted, ok := typ.(BACnetServiceAckConfirmedPrivateTransfer); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckConfirmedPrivateTransfer); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckConfirmedPrivateTransfer(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckConfirmedPrivateTransfer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetServiceAckConfirmedPrivateTransfer"
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckConfirmedPrivateTransferParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckConfirmedPrivateTransfer"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetServiceAckConfirmedPrivateTransfer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckConfirmedPrivateTransfer{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child.BACnetServiceAck, nil
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckConfirmedPrivateTransfer"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckConfirmedPrivateTransfer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckConfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
