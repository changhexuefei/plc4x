/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum LldpSystemCapabilitiesMap {
  lldpSystemCapabilitiesMapNone((long) 0L),
  lldpSystemCapabilitiesMapOther((long) 1L),
  lldpSystemCapabilitiesMapRepeater((long) 2L),
  lldpSystemCapabilitiesMapBridge((long) 4L),
  lldpSystemCapabilitiesMapWlanAccessPoint((long) 8L),
  lldpSystemCapabilitiesMapRouter((long) 16L),
  lldpSystemCapabilitiesMapTelephone((long) 32L),
  lldpSystemCapabilitiesMapDocsisCableDevice((long) 64L),
  lldpSystemCapabilitiesMapStationOnly((long) 128L),
  lldpSystemCapabilitiesMapCvlanComponent((long) 256L),
  lldpSystemCapabilitiesMapSvlanComponent((long) 512L),
  lldpSystemCapabilitiesMapTwoPortMacRelay((long) 1024L);
  private static final Map<Long, LldpSystemCapabilitiesMap> map;

  static {
    map = new HashMap<>();
    for (LldpSystemCapabilitiesMap value : LldpSystemCapabilitiesMap.values()) {
      map.put((long) value.getValue(), value);
    }
  }

  private final long value;

  LldpSystemCapabilitiesMap(long value) {
    this.value = value;
  }

  public long getValue() {
    return value;
  }

  public static LldpSystemCapabilitiesMap enumForValue(long value) {
    return map.get(value);
  }

  public static Boolean isDefined(long value) {
    return map.containsKey(value);
  }
}
