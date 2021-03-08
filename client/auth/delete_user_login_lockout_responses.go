// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// DeleteUserLoginLockoutReader is a Reader for the DeleteUserLoginLockout structure.
type DeleteUserLoginLockoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserLoginLockoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserLoginLockoutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserLoginLockoutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserLoginLockoutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserLoginLockoutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserLoginLockoutNoContent creates a DeleteUserLoginLockoutNoContent with default headers values
func NewDeleteUserLoginLockoutNoContent() *DeleteUserLoginLockoutNoContent {
	return &DeleteUserLoginLockoutNoContent{}
}

/* DeleteUserLoginLockoutNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteUserLoginLockoutNoContent struct {
	Payload string
}

func (o *DeleteUserLoginLockoutNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user_login_lockout/{key}][%d] deleteUserLoginLockoutNoContent  %+v", 204, o.Payload)
}
func (o *DeleteUserLoginLockoutNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteUserLoginLockoutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserLoginLockoutBadRequest creates a DeleteUserLoginLockoutBadRequest with default headers values
func NewDeleteUserLoginLockoutBadRequest() *DeleteUserLoginLockoutBadRequest {
	return &DeleteUserLoginLockoutBadRequest{}
}

/* DeleteUserLoginLockoutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserLoginLockoutBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserLoginLockoutBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user_login_lockout/{key}][%d] deleteUserLoginLockoutBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserLoginLockoutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserLoginLockoutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserLoginLockoutNotFound creates a DeleteUserLoginLockoutNotFound with default headers values
func NewDeleteUserLoginLockoutNotFound() *DeleteUserLoginLockoutNotFound {
	return &DeleteUserLoginLockoutNotFound{}
}

/* DeleteUserLoginLockoutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserLoginLockoutNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserLoginLockoutNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user_login_lockout/{key}][%d] deleteUserLoginLockoutNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserLoginLockoutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserLoginLockoutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserLoginLockoutTooManyRequests creates a DeleteUserLoginLockoutTooManyRequests with default headers values
func NewDeleteUserLoginLockoutTooManyRequests() *DeleteUserLoginLockoutTooManyRequests {
	return &DeleteUserLoginLockoutTooManyRequests{}
}

/* DeleteUserLoginLockoutTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteUserLoginLockoutTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteUserLoginLockoutTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /user_login_lockout/{key}][%d] deleteUserLoginLockoutTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteUserLoginLockoutTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserLoginLockoutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}