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
type SearchResponse struct {
	*KnxNetIpMessage
	HpaiControlEndpoint *HPAIControlEndpoint
	DibDeviceInfo       *DIBDeviceInfo
	DibSuppSvcFamilies  *DIBSuppSvcFamilies
}

// The corresponding interface
type ISearchResponse interface {
	IKnxNetIpMessage
	// GetHpaiControlEndpoint returns HpaiControlEndpoint
	GetHpaiControlEndpoint() *HPAIControlEndpoint
	// GetDibDeviceInfo returns DibDeviceInfo
	GetDibDeviceInfo() *DIBDeviceInfo
	// GetDibSuppSvcFamilies returns DibSuppSvcFamilies
	GetDibSuppSvcFamilies() *DIBSuppSvcFamilies
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
func (m *SearchResponse) MsgType() uint16 {
	return 0x0202
}

func (m *SearchResponse) GetMsgType() uint16 {
	return 0x0202
}

func (m *SearchResponse) InitializeParent(parent *KnxNetIpMessage) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *SearchResponse) GetHpaiControlEndpoint() *HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

func (m *SearchResponse) GetDibDeviceInfo() *DIBDeviceInfo {
	return m.DibDeviceInfo
}

func (m *SearchResponse) GetDibSuppSvcFamilies() *DIBSuppSvcFamilies {
	return m.DibSuppSvcFamilies
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewSearchResponse factory function for SearchResponse
func NewSearchResponse(hpaiControlEndpoint *HPAIControlEndpoint, dibDeviceInfo *DIBDeviceInfo, dibSuppSvcFamilies *DIBSuppSvcFamilies) *KnxNetIpMessage {
	child := &SearchResponse{
		HpaiControlEndpoint: hpaiControlEndpoint,
		DibDeviceInfo:       dibDeviceInfo,
		DibSuppSvcFamilies:  dibSuppSvcFamilies,
		KnxNetIpMessage:     NewKnxNetIpMessage(),
	}
	child.Child = child
	return child.KnxNetIpMessage
}

func CastSearchResponse(structType interface{}) *SearchResponse {
	castFunc := func(typ interface{}) *SearchResponse {
		if casted, ok := typ.(SearchResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*SearchResponse); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastSearchResponse(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastSearchResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SearchResponse) GetTypeName() string {
	return "SearchResponse"
}

func (m *SearchResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SearchResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits()

	// Simple field (dibDeviceInfo)
	lengthInBits += m.DibDeviceInfo.GetLengthInBits()

	// Simple field (dibSuppSvcFamilies)
	lengthInBits += m.DibSuppSvcFamilies.GetLengthInBits()

	return lengthInBits
}

func (m *SearchResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SearchResponseParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("SearchResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (hpaiControlEndpoint)
	if pullErr := readBuffer.PullContext("hpaiControlEndpoint"); pullErr != nil {
		return nil, pullErr
	}
	_hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParse(readBuffer)
	if _hpaiControlEndpointErr != nil {
		return nil, errors.Wrap(_hpaiControlEndpointErr, "Error parsing 'hpaiControlEndpoint' field")
	}
	hpaiControlEndpoint := CastHPAIControlEndpoint(_hpaiControlEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiControlEndpoint"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (dibDeviceInfo)
	if pullErr := readBuffer.PullContext("dibDeviceInfo"); pullErr != nil {
		return nil, pullErr
	}
	_dibDeviceInfo, _dibDeviceInfoErr := DIBDeviceInfoParse(readBuffer)
	if _dibDeviceInfoErr != nil {
		return nil, errors.Wrap(_dibDeviceInfoErr, "Error parsing 'dibDeviceInfo' field")
	}
	dibDeviceInfo := CastDIBDeviceInfo(_dibDeviceInfo)
	if closeErr := readBuffer.CloseContext("dibDeviceInfo"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (dibSuppSvcFamilies)
	if pullErr := readBuffer.PullContext("dibSuppSvcFamilies"); pullErr != nil {
		return nil, pullErr
	}
	_dibSuppSvcFamilies, _dibSuppSvcFamiliesErr := DIBSuppSvcFamiliesParse(readBuffer)
	if _dibSuppSvcFamiliesErr != nil {
		return nil, errors.Wrap(_dibSuppSvcFamiliesErr, "Error parsing 'dibSuppSvcFamilies' field")
	}
	dibSuppSvcFamilies := CastDIBSuppSvcFamilies(_dibSuppSvcFamilies)
	if closeErr := readBuffer.CloseContext("dibSuppSvcFamilies"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SearchResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SearchResponse{
		HpaiControlEndpoint: CastHPAIControlEndpoint(hpaiControlEndpoint),
		DibDeviceInfo:       CastDIBDeviceInfo(dibDeviceInfo),
		DibSuppSvcFamilies:  CastDIBSuppSvcFamilies(dibSuppSvcFamilies),
		KnxNetIpMessage:     &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child.KnxNetIpMessage, nil
}

func (m *SearchResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (hpaiControlEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiControlEndpoint"); pushErr != nil {
			return pushErr
		}
		_hpaiControlEndpointErr := m.HpaiControlEndpoint.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("hpaiControlEndpoint"); popErr != nil {
			return popErr
		}
		if _hpaiControlEndpointErr != nil {
			return errors.Wrap(_hpaiControlEndpointErr, "Error serializing 'hpaiControlEndpoint' field")
		}

		// Simple Field (dibDeviceInfo)
		if pushErr := writeBuffer.PushContext("dibDeviceInfo"); pushErr != nil {
			return pushErr
		}
		_dibDeviceInfoErr := m.DibDeviceInfo.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dibDeviceInfo"); popErr != nil {
			return popErr
		}
		if _dibDeviceInfoErr != nil {
			return errors.Wrap(_dibDeviceInfoErr, "Error serializing 'dibDeviceInfo' field")
		}

		// Simple Field (dibSuppSvcFamilies)
		if pushErr := writeBuffer.PushContext("dibSuppSvcFamilies"); pushErr != nil {
			return pushErr
		}
		_dibSuppSvcFamiliesErr := m.DibSuppSvcFamilies.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dibSuppSvcFamilies"); popErr != nil {
			return popErr
		}
		if _dibSuppSvcFamiliesErr != nil {
			return errors.Wrap(_dibSuppSvcFamiliesErr, "Error serializing 'dibSuppSvcFamilies' field")
		}

		if popErr := writeBuffer.PopContext("SearchResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SearchResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
