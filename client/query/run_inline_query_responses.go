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

// RunInlineQueryReader is a Reader for the RunInlineQuery structure.
type RunInlineQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunInlineQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunInlineQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunInlineQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunInlineQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRunInlineQueryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRunInlineQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRunInlineQueryOK creates a RunInlineQueryOK with default headers values
func NewRunInlineQueryOK() *RunInlineQueryOK {
	return &RunInlineQueryOK{}
}

/* RunInlineQueryOK describes a response with status code 200, with default header values.

Query Result
*/
type RunInlineQueryOK struct {
	Payload string
}

func (o *RunInlineQueryOK) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryOK  %+v", 200, o.Payload)
}
func (o *RunInlineQueryOK) GetPayload() string {
	return o.Payload
}

func (o *RunInlineQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunInlineQueryBadRequest creates a RunInlineQueryBadRequest with default headers values
func NewRunInlineQueryBadRequest() *RunInlineQueryBadRequest {
	return &RunInlineQueryBadRequest{}
}

/* RunInlineQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RunInlineQueryBadRequest struct {
	Payload *models.Error
}

func (o *RunInlineQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryBadRequest  %+v", 400, o.Payload)
}
func (o *RunInlineQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunInlineQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunInlineQueryNotFound creates a RunInlineQueryNotFound with default headers values
func NewRunInlineQueryNotFound() *RunInlineQueryNotFound {
	return &RunInlineQueryNotFound{}
}

/* RunInlineQueryNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RunInlineQueryNotFound struct {
	Payload *models.Error
}

func (o *RunInlineQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryNotFound  %+v", 404, o.Payload)
}
func (o *RunInlineQueryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunInlineQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunInlineQueryUnprocessableEntity creates a RunInlineQueryUnprocessableEntity with default headers values
func NewRunInlineQueryUnprocessableEntity() *RunInlineQueryUnprocessableEntity {
	return &RunInlineQueryUnprocessableEntity{}
}

/* RunInlineQueryUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type RunInlineQueryUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *RunInlineQueryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *RunInlineQueryUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RunInlineQueryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunInlineQueryTooManyRequests creates a RunInlineQueryTooManyRequests with default headers values
func NewRunInlineQueryTooManyRequests() *RunInlineQueryTooManyRequests {
	return &RunInlineQueryTooManyRequests{}
}

/* RunInlineQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RunInlineQueryTooManyRequests struct {
	Payload *models.Error
}

func (o *RunInlineQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /queries/run/{result_format}][%d] runInlineQueryTooManyRequests  %+v", 429, o.Payload)
}
func (o *RunInlineQueryTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunInlineQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
