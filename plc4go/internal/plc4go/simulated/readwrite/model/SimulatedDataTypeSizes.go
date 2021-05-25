//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type SimulatedDataTypeSizes uint8

type ISimulatedDataTypeSizes interface {
	DataTypeSize() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	SimulatedDataTypeSizes_BOOL           SimulatedDataTypeSizes = 1
	SimulatedDataTypeSizes_BYTE           SimulatedDataTypeSizes = 2
	SimulatedDataTypeSizes_WORD           SimulatedDataTypeSizes = 3
	SimulatedDataTypeSizes_DWORD          SimulatedDataTypeSizes = 4
	SimulatedDataTypeSizes_LWORD          SimulatedDataTypeSizes = 5
	SimulatedDataTypeSizes_SINT           SimulatedDataTypeSizes = 6
	SimulatedDataTypeSizes_INT            SimulatedDataTypeSizes = 7
	SimulatedDataTypeSizes_DINT           SimulatedDataTypeSizes = 8
	SimulatedDataTypeSizes_LINT           SimulatedDataTypeSizes = 9
	SimulatedDataTypeSizes_USINT          SimulatedDataTypeSizes = 10
	SimulatedDataTypeSizes_UINT           SimulatedDataTypeSizes = 11
	SimulatedDataTypeSizes_UDINT          SimulatedDataTypeSizes = 12
	SimulatedDataTypeSizes_ULINT          SimulatedDataTypeSizes = 13
	SimulatedDataTypeSizes_REAL           SimulatedDataTypeSizes = 14
	SimulatedDataTypeSizes_LREAL          SimulatedDataTypeSizes = 15
	SimulatedDataTypeSizes_TIME           SimulatedDataTypeSizes = 16
	SimulatedDataTypeSizes_LTIME          SimulatedDataTypeSizes = 17
	SimulatedDataTypeSizes_DATE           SimulatedDataTypeSizes = 18
	SimulatedDataTypeSizes_LDATE          SimulatedDataTypeSizes = 19
	SimulatedDataTypeSizes_TIME_OF_DAY    SimulatedDataTypeSizes = 20
	SimulatedDataTypeSizes_LTIME_OF_DAY   SimulatedDataTypeSizes = 21
	SimulatedDataTypeSizes_DATE_AND_TIME  SimulatedDataTypeSizes = 22
	SimulatedDataTypeSizes_LDATE_AND_TIME SimulatedDataTypeSizes = 23
	SimulatedDataTypeSizes_CHAR           SimulatedDataTypeSizes = 24
	SimulatedDataTypeSizes_WCHAR          SimulatedDataTypeSizes = 25
	SimulatedDataTypeSizes_STRING         SimulatedDataTypeSizes = 26
	SimulatedDataTypeSizes_WSTRING        SimulatedDataTypeSizes = 27
)

var SimulatedDataTypeSizesValues []SimulatedDataTypeSizes

func init() {
	_ = errors.New
	SimulatedDataTypeSizesValues = []SimulatedDataTypeSizes{
		SimulatedDataTypeSizes_BOOL,
		SimulatedDataTypeSizes_BYTE,
		SimulatedDataTypeSizes_WORD,
		SimulatedDataTypeSizes_DWORD,
		SimulatedDataTypeSizes_LWORD,
		SimulatedDataTypeSizes_SINT,
		SimulatedDataTypeSizes_INT,
		SimulatedDataTypeSizes_DINT,
		SimulatedDataTypeSizes_LINT,
		SimulatedDataTypeSizes_USINT,
		SimulatedDataTypeSizes_UINT,
		SimulatedDataTypeSizes_UDINT,
		SimulatedDataTypeSizes_ULINT,
		SimulatedDataTypeSizes_REAL,
		SimulatedDataTypeSizes_LREAL,
		SimulatedDataTypeSizes_TIME,
		SimulatedDataTypeSizes_LTIME,
		SimulatedDataTypeSizes_DATE,
		SimulatedDataTypeSizes_LDATE,
		SimulatedDataTypeSizes_TIME_OF_DAY,
		SimulatedDataTypeSizes_LTIME_OF_DAY,
		SimulatedDataTypeSizes_DATE_AND_TIME,
		SimulatedDataTypeSizes_LDATE_AND_TIME,
		SimulatedDataTypeSizes_CHAR,
		SimulatedDataTypeSizes_WCHAR,
		SimulatedDataTypeSizes_STRING,
		SimulatedDataTypeSizes_WSTRING,
	}
}

func (e SimulatedDataTypeSizes) DataTypeSize() uint8 {
	switch e {
	case 1:
		{ /* '1' */
			return 1
		}
	case 10:
		{ /* '10' */
			return 1
		}
	case 11:
		{ /* '11' */
			return 2
		}
	case 12:
		{ /* '12' */
			return 4
		}
	case 13:
		{ /* '13' */
			return 8
		}
	case 14:
		{ /* '14' */
			return 4
		}
	case 15:
		{ /* '15' */
			return 8
		}
	case 16:
		{ /* '16' */
			return 8
		}
	case 17:
		{ /* '17' */
			return 8
		}
	case 18:
		{ /* '18' */
			return 8
		}
	case 19:
		{ /* '19' */
			return 8
		}
	case 2:
		{ /* '2' */
			return 1
		}
	case 20:
		{ /* '20' */
			return 8
		}
	case 21:
		{ /* '21' */
			return 8
		}
	case 22:
		{ /* '22' */
			return 8
		}
	case 23:
		{ /* '23' */
			return 8
		}
	case 24:
		{ /* '24' */
			return 1
		}
	case 25:
		{ /* '25' */
			return 2
		}
	case 26:
		{ /* '26' */
			return 255
		}
	case 27:
		{ /* '27' */
			return 127
		}
	case 3:
		{ /* '3' */
			return 2
		}
	case 4:
		{ /* '4' */
			return 4
		}
	case 5:
		{ /* '5' */
			return 8
		}
	case 6:
		{ /* '6' */
			return 1
		}
	case 7:
		{ /* '7' */
			return 2
		}
	case 8:
		{ /* '8' */
			return 4
		}
	case 9:
		{ /* '9' */
			return 8
		}
	default:
		{
			return 0
		}
	}
}

func SimulatedDataTypeSizesFirstEnumForFieldDataTypeSize(value uint8) (SimulatedDataTypeSizes, error) {
	for _, sizeValue := range SimulatedDataTypeSizesValues {
		if sizeValue.DataTypeSize() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing DataTypeSize not found", value)
}
func SimulatedDataTypeSizesByValue(value uint8) SimulatedDataTypeSizes {
	switch value {
	case 1:
		return SimulatedDataTypeSizes_BOOL
	case 10:
		return SimulatedDataTypeSizes_USINT
	case 11:
		return SimulatedDataTypeSizes_UINT
	case 12:
		return SimulatedDataTypeSizes_UDINT
	case 13:
		return SimulatedDataTypeSizes_ULINT
	case 14:
		return SimulatedDataTypeSizes_REAL
	case 15:
		return SimulatedDataTypeSizes_LREAL
	case 16:
		return SimulatedDataTypeSizes_TIME
	case 17:
		return SimulatedDataTypeSizes_LTIME
	case 18:
		return SimulatedDataTypeSizes_DATE
	case 19:
		return SimulatedDataTypeSizes_LDATE
	case 2:
		return SimulatedDataTypeSizes_BYTE
	case 20:
		return SimulatedDataTypeSizes_TIME_OF_DAY
	case 21:
		return SimulatedDataTypeSizes_LTIME_OF_DAY
	case 22:
		return SimulatedDataTypeSizes_DATE_AND_TIME
	case 23:
		return SimulatedDataTypeSizes_LDATE_AND_TIME
	case 24:
		return SimulatedDataTypeSizes_CHAR
	case 25:
		return SimulatedDataTypeSizes_WCHAR
	case 26:
		return SimulatedDataTypeSizes_STRING
	case 27:
		return SimulatedDataTypeSizes_WSTRING
	case 3:
		return SimulatedDataTypeSizes_WORD
	case 4:
		return SimulatedDataTypeSizes_DWORD
	case 5:
		return SimulatedDataTypeSizes_LWORD
	case 6:
		return SimulatedDataTypeSizes_SINT
	case 7:
		return SimulatedDataTypeSizes_INT
	case 8:
		return SimulatedDataTypeSizes_DINT
	case 9:
		return SimulatedDataTypeSizes_LINT
	}
	return 0
}

func SimulatedDataTypeSizesByName(value string) SimulatedDataTypeSizes {
	switch value {
	case "BOOL":
		return SimulatedDataTypeSizes_BOOL
	case "USINT":
		return SimulatedDataTypeSizes_USINT
	case "UINT":
		return SimulatedDataTypeSizes_UINT
	case "UDINT":
		return SimulatedDataTypeSizes_UDINT
	case "ULINT":
		return SimulatedDataTypeSizes_ULINT
	case "REAL":
		return SimulatedDataTypeSizes_REAL
	case "LREAL":
		return SimulatedDataTypeSizes_LREAL
	case "TIME":
		return SimulatedDataTypeSizes_TIME
	case "LTIME":
		return SimulatedDataTypeSizes_LTIME
	case "DATE":
		return SimulatedDataTypeSizes_DATE
	case "LDATE":
		return SimulatedDataTypeSizes_LDATE
	case "BYTE":
		return SimulatedDataTypeSizes_BYTE
	case "TIME_OF_DAY":
		return SimulatedDataTypeSizes_TIME_OF_DAY
	case "LTIME_OF_DAY":
		return SimulatedDataTypeSizes_LTIME_OF_DAY
	case "DATE_AND_TIME":
		return SimulatedDataTypeSizes_DATE_AND_TIME
	case "LDATE_AND_TIME":
		return SimulatedDataTypeSizes_LDATE_AND_TIME
	case "CHAR":
		return SimulatedDataTypeSizes_CHAR
	case "WCHAR":
		return SimulatedDataTypeSizes_WCHAR
	case "STRING":
		return SimulatedDataTypeSizes_STRING
	case "WSTRING":
		return SimulatedDataTypeSizes_WSTRING
	case "WORD":
		return SimulatedDataTypeSizes_WORD
	case "DWORD":
		return SimulatedDataTypeSizes_DWORD
	case "LWORD":
		return SimulatedDataTypeSizes_LWORD
	case "SINT":
		return SimulatedDataTypeSizes_SINT
	case "INT":
		return SimulatedDataTypeSizes_INT
	case "DINT":
		return SimulatedDataTypeSizes_DINT
	case "LINT":
		return SimulatedDataTypeSizes_LINT
	}
	return 0
}

func CastSimulatedDataTypeSizes(structType interface{}) SimulatedDataTypeSizes {
	castFunc := func(typ interface{}) SimulatedDataTypeSizes {
		if sSimulatedDataTypeSizes, ok := typ.(SimulatedDataTypeSizes); ok {
			return sSimulatedDataTypeSizes
		}
		return 0
	}
	return castFunc(structType)
}

func (m SimulatedDataTypeSizes) LengthInBits() uint16 {
	return 8
}

func (m SimulatedDataTypeSizes) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SimulatedDataTypeSizesParse(readBuffer utils.ReadBuffer) (SimulatedDataTypeSizes, error) {
	val, err := readBuffer.ReadUint8("SimulatedDataTypeSizes", 8)
	if err != nil {
		return 0, nil
	}
	return SimulatedDataTypeSizesByValue(val), nil
}

func (e SimulatedDataTypeSizes) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("SimulatedDataTypeSizes", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e SimulatedDataTypeSizes) name() string {
	switch e {
	case SimulatedDataTypeSizes_BOOL:
		return "BOOL"
	case SimulatedDataTypeSizes_USINT:
		return "USINT"
	case SimulatedDataTypeSizes_UINT:
		return "UINT"
	case SimulatedDataTypeSizes_UDINT:
		return "UDINT"
	case SimulatedDataTypeSizes_ULINT:
		return "ULINT"
	case SimulatedDataTypeSizes_REAL:
		return "REAL"
	case SimulatedDataTypeSizes_LREAL:
		return "LREAL"
	case SimulatedDataTypeSizes_TIME:
		return "TIME"
	case SimulatedDataTypeSizes_LTIME:
		return "LTIME"
	case SimulatedDataTypeSizes_DATE:
		return "DATE"
	case SimulatedDataTypeSizes_LDATE:
		return "LDATE"
	case SimulatedDataTypeSizes_BYTE:
		return "BYTE"
	case SimulatedDataTypeSizes_TIME_OF_DAY:
		return "TIME_OF_DAY"
	case SimulatedDataTypeSizes_LTIME_OF_DAY:
		return "LTIME_OF_DAY"
	case SimulatedDataTypeSizes_DATE_AND_TIME:
		return "DATE_AND_TIME"
	case SimulatedDataTypeSizes_LDATE_AND_TIME:
		return "LDATE_AND_TIME"
	case SimulatedDataTypeSizes_CHAR:
		return "CHAR"
	case SimulatedDataTypeSizes_WCHAR:
		return "WCHAR"
	case SimulatedDataTypeSizes_STRING:
		return "STRING"
	case SimulatedDataTypeSizes_WSTRING:
		return "WSTRING"
	case SimulatedDataTypeSizes_WORD:
		return "WORD"
	case SimulatedDataTypeSizes_DWORD:
		return "DWORD"
	case SimulatedDataTypeSizes_LWORD:
		return "LWORD"
	case SimulatedDataTypeSizes_SINT:
		return "SINT"
	case SimulatedDataTypeSizes_INT:
		return "INT"
	case SimulatedDataTypeSizes_DINT:
		return "DINT"
	case SimulatedDataTypeSizes_LINT:
		return "LINT"
	}
	return ""
}

func (e SimulatedDataTypeSizes) String() string {
	return e.name()
}