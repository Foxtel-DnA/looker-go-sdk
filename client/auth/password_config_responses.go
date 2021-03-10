// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// PasswordConfigReader is a Reader for the PasswordConfig structure.
type PasswordConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PasswordConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPasswordConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPasswordConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPasswordConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPasswordConfigOK creates a PasswordConfigOK with default headers values
func NewPasswordConfigOK() *PasswordConfigOK {
	return &PasswordConfigOK{}
}

/* PasswordConfigOK describes a response with status code 200, with default header values.

Password Config
*/
type PasswordConfigOK struct {
	Payload *models.PasswordConfig
}

func (o *PasswordConfigOK) Error() string {
	return fmt.Sprintf("[GET /password_config][%d] passwordConfigOK  %+v", 200, o.Payload)
}
func (o *PasswordConfigOK) GetPayload() *models.PasswordConfig {
	return o.Payload
}

func (o *PasswordConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PasswordConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPasswordConfigBadRequest creates a PasswordConfigBadRequest with default headers values
func NewPasswordConfigBadRequest() *PasswordConfigBadRequest {
	return &PasswordConfigBadRequest{}
}

/* PasswordConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PasswordConfigBadRequest struct {
	Payload *models.Error
}

func (o *PasswordConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /password_config][%d] passwordConfigBadRequest  %+v", 400, o.Payload)
}
func (o *PasswordConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PasswordConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPasswordConfigNotFound creates a PasswordConfigNotFound with default headers values
func NewPasswordConfigNotFound() *PasswordConfigNotFound {
	return &PasswordConfigNotFound{}
}

/* PasswordConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PasswordConfigNotFound struct {
	Payload *models.Error
}

func (o *PasswordConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /password_config][%d] passwordConfigNotFound  %+v", 404, o.Payload)
}
func (o *PasswordConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PasswordConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
