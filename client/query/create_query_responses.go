// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// CreateQueryReader is a Reader for the CreateQuery structure.
type CreateQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateQueryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateQueryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateQueryOK creates a CreateQueryOK with default headers values
func NewCreateQueryOK() *CreateQueryOK {
	return &CreateQueryOK{}
}

/* CreateQueryOK describes a response with status code 200, with default header values.

Query
*/
type CreateQueryOK struct {
	Payload *models.Query
}

func (o *CreateQueryOK) Error() string {
	return fmt.Sprintf("[POST /queries][%d] createQueryOK  %+v", 200, o.Payload)
}
func (o *CreateQueryOK) GetPayload() *models.Query {
	return o.Payload
}

func (o *CreateQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Query)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryBadRequest creates a CreateQueryBadRequest with default headers values
func NewCreateQueryBadRequest() *CreateQueryBadRequest {
	return &CreateQueryBadRequest{}
}

/* CreateQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateQueryBadRequest struct {
	Payload *models.Error
}

func (o *CreateQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /queries][%d] createQueryBadRequest  %+v", 400, o.Payload)
}
func (o *CreateQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryNotFound creates a CreateQueryNotFound with default headers values
func NewCreateQueryNotFound() *CreateQueryNotFound {
	return &CreateQueryNotFound{}
}

/* CreateQueryNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateQueryNotFound struct {
	Payload *models.Error
}

func (o *CreateQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /queries][%d] createQueryNotFound  %+v", 404, o.Payload)
}
func (o *CreateQueryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryConflict creates a CreateQueryConflict with default headers values
func NewCreateQueryConflict() *CreateQueryConflict {
	return &CreateQueryConflict{}
}

/* CreateQueryConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateQueryConflict struct {
	Payload *models.Error
}

func (o *CreateQueryConflict) Error() string {
	return fmt.Sprintf("[POST /queries][%d] createQueryConflict  %+v", 409, o.Payload)
}
func (o *CreateQueryConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryUnprocessableEntity creates a CreateQueryUnprocessableEntity with default headers values
func NewCreateQueryUnprocessableEntity() *CreateQueryUnprocessableEntity {
	return &CreateQueryUnprocessableEntity{}
}

/* CreateQueryUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateQueryUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateQueryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /queries][%d] createQueryUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateQueryUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateQueryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryTooManyRequests creates a CreateQueryTooManyRequests with default headers values
func NewCreateQueryTooManyRequests() *CreateQueryTooManyRequests {
	return &CreateQueryTooManyRequests{}
}

/* CreateQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateQueryTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /queries][%d] createQueryTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateQueryTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
