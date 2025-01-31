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

// DeleteUserCredentialsGoogleReader is a Reader for the DeleteUserCredentialsGoogle structure.
type DeleteUserCredentialsGoogleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserCredentialsGoogleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserCredentialsGoogleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserCredentialsGoogleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserCredentialsGoogleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserCredentialsGoogleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserCredentialsGoogleNoContent creates a DeleteUserCredentialsGoogleNoContent with default headers values
func NewDeleteUserCredentialsGoogleNoContent() *DeleteUserCredentialsGoogleNoContent {
	return &DeleteUserCredentialsGoogleNoContent{}
}

/* DeleteUserCredentialsGoogleNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsGoogleNoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsGoogleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_google][%d] deleteUserCredentialsGoogleNoContent  %+v", 204, o.Payload)
}
func (o *DeleteUserCredentialsGoogleNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteUserCredentialsGoogleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsGoogleBadRequest creates a DeleteUserCredentialsGoogleBadRequest with default headers values
func NewDeleteUserCredentialsGoogleBadRequest() *DeleteUserCredentialsGoogleBadRequest {
	return &DeleteUserCredentialsGoogleBadRequest{}
}

/* DeleteUserCredentialsGoogleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserCredentialsGoogleBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsGoogleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_google][%d] deleteUserCredentialsGoogleBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserCredentialsGoogleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsGoogleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsGoogleNotFound creates a DeleteUserCredentialsGoogleNotFound with default headers values
func NewDeleteUserCredentialsGoogleNotFound() *DeleteUserCredentialsGoogleNotFound {
	return &DeleteUserCredentialsGoogleNotFound{}
}

/* DeleteUserCredentialsGoogleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserCredentialsGoogleNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsGoogleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_google][%d] deleteUserCredentialsGoogleNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserCredentialsGoogleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsGoogleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsGoogleTooManyRequests creates a DeleteUserCredentialsGoogleTooManyRequests with default headers values
func NewDeleteUserCredentialsGoogleTooManyRequests() *DeleteUserCredentialsGoogleTooManyRequests {
	return &DeleteUserCredentialsGoogleTooManyRequests{}
}

/* DeleteUserCredentialsGoogleTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteUserCredentialsGoogleTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsGoogleTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_google][%d] deleteUserCredentialsGoogleTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteUserCredentialsGoogleTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsGoogleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
