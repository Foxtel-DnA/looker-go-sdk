// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// DeleteUserSessionReader is a Reader for the DeleteUserSession structure.
type DeleteUserSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserSessionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserSessionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserSessionNoContent creates a DeleteUserSessionNoContent with default headers values
func NewDeleteUserSessionNoContent() *DeleteUserSessionNoContent {
	return &DeleteUserSessionNoContent{}
}

/* DeleteUserSessionNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteUserSessionNoContent struct {
	Payload string
}

func (o *DeleteUserSessionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/sessions/{session_id}][%d] deleteUserSessionNoContent  %+v", 204, o.Payload)
}
func (o *DeleteUserSessionNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteUserSessionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserSessionBadRequest creates a DeleteUserSessionBadRequest with default headers values
func NewDeleteUserSessionBadRequest() *DeleteUserSessionBadRequest {
	return &DeleteUserSessionBadRequest{}
}

/* DeleteUserSessionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserSessionBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserSessionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/sessions/{session_id}][%d] deleteUserSessionBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserSessionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserSessionNotFound creates a DeleteUserSessionNotFound with default headers values
func NewDeleteUserSessionNotFound() *DeleteUserSessionNotFound {
	return &DeleteUserSessionNotFound{}
}

/* DeleteUserSessionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserSessionNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserSessionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/sessions/{session_id}][%d] deleteUserSessionNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserSessionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserSessionTooManyRequests creates a DeleteUserSessionTooManyRequests with default headers values
func NewDeleteUserSessionTooManyRequests() *DeleteUserSessionTooManyRequests {
	return &DeleteUserSessionTooManyRequests{}
}

/* DeleteUserSessionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteUserSessionTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteUserSessionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/sessions/{session_id}][%d] deleteUserSessionTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteUserSessionTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserSessionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
