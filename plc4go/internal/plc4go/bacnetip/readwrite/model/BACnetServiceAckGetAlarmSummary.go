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
type BACnetServiceAckGetAlarmSummary struct {
	*BACnetServiceAck
}

// The corresponding interface
type IBACnetServiceAckGetAlarmSummary interface {
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
func (m *BACnetServiceAckGetAlarmSummary) ServiceChoice() uint8 {
	return 0x03
}

func (m *BACnetServiceAckGetAlarmSummary) GetServiceChoice() uint8 {
	return 0x03
}

func (m *BACnetServiceAckGetAlarmSummary) InitializeParent(parent *BACnetServiceAck) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetAlarmSummary factory function for BACnetServiceAckGetAlarmSummary
func NewBACnetServiceAckGetAlarmSummary() *BACnetServiceAck {
	child := &BACnetServiceAckGetAlarmSummary{
		BACnetServiceAck: NewBACnetServiceAck(),
	}
	child.Child = child
	return child.BACnetServiceAck
}

func CastBACnetServiceAckGetAlarmSummary(structType interface{}) *BACnetServiceAckGetAlarmSummary {
	castFunc := func(typ interface{}) *BACnetServiceAckGetAlarmSummary {
		if casted, ok := typ.(BACnetServiceAckGetAlarmSummary); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckGetAlarmSummary); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckGetAlarmSummary(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckGetAlarmSummary(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckGetAlarmSummary) GetTypeName() string {
	return "BACnetServiceAckGetAlarmSummary"
}

func (m *BACnetServiceAckGetAlarmSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckGetAlarmSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetServiceAckGetAlarmSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckGetAlarmSummaryParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetAlarmSummary"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetAlarmSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckGetAlarmSummary{
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child.BACnetServiceAck, nil
}

func (m *BACnetServiceAckGetAlarmSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetAlarmSummary"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetAlarmSummary"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckGetAlarmSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
