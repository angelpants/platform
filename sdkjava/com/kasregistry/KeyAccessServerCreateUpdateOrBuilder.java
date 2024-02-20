// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kasregistry/key_access_server_registry.proto

// Protobuf Java Version: 3.25.3
package com.kasregistry;

public interface KeyAccessServerCreateUpdateOrBuilder extends
    // @@protoc_insertion_point(interface_extends:kasregistry.KeyAccessServerCreateUpdate)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Optional metadata for the attribute definition
   * </pre>
   *
   * <code>.common.MetadataMutable metadata = 1 [json_name = "metadata"];</code>
   * @return Whether the metadata field is set.
   */
  boolean hasMetadata();
  /**
   * <pre>
   * Optional metadata for the attribute definition
   * </pre>
   *
   * <code>.common.MetadataMutable metadata = 1 [json_name = "metadata"];</code>
   * @return The metadata.
   */
  com.common.MetadataMutable getMetadata();
  /**
   * <pre>
   * Optional metadata for the attribute definition
   * </pre>
   *
   * <code>.common.MetadataMutable metadata = 1 [json_name = "metadata"];</code>
   */
  com.common.MetadataMutableOrBuilder getMetadataOrBuilder();

  /**
   * <pre>
   * Address of a KAS instance
   * </pre>
   *
   * <code>string uri = 2 [json_name = "uri", (.buf.validate.field) = { ... }</code>
   * @return The uri.
   */
  java.lang.String getUri();
  /**
   * <pre>
   * Address of a KAS instance
   * </pre>
   *
   * <code>string uri = 2 [json_name = "uri", (.buf.validate.field) = { ... }</code>
   * @return The bytes for uri.
   */
  com.google.protobuf.ByteString
      getUriBytes();

  /**
   * <code>.kasregistry.PublicKey public_key = 3 [json_name = "publicKey", (.buf.validate.field) = { ... }</code>
   * @return Whether the publicKey field is set.
   */
  boolean hasPublicKey();
  /**
   * <code>.kasregistry.PublicKey public_key = 3 [json_name = "publicKey", (.buf.validate.field) = { ... }</code>
   * @return The publicKey.
   */
  com.kasregistry.PublicKey getPublicKey();
  /**
   * <code>.kasregistry.PublicKey public_key = 3 [json_name = "publicKey", (.buf.validate.field) = { ... }</code>
   */
  com.kasregistry.PublicKeyOrBuilder getPublicKeyOrBuilder();
}
