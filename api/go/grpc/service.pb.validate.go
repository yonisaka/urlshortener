// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/service.proto

package grpc

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

// Validate checks the field values on CreateURLShortenerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateURLShortenerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateURLShortenerRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateURLShortenerRequestMultiError, or nil if none found.
func (m *CreateURLShortenerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateURLShortenerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := CreateURLShortenerRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUrl()); l < 1 || l > 2048 {
		err := CreateURLShortenerRequestValidationError{
			field:  "Url",
			reason: "value length must be between 1 and 2048 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetDatetime() == nil {
		err := CreateURLShortenerRequestValidationError{
			field:  "Datetime",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateURLShortenerRequestMultiError(errors)
	}

	return nil
}

// CreateURLShortenerRequestMultiError is an error wrapping multiple validation
// errors returned by CreateURLShortenerRequest.ValidateAll() if the
// designated constraints aren't met.
type CreateURLShortenerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateURLShortenerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateURLShortenerRequestMultiError) AllErrors() []error { return m }

// CreateURLShortenerRequestValidationError is the validation error returned by
// CreateURLShortenerRequest.Validate if the designated constraints aren't met.
type CreateURLShortenerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateURLShortenerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateURLShortenerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateURLShortenerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateURLShortenerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateURLShortenerRequestValidationError) ErrorName() string {
	return "CreateURLShortenerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateURLShortenerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateURLShortenerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateURLShortenerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateURLShortenerRequestValidationError{}

// Validate checks the field values on ListURLShortenerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListURLShortenerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListURLShortenerRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListURLShortenerRequestMultiError, or nil if none found.
func (m *ListURLShortenerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListURLShortenerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := ListURLShortenerRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStartDatetime() == nil {
		err := ListURLShortenerRequestValidationError{
			field:  "StartDatetime",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetEndDatetime() == nil {
		err := ListURLShortenerRequestValidationError{
			field:  "EndDatetime",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListURLShortenerRequestMultiError(errors)
	}

	return nil
}

// ListURLShortenerRequestMultiError is an error wrapping multiple validation
// errors returned by ListURLShortenerRequest.ValidateAll() if the designated
// constraints aren't met.
type ListURLShortenerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListURLShortenerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListURLShortenerRequestMultiError) AllErrors() []error { return m }

// ListURLShortenerRequestValidationError is the validation error returned by
// ListURLShortenerRequest.Validate if the designated constraints aren't met.
type ListURLShortenerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListURLShortenerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListURLShortenerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListURLShortenerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListURLShortenerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListURLShortenerRequestValidationError) ErrorName() string {
	return "ListURLShortenerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListURLShortenerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListURLShortenerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListURLShortenerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListURLShortenerRequestValidationError{}

// Validate checks the field values on GetShortenedURLRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetShortenedURLRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetShortenedURLRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetShortenedURLRequestMultiError, or nil if none found.
func (m *GetShortenedURLRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetShortenedURLRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUrl()); l < 1 || l > 2048 {
		err := GetShortenedURLRequestValidationError{
			field:  "Url",
			reason: "value length must be between 1 and 2048 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetShortenedURLRequestMultiError(errors)
	}

	return nil
}

// GetShortenedURLRequestMultiError is an error wrapping multiple validation
// errors returned by GetShortenedURLRequest.ValidateAll() if the designated
// constraints aren't met.
type GetShortenedURLRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetShortenedURLRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetShortenedURLRequestMultiError) AllErrors() []error { return m }

// GetShortenedURLRequestValidationError is the validation error returned by
// GetShortenedURLRequest.Validate if the designated constraints aren't met.
type GetShortenedURLRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShortenedURLRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShortenedURLRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShortenedURLRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShortenedURLRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShortenedURLRequestValidationError) ErrorName() string {
	return "GetShortenedURLRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetShortenedURLRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShortenedURLRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShortenedURLRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShortenedURLRequestValidationError{}

// Validate checks the field values on ListURLShortenerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListURLShortenerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListURLShortenerResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListURLShortenerResponseMultiError, or nil if none found.
func (m *ListURLShortenerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListURLShortenerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUrlShortener() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListURLShortenerResponseValidationError{
						field:  fmt.Sprintf("UrlShortener[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListURLShortenerResponseValidationError{
						field:  fmt.Sprintf("UrlShortener[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListURLShortenerResponseValidationError{
					field:  fmt.Sprintf("UrlShortener[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListURLShortenerResponseMultiError(errors)
	}

	return nil
}

// ListURLShortenerResponseMultiError is an error wrapping multiple validation
// errors returned by ListURLShortenerResponse.ValidateAll() if the designated
// constraints aren't met.
type ListURLShortenerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListURLShortenerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListURLShortenerResponseMultiError) AllErrors() []error { return m }

// ListURLShortenerResponseValidationError is the validation error returned by
// ListURLShortenerResponse.Validate if the designated constraints aren't met.
type ListURLShortenerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListURLShortenerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListURLShortenerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListURLShortenerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListURLShortenerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListURLShortenerResponseValidationError) ErrorName() string {
	return "ListURLShortenerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListURLShortenerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListURLShortenerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListURLShortenerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListURLShortenerResponseValidationError{}
