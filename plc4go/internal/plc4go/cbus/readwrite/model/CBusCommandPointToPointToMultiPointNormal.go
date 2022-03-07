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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CBusCommandPointToPointToMultiPointNormal_CR byte = 0xD

// The data-structure of this message
type CBusCommandPointToPointToMultiPointNormal struct {
	*CBusPointToPointToMultipointCommand
	Application ApplicationIdContainer
	SalData     *SALData
	Crc         *Checksum
	PeekAlpha   byte
	Alpha       *Alpha

	// Arguments.
	Srchk bool
}

// The corresponding interface
type ICBusCommandPointToPointToMultiPointNormal interface {
	ICBusPointToPointToMultipointCommand
	// GetApplication returns Application
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData
	GetSalData() *SALData
	// GetCrc returns Crc
	GetCrc() *Checksum
	// GetPeekAlpha returns PeekAlpha
	GetPeekAlpha() byte
	// GetAlpha returns Alpha
	GetAlpha() *Alpha
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

func (m *CBusCommandPointToPointToMultiPointNormal) InitializeParent(parent *CBusPointToPointToMultipointCommand, bridgeAddress *BridgeAddress, networkRoute *NetworkRoute, peekedApplication byte) {
	m.CBusPointToPointToMultipointCommand.BridgeAddress = bridgeAddress
	m.CBusPointToPointToMultipointCommand.NetworkRoute = networkRoute
	m.CBusPointToPointToMultipointCommand.PeekedApplication = peekedApplication
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CBusCommandPointToPointToMultiPointNormal) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetSalData() *SALData {
	return m.SalData
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetCrc() *Checksum {
	return m.Crc
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetPeekAlpha() byte {
	return m.PeekAlpha
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetAlpha() *Alpha {
	return m.Alpha
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCBusCommandPointToPointToMultiPointNormal factory function for CBusCommandPointToPointToMultiPointNormal
func NewCBusCommandPointToPointToMultiPointNormal(application ApplicationIdContainer, salData *SALData, crc *Checksum, peekAlpha byte, alpha *Alpha, bridgeAddress *BridgeAddress, networkRoute *NetworkRoute, peekedApplication byte, srchk bool) *CBusPointToPointToMultipointCommand {
	child := &CBusCommandPointToPointToMultiPointNormal{
		Application:                         application,
		SalData:                             salData,
		Crc:                                 crc,
		PeekAlpha:                           peekAlpha,
		Alpha:                               alpha,
		CBusPointToPointToMultipointCommand: NewCBusPointToPointToMultipointCommand(bridgeAddress, networkRoute, peekedApplication, srchk),
	}
	child.Child = child
	return child.CBusPointToPointToMultipointCommand
}

func CastCBusCommandPointToPointToMultiPointNormal(structType interface{}) *CBusCommandPointToPointToMultiPointNormal {
	castFunc := func(typ interface{}) *CBusCommandPointToPointToMultiPointNormal {
		if casted, ok := typ.(CBusCommandPointToPointToMultiPointNormal); ok {
			return &casted
		}
		if casted, ok := typ.(*CBusCommandPointToPointToMultiPointNormal); ok {
			return casted
		}
		if casted, ok := typ.(CBusPointToPointToMultipointCommand); ok {
			return CastCBusCommandPointToPointToMultiPointNormal(casted.Child)
		}
		if casted, ok := typ.(*CBusPointToPointToMultipointCommand); ok {
			return CastCBusCommandPointToPointToMultiPointNormal(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetTypeName() string {
	return "CBusCommandPointToPointToMultiPointNormal"
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (application)
	lengthInBits += 8

	// Simple field (salData)
	lengthInBits += m.SalData.GetLengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += (*m.Crc).GetLengthInBits()
	}

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += (*m.Alpha).GetLengthInBits()
	}

	// Const Field (cr)
	lengthInBits += 8

	return lengthInBits
}

func (m *CBusCommandPointToPointToMultiPointNormal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusCommandPointToPointToMultiPointNormalParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusPointToPointToMultipointCommand, error) {
	if pullErr := readBuffer.PullContext("CBusCommandPointToPointToMultiPointNormal"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, pullErr
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (salData)
	if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
		return nil, pullErr
	}
	_salData, _salDataErr := SALDataParse(readBuffer)
	if _salDataErr != nil {
		return nil, errors.Wrap(_salDataErr, "Error parsing 'salData' field")
	}
	salData := CastSALData(_salData)
	if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc *Checksum = nil
	if srchk {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("crc"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := ChecksumParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'crc' field")
		default:
			crc = CastChecksum(_val)
			if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Peek Field (peekAlpha)
	currentPos = readBuffer.GetPos()
	peekAlpha, _err := readBuffer.ReadByte("peekAlpha")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekAlpha' field")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha *Alpha = nil
	if bool(bool(bool((peekAlpha) >= (0x67)))) && bool(bool(bool((peekAlpha) <= (0x7A)))) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field")
		default:
			alpha = CastAlpha(_val)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != CBusCommandPointToPointToMultiPointNormal_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CBusCommandPointToPointToMultiPointNormal_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	if closeErr := readBuffer.CloseContext("CBusCommandPointToPointToMultiPointNormal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CBusCommandPointToPointToMultiPointNormal{
		Application:                         application,
		SalData:                             CastSALData(salData),
		Crc:                                 CastChecksum(crc),
		PeekAlpha:                           peekAlpha,
		Alpha:                               CastAlpha(alpha),
		CBusPointToPointToMultipointCommand: &CBusPointToPointToMultipointCommand{},
	}
	_child.CBusPointToPointToMultipointCommand.Child = _child
	return _child.CBusPointToPointToMultipointCommand, nil
}

func (m *CBusCommandPointToPointToMultiPointNormal) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToPointToMultiPointNormal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return pushErr
		}
		_applicationErr := m.Application.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return popErr
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Simple Field (salData)
		if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
			return pushErr
		}
		_salDataErr := m.SalData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("salData"); popErr != nil {
			return popErr
		}
		if _salDataErr != nil {
			return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
		}

		// Optional Field (crc) (Can be skipped, if the value is null)
		var crc *Checksum = nil
		if m.Crc != nil {
			if pushErr := writeBuffer.PushContext("crc"); pushErr != nil {
				return pushErr
			}
			crc = m.Crc
			_crcErr := crc.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("crc"); popErr != nil {
				return popErr
			}
			if _crcErr != nil {
				return errors.Wrap(_crcErr, "Error serializing 'crc' field")
			}
		}

		// Optional Field (alpha) (Can be skipped, if the value is null)
		var alpha *Alpha = nil
		if m.Alpha != nil {
			if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
				return pushErr
			}
			alpha = m.Alpha
			_alphaErr := alpha.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
				return popErr
			}
			if _alphaErr != nil {
				return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
			}
		}

		// Const Field (cr)
		_crErr := writeBuffer.WriteByte("cr", 0xD)
		if _crErr != nil {
			return errors.Wrap(_crErr, "Error serializing 'cr' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToPointToMultiPointNormal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CBusCommandPointToPointToMultiPointNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
