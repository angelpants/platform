// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: policy/selectors.proto

package policy

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
)

// Validate checks the field values on AttributeNamespaceSelector with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeNamespaceSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeNamespaceSelector with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeNamespaceSelectorMultiError, or nil if none found.
func (m *AttributeNamespaceSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeNamespaceSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWithAttributes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeNamespaceSelectorValidationError{
					field:  "WithAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeNamespaceSelectorValidationError{
					field:  "WithAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWithAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeNamespaceSelectorValidationError{
				field:  "WithAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeNamespaceSelectorMultiError(errors)
	}

	return nil
}

// AttributeNamespaceSelectorMultiError is an error wrapping multiple
// validation errors returned by AttributeNamespaceSelector.ValidateAll() if
// the designated constraints aren't met.
type AttributeNamespaceSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeNamespaceSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeNamespaceSelectorMultiError) AllErrors() []error { return m }

// AttributeNamespaceSelectorValidationError is the validation error returned
// by AttributeNamespaceSelector.Validate if the designated constraints aren't met.
type AttributeNamespaceSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeNamespaceSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeNamespaceSelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeNamespaceSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeNamespaceSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeNamespaceSelectorValidationError) ErrorName() string {
	return "AttributeNamespaceSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeNamespaceSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeNamespaceSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeNamespaceSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeNamespaceSelectorValidationError{}

// Validate checks the field values on AttributeDefinitionSelector with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeDefinitionSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeDefinitionSelector with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeDefinitionSelectorMultiError, or nil if none found.
func (m *AttributeDefinitionSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeDefinitionSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WithKeyAccessGrants

	if all {
		switch v := interface{}(m.GetWithNamespace()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeDefinitionSelectorValidationError{
					field:  "WithNamespace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeDefinitionSelectorValidationError{
					field:  "WithNamespace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWithNamespace()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeDefinitionSelectorValidationError{
				field:  "WithNamespace",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWithValues()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeDefinitionSelectorValidationError{
					field:  "WithValues",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeDefinitionSelectorValidationError{
					field:  "WithValues",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWithValues()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeDefinitionSelectorValidationError{
				field:  "WithValues",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeDefinitionSelectorMultiError(errors)
	}

	return nil
}

// AttributeDefinitionSelectorMultiError is an error wrapping multiple
// validation errors returned by AttributeDefinitionSelector.ValidateAll() if
// the designated constraints aren't met.
type AttributeDefinitionSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeDefinitionSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeDefinitionSelectorMultiError) AllErrors() []error { return m }

// AttributeDefinitionSelectorValidationError is the validation error returned
// by AttributeDefinitionSelector.Validate if the designated constraints
// aren't met.
type AttributeDefinitionSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeDefinitionSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeDefinitionSelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeDefinitionSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeDefinitionSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeDefinitionSelectorValidationError) ErrorName() string {
	return "AttributeDefinitionSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeDefinitionSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeDefinitionSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeDefinitionSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeDefinitionSelectorValidationError{}

// Validate checks the field values on AttributeValueSelector with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeValueSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttributeValueSelector with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeValueSelectorMultiError, or nil if none found.
func (m *AttributeValueSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeValueSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WithKeyAccessGrants

	// no validation rules for WithSubjectMaps

	// no validation rules for WithResourceMaps

	if all {
		switch v := interface{}(m.GetWithAttribute()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeValueSelectorValidationError{
					field:  "WithAttribute",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeValueSelectorValidationError{
					field:  "WithAttribute",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWithAttribute()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeValueSelectorValidationError{
				field:  "WithAttribute",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeValueSelectorMultiError(errors)
	}

	return nil
}

// AttributeValueSelectorMultiError is an error wrapping multiple validation
// errors returned by AttributeValueSelector.ValidateAll() if the designated
// constraints aren't met.
type AttributeValueSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeValueSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeValueSelectorMultiError) AllErrors() []error { return m }

// AttributeValueSelectorValidationError is the validation error returned by
// AttributeValueSelector.Validate if the designated constraints aren't met.
type AttributeValueSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeValueSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeValueSelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeValueSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeValueSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeValueSelectorValidationError) ErrorName() string {
	return "AttributeValueSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeValueSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeValueSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeValueSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeValueSelectorValidationError{}

// Validate checks the field values on
// AttributeNamespaceSelector_AttributeSelector with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AttributeNamespaceSelector_AttributeSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttributeNamespaceSelector_AttributeSelector with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// AttributeNamespaceSelector_AttributeSelectorMultiError, or nil if none found.
func (m *AttributeNamespaceSelector_AttributeSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeNamespaceSelector_AttributeSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WithKeyAccessGrants

	if all {
		switch v := interface{}(m.GetWithValues()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeNamespaceSelector_AttributeSelectorValidationError{
					field:  "WithValues",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeNamespaceSelector_AttributeSelectorValidationError{
					field:  "WithValues",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWithValues()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeNamespaceSelector_AttributeSelectorValidationError{
				field:  "WithValues",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeNamespaceSelector_AttributeSelectorMultiError(errors)
	}

	return nil
}

// AttributeNamespaceSelector_AttributeSelectorMultiError is an error wrapping
// multiple validation errors returned by
// AttributeNamespaceSelector_AttributeSelector.ValidateAll() if the
// designated constraints aren't met.
type AttributeNamespaceSelector_AttributeSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeNamespaceSelector_AttributeSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeNamespaceSelector_AttributeSelectorMultiError) AllErrors() []error { return m }

// AttributeNamespaceSelector_AttributeSelectorValidationError is the
// validation error returned by
// AttributeNamespaceSelector_AttributeSelector.Validate if the designated
// constraints aren't met.
type AttributeNamespaceSelector_AttributeSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeNamespaceSelector_AttributeSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeNamespaceSelector_AttributeSelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeNamespaceSelector_AttributeSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeNamespaceSelector_AttributeSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeNamespaceSelector_AttributeSelectorValidationError) ErrorName() string {
	return "AttributeNamespaceSelector_AttributeSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeNamespaceSelector_AttributeSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeNamespaceSelector_AttributeSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeNamespaceSelector_AttributeSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeNamespaceSelector_AttributeSelectorValidationError{}

// Validate checks the field values on
// AttributeNamespaceSelector_AttributeSelector_ValueSelector with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeNamespaceSelector_AttributeSelector_ValueSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttributeNamespaceSelector_AttributeSelector_ValueSelector with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeNamespaceSelector_AttributeSelector_ValueSelectorMultiError, or
// nil if none found.
func (m *AttributeNamespaceSelector_AttributeSelector_ValueSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeNamespaceSelector_AttributeSelector_ValueSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WithKeyAccessGrants

	// no validation rules for WithSubjectMaps

	// no validation rules for WithResourceMaps

	if len(errors) > 0 {
		return AttributeNamespaceSelector_AttributeSelector_ValueSelectorMultiError(errors)
	}

	return nil
}

// AttributeNamespaceSelector_AttributeSelector_ValueSelectorMultiError is an
// error wrapping multiple validation errors returned by
// AttributeNamespaceSelector_AttributeSelector_ValueSelector.ValidateAll() if
// the designated constraints aren't met.
type AttributeNamespaceSelector_AttributeSelector_ValueSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeNamespaceSelector_AttributeSelector_ValueSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeNamespaceSelector_AttributeSelector_ValueSelectorMultiError) AllErrors() []error {
	return m
}

// AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError is
// the validation error returned by
// AttributeNamespaceSelector_AttributeSelector_ValueSelector.Validate if the
// designated constraints aren't met.
type AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError) ErrorName() string {
	return "AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeNamespaceSelector_AttributeSelector_ValueSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeNamespaceSelector_AttributeSelector_ValueSelectorValidationError{}

// Validate checks the field values on
// AttributeDefinitionSelector_NamespaceSelector with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AttributeDefinitionSelector_NamespaceSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttributeDefinitionSelector_NamespaceSelector with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// AttributeDefinitionSelector_NamespaceSelectorMultiError, or nil if none found.
func (m *AttributeDefinitionSelector_NamespaceSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeDefinitionSelector_NamespaceSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AttributeDefinitionSelector_NamespaceSelectorMultiError(errors)
	}

	return nil
}

// AttributeDefinitionSelector_NamespaceSelectorMultiError is an error wrapping
// multiple validation errors returned by
// AttributeDefinitionSelector_NamespaceSelector.ValidateAll() if the
// designated constraints aren't met.
type AttributeDefinitionSelector_NamespaceSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeDefinitionSelector_NamespaceSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeDefinitionSelector_NamespaceSelectorMultiError) AllErrors() []error { return m }

// AttributeDefinitionSelector_NamespaceSelectorValidationError is the
// validation error returned by
// AttributeDefinitionSelector_NamespaceSelector.Validate if the designated
// constraints aren't met.
type AttributeDefinitionSelector_NamespaceSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeDefinitionSelector_NamespaceSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeDefinitionSelector_NamespaceSelectorValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e AttributeDefinitionSelector_NamespaceSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeDefinitionSelector_NamespaceSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeDefinitionSelector_NamespaceSelectorValidationError) ErrorName() string {
	return "AttributeDefinitionSelector_NamespaceSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeDefinitionSelector_NamespaceSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeDefinitionSelector_NamespaceSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeDefinitionSelector_NamespaceSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeDefinitionSelector_NamespaceSelectorValidationError{}

// Validate checks the field values on
// AttributeDefinitionSelector_ValueSelector with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AttributeDefinitionSelector_ValueSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttributeDefinitionSelector_ValueSelector with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// AttributeDefinitionSelector_ValueSelectorMultiError, or nil if none found.
func (m *AttributeDefinitionSelector_ValueSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeDefinitionSelector_ValueSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WithKeyAccessGrants

	// no validation rules for WithSubjectMaps

	// no validation rules for WithResourceMaps

	if len(errors) > 0 {
		return AttributeDefinitionSelector_ValueSelectorMultiError(errors)
	}

	return nil
}

// AttributeDefinitionSelector_ValueSelectorMultiError is an error wrapping
// multiple validation errors returned by
// AttributeDefinitionSelector_ValueSelector.ValidateAll() if the designated
// constraints aren't met.
type AttributeDefinitionSelector_ValueSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeDefinitionSelector_ValueSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeDefinitionSelector_ValueSelectorMultiError) AllErrors() []error { return m }

// AttributeDefinitionSelector_ValueSelectorValidationError is the validation
// error returned by AttributeDefinitionSelector_ValueSelector.Validate if the
// designated constraints aren't met.
type AttributeDefinitionSelector_ValueSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeDefinitionSelector_ValueSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeDefinitionSelector_ValueSelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeDefinitionSelector_ValueSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeDefinitionSelector_ValueSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeDefinitionSelector_ValueSelectorValidationError) ErrorName() string {
	return "AttributeDefinitionSelector_ValueSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeDefinitionSelector_ValueSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeDefinitionSelector_ValueSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeDefinitionSelector_ValueSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeDefinitionSelector_ValueSelectorValidationError{}

// Validate checks the field values on AttributeValueSelector_AttributeSelector
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *AttributeValueSelector_AttributeSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttributeValueSelector_AttributeSelector with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// AttributeValueSelector_AttributeSelectorMultiError, or nil if none found.
func (m *AttributeValueSelector_AttributeSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeValueSelector_AttributeSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WithKeyAccessGrants

	if all {
		switch v := interface{}(m.GetWithNamespace()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttributeValueSelector_AttributeSelectorValidationError{
					field:  "WithNamespace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttributeValueSelector_AttributeSelectorValidationError{
					field:  "WithNamespace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWithNamespace()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeValueSelector_AttributeSelectorValidationError{
				field:  "WithNamespace",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttributeValueSelector_AttributeSelectorMultiError(errors)
	}

	return nil
}

// AttributeValueSelector_AttributeSelectorMultiError is an error wrapping
// multiple validation errors returned by
// AttributeValueSelector_AttributeSelector.ValidateAll() if the designated
// constraints aren't met.
type AttributeValueSelector_AttributeSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeValueSelector_AttributeSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeValueSelector_AttributeSelectorMultiError) AllErrors() []error { return m }

// AttributeValueSelector_AttributeSelectorValidationError is the validation
// error returned by AttributeValueSelector_AttributeSelector.Validate if the
// designated constraints aren't met.
type AttributeValueSelector_AttributeSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeValueSelector_AttributeSelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeValueSelector_AttributeSelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeValueSelector_AttributeSelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeValueSelector_AttributeSelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeValueSelector_AttributeSelectorValidationError) ErrorName() string {
	return "AttributeValueSelector_AttributeSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeValueSelector_AttributeSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeValueSelector_AttributeSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeValueSelector_AttributeSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeValueSelector_AttributeSelectorValidationError{}

// Validate checks the field values on
// AttributeValueSelector_AttributeSelector_NamespaceSelector with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttributeValueSelector_AttributeSelector_NamespaceSelector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttributeValueSelector_AttributeSelector_NamespaceSelector with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttributeValueSelector_AttributeSelector_NamespaceSelectorMultiError, or
// nil if none found.
func (m *AttributeValueSelector_AttributeSelector_NamespaceSelector) ValidateAll() error {
	return m.validate(true)
}

func (m *AttributeValueSelector_AttributeSelector_NamespaceSelector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AttributeValueSelector_AttributeSelector_NamespaceSelectorMultiError(errors)
	}

	return nil
}

// AttributeValueSelector_AttributeSelector_NamespaceSelectorMultiError is an
// error wrapping multiple validation errors returned by
// AttributeValueSelector_AttributeSelector_NamespaceSelector.ValidateAll() if
// the designated constraints aren't met.
type AttributeValueSelector_AttributeSelector_NamespaceSelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttributeValueSelector_AttributeSelector_NamespaceSelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttributeValueSelector_AttributeSelector_NamespaceSelectorMultiError) AllErrors() []error {
	return m
}

// AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError is
// the validation error returned by
// AttributeValueSelector_AttributeSelector_NamespaceSelector.Validate if the
// designated constraints aren't met.
type AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError) ErrorName() string {
	return "AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeValueSelector_AttributeSelector_NamespaceSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeValueSelector_AttributeSelector_NamespaceSelectorValidationError{}
