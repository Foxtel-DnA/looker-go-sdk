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

// AllRunningQueriesReader is a Reader for the AllRunningQueries structure.
type AllRunningQueriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllRunningQueriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllRunningQueriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAllRunningQueriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllRunningQueriesOK creates a AllRunningQueriesOK with default headers values
func NewAllRunningQueriesOK() *AllRunningQueriesOK {
	return &AllRunningQueriesOK{}
}

/* AllRunningQueriesOK describes a response with status code 200, with default header values.

Running Queries.
*/
type AllRunningQueriesOK struct {
	Payload []*models.RunningQueries
}

func (o *AllRunningQueriesOK) Error() string {
	return fmt.Sprintf("[GET /running_queries][%d] allRunningQueriesOK  %+v", 200, o.Payload)
}
func (o *AllRunningQueriesOK) GetPayload() []*models.RunningQueries {
	return o.Payload
}

func (o *AllRunningQueriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllRunningQueriesNotFound creates a AllRunningQueriesNotFound with default headers values
func NewAllRunningQueriesNotFound() *AllRunningQueriesNotFound {
	return &AllRunningQueriesNotFound{}
}

/* AllRunningQueriesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllRunningQueriesNotFound struct {
	Payload *models.Error
}

func (o *AllRunningQueriesNotFound) Error() string {
	return fmt.Sprintf("[GET /running_queries][%d] allRunningQueriesNotFound  %+v", 404, o.Payload)
}
func (o *AllRunningQueriesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllRunningQueriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
