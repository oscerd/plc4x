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
type BVLCWriteBroadcastDistributionTable struct {
	*BVLC
	Table []*BVLCWriteBroadcastDistributionTableEntry
}

// The corresponding interface
type IBVLCWriteBroadcastDistributionTable interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCWriteBroadcastDistributionTable) BvlcFunction() uint8 {
return 0x01}


func (m *BVLCWriteBroadcastDistributionTable) InitializeParent(parent *BVLCbvlcPayloadLength uint16) {
}

func NewBVLCWriteBroadcastDistributionTable( table []*BVLCWriteBroadcastDistributionTableEntry , bvlcPayloadLength uint16 ) *BVLC {	child := &BVLCWriteBroadcastDistributionTable{
		Table: table,
    	BVLC: NewBVLC(bvlcPayloadLength),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCWriteBroadcastDistributionTable(structType interface{}) *BVLCWriteBroadcastDistributionTable {
	castFunc := func(typ interface{}) *BVLCWriteBroadcastDistributionTable {
		if casted, ok := typ.(BVLCWriteBroadcastDistributionTable); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCWriteBroadcastDistributionTable); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCWriteBroadcastDistributionTable(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCWriteBroadcastDistributionTable(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCWriteBroadcastDistributionTable) GetTypeName() string {
	return "BVLCWriteBroadcastDistributionTable"
}

func (m *BVLCWriteBroadcastDistributionTable) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCWriteBroadcastDistributionTable) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Array field
	if len(m.Table) > 0 {
		for _, element := range m.Table {
			lengthInBits += element.LengthInBits()
		}
	}

	return lengthInBits
}


func (m *BVLCWriteBroadcastDistributionTable) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCWriteBroadcastDistributionTableParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCWriteBroadcastDistributionTable"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (table)
	if pullErr := readBuffer.PullContext("table", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	table := make([]*BVLCWriteBroadcastDistributionTableEntry, 0)
	{
		_tableLength := bvlcPayloadLength
		_tableEndPos := readBuffer.GetPos() + uint16(_tableLength)
		for ;readBuffer.GetPos() < _tableEndPos; {
			_item, _err := BVLCWriteBroadcastDistributionTableEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'table' field")
			}
			table = append(table, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("table", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BVLCWriteBroadcastDistributionTable"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCWriteBroadcastDistributionTable{
		Table: table,
        BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCWriteBroadcastDistributionTable) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCWriteBroadcastDistributionTable"); pushErr != nil {
			return pushErr
		}

	// Array Field (table)
	if m.Table != nil {
		if pushErr := writeBuffer.PushContext("table", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.Table {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'table' field")
			}
		}
		if popErr := writeBuffer.PopContext("table", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

		if popErr := writeBuffer.PopContext("BVLCWriteBroadcastDistributionTable"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCWriteBroadcastDistributionTable) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}



