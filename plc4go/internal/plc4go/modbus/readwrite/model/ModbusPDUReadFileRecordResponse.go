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
type ModbusPDUReadFileRecordResponse struct {
	*ModbusPDU
	Items []*ModbusPDUReadFileRecordResponseItem
}

// The corresponding interface
type IModbusPDUReadFileRecordResponse interface {
	IModbusPDU
	// GetItems returns Items
	GetItems() []*ModbusPDUReadFileRecordResponseItem
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
func (m *ModbusPDUReadFileRecordResponse) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadFileRecordResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadFileRecordResponse) FunctionFlag() uint8 {
	return 0x14
}

func (m *ModbusPDUReadFileRecordResponse) GetFunctionFlag() uint8 {
	return 0x14
}

func (m *ModbusPDUReadFileRecordResponse) Response() bool {
	return bool(true)
}

func (m *ModbusPDUReadFileRecordResponse) GetResponse() bool {
	return bool(true)
}

func (m *ModbusPDUReadFileRecordResponse) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadFileRecordResponse) GetItems() []*ModbusPDUReadFileRecordResponseItem {
	return m.Items
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUReadFileRecordResponse factory function for ModbusPDUReadFileRecordResponse
func NewModbusPDUReadFileRecordResponse(items []*ModbusPDUReadFileRecordResponseItem) *ModbusPDU {
	child := &ModbusPDUReadFileRecordResponse{
		Items:     items,
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadFileRecordResponse(structType interface{}) *ModbusPDUReadFileRecordResponse {
	castFunc := func(typ interface{}) *ModbusPDUReadFileRecordResponse {
		if casted, ok := typ.(ModbusPDUReadFileRecordResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadFileRecordResponse); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadFileRecordResponse(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadFileRecordResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadFileRecordResponse) GetTypeName() string {
	return "ModbusPDUReadFileRecordResponse"
}

func (m *ModbusPDUReadFileRecordResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadFileRecordResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *ModbusPDUReadFileRecordResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadFileRecordResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	items := make([]*ModbusPDUReadFileRecordResponseItem, 0)
	{
		_itemsLength := byteCount
		_itemsEndPos := readBuffer.GetPos() + uint16(_itemsLength)
		for readBuffer.GetPos() < _itemsEndPos {
			_item, _err := ModbusPDUReadFileRecordResponseItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items = append(items, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadFileRecordResponse{
		Items:     items,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadFileRecordResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	itemsArraySizeInBytes := func(items []*ModbusPDUReadFileRecordResponseItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes())
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadFileRecordResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
