/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.profinet.types;

/**
 * (Values determined by evaluating generated ".pcapng" files)
 */
public enum MemoryArea {
    SYSTEM_INFO_OF_S200_FAMILY((byte) 0x03), /* Renamed from "System info of 200 family" */
    SYSTEM_FLAGS_OF_S200_FAMILY((byte) 0x05), /* Renamed from "System flags of 200 family" */
    ANALOG_INPUTS_OF_S200_FAMILY((byte) 0x06), /* Renamed from "System inputs of 200 family" */
    ANALOG_OUTPUTS_OF_S200_FAMILY((byte) 0x07), /* Renamed from "System outputs of 200 family" */
    S7_COUNTERS((byte) 0x1C),
    S7_TIMERS((byte) 0x1D),
    IEC_COUNTERS_OF_S200_FAMILY((byte) 0x1E), /* Renamed from "IEC counters (200 family)" */
    IEC_TIMERS_OF_S200_FAMILY((byte) 0x1F), /* Renamed from "IEC timers (200 family)" */
    DIRECT_PERIPHERAL_ACCESS((byte) 0x80),
    INPUTS((byte) 0x81),
    OUTPUTS((byte) 0x82),
    FLAGS((byte) 0x83),
    DATA_BLOCKS((byte) 0x84),
    INSTANCE_DATA_BLOCKS((byte) 0x85),
    LOCAL_DATA((byte) 0x86);

    private byte code;

    MemoryArea(byte code) {
        this.code = code;
    }

    public byte getCode() {
        return code;
    }

}