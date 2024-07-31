// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: policy/resourcemapping/resource_mapping.proto

package resourcemapping

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	common "github.com/opentdf/platform/protocol/go/common"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = common.MetadataUpdateEnum(0)
)

// Validate checks the field values on ListResourceMappingsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourceMappingsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourceMappingsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourceMappingsRequestMultiError, or nil if none found.
func (m *ListResourceMappingsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceMappingsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListResourceMappingsRequestMultiError(errors)
	}

	return nil
}

// ListResourceMappingsRequestMultiError is an error wrapping multiple
// validation errors returned by ListResourceMappingsRequest.ValidateAll() if
// the designated constraints aren't met.
type ListResourceMappingsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceMappingsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceMappingsRequestMultiError) AllErrors() []error { return m }

// ListResourceMappingsRequestValidationError is the validation error returned
// by ListResourceMappingsRequest.Validate if the designated constraints
// aren't met.
type ListResourceMappingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceMappingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceMappingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceMappingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceMappingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceMappingsRequestValidationError) ErrorName() string {
	return "ListResourceMappingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceMappingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceMappingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceMappingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceMappingsRequestValidationError{}

// Validate checks the field values on ListResourceMappingsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourceMappingsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourceMappingsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourceMappingsResponseMultiError, or nil if none found.
func (m *ListResourceMappingsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceMappingsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResourceMappings() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResourceMappingsResponseValidationError{
						field:  fmt.Sprintf("ResourceMappings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResourceMappingsResponseValidationError{
						field:  fmt.Sprintf("ResourceMappings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResourceMappingsResponseValidationError{
					field:  fmt.Sprintf("ResourceMappings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListResourceMappingsResponseMultiError(errors)
	}

	return nil
}

// ListResourceMappingsResponseMultiError is an error wrapping multiple
// validation errors returned by ListResourceMappingsResponse.ValidateAll() if
// the designated constraints aren't met.
type ListResourceMappingsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceMappingsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceMappingsResponseMultiError) AllErrors() []error { return m }

// ListResourceMappingsResponseValidationError is the validation error returned
// by ListResourceMappingsResponse.Validate if the designated constraints
// aren't met.
type ListResourceMappingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceMappingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceMappingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceMappingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceMappingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceMappingsResponseValidationError) ErrorName() string {
	return "ListResourceMappingsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceMappingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceMappingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceMappingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceMappingsResponseValidationError{}

// Validate checks the field values on GetResourceMappingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetResourceMappingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourceMappingRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourceMappingRequestMultiError, or nil if none found.
func (m *GetResourceMappingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourceMappingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetResourceMappingRequestMultiError(errors)
	}

	return nil
}

// GetResourceMappingRequestMultiError is an error wrapping multiple validation
// errors returned by GetResourceMappingRequest.ValidateAll() if the
// designated constraints aren't met.
type GetResourceMappingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourceMappingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourceMappingRequestMultiError) AllErrors() []error { return m }

// GetResourceMappingRequestValidationError is the validation error returned by
// GetResourceMappingRequest.Validate if the designated constraints aren't met.
type GetResourceMappingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourceMappingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourceMappingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourceMappingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourceMappingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourceMappingRequestValidationError) ErrorName() string {
	return "GetResourceMappingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetResourceMappingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourceMappingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourceMappingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourceMappingRequestValidationError{}

// Validate checks the field values on GetResourceMappingResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetResourceMappingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourceMappingResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourceMappingResponseMultiError, or nil if none found.
func (m *GetResourceMappingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourceMappingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResourceMapping()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResourceMapping()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResourceMappingResponseValidationError{
				field:  "ResourceMapping",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetResourceMappingResponseMultiError(errors)
	}

	return nil
}

// GetResourceMappingResponseMultiError is an error wrapping multiple
// validation errors returned by GetResourceMappingResponse.ValidateAll() if
// the designated constraints aren't met.
type GetResourceMappingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourceMappingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourceMappingResponseMultiError) AllErrors() []error { return m }

// GetResourceMappingResponseValidationError is the validation error returned
// by GetResourceMappingResponse.Validate if the designated constraints aren't met.
type GetResourceMappingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourceMappingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourceMappingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourceMappingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourceMappingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourceMappingResponseValidationError) ErrorName() string {
	return "GetResourceMappingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetResourceMappingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourceMappingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourceMappingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourceMappingResponseValidationError{}

// Validate checks the field values on CreateResourceMappingRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateResourceMappingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResourceMappingRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateResourceMappingRequestMultiError, or nil if none found.
func (m *CreateResourceMappingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResourceMappingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AttributeValueId

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateResourceMappingRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateResourceMappingRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateResourceMappingRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateResourceMappingRequestMultiError(errors)
	}

	return nil
}

// CreateResourceMappingRequestMultiError is an error wrapping multiple
// validation errors returned by CreateResourceMappingRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateResourceMappingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResourceMappingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResourceMappingRequestMultiError) AllErrors() []error { return m }

// CreateResourceMappingRequestValidationError is the validation error returned
// by CreateResourceMappingRequest.Validate if the designated constraints
// aren't met.
type CreateResourceMappingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResourceMappingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResourceMappingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResourceMappingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResourceMappingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResourceMappingRequestValidationError) ErrorName() string {
	return "CreateResourceMappingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResourceMappingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResourceMappingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResourceMappingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResourceMappingRequestValidationError{}

// Validate checks the field values on CreateResourceMappingResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateResourceMappingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResourceMappingResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateResourceMappingResponseMultiError, or nil if none found.
func (m *CreateResourceMappingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResourceMappingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResourceMapping()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResourceMapping()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateResourceMappingResponseValidationError{
				field:  "ResourceMapping",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateResourceMappingResponseMultiError(errors)
	}

	return nil
}

// CreateResourceMappingResponseMultiError is an error wrapping multiple
// validation errors returned by CreateResourceMappingResponse.ValidateAll()
// if the designated constraints aren't met.
type CreateResourceMappingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResourceMappingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResourceMappingResponseMultiError) AllErrors() []error { return m }

// CreateResourceMappingResponseValidationError is the validation error
// returned by CreateResourceMappingResponse.Validate if the designated
// constraints aren't met.
type CreateResourceMappingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResourceMappingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResourceMappingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResourceMappingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResourceMappingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResourceMappingResponseValidationError) ErrorName() string {
	return "CreateResourceMappingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResourceMappingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResourceMappingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResourceMappingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResourceMappingResponseValidationError{}

// Validate checks the field values on UpdateResourceMappingRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceMappingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceMappingRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateResourceMappingRequestMultiError, or nil if none found.
func (m *UpdateResourceMappingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceMappingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for AttributeValueId

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateResourceMappingRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateResourceMappingRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateResourceMappingRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MetadataUpdateBehavior

	if len(errors) > 0 {
		return UpdateResourceMappingRequestMultiError(errors)
	}

	return nil
}

// UpdateResourceMappingRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateResourceMappingRequest.ValidateAll() if
// the designated constraints aren't met.
type UpdateResourceMappingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceMappingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceMappingRequestMultiError) AllErrors() []error { return m }

// UpdateResourceMappingRequestValidationError is the validation error returned
// by UpdateResourceMappingRequest.Validate if the designated constraints
// aren't met.
type UpdateResourceMappingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceMappingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceMappingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceMappingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceMappingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceMappingRequestValidationError) ErrorName() string {
	return "UpdateResourceMappingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceMappingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceMappingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceMappingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceMappingRequestValidationError{}

// Validate checks the field values on UpdateResourceMappingResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceMappingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceMappingResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateResourceMappingResponseMultiError, or nil if none found.
func (m *UpdateResourceMappingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceMappingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResourceMapping()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResourceMapping()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateResourceMappingResponseValidationError{
				field:  "ResourceMapping",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateResourceMappingResponseMultiError(errors)
	}

	return nil
}

// UpdateResourceMappingResponseMultiError is an error wrapping multiple
// validation errors returned by UpdateResourceMappingResponse.ValidateAll()
// if the designated constraints aren't met.
type UpdateResourceMappingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceMappingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceMappingResponseMultiError) AllErrors() []error { return m }

// UpdateResourceMappingResponseValidationError is the validation error
// returned by UpdateResourceMappingResponse.Validate if the designated
// constraints aren't met.
type UpdateResourceMappingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceMappingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceMappingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceMappingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceMappingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceMappingResponseValidationError) ErrorName() string {
	return "UpdateResourceMappingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceMappingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceMappingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceMappingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceMappingResponseValidationError{}

// Validate checks the field values on DeleteResourceMappingRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteResourceMappingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResourceMappingRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteResourceMappingRequestMultiError, or nil if none found.
func (m *DeleteResourceMappingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResourceMappingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteResourceMappingRequestMultiError(errors)
	}

	return nil
}

// DeleteResourceMappingRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteResourceMappingRequest.ValidateAll() if
// the designated constraints aren't met.
type DeleteResourceMappingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResourceMappingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResourceMappingRequestMultiError) AllErrors() []error { return m }

// DeleteResourceMappingRequestValidationError is the validation error returned
// by DeleteResourceMappingRequest.Validate if the designated constraints
// aren't met.
type DeleteResourceMappingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResourceMappingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResourceMappingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResourceMappingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResourceMappingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResourceMappingRequestValidationError) ErrorName() string {
	return "DeleteResourceMappingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteResourceMappingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResourceMappingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResourceMappingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResourceMappingRequestValidationError{}

// Validate checks the field values on DeleteResourceMappingResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteResourceMappingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResourceMappingResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeleteResourceMappingResponseMultiError, or nil if none found.
func (m *DeleteResourceMappingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResourceMappingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResourceMapping()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteResourceMappingResponseValidationError{
					field:  "ResourceMapping",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResourceMapping()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteResourceMappingResponseValidationError{
				field:  "ResourceMapping",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeleteResourceMappingResponseMultiError(errors)
	}

	return nil
}

// DeleteResourceMappingResponseMultiError is an error wrapping multiple
// validation errors returned by DeleteResourceMappingResponse.ValidateAll()
// if the designated constraints aren't met.
type DeleteResourceMappingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResourceMappingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResourceMappingResponseMultiError) AllErrors() []error { return m }

// DeleteResourceMappingResponseValidationError is the validation error
// returned by DeleteResourceMappingResponse.Validate if the designated
// constraints aren't met.
type DeleteResourceMappingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResourceMappingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResourceMappingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResourceMappingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResourceMappingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResourceMappingResponseValidationError) ErrorName() string {
	return "DeleteResourceMappingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteResourceMappingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResourceMappingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResourceMappingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResourceMappingResponseValidationError{}
