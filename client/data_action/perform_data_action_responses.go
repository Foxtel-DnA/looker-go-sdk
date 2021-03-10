// Code generated by go-swagger; DO NOT EDIT.

package data_action

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// PerformDataActionReader is a Reader for the PerformDataAction structure.
type PerformDataActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformDataActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformDataActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformDataActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformDataActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPerformDataActionOK creates a PerformDataActionOK with default headers values
func NewPerformDataActionOK() *PerformDataActionOK {
	return &PerformDataActionOK{}
}

/* PerformDataActionOK describes a response with status code 200, with default header values.

Data Action Response
*/
type PerformDataActionOK struct {
	Payload *models.DataActionResponse
}

func (o *PerformDataActionOK) Error() string {
	return fmt.Sprintf("[POST /data_actions][%d] performDataActionOK  %+v", 200, o.Payload)
}
func (o *PerformDataActionOK) GetPayload() *models.DataActionResponse {
	return o.Payload
}

func (o *PerformDataActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataActionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDataActionBadRequest creates a PerformDataActionBadRequest with default headers values
func NewPerformDataActionBadRequest() *PerformDataActionBadRequest {
	return &PerformDataActionBadRequest{}
}

/* PerformDataActionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PerformDataActionBadRequest struct {
	Payload *models.Error
}

func (o *PerformDataActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /data_actions][%d] performDataActionBadRequest  %+v", 400, o.Payload)
}
func (o *PerformDataActionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PerformDataActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDataActionNotFound creates a PerformDataActionNotFound with default headers values
func NewPerformDataActionNotFound() *PerformDataActionNotFound {
	return &PerformDataActionNotFound{}
}

/* PerformDataActionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PerformDataActionNotFound struct {
	Payload *models.Error
}

func (o *PerformDataActionNotFound) Error() string {
	return fmt.Sprintf("[POST /data_actions][%d] performDataActionNotFound  %+v", 404, o.Payload)
}
func (o *PerformDataActionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PerformDataActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
