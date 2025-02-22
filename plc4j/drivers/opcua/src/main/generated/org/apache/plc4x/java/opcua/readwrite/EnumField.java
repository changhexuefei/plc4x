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

public class EnumField extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 104;
  }

  // Properties.
  protected final long value;
  protected final LocalizedText displayName;
  protected final LocalizedText description;
  protected final PascalString name;

  public EnumField(
      long value, LocalizedText displayName, LocalizedText description, PascalString name) {
    super();
    this.value = value;
    this.displayName = displayName;
    this.description = description;
    this.name = name;
  }

  public long getValue() {
    return value;
  }

  public LocalizedText getDisplayName() {
    return displayName;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public PascalString getName() {
    return name;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("EnumField");

    // Simple Field (value)
    writeSimpleField("value", value, writeSignedLong(writeBuffer, 64));

    // Simple Field (displayName)
    writeSimpleField("displayName", displayName, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    writeBuffer.popContext("EnumField");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EnumField _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (value)
    lengthInBits += 64;

    // Simple field (displayName)
    lengthInBits += displayName.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("EnumField");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long value = readSimpleField("value", readSignedLong(readBuffer, 64));

    LocalizedText displayName =
        readSimpleField(
            "displayName", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    PascalString name =
        readSimpleField(
            "name", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("EnumField");
    // Create the instance
    return new EnumFieldBuilderImpl(value, displayName, description, name);
  }

  public static class EnumFieldBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final long value;
    private final LocalizedText displayName;
    private final LocalizedText description;
    private final PascalString name;

    public EnumFieldBuilderImpl(
        long value, LocalizedText displayName, LocalizedText description, PascalString name) {
      this.value = value;
      this.displayName = displayName;
      this.description = description;
      this.name = name;
    }

    public EnumField build() {
      EnumField enumField = new EnumField(value, displayName, description, name);
      return enumField;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EnumField)) {
      return false;
    }
    EnumField that = (EnumField) o;
    return (getValue() == that.getValue())
        && (getDisplayName() == that.getDisplayName())
        && (getDescription() == that.getDescription())
        && (getName() == that.getName())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getValue(), getDisplayName(), getDescription(), getName());
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
