// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// AllTimezonesReader is a Reader for the AllTimezones structure.
type AllTimezonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllTimezonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllTimezonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllTimezonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllTimezonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllTimezonesOK creates a AllTimezonesOK with default headers values
func NewAllTimezonesOK() *AllTimezonesOK {
	return &AllTimezonesOK{}
}

/*AllTimezonesOK handles this case with default header values.

Timezone
*/
type AllTimezonesOK struct {
	Payload []*models.Timezone
}

func (o *AllTimezonesOK) Error() string {
	return fmt.Sprintf("[GET /timezones][%d] allTimezonesOK  %+v", 200, o.Payload)
}

func (o *AllTimezonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllTimezonesBadRequest creates a AllTimezonesBadRequest with default headers values
func NewAllTimezonesBadRequest() *AllTimezonesBadRequest {
	return &AllTimezonesBadRequest{}
}

/*AllTimezonesBadRequest handles this case with default header values.

Bad Request
*/
type AllTimezonesBadRequest struct {
	Payload *models.Error
}

func (o *AllTimezonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /timezones][%d] allTimezonesBadRequest  %+v", 400, o.Payload)
}

func (o *AllTimezonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllTimezonesNotFound creates a AllTimezonesNotFound with default headers values
func NewAllTimezonesNotFound() *AllTimezonesNotFound {
	return &AllTimezonesNotFound{}
}

/*AllTimezonesNotFound handles this case with default header values.

Not Found
*/
type AllTimezonesNotFound struct {
	Payload *models.Error
}

func (o *AllTimezonesNotFound) Error() string {
	return fmt.Sprintf("[GET /timezones][%d] allTimezonesNotFound  %+v", 404, o.Payload)
}

func (o *AllTimezonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
