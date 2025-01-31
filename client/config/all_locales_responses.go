// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// AllLocalesReader is a Reader for the AllLocales structure.
type AllLocalesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllLocalesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllLocalesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllLocalesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllLocalesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllLocalesOK creates a AllLocalesOK with default headers values
func NewAllLocalesOK() *AllLocalesOK {
	return &AllLocalesOK{}
}

/* AllLocalesOK describes a response with status code 200, with default header values.

Locale
*/
type AllLocalesOK struct {
	Payload []*models.Locale
}

func (o *AllLocalesOK) Error() string {
	return fmt.Sprintf("[GET /locales][%d] allLocalesOK  %+v", 200, o.Payload)
}
func (o *AllLocalesOK) GetPayload() []*models.Locale {
	return o.Payload
}

func (o *AllLocalesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllLocalesBadRequest creates a AllLocalesBadRequest with default headers values
func NewAllLocalesBadRequest() *AllLocalesBadRequest {
	return &AllLocalesBadRequest{}
}

/* AllLocalesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllLocalesBadRequest struct {
	Payload *models.Error
}

func (o *AllLocalesBadRequest) Error() string {
	return fmt.Sprintf("[GET /locales][%d] allLocalesBadRequest  %+v", 400, o.Payload)
}
func (o *AllLocalesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllLocalesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllLocalesNotFound creates a AllLocalesNotFound with default headers values
func NewAllLocalesNotFound() *AllLocalesNotFound {
	return &AllLocalesNotFound{}
}

/* AllLocalesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllLocalesNotFound struct {
	Payload *models.Error
}

func (o *AllLocalesNotFound) Error() string {
	return fmt.Sprintf("[GET /locales][%d] allLocalesNotFound  %+v", 404, o.Payload)
}
func (o *AllLocalesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllLocalesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
