// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// CreateSpaceReader is a Reader for the CreateSpace structure.
type CreateSpaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSpaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSpaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateSpaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateSpaceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateSpaceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateSpaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSpaceOK creates a CreateSpaceOK with default headers values
func NewCreateSpaceOK() *CreateSpaceOK {
	return &CreateSpaceOK{}
}

/* CreateSpaceOK describes a response with status code 200, with default header values.

Space
*/
type CreateSpaceOK struct {
	Payload *models.Space
}

func (o *CreateSpaceOK) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceOK  %+v", 200, o.Payload)
}
func (o *CreateSpaceOK) GetPayload() *models.Space {
	return o.Payload
}

func (o *CreateSpaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Space)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpaceBadRequest creates a CreateSpaceBadRequest with default headers values
func NewCreateSpaceBadRequest() *CreateSpaceBadRequest {
	return &CreateSpaceBadRequest{}
}

/* CreateSpaceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateSpaceBadRequest struct {
	Payload *models.Error
}

func (o *CreateSpaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSpaceBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpaceNotFound creates a CreateSpaceNotFound with default headers values
func NewCreateSpaceNotFound() *CreateSpaceNotFound {
	return &CreateSpaceNotFound{}
}

/* CreateSpaceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateSpaceNotFound struct {
	Payload *models.Error
}

func (o *CreateSpaceNotFound) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceNotFound  %+v", 404, o.Payload)
}
func (o *CreateSpaceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpaceConflict creates a CreateSpaceConflict with default headers values
func NewCreateSpaceConflict() *CreateSpaceConflict {
	return &CreateSpaceConflict{}
}

/* CreateSpaceConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateSpaceConflict struct {
	Payload *models.Error
}

func (o *CreateSpaceConflict) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceConflict  %+v", 409, o.Payload)
}
func (o *CreateSpaceConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpaceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpaceUnprocessableEntity creates a CreateSpaceUnprocessableEntity with default headers values
func NewCreateSpaceUnprocessableEntity() *CreateSpaceUnprocessableEntity {
	return &CreateSpaceUnprocessableEntity{}
}

/* CreateSpaceUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateSpaceUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateSpaceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateSpaceUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateSpaceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpaceTooManyRequests creates a CreateSpaceTooManyRequests with default headers values
func NewCreateSpaceTooManyRequests() *CreateSpaceTooManyRequests {
	return &CreateSpaceTooManyRequests{}
}

/* CreateSpaceTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateSpaceTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateSpaceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /spaces][%d] createSpaceTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateSpaceTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
