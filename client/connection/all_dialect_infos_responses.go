// Code generated by go-swagger; DO NOT EDIT.

package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// AllDialectInfosReader is a Reader for the AllDialectInfos structure.
type AllDialectInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllDialectInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllDialectInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllDialectInfosBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllDialectInfosNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllDialectInfosOK creates a AllDialectInfosOK with default headers values
func NewAllDialectInfosOK() *AllDialectInfosOK {
	return &AllDialectInfosOK{}
}

/* AllDialectInfosOK describes a response with status code 200, with default header values.

Dialect Info
*/
type AllDialectInfosOK struct {
	Payload []*models.DialectInfo
}

func (o *AllDialectInfosOK) Error() string {
	return fmt.Sprintf("[GET /dialect_info][%d] allDialectInfosOK  %+v", 200, o.Payload)
}
func (o *AllDialectInfosOK) GetPayload() []*models.DialectInfo {
	return o.Payload
}

func (o *AllDialectInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllDialectInfosBadRequest creates a AllDialectInfosBadRequest with default headers values
func NewAllDialectInfosBadRequest() *AllDialectInfosBadRequest {
	return &AllDialectInfosBadRequest{}
}

/* AllDialectInfosBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllDialectInfosBadRequest struct {
	Payload *models.Error
}

func (o *AllDialectInfosBadRequest) Error() string {
	return fmt.Sprintf("[GET /dialect_info][%d] allDialectInfosBadRequest  %+v", 400, o.Payload)
}
func (o *AllDialectInfosBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllDialectInfosBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllDialectInfosNotFound creates a AllDialectInfosNotFound with default headers values
func NewAllDialectInfosNotFound() *AllDialectInfosNotFound {
	return &AllDialectInfosNotFound{}
}

/* AllDialectInfosNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllDialectInfosNotFound struct {
	Payload *models.Error
}

func (o *AllDialectInfosNotFound) Error() string {
	return fmt.Sprintf("[GET /dialect_info][%d] allDialectInfosNotFound  %+v", 404, o.Payload)
}
func (o *AllDialectInfosNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllDialectInfosNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
