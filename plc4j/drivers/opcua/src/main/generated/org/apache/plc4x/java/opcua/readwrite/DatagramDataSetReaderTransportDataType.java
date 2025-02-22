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

public class DatagramDataSetReaderTransportDataType extends ExtensionObjectDefinition
    implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 23616;
  }

  // Properties.
  protected final ExtensionObject address;
  protected final PascalString qosCategory;
  protected final List<ExtensionObject> datagramQos;
  protected final PascalString topic;

  public DatagramDataSetReaderTransportDataType(
      ExtensionObject address,
      PascalString qosCategory,
      List<ExtensionObject> datagramQos,
      PascalString topic) {
    super();
    this.address = address;
    this.qosCategory = qosCategory;
    this.datagramQos = datagramQos;
    this.topic = topic;
  }

  public ExtensionObject getAddress() {
    return address;
  }

  public PascalString getQosCategory() {
    return qosCategory;
  }

  public List<ExtensionObject> getDatagramQos() {
    return datagramQos;
  }

  public PascalString getTopic() {
    return topic;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("DatagramDataSetReaderTransportDataType");

    // Simple Field (address)
    writeSimpleField("address", address, writeComplex(writeBuffer));

    // Simple Field (qosCategory)
    writeSimpleField("qosCategory", qosCategory, writeComplex(writeBuffer));

    // Implicit Field (noOfDatagramQos) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfDatagramQos = (int) ((((getDatagramQos()) == (null)) ? -(1) : COUNT(getDatagramQos())));
    writeImplicitField("noOfDatagramQos", noOfDatagramQos, writeSignedInt(writeBuffer, 32));

    // Array Field (datagramQos)
    writeComplexTypeArrayField("datagramQos", datagramQos, writeBuffer);

    // Simple Field (topic)
    writeSimpleField("topic", topic, writeComplex(writeBuffer));

    writeBuffer.popContext("DatagramDataSetReaderTransportDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DatagramDataSetReaderTransportDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (address)
    lengthInBits += address.getLengthInBits();

    // Simple field (qosCategory)
    lengthInBits += qosCategory.getLengthInBits();

    // Implicit Field (noOfDatagramQos)
    lengthInBits += 32;

    // Array field
    if (datagramQos != null) {
      int i = 0;
      for (ExtensionObject element : datagramQos) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= datagramQos.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (topic)
    lengthInBits += topic.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("DatagramDataSetReaderTransportDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObject address =
        readSimpleField(
            "address",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    PascalString qosCategory =
        readSimpleField(
            "qosCategory", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfDatagramQos = readImplicitField("noOfDatagramQos", readSignedInt(readBuffer, 32));

    List<ExtensionObject> datagramQos =
        readCountArrayField(
            "datagramQos",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer),
            noOfDatagramQos);

    PascalString topic =
        readSimpleField(
            "topic", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("DatagramDataSetReaderTransportDataType");
    // Create the instance
    return new DatagramDataSetReaderTransportDataTypeBuilderImpl(
        address, qosCategory, datagramQos, topic);
  }

  public static class DatagramDataSetReaderTransportDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObject address;
    private final PascalString qosCategory;
    private final List<ExtensionObject> datagramQos;
    private final PascalString topic;

    public DatagramDataSetReaderTransportDataTypeBuilderImpl(
        ExtensionObject address,
        PascalString qosCategory,
        List<ExtensionObject> datagramQos,
        PascalString topic) {
      this.address = address;
      this.qosCategory = qosCategory;
      this.datagramQos = datagramQos;
      this.topic = topic;
    }

    public DatagramDataSetReaderTransportDataType build() {
      DatagramDataSetReaderTransportDataType datagramDataSetReaderTransportDataType =
          new DatagramDataSetReaderTransportDataType(address, qosCategory, datagramQos, topic);
      return datagramDataSetReaderTransportDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DatagramDataSetReaderTransportDataType)) {
      return false;
    }
    DatagramDataSetReaderTransportDataType that = (DatagramDataSetReaderTransportDataType) o;
    return (getAddress() == that.getAddress())
        && (getQosCategory() == that.getQosCategory())
        && (getDatagramQos() == that.getDatagramQos())
        && (getTopic() == that.getTopic())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getAddress(), getQosCategory(), getDatagramQos(), getTopic());
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
