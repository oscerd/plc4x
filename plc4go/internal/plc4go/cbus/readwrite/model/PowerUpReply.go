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
type PowerUpReply struct {
	*Reply
	IsA *PowerUp
}

// The corresponding interface
type IPowerUpReply interface {
	IReply
	// GetIsA returns IsA
	GetIsA() *PowerUp
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

func (m *PowerUpReply) InitializeParent(parent *Reply, magicByte byte) {
	m.Reply.MagicByte = magicByte
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *PowerUpReply) GetIsA() *PowerUp {
	return m.IsA
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewPowerUpReply factory function for PowerUpReply
func NewPowerUpReply(isA *PowerUp, magicByte byte) *Reply {
	child := &PowerUpReply{
		IsA:   isA,
		Reply: NewReply(magicByte),
	}
	child.Child = child
	return child.Reply
}

func CastPowerUpReply(structType interface{}) *PowerUpReply {
	castFunc := func(typ interface{}) *PowerUpReply {
		if casted, ok := typ.(PowerUpReply); ok {
			return &casted
		}
		if casted, ok := typ.(*PowerUpReply); ok {
			return casted
		}
		if casted, ok := typ.(Reply); ok {
			return CastPowerUpReply(casted.Child)
		}
		if casted, ok := typ.(*Reply); ok {
			return CastPowerUpReply(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *PowerUpReply) GetTypeName() string {
	return "PowerUpReply"
}

func (m *PowerUpReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *PowerUpReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (isA)
	lengthInBits += m.IsA.GetLengthInBits()

	return lengthInBits
}

func (m *PowerUpReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func PowerUpReplyParse(readBuffer utils.ReadBuffer) (*Reply, error) {
	if pullErr := readBuffer.PullContext("PowerUpReply"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (isA)
	if pullErr := readBuffer.PullContext("isA"); pullErr != nil {
		return nil, pullErr
	}
	_isA, _isAErr := PowerUpParse(readBuffer)
	if _isAErr != nil {
		return nil, errors.Wrap(_isAErr, "Error parsing 'isA' field")
	}
	isA := CastPowerUp(_isA)
	if closeErr := readBuffer.CloseContext("isA"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("PowerUpReply"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &PowerUpReply{
		IsA:   CastPowerUp(isA),
		Reply: &Reply{},
	}
	_child.Reply.Child = _child
	return _child.Reply, nil
}

func (m *PowerUpReply) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PowerUpReply"); pushErr != nil {
			return pushErr
		}

		// Simple Field (isA)
		if pushErr := writeBuffer.PushContext("isA"); pushErr != nil {
			return pushErr
		}
		_isAErr := m.IsA.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("isA"); popErr != nil {
			return popErr
		}
		if _isAErr != nil {
			return errors.Wrap(_isAErr, "Error serializing 'isA' field")
		}

		if popErr := writeBuffer.PopContext("PowerUpReply"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *PowerUpReply) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
