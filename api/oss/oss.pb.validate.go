// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: oss-proto/oss.proto

package oss

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

// Validate checks the field values on PreSignedRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PreSignedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PreSignedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PreSignedRequestMultiError, or nil if none found.
func (m *PreSignedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PreSignedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Category

	// no validation rules for Name

	if len(errors) > 0 {
		return PreSignedRequestMultiError(errors)
	}

	return nil
}

// PreSignedRequestMultiError is an error wrapping multiple validation errors
// returned by PreSignedRequest.ValidateAll() if the designated constraints
// aren't met.
type PreSignedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PreSignedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PreSignedRequestMultiError) AllErrors() []error { return m }

// PreSignedRequestValidationError is the validation error returned by
// PreSignedRequest.Validate if the designated constraints aren't met.
type PreSignedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PreSignedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PreSignedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PreSignedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PreSignedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PreSignedRequestValidationError) ErrorName() string { return "PreSignedRequestValidationError" }

// Error satisfies the builtin error interface
func (e PreSignedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPreSignedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PreSignedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PreSignedRequestValidationError{}

// Validate checks the field values on PreSignedReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PreSignedReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PreSignedReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PreSignedReplyMultiError,
// or nil if none found.
func (m *PreSignedReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PreSignedReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uri

	if len(errors) > 0 {
		return PreSignedReplyMultiError(errors)
	}

	return nil
}

// PreSignedReplyMultiError is an error wrapping multiple validation errors
// returned by PreSignedReply.ValidateAll() if the designated constraints
// aren't met.
type PreSignedReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PreSignedReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PreSignedReplyMultiError) AllErrors() []error { return m }

// PreSignedReplyValidationError is the validation error returned by
// PreSignedReply.Validate if the designated constraints aren't met.
type PreSignedReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PreSignedReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PreSignedReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PreSignedReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PreSignedReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PreSignedReplyValidationError) ErrorName() string { return "PreSignedReplyValidationError" }

// Error satisfies the builtin error interface
func (e PreSignedReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPreSignedReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PreSignedReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PreSignedReplyValidationError{}

// Validate checks the field values on OcrRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OcrRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OcrRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OcrRequestMultiError, or
// nil if none found.
func (m *OcrRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OcrRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OcrRequestMultiError(errors)
	}

	return nil
}

// OcrRequestMultiError is an error wrapping multiple validation errors
// returned by OcrRequest.ValidateAll() if the designated constraints aren't met.
type OcrRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OcrRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OcrRequestMultiError) AllErrors() []error { return m }

// OcrRequestValidationError is the validation error returned by
// OcrRequest.Validate if the designated constraints aren't met.
type OcrRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OcrRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OcrRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OcrRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OcrRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OcrRequestValidationError) ErrorName() string { return "OcrRequestValidationError" }

// Error satisfies the builtin error interface
func (e OcrRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOcrRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OcrRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OcrRequestValidationError{}

// Validate checks the field values on OcrResult with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OcrResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OcrResult with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OcrResultMultiError, or nil
// if none found.
func (m *OcrResult) ValidateAll() error {
	return m.validate(true)
}

func (m *OcrResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Original

	// no validation rules for Msg

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OcrResultValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OcrResultValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OcrResultValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ParseLatency

	// no validation rules for OcrLatency

	if len(errors) > 0 {
		return OcrResultMultiError(errors)
	}

	return nil
}

// OcrResultMultiError is an error wrapping multiple validation errors returned
// by OcrResult.ValidateAll() if the designated constraints aren't met.
type OcrResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OcrResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OcrResultMultiError) AllErrors() []error { return m }

// OcrResultValidationError is the validation error returned by
// OcrResult.Validate if the designated constraints aren't met.
type OcrResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OcrResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OcrResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OcrResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OcrResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OcrResultValidationError) ErrorName() string { return "OcrResultValidationError" }

// Error satisfies the builtin error interface
func (e OcrResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOcrResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OcrResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OcrResultValidationError{}

// Validate checks the field values on OcrItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OcrItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OcrItem with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OcrItemMultiError, or nil if none found.
func (m *OcrItem) ValidateAll() error {
	return m.validate(true)
}

func (m *OcrItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Boxes

	// no validation rules for Confidence

	// no validation rules for Text

	if len(errors) > 0 {
		return OcrItemMultiError(errors)
	}

	return nil
}

// OcrItemMultiError is an error wrapping multiple validation errors returned
// by OcrItem.ValidateAll() if the designated constraints aren't met.
type OcrItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OcrItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OcrItemMultiError) AllErrors() []error { return m }

// OcrItemValidationError is the validation error returned by OcrItem.Validate
// if the designated constraints aren't met.
type OcrItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OcrItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OcrItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OcrItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OcrItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OcrItemValidationError) ErrorName() string { return "OcrItemValidationError" }

// Error satisfies the builtin error interface
func (e OcrItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOcrItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OcrItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OcrItemValidationError{}

// Validate checks the field values on OcrReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OcrReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OcrReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OcrReplyMultiError, or nil
// if none found.
func (m *OcrReply) ValidateAll() error {
	return m.validate(true)
}

func (m *OcrReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OcrReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OcrReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OcrReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Latency

	if len(errors) > 0 {
		return OcrReplyMultiError(errors)
	}

	return nil
}

// OcrReplyMultiError is an error wrapping multiple validation errors returned
// by OcrReply.ValidateAll() if the designated constraints aren't met.
type OcrReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OcrReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OcrReplyMultiError) AllErrors() []error { return m }

// OcrReplyValidationError is the validation error returned by
// OcrReply.Validate if the designated constraints aren't met.
type OcrReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OcrReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OcrReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OcrReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OcrReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OcrReplyValidationError) ErrorName() string { return "OcrReplyValidationError" }

// Error satisfies the builtin error interface
func (e OcrReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOcrReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OcrReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OcrReplyValidationError{}
