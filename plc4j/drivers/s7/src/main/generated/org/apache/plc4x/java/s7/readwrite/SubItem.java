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
package org.apache.plc4x.java.s7.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class SubItem implements Message {

  // Properties.
  protected final short bytesToRead;
  protected final int dbNumber;
  protected final int startAddress;

  public SubItem(short bytesToRead, int dbNumber, int startAddress) {
    super();
    this.bytesToRead = bytesToRead;
    this.dbNumber = dbNumber;
    this.startAddress = startAddress;
  }

  public short getBytesToRead() {
    return bytesToRead;
  }

  public int getDbNumber() {
    return dbNumber;
  }

  public int getStartAddress() {
    return startAddress;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SubItem");

    // Simple Field (bytesToRead)
    writeSimpleField("bytesToRead", bytesToRead, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (dbNumber)
    writeSimpleField("dbNumber", dbNumber, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (startAddress)
    writeSimpleField("startAddress", startAddress, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("SubItem");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    SubItem _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (bytesToRead)
    lengthInBits += 8;

    // Simple field (dbNumber)
    lengthInBits += 16;

    // Simple field (startAddress)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static SubItem staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("SubItem");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short bytesToRead = readSimpleField("bytesToRead", readUnsignedShort(readBuffer, 8));

    int dbNumber = readSimpleField("dbNumber", readUnsignedInt(readBuffer, 16));

    int startAddress = readSimpleField("startAddress", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("SubItem");
    // Create the instance
    SubItem _subItem;
    _subItem = new SubItem(bytesToRead, dbNumber, startAddress);
    return _subItem;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SubItem)) {
      return false;
    }
    SubItem that = (SubItem) o;
    return (getBytesToRead() == that.getBytesToRead())
        && (getDbNumber() == that.getDbNumber())
        && (getStartAddress() == that.getStartAddress())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getBytesToRead(), getDbNumber(), getStartAddress());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
