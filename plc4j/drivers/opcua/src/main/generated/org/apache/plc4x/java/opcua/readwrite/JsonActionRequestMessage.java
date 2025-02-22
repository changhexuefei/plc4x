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

public class JsonActionRequestMessage extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 19323;
  }

  // Properties.
  protected final int dataSetWriterId;
  protected final int actionTargetId;
  protected final PascalString dataSetWriterName;
  protected final PascalString writerGroupName;
  protected final ConfigurationVersionDataType metaDataVersion;
  protected final long minorVersion;
  protected final long timestamp;
  protected final PascalString messageType;
  protected final int requestId;
  protected final ActionState actionState;
  protected final ExtensionObject payload;

  public JsonActionRequestMessage(
      int dataSetWriterId,
      int actionTargetId,
      PascalString dataSetWriterName,
      PascalString writerGroupName,
      ConfigurationVersionDataType metaDataVersion,
      long minorVersion,
      long timestamp,
      PascalString messageType,
      int requestId,
      ActionState actionState,
      ExtensionObject payload) {
    super();
    this.dataSetWriterId = dataSetWriterId;
    this.actionTargetId = actionTargetId;
    this.dataSetWriterName = dataSetWriterName;
    this.writerGroupName = writerGroupName;
    this.metaDataVersion = metaDataVersion;
    this.minorVersion = minorVersion;
    this.timestamp = timestamp;
    this.messageType = messageType;
    this.requestId = requestId;
    this.actionState = actionState;
    this.payload = payload;
  }

  public int getDataSetWriterId() {
    return dataSetWriterId;
  }

  public int getActionTargetId() {
    return actionTargetId;
  }

  public PascalString getDataSetWriterName() {
    return dataSetWriterName;
  }

  public PascalString getWriterGroupName() {
    return writerGroupName;
  }

  public ConfigurationVersionDataType getMetaDataVersion() {
    return metaDataVersion;
  }

  public long getMinorVersion() {
    return minorVersion;
  }

  public long getTimestamp() {
    return timestamp;
  }

  public PascalString getMessageType() {
    return messageType;
  }

  public int getRequestId() {
    return requestId;
  }

  public ActionState getActionState() {
    return actionState;
  }

  public ExtensionObject getPayload() {
    return payload;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("JsonActionRequestMessage");

    // Simple Field (dataSetWriterId)
    writeSimpleField("dataSetWriterId", dataSetWriterId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (actionTargetId)
    writeSimpleField("actionTargetId", actionTargetId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (dataSetWriterName)
    writeSimpleField("dataSetWriterName", dataSetWriterName, writeComplex(writeBuffer));

    // Simple Field (writerGroupName)
    writeSimpleField("writerGroupName", writerGroupName, writeComplex(writeBuffer));

    // Simple Field (metaDataVersion)
    writeSimpleField("metaDataVersion", metaDataVersion, writeComplex(writeBuffer));

    // Simple Field (minorVersion)
    writeSimpleField("minorVersion", minorVersion, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (timestamp)
    writeSimpleField("timestamp", timestamp, writeSignedLong(writeBuffer, 64));

    // Simple Field (messageType)
    writeSimpleField("messageType", messageType, writeComplex(writeBuffer));

    // Simple Field (requestId)
    writeSimpleField("requestId", requestId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (actionState)
    writeSimpleEnumField(
        "actionState",
        "ActionState",
        actionState,
        writeEnum(ActionState::getValue, ActionState::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (payload)
    writeSimpleField("payload", payload, writeComplex(writeBuffer));

    writeBuffer.popContext("JsonActionRequestMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    JsonActionRequestMessage _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dataSetWriterId)
    lengthInBits += 16;

    // Simple field (actionTargetId)
    lengthInBits += 16;

    // Simple field (dataSetWriterName)
    lengthInBits += dataSetWriterName.getLengthInBits();

    // Simple field (writerGroupName)
    lengthInBits += writerGroupName.getLengthInBits();

    // Simple field (metaDataVersion)
    lengthInBits += metaDataVersion.getLengthInBits();

    // Simple field (minorVersion)
    lengthInBits += 32;

    // Simple field (timestamp)
    lengthInBits += 64;

    // Simple field (messageType)
    lengthInBits += messageType.getLengthInBits();

    // Simple field (requestId)
    lengthInBits += 16;

    // Simple field (actionState)
    lengthInBits += 32;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("JsonActionRequestMessage");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int dataSetWriterId = readSimpleField("dataSetWriterId", readUnsignedInt(readBuffer, 16));

    int actionTargetId = readSimpleField("actionTargetId", readUnsignedInt(readBuffer, 16));

    PascalString dataSetWriterName =
        readSimpleField(
            "dataSetWriterName",
            readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString writerGroupName =
        readSimpleField(
            "writerGroupName", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    ConfigurationVersionDataType metaDataVersion =
        readSimpleField(
            "metaDataVersion",
            readComplex(
                () ->
                    (ConfigurationVersionDataType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (14595)),
                readBuffer));

    long minorVersion = readSimpleField("minorVersion", readUnsignedLong(readBuffer, 32));

    long timestamp = readSimpleField("timestamp", readSignedLong(readBuffer, 64));

    PascalString messageType =
        readSimpleField(
            "messageType", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    int requestId = readSimpleField("requestId", readUnsignedInt(readBuffer, 16));

    ActionState actionState =
        readEnumField(
            "actionState",
            "ActionState",
            readEnum(ActionState::enumForValue, readUnsignedLong(readBuffer, 32)));

    ExtensionObject payload =
        readSimpleField(
            "payload",
            readComplex(
                () -> ExtensionObject.staticParse(readBuffer, (boolean) (true)), readBuffer));

    readBuffer.closeContext("JsonActionRequestMessage");
    // Create the instance
    return new JsonActionRequestMessageBuilderImpl(
        dataSetWriterId,
        actionTargetId,
        dataSetWriterName,
        writerGroupName,
        metaDataVersion,
        minorVersion,
        timestamp,
        messageType,
        requestId,
        actionState,
        payload);
  }

  public static class JsonActionRequestMessageBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int dataSetWriterId;
    private final int actionTargetId;
    private final PascalString dataSetWriterName;
    private final PascalString writerGroupName;
    private final ConfigurationVersionDataType metaDataVersion;
    private final long minorVersion;
    private final long timestamp;
    private final PascalString messageType;
    private final int requestId;
    private final ActionState actionState;
    private final ExtensionObject payload;

    public JsonActionRequestMessageBuilderImpl(
        int dataSetWriterId,
        int actionTargetId,
        PascalString dataSetWriterName,
        PascalString writerGroupName,
        ConfigurationVersionDataType metaDataVersion,
        long minorVersion,
        long timestamp,
        PascalString messageType,
        int requestId,
        ActionState actionState,
        ExtensionObject payload) {
      this.dataSetWriterId = dataSetWriterId;
      this.actionTargetId = actionTargetId;
      this.dataSetWriterName = dataSetWriterName;
      this.writerGroupName = writerGroupName;
      this.metaDataVersion = metaDataVersion;
      this.minorVersion = minorVersion;
      this.timestamp = timestamp;
      this.messageType = messageType;
      this.requestId = requestId;
      this.actionState = actionState;
      this.payload = payload;
    }

    public JsonActionRequestMessage build() {
      JsonActionRequestMessage jsonActionRequestMessage =
          new JsonActionRequestMessage(
              dataSetWriterId,
              actionTargetId,
              dataSetWriterName,
              writerGroupName,
              metaDataVersion,
              minorVersion,
              timestamp,
              messageType,
              requestId,
              actionState,
              payload);
      return jsonActionRequestMessage;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof JsonActionRequestMessage)) {
      return false;
    }
    JsonActionRequestMessage that = (JsonActionRequestMessage) o;
    return (getDataSetWriterId() == that.getDataSetWriterId())
        && (getActionTargetId() == that.getActionTargetId())
        && (getDataSetWriterName() == that.getDataSetWriterName())
        && (getWriterGroupName() == that.getWriterGroupName())
        && (getMetaDataVersion() == that.getMetaDataVersion())
        && (getMinorVersion() == that.getMinorVersion())
        && (getTimestamp() == that.getTimestamp())
        && (getMessageType() == that.getMessageType())
        && (getRequestId() == that.getRequestId())
        && (getActionState() == that.getActionState())
        && (getPayload() == that.getPayload())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getDataSetWriterId(),
        getActionTargetId(),
        getDataSetWriterName(),
        getWriterGroupName(),
        getMetaDataVersion(),
        getMinorVersion(),
        getTimestamp(),
        getMessageType(),
        getRequestId(),
        getActionState(),
        getPayload());
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
