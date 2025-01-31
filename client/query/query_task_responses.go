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

// QueryTaskReader is a Reader for the QueryTask structure.
type QueryTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryTaskOK creates a QueryTaskOK with default headers values
func NewQueryTaskOK() *QueryTaskOK {
	return &QueryTaskOK{}
}

/* QueryTaskOK describes a response with status code 200, with default header values.

query_task
*/
type QueryTaskOK struct {
	Payload *models.QueryTask
}

func (o *QueryTaskOK) Error() string {
	return fmt.Sprintf("[GET /query_tasks/{query_task_id}][%d] queryTaskOK  %+v", 200, o.Payload)
}
func (o *QueryTaskOK) GetPayload() *models.QueryTask {
	return o.Payload
}

func (o *QueryTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryTaskBadRequest creates a QueryTaskBadRequest with default headers values
func NewQueryTaskBadRequest() *QueryTaskBadRequest {
	return &QueryTaskBadRequest{}
}

/* QueryTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryTaskBadRequest struct {
	Payload *models.Error
}

func (o *QueryTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /query_tasks/{query_task_id}][%d] queryTaskBadRequest  %+v", 400, o.Payload)
}
func (o *QueryTaskBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *QueryTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryTaskNotFound creates a QueryTaskNotFound with default headers values
func NewQueryTaskNotFound() *QueryTaskNotFound {
	return &QueryTaskNotFound{}
}

/* QueryTaskNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryTaskNotFound struct {
	Payload *models.Error
}

func (o *QueryTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /query_tasks/{query_task_id}][%d] queryTaskNotFound  %+v", 404, o.Payload)
}
func (o *QueryTaskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *QueryTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
