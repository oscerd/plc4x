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

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by code-generation. DO NOT EDIT.

type BacnetipXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m BacnetipXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "BACnetContextTag":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumberArgument := uint8(parsedUint0)
		dataType := model.BACnetDataTypeByName(parserArguments[1])
		return model.BACnetContextTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumberArgument, dataType)
	case "BACnetStatusFlags":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetStatusFlagsParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "NLM":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		apduLength := uint16(parsedUint0)
		return model.NLMParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), apduLength)
	case "BACnetNotificationParametersExtendedParameters":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetNotificationParametersExtendedParametersParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetNotificationParametersChangeOfValueNewValue":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetNotificationParametersChangeOfValueNewValueParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetPropertyStates":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetPropertyStatesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetConstructedData":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetConstructedDataParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetSegmentation":
		return model.BACnetSegmentationParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetConfirmedServiceACK":
		return model.BACnetConfirmedServiceACKParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetUnconfirmedServiceRequest":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		len := uint16(parsedUint0)
		return model.BACnetUnconfirmedServiceRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), len)
	case "BVLC":
		return model.BVLCParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetDateTime":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetDateTimeParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "APDU":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		apduLength := uint16(parsedUint0)
		return model.APDUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), apduLength)
	case "BACnetError":
		return model.BACnetErrorParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetTimeStamp":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetTimeStampParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetNotificationParameters":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetNotificationParametersParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetConfirmedServiceRequest":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		len := uint16(parsedUint0)
		return model.BACnetConfirmedServiceRequestParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), len)
	case "BACnetAddress":
		return model.BACnetAddressParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetApplicationTag":
		return model.BACnetApplicationTagParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetDeviceObjectPropertyReference":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetDeviceObjectPropertyReferenceParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetConstructedDataElement":
		return model.BACnetConstructedDataElementParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetPropertyValues":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetPropertyValuesParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "BACnetTagHeader":
		return model.BACnetTagHeaderParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetPropertyValue":
		return model.BACnetPropertyValueParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NLMInitalizeRoutingTablePortMapping":
		return model.NLMInitalizeRoutingTablePortMappingParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetServiceAck":
		return model.BACnetServiceAckParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BACnetAction":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		tagNumber := uint8(parsedUint0)
		return model.BACnetActionParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), tagNumber)
	case "NPDU":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 16)
		if err != nil {
			return nil, err
		}
		npduLength := uint16(parsedUint0)
		return model.NPDUParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), npduLength)
	case "BVLCWriteBroadcastDistributionTableEntry":
		return model.BVLCWriteBroadcastDistributionTableEntryParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
