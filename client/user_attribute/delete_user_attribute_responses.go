// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// DeleteUserAttributeReader is a Reader for the DeleteUserAttribute structure.
type DeleteUserAttributeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserAttributeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserAttributeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserAttributeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserAttributeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserAttributeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserAttributeNoContent creates a DeleteUserAttributeNoContent with default headers values
func NewDeleteUserAttributeNoContent() *DeleteUserAttributeNoContent {
	return &DeleteUserAttributeNoContent{}
}

/* DeleteUserAttributeNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteUserAttributeNoContent struct {
	Payload string
}

func (o *DeleteUserAttributeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user_attributes/{user_attribute_id}][%d] deleteUserAttributeNoContent  %+v", 204, o.Payload)
}
func (o *DeleteUserAttributeNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteUserAttributeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserAttributeBadRequest creates a DeleteUserAttributeBadRequest with default headers values
func NewDeleteUserAttributeBadRequest() *DeleteUserAttributeBadRequest {
	return &DeleteUserAttributeBadRequest{}
}

/* DeleteUserAttributeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserAttributeBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserAttributeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user_attributes/{user_attribute_id}][%d] deleteUserAttributeBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserAttributeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserAttributeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserAttributeNotFound creates a DeleteUserAttributeNotFound with default headers values
func NewDeleteUserAttributeNotFound() *DeleteUserAttributeNotFound {
	return &DeleteUserAttributeNotFound{}
}

/* DeleteUserAttributeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserAttributeNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserAttributeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user_attributes/{user_attribute_id}][%d] deleteUserAttributeNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserAttributeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserAttributeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserAttributeTooManyRequests creates a DeleteUserAttributeTooManyRequests with default headers values
func NewDeleteUserAttributeTooManyRequests() *DeleteUserAttributeTooManyRequests {
	return &DeleteUserAttributeTooManyRequests{}
}

/* DeleteUserAttributeTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteUserAttributeTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteUserAttributeTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /user_attributes/{user_attribute_id}][%d] deleteUserAttributeTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteUserAttributeTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserAttributeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
