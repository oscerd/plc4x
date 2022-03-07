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
type BACnetConfirmedServiceRequestAtomicReadFile struct {
	*BACnetConfirmedServiceRequest
	FileIdentifier *BACnetApplicationTagObjectIdentifier
	AccessMethod   *BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetConfirmedServiceRequestAtomicReadFile interface {
	IBACnetConfirmedServiceRequest
	// GetFileIdentifier returns FileIdentifier
	GetFileIdentifier() *BACnetApplicationTagObjectIdentifier
	// GetAccessMethod returns AccessMethod
	GetAccessMethod() *BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
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
func (m *BACnetConfirmedServiceRequestAtomicReadFile) ServiceChoice() uint8 {
	return 0x06
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) GetServiceChoice() uint8 {
	return 0x06
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestAtomicReadFile) GetFileIdentifier() *BACnetApplicationTagObjectIdentifier {
	return m.FileIdentifier
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) GetAccessMethod() *BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	return m.AccessMethod
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicReadFile factory function for BACnetConfirmedServiceRequestAtomicReadFile
func NewBACnetConfirmedServiceRequestAtomicReadFile(fileIdentifier *BACnetApplicationTagObjectIdentifier, accessMethod *BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, len uint16) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestAtomicReadFile{
		FileIdentifier:                fileIdentifier,
		AccessMethod:                  accessMethod,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestAtomicReadFile(structType interface{}) *BACnetConfirmedServiceRequestAtomicReadFile {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestAtomicReadFile {
		if casted, ok := typ.(BACnetConfirmedServiceRequestAtomicReadFile); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestAtomicReadFile); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestAtomicReadFile(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestAtomicReadFile(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFile"
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileIdentifier)
	lengthInBits += m.FileIdentifier.GetLengthInBits()

	// Simple field (accessMethod)
	lengthInBits += m.AccessMethod.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicReadFileParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFile"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (fileIdentifier)
	if pullErr := readBuffer.PullContext("fileIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_fileIdentifier, _fileIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _fileIdentifierErr != nil {
		return nil, errors.Wrap(_fileIdentifierErr, "Error parsing 'fileIdentifier' field")
	}
	fileIdentifier := CastBACnetApplicationTagObjectIdentifier(_fileIdentifier)
	if closeErr := readBuffer.CloseContext("fileIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (accessMethod)
	if pullErr := readBuffer.PullContext("accessMethod"); pullErr != nil {
		return nil, pullErr
	}
	_accessMethod, _accessMethodErr := BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParse(readBuffer)
	if _accessMethodErr != nil {
		return nil, errors.Wrap(_accessMethodErr, "Error parsing 'accessMethod' field")
	}
	accessMethod := CastBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(_accessMethod)
	if closeErr := readBuffer.CloseContext("accessMethod"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFile"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestAtomicReadFile{
		FileIdentifier:                CastBACnetApplicationTagObjectIdentifier(fileIdentifier),
		AccessMethod:                  CastBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(accessMethod),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFile"); pushErr != nil {
			return pushErr
		}

		// Simple Field (fileIdentifier)
		if pushErr := writeBuffer.PushContext("fileIdentifier"); pushErr != nil {
			return pushErr
		}
		_fileIdentifierErr := m.FileIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fileIdentifier"); popErr != nil {
			return popErr
		}
		if _fileIdentifierErr != nil {
			return errors.Wrap(_fileIdentifierErr, "Error serializing 'fileIdentifier' field")
		}

		// Simple Field (accessMethod)
		if pushErr := writeBuffer.PushContext("accessMethod"); pushErr != nil {
			return pushErr
		}
		_accessMethodErr := m.AccessMethod.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("accessMethod"); popErr != nil {
			return popErr
		}
		if _accessMethodErr != nil {
			return errors.Wrap(_accessMethodErr, "Error serializing 'accessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFile"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestAtomicReadFile) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
