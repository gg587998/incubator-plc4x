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
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
)

// The data-structure of this message
type BACnetErrorVTOpen struct {
	BACnetError
}

// The corresponding interface
type IBACnetErrorVTOpen interface {
	IBACnetError
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m BACnetErrorVTOpen) ServiceChoice() uint8 {
	return 0x15
}

func (m BACnetErrorVTOpen) initialize() spi.Message {
	return m
}

func NewBACnetErrorVTOpen() BACnetErrorInitializer {
	return &BACnetErrorVTOpen{}
}

func CastIBACnetErrorVTOpen(structType interface{}) IBACnetErrorVTOpen {
	castFunc := func(typ interface{}) IBACnetErrorVTOpen {
		if iBACnetErrorVTOpen, ok := typ.(IBACnetErrorVTOpen); ok {
			return iBACnetErrorVTOpen
		}
		return nil
	}
	return castFunc(structType)
}

func CastBACnetErrorVTOpen(structType interface{}) BACnetErrorVTOpen {
	castFunc := func(typ interface{}) BACnetErrorVTOpen {
		if sBACnetErrorVTOpen, ok := typ.(BACnetErrorVTOpen); ok {
			return sBACnetErrorVTOpen
		}
		return BACnetErrorVTOpen{}
	}
	return castFunc(structType)
}

func (m BACnetErrorVTOpen) LengthInBits() uint16 {
	var lengthInBits uint16 = m.BACnetError.LengthInBits()

	return lengthInBits
}

func (m BACnetErrorVTOpen) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorVTOpenParse(io spi.ReadBuffer) (BACnetErrorInitializer, error) {

	// Create the instance
	return NewBACnetErrorVTOpen(), nil
}

func (m BACnetErrorVTOpen) Serialize(io spi.WriteBuffer) {
	ser := func() {

	}
	BACnetErrorSerialize(io, m.BACnetError, CastIBACnetError(m), ser)
}
