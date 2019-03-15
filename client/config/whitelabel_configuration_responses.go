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

// WhitelabelConfigurationReader is a Reader for the WhitelabelConfiguration structure.
type WhitelabelConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WhitelabelConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWhitelabelConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewWhitelabelConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWhitelabelConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWhitelabelConfigurationOK creates a WhitelabelConfigurationOK with default headers values
func NewWhitelabelConfigurationOK() *WhitelabelConfigurationOK {
	return &WhitelabelConfigurationOK{}
}

/*WhitelabelConfigurationOK handles this case with default header values.

Whitelabel configuration
*/
type WhitelabelConfigurationOK struct {
	Payload *models.WhitelabelConfiguration
}

func (o *WhitelabelConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /whitelabel_configuration][%d] whitelabelConfigurationOK  %+v", 200, o.Payload)
}

func (o *WhitelabelConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WhitelabelConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWhitelabelConfigurationBadRequest creates a WhitelabelConfigurationBadRequest with default headers values
func NewWhitelabelConfigurationBadRequest() *WhitelabelConfigurationBadRequest {
	return &WhitelabelConfigurationBadRequest{}
}

/*WhitelabelConfigurationBadRequest handles this case with default header values.

Bad Request
*/
type WhitelabelConfigurationBadRequest struct {
	Payload *models.Error
}

func (o *WhitelabelConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /whitelabel_configuration][%d] whitelabelConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *WhitelabelConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWhitelabelConfigurationNotFound creates a WhitelabelConfigurationNotFound with default headers values
func NewWhitelabelConfigurationNotFound() *WhitelabelConfigurationNotFound {
	return &WhitelabelConfigurationNotFound{}
}

/*WhitelabelConfigurationNotFound handles this case with default header values.

Not Found
*/
type WhitelabelConfigurationNotFound struct {
	Payload *models.Error
}

func (o *WhitelabelConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /whitelabel_configuration][%d] whitelabelConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *WhitelabelConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}