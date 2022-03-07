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
type MultipleServiceResponse struct {
	*CipService
	Status       uint8
	ExtStatus    uint8
	ServiceNb    uint16
	Offsets      []uint16
	ServicesData []byte

	// Arguments.
	ServiceLen uint16
}

// The corresponding interface
type IMultipleServiceResponse interface {
	ICipService
	// GetStatus returns Status
	GetStatus() uint8
	// GetExtStatus returns ExtStatus
	GetExtStatus() uint8
	// GetServiceNb returns ServiceNb
	GetServiceNb() uint16
	// GetOffsets returns Offsets
	GetOffsets() []uint16
	// GetServicesData returns ServicesData
	GetServicesData() []byte
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
func (m *MultipleServiceResponse) Service() uint8 {
	return 0x8A
}

func (m *MultipleServiceResponse) GetService() uint8 {
	return 0x8A
}

func (m *MultipleServiceResponse) InitializeParent(parent *CipService) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *MultipleServiceResponse) GetStatus() uint8 {
	return m.Status
}

func (m *MultipleServiceResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

func (m *MultipleServiceResponse) GetServiceNb() uint16 {
	return m.ServiceNb
}

func (m *MultipleServiceResponse) GetOffsets() []uint16 {
	return m.Offsets
}

func (m *MultipleServiceResponse) GetServicesData() []byte {
	return m.ServicesData
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewMultipleServiceResponse factory function for MultipleServiceResponse
func NewMultipleServiceResponse(status uint8, extStatus uint8, serviceNb uint16, offsets []uint16, servicesData []byte, serviceLen uint16) *CipService {
	child := &MultipleServiceResponse{
		Status:       status,
		ExtStatus:    extStatus,
		ServiceNb:    serviceNb,
		Offsets:      offsets,
		ServicesData: servicesData,
		CipService:   NewCipService(serviceLen),
	}
	child.Child = child
	return child.CipService
}

func CastMultipleServiceResponse(structType interface{}) *MultipleServiceResponse {
	castFunc := func(typ interface{}) *MultipleServiceResponse {
		if casted, ok := typ.(MultipleServiceResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*MultipleServiceResponse); ok {
			return casted
		}
		if casted, ok := typ.(CipService); ok {
			return CastMultipleServiceResponse(casted.Child)
		}
		if casted, ok := typ.(*CipService); ok {
			return CastMultipleServiceResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MultipleServiceResponse) GetTypeName() string {
	return "MultipleServiceResponse"
}

func (m *MultipleServiceResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MultipleServiceResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	// Simple field (serviceNb)
	lengthInBits += 16

	// Array field
	if len(m.Offsets) > 0 {
		lengthInBits += 16 * uint16(len(m.Offsets))
	}

	// Array field
	if len(m.ServicesData) > 0 {
		lengthInBits += 8 * uint16(len(m.ServicesData))
	}

	return lengthInBits
}

func (m *MultipleServiceResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MultipleServiceResponseParse(readBuffer utils.ReadBuffer, serviceLen uint16) (*CipService, error) {
	if pullErr := readBuffer.PullContext("MultipleServiceResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status

	// Simple Field (extStatus)
	_extStatus, _extStatusErr := readBuffer.ReadUint8("extStatus", 8)
	if _extStatusErr != nil {
		return nil, errors.Wrap(_extStatusErr, "Error parsing 'extStatus' field")
	}
	extStatus := _extStatus

	// Simple Field (serviceNb)
	_serviceNb, _serviceNbErr := readBuffer.ReadUint16("serviceNb", 16)
	if _serviceNbErr != nil {
		return nil, errors.Wrap(_serviceNbErr, "Error parsing 'serviceNb' field")
	}
	serviceNb := _serviceNb

	// Array field (offsets)
	if pullErr := readBuffer.PullContext("offsets", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	offsets := make([]uint16, serviceNb)
	{
		for curItem := uint16(0); curItem < uint16(serviceNb); curItem++ {
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'offsets' field")
			}
			offsets[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("offsets", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}
	// Byte Array field (servicesData)
	numberOfBytesservicesData := int(uint16(uint16(serviceLen)-uint16(uint16(6))) - uint16(uint16(uint16(uint16(2))*uint16(serviceNb))))
	servicesData, _readArrayErr := readBuffer.ReadByteArray("servicesData", numberOfBytesservicesData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'servicesData' field")
	}

	if closeErr := readBuffer.CloseContext("MultipleServiceResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MultipleServiceResponse{
		Status:       status,
		ExtStatus:    extStatus,
		ServiceNb:    serviceNb,
		Offsets:      offsets,
		ServicesData: servicesData,
		CipService:   &CipService{},
	}
	_child.CipService.Child = _child
	return _child.CipService, nil
}

func (m *MultipleServiceResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MultipleServiceResponse"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (status)
		status := uint8(m.Status)
		_statusErr := writeBuffer.WriteUint8("status", 8, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (extStatus)
		extStatus := uint8(m.ExtStatus)
		_extStatusErr := writeBuffer.WriteUint8("extStatus", 8, (extStatus))
		if _extStatusErr != nil {
			return errors.Wrap(_extStatusErr, "Error serializing 'extStatus' field")
		}

		// Simple Field (serviceNb)
		serviceNb := uint16(m.ServiceNb)
		_serviceNbErr := writeBuffer.WriteUint16("serviceNb", 16, (serviceNb))
		if _serviceNbErr != nil {
			return errors.Wrap(_serviceNbErr, "Error serializing 'serviceNb' field")
		}

		// Array Field (offsets)
		if m.Offsets != nil {
			if pushErr := writeBuffer.PushContext("offsets", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Offsets {
				_elementErr := writeBuffer.WriteUint16("", 16, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'offsets' field")
				}
			}
			if popErr := writeBuffer.PopContext("offsets", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Array Field (servicesData)
		if m.ServicesData != nil {
			// Byte Array field (servicesData)
			_writeArrayErr := writeBuffer.WriteByteArray("servicesData", m.ServicesData)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'servicesData' field")
			}
		}

		if popErr := writeBuffer.PopContext("MultipleServiceResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MultipleServiceResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
