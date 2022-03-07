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
type ConnectionResponseDataBlockTunnelConnection struct {
	*ConnectionResponseDataBlock
	KnxAddress *KnxAddress
}

// The corresponding interface
type IConnectionResponseDataBlockTunnelConnection interface {
	IConnectionResponseDataBlock
	// GetKnxAddress returns KnxAddress
	GetKnxAddress() *KnxAddress
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
func (m *ConnectionResponseDataBlockTunnelConnection) ConnectionType() uint8 {
	return 0x04
}

func (m *ConnectionResponseDataBlockTunnelConnection) GetConnectionType() uint8 {
	return 0x04
}

func (m *ConnectionResponseDataBlockTunnelConnection) InitializeParent(parent *ConnectionResponseDataBlock) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ConnectionResponseDataBlockTunnelConnection) GetKnxAddress() *KnxAddress {
	return m.KnxAddress
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewConnectionResponseDataBlockTunnelConnection factory function for ConnectionResponseDataBlockTunnelConnection
func NewConnectionResponseDataBlockTunnelConnection(knxAddress *KnxAddress) *ConnectionResponseDataBlock {
	child := &ConnectionResponseDataBlockTunnelConnection{
		KnxAddress:                  knxAddress,
		ConnectionResponseDataBlock: NewConnectionResponseDataBlock(),
	}
	child.Child = child
	return child.ConnectionResponseDataBlock
}

func CastConnectionResponseDataBlockTunnelConnection(structType interface{}) *ConnectionResponseDataBlockTunnelConnection {
	castFunc := func(typ interface{}) *ConnectionResponseDataBlockTunnelConnection {
		if casted, ok := typ.(ConnectionResponseDataBlockTunnelConnection); ok {
			return &casted
		}
		if casted, ok := typ.(*ConnectionResponseDataBlockTunnelConnection); ok {
			return casted
		}
		if casted, ok := typ.(ConnectionResponseDataBlock); ok {
			return CastConnectionResponseDataBlockTunnelConnection(casted.Child)
		}
		if casted, ok := typ.(*ConnectionResponseDataBlock); ok {
			return CastConnectionResponseDataBlockTunnelConnection(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ConnectionResponseDataBlockTunnelConnection) GetTypeName() string {
	return "ConnectionResponseDataBlockTunnelConnection"
}

func (m *ConnectionResponseDataBlockTunnelConnection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ConnectionResponseDataBlockTunnelConnection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (knxAddress)
	lengthInBits += m.KnxAddress.GetLengthInBits()

	return lengthInBits
}

func (m *ConnectionResponseDataBlockTunnelConnection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConnectionResponseDataBlockTunnelConnectionParse(readBuffer utils.ReadBuffer) (*ConnectionResponseDataBlock, error) {
	if pullErr := readBuffer.PullContext("ConnectionResponseDataBlockTunnelConnection"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (knxAddress)
	if pullErr := readBuffer.PullContext("knxAddress"); pullErr != nil {
		return nil, pullErr
	}
	_knxAddress, _knxAddressErr := KnxAddressParse(readBuffer)
	if _knxAddressErr != nil {
		return nil, errors.Wrap(_knxAddressErr, "Error parsing 'knxAddress' field")
	}
	knxAddress := CastKnxAddress(_knxAddress)
	if closeErr := readBuffer.CloseContext("knxAddress"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ConnectionResponseDataBlockTunnelConnection"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ConnectionResponseDataBlockTunnelConnection{
		KnxAddress:                  CastKnxAddress(knxAddress),
		ConnectionResponseDataBlock: &ConnectionResponseDataBlock{},
	}
	_child.ConnectionResponseDataBlock.Child = _child
	return _child.ConnectionResponseDataBlock, nil
}

func (m *ConnectionResponseDataBlockTunnelConnection) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionResponseDataBlockTunnelConnection"); pushErr != nil {
			return pushErr
		}

		// Simple Field (knxAddress)
		if pushErr := writeBuffer.PushContext("knxAddress"); pushErr != nil {
			return pushErr
		}
		_knxAddressErr := m.KnxAddress.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("knxAddress"); popErr != nil {
			return popErr
		}
		if _knxAddressErr != nil {
			return errors.Wrap(_knxAddressErr, "Error serializing 'knxAddress' field")
		}

		if popErr := writeBuffer.PopContext("ConnectionResponseDataBlockTunnelConnection"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ConnectionResponseDataBlockTunnelConnection) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
