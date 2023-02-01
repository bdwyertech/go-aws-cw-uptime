// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// More than one process tried to modify a resource at the same time.
type ConcurrentModificationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ConcurrentModificationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Some part of the dashboard data is invalid.
type DashboardInvalidInputError struct {
	Message *string

	ErrorCodeOverride *string

	DashboardValidationMessages []DashboardValidationMessage

	noSmithyDocumentSerde
}

func (e *DashboardInvalidInputError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DashboardInvalidInputError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DashboardInvalidInputError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidParameterInput"
	}
	return *e.ErrorCodeOverride
}
func (e *DashboardInvalidInputError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified dashboard does not exist.
type DashboardNotFoundError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DashboardNotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DashboardNotFoundError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DashboardNotFoundError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *DashboardNotFoundError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Request processing has failed due to some unknown error, exception, or failure.
type InternalServiceFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServiceFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceFault) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalServiceError"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceFault) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Data was not syntactically valid JSON.
type InvalidFormatFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidFormatFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFormatFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFormatFault) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidFormat"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidFormatFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The next token specified is invalid.
type InvalidNextToken struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNextToken) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextToken) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextToken) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidNextToken"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextToken) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Parameters were used together that cannot be used together.
type InvalidParameterCombinationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterCombinationException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidParameterCombination"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterCombinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The value of an input parameter is bad or out-of-range.
type InvalidParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidParameterValue"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation exceeded one or more limits.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The quota for alarms for this customer has already been reached.
type LimitExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededFault) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An input parameter that is required is missing.
type MissingRequiredParameterException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MissingRequiredParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingRequiredParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingRequiredParameterException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "MissingParameter"
	}
	return *e.ErrorCodeOverride
}
func (e *MissingRequiredParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The named resource does not exist.
type ResourceNotFound struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The named resource does not exist.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType *string
	ResourceId   *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
