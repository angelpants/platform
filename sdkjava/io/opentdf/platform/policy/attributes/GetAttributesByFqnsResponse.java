// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: policy/attributes/attributes.proto

// Protobuf Java Version: 3.25.3
package io.opentdf.platform.policy.attributes;

/**
 * Protobuf type {@code policy.attributes.GetAttributesByFqnsResponse}
 */
public final class GetAttributesByFqnsResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:policy.attributes.GetAttributesByFqnsResponse)
    GetAttributesByFqnsResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetAttributesByFqnsResponse.newBuilder() to construct.
  private GetAttributesByFqnsResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetAttributesByFqnsResponse() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new GetAttributesByFqnsResponse();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.opentdf.platform.policy.attributes.AttributesProto.internal_static_policy_attributes_GetAttributesByFqnsResponse_descriptor;
  }

  @SuppressWarnings({"rawtypes"})
  @java.lang.Override
  protected com.google.protobuf.MapFieldReflectionAccessor internalGetMapFieldReflection(
      int number) {
    switch (number) {
      case 1:
        return internalGetAttributes();
      default:
        throw new RuntimeException(
            "Invalid map field number: " + number);
    }
  }
  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.opentdf.platform.policy.attributes.AttributesProto.internal_static_policy_attributes_GetAttributesByFqnsResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse.class, io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse.Builder.class);
  }

  public static final int ATTRIBUTES_FIELD_NUMBER = 1;
  private static final class AttributesDefaultEntryHolder {
    static final com.google.protobuf.MapEntry<
        java.lang.String, io.opentdf.platform.policy.attributes.Attribute> defaultEntry =
            com.google.protobuf.MapEntry
            .<java.lang.String, io.opentdf.platform.policy.attributes.Attribute>newDefaultInstance(
                io.opentdf.platform.policy.attributes.AttributesProto.internal_static_policy_attributes_GetAttributesByFqnsResponse_AttributesEntry_descriptor, 
                com.google.protobuf.WireFormat.FieldType.STRING,
                "",
                com.google.protobuf.WireFormat.FieldType.MESSAGE,
                io.opentdf.platform.policy.attributes.Attribute.getDefaultInstance());
  }
  @SuppressWarnings("serial")
  private com.google.protobuf.MapField<
      java.lang.String, io.opentdf.platform.policy.attributes.Attribute> attributes_;
  private com.google.protobuf.MapField<java.lang.String, io.opentdf.platform.policy.attributes.Attribute>
  internalGetAttributes() {
    if (attributes_ == null) {
      return com.google.protobuf.MapField.emptyMapField(
          AttributesDefaultEntryHolder.defaultEntry);
    }
    return attributes_;
  }
  public int getAttributesCount() {
    return internalGetAttributes().getMap().size();
  }
  /**
   * <pre>
   * map of fqns as keys to attributes as values
   * </pre>
   *
   * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
   */
  @java.lang.Override
  public boolean containsAttributes(
      java.lang.String key) {
    if (key == null) { throw new NullPointerException("map key"); }
    return internalGetAttributes().getMap().containsKey(key);
  }
  /**
   * Use {@link #getAttributesMap()} instead.
   */
  @java.lang.Override
  @java.lang.Deprecated
  public java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> getAttributes() {
    return getAttributesMap();
  }
  /**
   * <pre>
   * map of fqns as keys to attributes as values
   * </pre>
   *
   * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
   */
  @java.lang.Override
  public java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> getAttributesMap() {
    return internalGetAttributes().getMap();
  }
  /**
   * <pre>
   * map of fqns as keys to attributes as values
   * </pre>
   *
   * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
   */
  @java.lang.Override
  public /* nullable */
io.opentdf.platform.policy.attributes.Attribute getAttributesOrDefault(
      java.lang.String key,
      /* nullable */
io.opentdf.platform.policy.attributes.Attribute defaultValue) {
    if (key == null) { throw new NullPointerException("map key"); }
    java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> map =
        internalGetAttributes().getMap();
    return map.containsKey(key) ? map.get(key) : defaultValue;
  }
  /**
   * <pre>
   * map of fqns as keys to attributes as values
   * </pre>
   *
   * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
   */
  @java.lang.Override
  public io.opentdf.platform.policy.attributes.Attribute getAttributesOrThrow(
      java.lang.String key) {
    if (key == null) { throw new NullPointerException("map key"); }
    java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> map =
        internalGetAttributes().getMap();
    if (!map.containsKey(key)) {
      throw new java.lang.IllegalArgumentException();
    }
    return map.get(key);
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    com.google.protobuf.GeneratedMessageV3
      .serializeStringMapTo(
        output,
        internalGetAttributes(),
        AttributesDefaultEntryHolder.defaultEntry,
        1);
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    for (java.util.Map.Entry<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> entry
         : internalGetAttributes().getMap().entrySet()) {
      com.google.protobuf.MapEntry<java.lang.String, io.opentdf.platform.policy.attributes.Attribute>
      attributes__ = AttributesDefaultEntryHolder.defaultEntry.newBuilderForType()
          .setKey(entry.getKey())
          .setValue(entry.getValue())
          .build();
      size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, attributes__);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse)) {
      return super.equals(obj);
    }
    io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse other = (io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse) obj;

    if (!internalGetAttributes().equals(
        other.internalGetAttributes())) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (!internalGetAttributes().getMap().isEmpty()) {
      hash = (37 * hash) + ATTRIBUTES_FIELD_NUMBER;
      hash = (53 * hash) + internalGetAttributes().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code policy.attributes.GetAttributesByFqnsResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:policy.attributes.GetAttributesByFqnsResponse)
      io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.opentdf.platform.policy.attributes.AttributesProto.internal_static_policy_attributes_GetAttributesByFqnsResponse_descriptor;
    }

    @SuppressWarnings({"rawtypes"})
    protected com.google.protobuf.MapFieldReflectionAccessor internalGetMapFieldReflection(
        int number) {
      switch (number) {
        case 1:
          return internalGetAttributes();
        default:
          throw new RuntimeException(
              "Invalid map field number: " + number);
      }
    }
    @SuppressWarnings({"rawtypes"})
    protected com.google.protobuf.MapFieldReflectionAccessor internalGetMutableMapFieldReflection(
        int number) {
      switch (number) {
        case 1:
          return internalGetMutableAttributes();
        default:
          throw new RuntimeException(
              "Invalid map field number: " + number);
      }
    }
    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.opentdf.platform.policy.attributes.AttributesProto.internal_static_policy_attributes_GetAttributesByFqnsResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse.class, io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse.Builder.class);
    }

    // Construct using io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      internalGetMutableAttributes().clear();
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.opentdf.platform.policy.attributes.AttributesProto.internal_static_policy_attributes_GetAttributesByFqnsResponse_descriptor;
    }

    @java.lang.Override
    public io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse getDefaultInstanceForType() {
      return io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse.getDefaultInstance();
    }

    @java.lang.Override
    public io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse build() {
      io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse buildPartial() {
      io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse result = new io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.attributes_ = internalGetAttributes().build(AttributesDefaultEntryHolder.defaultEntry);
      }
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse) {
        return mergeFrom((io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse other) {
      if (other == io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse.getDefaultInstance()) return this;
      internalGetMutableAttributes().mergeFrom(
          other.internalGetAttributes());
      bitField0_ |= 0x00000001;
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              com.google.protobuf.MapEntry<java.lang.String, io.opentdf.platform.policy.attributes.Attribute>
              attributes__ = input.readMessage(
                  AttributesDefaultEntryHolder.defaultEntry.getParserForType(), extensionRegistry);
              internalGetMutableAttributes().ensureBuilderMap().put(
                  attributes__.getKey(), attributes__.getValue());
              bitField0_ |= 0x00000001;
              break;
            } // case 10
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private static final class AttributesConverter implements com.google.protobuf.MapFieldBuilder.Converter<java.lang.String, io.opentdf.platform.policy.attributes.AttributeOrBuilder, io.opentdf.platform.policy.attributes.Attribute> {
      @java.lang.Override
      public io.opentdf.platform.policy.attributes.Attribute build(io.opentdf.platform.policy.attributes.AttributeOrBuilder val) {
        if (val instanceof io.opentdf.platform.policy.attributes.Attribute) { return (io.opentdf.platform.policy.attributes.Attribute) val; }
        return ((io.opentdf.platform.policy.attributes.Attribute.Builder) val).build();
      }

      @java.lang.Override
      public com.google.protobuf.MapEntry<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> defaultEntry() {
        return AttributesDefaultEntryHolder.defaultEntry;
      }
    };
    private static final AttributesConverter attributesConverter = new AttributesConverter();

    private com.google.protobuf.MapFieldBuilder<
        java.lang.String, io.opentdf.platform.policy.attributes.AttributeOrBuilder, io.opentdf.platform.policy.attributes.Attribute, io.opentdf.platform.policy.attributes.Attribute.Builder> attributes_;
    private com.google.protobuf.MapFieldBuilder<java.lang.String, io.opentdf.platform.policy.attributes.AttributeOrBuilder, io.opentdf.platform.policy.attributes.Attribute, io.opentdf.platform.policy.attributes.Attribute.Builder>
        internalGetAttributes() {
      if (attributes_ == null) {
        return new com.google.protobuf.MapFieldBuilder<>(attributesConverter);
      }
      return attributes_;
    }
    private com.google.protobuf.MapFieldBuilder<java.lang.String, io.opentdf.platform.policy.attributes.AttributeOrBuilder, io.opentdf.platform.policy.attributes.Attribute, io.opentdf.platform.policy.attributes.Attribute.Builder>
        internalGetMutableAttributes() {
      if (attributes_ == null) {
        attributes_ = new com.google.protobuf.MapFieldBuilder<>(attributesConverter);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return attributes_;
    }
    public int getAttributesCount() {
      return internalGetAttributes().ensureBuilderMap().size();
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    @java.lang.Override
    public boolean containsAttributes(
        java.lang.String key) {
      if (key == null) { throw new NullPointerException("map key"); }
      return internalGetAttributes().ensureBuilderMap().containsKey(key);
    }
    /**
     * Use {@link #getAttributesMap()} instead.
     */
    @java.lang.Override
    @java.lang.Deprecated
    public java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> getAttributes() {
      return getAttributesMap();
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    @java.lang.Override
    public java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> getAttributesMap() {
      return internalGetAttributes().getImmutableMap();
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    @java.lang.Override
    public /* nullable */
io.opentdf.platform.policy.attributes.Attribute getAttributesOrDefault(
        java.lang.String key,
        /* nullable */
io.opentdf.platform.policy.attributes.Attribute defaultValue) {
      if (key == null) { throw new NullPointerException("map key"); }
      java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.AttributeOrBuilder> map = internalGetMutableAttributes().ensureBuilderMap();
      return map.containsKey(key) ? attributesConverter.build(map.get(key)) : defaultValue;
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    @java.lang.Override
    public io.opentdf.platform.policy.attributes.Attribute getAttributesOrThrow(
        java.lang.String key) {
      if (key == null) { throw new NullPointerException("map key"); }
      java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.AttributeOrBuilder> map = internalGetMutableAttributes().ensureBuilderMap();
      if (!map.containsKey(key)) {
        throw new java.lang.IllegalArgumentException();
      }
      return attributesConverter.build(map.get(key));
    }
    public Builder clearAttributes() {
      bitField0_ = (bitField0_ & ~0x00000001);
      internalGetMutableAttributes().clear();
      return this;
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    public Builder removeAttributes(
        java.lang.String key) {
      if (key == null) { throw new NullPointerException("map key"); }
      internalGetMutableAttributes().ensureBuilderMap()
          .remove(key);
      return this;
    }
    /**
     * Use alternate mutation accessors instead.
     */
    @java.lang.Deprecated
    public java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute>
        getMutableAttributes() {
      bitField0_ |= 0x00000001;
      return internalGetMutableAttributes().ensureMessageMap();
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    public Builder putAttributes(
        java.lang.String key,
        io.opentdf.platform.policy.attributes.Attribute value) {
      if (key == null) { throw new NullPointerException("map key"); }
      if (value == null) { throw new NullPointerException("map value"); }
      internalGetMutableAttributes().ensureBuilderMap()
          .put(key, value);
      bitField0_ |= 0x00000001;
      return this;
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    public Builder putAllAttributes(
        java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> values) {
      for (java.util.Map.Entry<java.lang.String, io.opentdf.platform.policy.attributes.Attribute> e : values.entrySet()) {
        if (e.getKey() == null || e.getValue() == null) {
          throw new NullPointerException();
        }
      }
      internalGetMutableAttributes().ensureBuilderMap()
          .putAll(values);
      bitField0_ |= 0x00000001;
      return this;
    }
    /**
     * <pre>
     * map of fqns as keys to attributes as values
     * </pre>
     *
     * <code>map&lt;string, .policy.attributes.Attribute&gt; attributes = 1 [json_name = "attributes"];</code>
     */
    public io.opentdf.platform.policy.attributes.Attribute.Builder putAttributesBuilderIfAbsent(
        java.lang.String key) {
      java.util.Map<java.lang.String, io.opentdf.platform.policy.attributes.AttributeOrBuilder> builderMap = internalGetMutableAttributes().ensureBuilderMap();
      io.opentdf.platform.policy.attributes.AttributeOrBuilder entry = builderMap.get(key);
      if (entry == null) {
        entry = io.opentdf.platform.policy.attributes.Attribute.newBuilder();
        builderMap.put(key, entry);
      }
      if (entry instanceof io.opentdf.platform.policy.attributes.Attribute) {
        entry = ((io.opentdf.platform.policy.attributes.Attribute) entry).toBuilder();
        builderMap.put(key, entry);
      }
      return (io.opentdf.platform.policy.attributes.Attribute.Builder) entry;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:policy.attributes.GetAttributesByFqnsResponse)
  }

  // @@protoc_insertion_point(class_scope:policy.attributes.GetAttributesByFqnsResponse)
  private static final io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse();
  }

  public static io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetAttributesByFqnsResponse>
      PARSER = new com.google.protobuf.AbstractParser<GetAttributesByFqnsResponse>() {
    @java.lang.Override
    public GetAttributesByFqnsResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<GetAttributesByFqnsResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetAttributesByFqnsResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.opentdf.platform.policy.attributes.GetAttributesByFqnsResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

