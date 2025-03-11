// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: file.proto

package pb

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

// Validate checks the field values on UploadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UploadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UploadRequestMultiError, or
// nil if none found.
func (m *UploadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetFilename()); l < 1 || l > 1000 {
		err := UploadRequestValidationError{
			field:  "Filename",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetData()) < 1 {
		err := UploadRequestValidationError{
			field:  "Data",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UploadRequestMultiError(errors)
	}

	return nil
}

// UploadRequestMultiError is an error wrapping multiple validation errors
// returned by UploadRequest.ValidateAll() if the designated constraints
// aren't met.
type UploadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadRequestMultiError) AllErrors() []error { return m }

// UploadRequestValidationError is the validation error returned by
// UploadRequest.Validate if the designated constraints aren't met.
type UploadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadRequestValidationError) ErrorName() string { return "UploadRequestValidationError" }

// Error satisfies the builtin error interface
func (e UploadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadRequestValidationError{}

// Validate checks the field values on UploadResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UploadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UploadResponseMultiError,
// or nil if none found.
func (m *UploadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return UploadResponseMultiError(errors)
	}

	return nil
}

// UploadResponseMultiError is an error wrapping multiple validation errors
// returned by UploadResponse.ValidateAll() if the designated constraints
// aren't met.
type UploadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadResponseMultiError) AllErrors() []error { return m }

// UploadResponseValidationError is the validation error returned by
// UploadResponse.Validate if the designated constraints aren't met.
type UploadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadResponseValidationError) ErrorName() string { return "UploadResponseValidationError" }

// Error satisfies the builtin error interface
func (e UploadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadResponseValidationError{}

// Validate checks the field values on DownloadRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DownloadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadRequestMultiError, or nil if none found.
func (m *DownloadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetFilename()); l < 1 || l > 1000 {
		err := DownloadRequestValidationError{
			field:  "Filename",
			reason: "value length must be between 1 and 1000 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DownloadRequestMultiError(errors)
	}

	return nil
}

// DownloadRequestMultiError is an error wrapping multiple validation errors
// returned by DownloadRequest.ValidateAll() if the designated constraints
// aren't met.
type DownloadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadRequestMultiError) AllErrors() []error { return m }

// DownloadRequestValidationError is the validation error returned by
// DownloadRequest.Validate if the designated constraints aren't met.
type DownloadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadRequestValidationError) ErrorName() string { return "DownloadRequestValidationError" }

// Error satisfies the builtin error interface
func (e DownloadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadRequestValidationError{}

// Validate checks the field values on DownloadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DownloadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadResponseMultiError, or nil if none found.
func (m *DownloadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return DownloadResponseMultiError(errors)
	}

	return nil
}

// DownloadResponseMultiError is an error wrapping multiple validation errors
// returned by DownloadResponse.ValidateAll() if the designated constraints
// aren't met.
type DownloadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadResponseMultiError) AllErrors() []error { return m }

// DownloadResponseValidationError is the validation error returned by
// DownloadResponse.Validate if the designated constraints aren't met.
type DownloadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadResponseValidationError) ErrorName() string { return "DownloadResponseValidationError" }

// Error satisfies the builtin error interface
func (e DownloadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadResponseValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on FileInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FileInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FileInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FileInfoMultiError, or nil
// if none found.
func (m *FileInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *FileInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Filename

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return FileInfoMultiError(errors)
	}

	return nil
}

// FileInfoMultiError is an error wrapping multiple validation errors returned
// by FileInfo.ValidateAll() if the designated constraints aren't met.
type FileInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FileInfoMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FileInfoMultiError) AllErrors() []error { return m }

// FileInfoValidationError is the validation error returned by
// FileInfo.Validate if the designated constraints aren't met.
type FileInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileInfoValidationError) ErrorName() string { return "FileInfoValidationError" }

// Error satisfies the builtin error interface
func (e FileInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileInfoValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListResponseMultiError, or
// nil if none found.
func (m *ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetFiles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Files[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Files[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Files[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListResponseMultiError(errors)
	}

	return nil
}

// ListResponseMultiError is an error wrapping multiple validation errors
// returned by ListResponse.ValidateAll() if the designated constraints aren't met.
type ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResponseMultiError) AllErrors() []error { return m }

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}
