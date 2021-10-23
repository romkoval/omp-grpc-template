// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/omp_template_api/omp_template_api.proto

package omp_template_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Template with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Template) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Foo

	return nil
}

// TemplateValidationError is the validation error returned by
// Template.Validate if the designated constraints aren't met.
type TemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TemplateValidationError) ErrorName() string { return "TemplateValidationError" }

// Error satisfies the builtin error interface
func (e TemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TemplateValidationError{}

// Validate checks the field values on DescribeTemplateV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeTemplateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return DescribeTemplateV1RequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeTemplateV1RequestValidationError is the validation error returned by
// DescribeTemplateV1Request.Validate if the designated constraints aren't met.
type DescribeTemplateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeTemplateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeTemplateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeTemplateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeTemplateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeTemplateV1RequestValidationError) ErrorName() string {
	return "DescribeTemplateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeTemplateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeTemplateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeTemplateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeTemplateV1RequestValidationError{}

// Validate checks the field values on DescribeTemplateV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeTemplateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeTemplateV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeTemplateV1ResponseValidationError is the validation error returned
// by DescribeTemplateV1Response.Validate if the designated constraints aren't met.
type DescribeTemplateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeTemplateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeTemplateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeTemplateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeTemplateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeTemplateV1ResponseValidationError) ErrorName() string {
	return "DescribeTemplateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeTemplateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeTemplateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeTemplateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeTemplateV1ResponseValidationError{}
