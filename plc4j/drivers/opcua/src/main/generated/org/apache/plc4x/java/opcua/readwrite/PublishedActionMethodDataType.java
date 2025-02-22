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

public class PublishedActionMethodDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 18795;
  }

  // Properties.
  protected final DataSetMetaDataType requestDataSetMetaData;
  protected final List<ActionTargetDataType> actionTargets;
  protected final List<ActionMethodDataType> actionMethods;

  public PublishedActionMethodDataType(
      DataSetMetaDataType requestDataSetMetaData,
      List<ActionTargetDataType> actionTargets,
      List<ActionMethodDataType> actionMethods) {
    super();
    this.requestDataSetMetaData = requestDataSetMetaData;
    this.actionTargets = actionTargets;
    this.actionMethods = actionMethods;
  }

  public DataSetMetaDataType getRequestDataSetMetaData() {
    return requestDataSetMetaData;
  }

  public List<ActionTargetDataType> getActionTargets() {
    return actionTargets;
  }

  public List<ActionMethodDataType> getActionMethods() {
    return actionMethods;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PublishedActionMethodDataType");

    // Simple Field (requestDataSetMetaData)
    writeSimpleField("requestDataSetMetaData", requestDataSetMetaData, writeComplex(writeBuffer));

    // Implicit Field (noOfActionTargets) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfActionTargets =
        (int) ((((getActionTargets()) == (null)) ? -(1) : COUNT(getActionTargets())));
    writeImplicitField("noOfActionTargets", noOfActionTargets, writeSignedInt(writeBuffer, 32));

    // Array Field (actionTargets)
    writeComplexTypeArrayField("actionTargets", actionTargets, writeBuffer);

    // Implicit Field (noOfActionMethods) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfActionMethods =
        (int) ((((getActionMethods()) == (null)) ? -(1) : COUNT(getActionMethods())));
    writeImplicitField("noOfActionMethods", noOfActionMethods, writeSignedInt(writeBuffer, 32));

    // Array Field (actionMethods)
    writeComplexTypeArrayField("actionMethods", actionMethods, writeBuffer);

    writeBuffer.popContext("PublishedActionMethodDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PublishedActionMethodDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestDataSetMetaData)
    lengthInBits += requestDataSetMetaData.getLengthInBits();

    // Implicit Field (noOfActionTargets)
    lengthInBits += 32;

    // Array field
    if (actionTargets != null) {
      int i = 0;
      for (ActionTargetDataType element : actionTargets) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= actionTargets.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Implicit Field (noOfActionMethods)
    lengthInBits += 32;

    // Array field
    if (actionMethods != null) {
      int i = 0;
      for (ActionMethodDataType element : actionMethods) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= actionMethods.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("PublishedActionMethodDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    DataSetMetaDataType requestDataSetMetaData =
        readSimpleField(
            "requestDataSetMetaData",
            readComplex(
                () ->
                    (DataSetMetaDataType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (14525)),
                readBuffer));

    int noOfActionTargets = readImplicitField("noOfActionTargets", readSignedInt(readBuffer, 32));

    List<ActionTargetDataType> actionTargets =
        readCountArrayField(
            "actionTargets",
            readComplex(
                () ->
                    (ActionTargetDataType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (18595)),
                readBuffer),
            noOfActionTargets);

    int noOfActionMethods = readImplicitField("noOfActionMethods", readSignedInt(readBuffer, 32));

    List<ActionMethodDataType> actionMethods =
        readCountArrayField(
            "actionMethods",
            readComplex(
                () ->
                    (ActionMethodDataType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (18599)),
                readBuffer),
            noOfActionMethods);

    readBuffer.closeContext("PublishedActionMethodDataType");
    // Create the instance
    return new PublishedActionMethodDataTypeBuilderImpl(
        requestDataSetMetaData, actionTargets, actionMethods);
  }

  public static class PublishedActionMethodDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final DataSetMetaDataType requestDataSetMetaData;
    private final List<ActionTargetDataType> actionTargets;
    private final List<ActionMethodDataType> actionMethods;

    public PublishedActionMethodDataTypeBuilderImpl(
        DataSetMetaDataType requestDataSetMetaData,
        List<ActionTargetDataType> actionTargets,
        List<ActionMethodDataType> actionMethods) {
      this.requestDataSetMetaData = requestDataSetMetaData;
      this.actionTargets = actionTargets;
      this.actionMethods = actionMethods;
    }

    public PublishedActionMethodDataType build() {
      PublishedActionMethodDataType publishedActionMethodDataType =
          new PublishedActionMethodDataType(requestDataSetMetaData, actionTargets, actionMethods);
      return publishedActionMethodDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PublishedActionMethodDataType)) {
      return false;
    }
    PublishedActionMethodDataType that = (PublishedActionMethodDataType) o;
    return (getRequestDataSetMetaData() == that.getRequestDataSetMetaData())
        && (getActionTargets() == that.getActionTargets())
        && (getActionMethods() == that.getActionMethods())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getRequestDataSetMetaData(), getActionTargets(), getActionMethods());
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
