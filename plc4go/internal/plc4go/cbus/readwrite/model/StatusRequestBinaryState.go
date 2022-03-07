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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type StatusRequestBinaryState struct {
	*StatusRequest
	Application byte
}

// The corresponding interface
type IStatusRequestBinaryState interface {
	IStatusRequest
	// GetApplication returns Application
	GetApplication() byte
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

func (m *StatusRequestBinaryState) InitializeParent(parent *StatusRequest, statusType byte) {
	m.StatusRequest.StatusType = statusType
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *StatusRequestBinaryState) GetApplication() byte {
	return m.Application
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewStatusRequestBinaryState factory function for StatusRequestBinaryState
func NewStatusRequestBinaryState(application byte, statusType byte) *StatusRequest {
	child := &StatusRequestBinaryState{
		Application:   application,
		StatusRequest: NewStatusRequest(statusType),
	}
	child.Child = child
	return child.StatusRequest
}

func CastStatusRequestBinaryState(structType interface{}) *StatusRequestBinaryState {
	castFunc := func(typ interface{}) *StatusRequestBinaryState {
		if casted, ok := typ.(StatusRequestBinaryState); ok {
			return &casted
		}
		if casted, ok := typ.(*StatusRequestBinaryState); ok {
			return casted
		}
		if casted, ok := typ.(StatusRequest); ok {
			return CastStatusRequestBinaryState(casted.Child)
		}
		if casted, ok := typ.(*StatusRequest); ok {
			return CastStatusRequestBinaryState(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *StatusRequestBinaryState) GetTypeName() string {
	return "StatusRequestBinaryState"
}

func (m *StatusRequestBinaryState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *StatusRequestBinaryState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *StatusRequestBinaryState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusRequestBinaryStateParse(readBuffer utils.ReadBuffer) (*StatusRequest, error) {
	if pullErr := readBuffer.PullContext("StatusRequestBinaryState"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x7A) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x7A),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (application)
	_application, _applicationErr := readBuffer.ReadByte("application")
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("StatusRequestBinaryState"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &StatusRequestBinaryState{
		Application:   application,
		StatusRequest: &StatusRequest{},
	}
	_child.StatusRequest.Child = _child
	return _child.StatusRequest, nil
}

func (m *StatusRequestBinaryState) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusRequestBinaryState"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x7A))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (application)
		application := byte(m.Application)
		_applicationErr := writeBuffer.WriteByte("application", (application))
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("StatusRequestBinaryState"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *StatusRequestBinaryState) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
