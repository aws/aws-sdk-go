// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshiftserverless

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The submitted action has conflicts.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInsufficientCapacityException for service response error code
	// "InsufficientCapacityException".
	//
	// There is an insufficient capacity to perform the action.
	ErrCodeInsufficientCapacityException = "InsufficientCapacityException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The request processing has failed because of an unknown error, exception
	// or failure.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidPaginationException for service response error code
	// "InvalidPaginationException".
	//
	// The provided pagination token is invalid.
	ErrCodeInvalidPaginationException = "InvalidPaginationException"

	// ErrCodeIpv6CidrBlockNotFoundException for service response error code
	// "Ipv6CidrBlockNotFoundException".
	//
	// There are no subnets in your VPC with associated IPv6 CIDR blocks. To use
	// dual-stack mode, associate an IPv6 CIDR block with each subnet in your VPC.
	ErrCodeIpv6CidrBlockNotFoundException = "Ipv6CidrBlockNotFoundException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The service limit was exceeded.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The request exceeded the number of tags allowed for a resource.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input failed to satisfy the constraints specified by an AWS service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":          newErrorAccessDeniedException,
	"ConflictException":              newErrorConflictException,
	"InsufficientCapacityException":  newErrorInsufficientCapacityException,
	"InternalServerException":        newErrorInternalServerException,
	"InvalidPaginationException":     newErrorInvalidPaginationException,
	"Ipv6CidrBlockNotFoundException": newErrorIpv6CidrBlockNotFoundException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"ServiceQuotaExceededException":  newErrorServiceQuotaExceededException,
	"ThrottlingException":            newErrorThrottlingException,
	"TooManyTagsException":           newErrorTooManyTagsException,
	"ValidationException":            newErrorValidationException,
}
