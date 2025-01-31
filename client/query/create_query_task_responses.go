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

// CreateQueryTaskReader is a Reader for the CreateQueryTask structure.
type CreateQueryTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueryTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateQueryTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateQueryTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateQueryTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateQueryTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateQueryTaskUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateQueryTaskTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateQueryTaskOK creates a CreateQueryTaskOK with default headers values
func NewCreateQueryTaskOK() *CreateQueryTaskOK {
	return &CreateQueryTaskOK{}
}

/* CreateQueryTaskOK describes a response with status code 200, with default header values.

query_task
*/
type CreateQueryTaskOK struct {
	Payload *models.QueryTask
}

func (o *CreateQueryTaskOK) Error() string {
	return fmt.Sprintf("[POST /query_tasks][%d] createQueryTaskOK  %+v", 200, o.Payload)
}
func (o *CreateQueryTaskOK) GetPayload() *models.QueryTask {
	return o.Payload
}

func (o *CreateQueryTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryTaskBadRequest creates a CreateQueryTaskBadRequest with default headers values
func NewCreateQueryTaskBadRequest() *CreateQueryTaskBadRequest {
	return &CreateQueryTaskBadRequest{}
}

/* CreateQueryTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateQueryTaskBadRequest struct {
	Payload *models.Error
}

func (o *CreateQueryTaskBadRequest) Error() string {
	return fmt.Sprintf("[POST /query_tasks][%d] createQueryTaskBadRequest  %+v", 400, o.Payload)
}
func (o *CreateQueryTaskBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryTaskNotFound creates a CreateQueryTaskNotFound with default headers values
func NewCreateQueryTaskNotFound() *CreateQueryTaskNotFound {
	return &CreateQueryTaskNotFound{}
}

/* CreateQueryTaskNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateQueryTaskNotFound struct {
	Payload *models.Error
}

func (o *CreateQueryTaskNotFound) Error() string {
	return fmt.Sprintf("[POST /query_tasks][%d] createQueryTaskNotFound  %+v", 404, o.Payload)
}
func (o *CreateQueryTaskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryTaskConflict creates a CreateQueryTaskConflict with default headers values
func NewCreateQueryTaskConflict() *CreateQueryTaskConflict {
	return &CreateQueryTaskConflict{}
}

/* CreateQueryTaskConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateQueryTaskConflict struct {
	Payload *models.Error
}

func (o *CreateQueryTaskConflict) Error() string {
	return fmt.Sprintf("[POST /query_tasks][%d] createQueryTaskConflict  %+v", 409, o.Payload)
}
func (o *CreateQueryTaskConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryTaskUnprocessableEntity creates a CreateQueryTaskUnprocessableEntity with default headers values
func NewCreateQueryTaskUnprocessableEntity() *CreateQueryTaskUnprocessableEntity {
	return &CreateQueryTaskUnprocessableEntity{}
}

/* CreateQueryTaskUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateQueryTaskUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateQueryTaskUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /query_tasks][%d] createQueryTaskUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateQueryTaskUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateQueryTaskUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryTaskTooManyRequests creates a CreateQueryTaskTooManyRequests with default headers values
func NewCreateQueryTaskTooManyRequests() *CreateQueryTaskTooManyRequests {
	return &CreateQueryTaskTooManyRequests{}
}

/* CreateQueryTaskTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateQueryTaskTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateQueryTaskTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /query_tasks][%d] createQueryTaskTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateQueryTaskTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateQueryTaskTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
