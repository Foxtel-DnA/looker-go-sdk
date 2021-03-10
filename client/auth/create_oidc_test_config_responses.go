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

// CreateOidcTestConfigReader is a Reader for the CreateOidcTestConfig structure.
type CreateOidcTestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOidcTestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOidcTestConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOidcTestConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOidcTestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateOidcTestConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOidcTestConfigOK creates a CreateOidcTestConfigOK with default headers values
func NewCreateOidcTestConfigOK() *CreateOidcTestConfigOK {
	return &CreateOidcTestConfigOK{}
}

/* CreateOidcTestConfigOK describes a response with status code 200, with default header values.

OIDC test config
*/
type CreateOidcTestConfigOK struct {
	Payload *models.OIDCConfig
}

func (o *CreateOidcTestConfigOK) Error() string {
	return fmt.Sprintf("[POST /oidc_test_configs][%d] createOidcTestConfigOK  %+v", 200, o.Payload)
}
func (o *CreateOidcTestConfigOK) GetPayload() *models.OIDCConfig {
	return o.Payload
}

func (o *CreateOidcTestConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OIDCConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOidcTestConfigBadRequest creates a CreateOidcTestConfigBadRequest with default headers values
func NewCreateOidcTestConfigBadRequest() *CreateOidcTestConfigBadRequest {
	return &CreateOidcTestConfigBadRequest{}
}

/* CreateOidcTestConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateOidcTestConfigBadRequest struct {
	Payload *models.Error
}

func (o *CreateOidcTestConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /oidc_test_configs][%d] createOidcTestConfigBadRequest  %+v", 400, o.Payload)
}
func (o *CreateOidcTestConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateOidcTestConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOidcTestConfigNotFound creates a CreateOidcTestConfigNotFound with default headers values
func NewCreateOidcTestConfigNotFound() *CreateOidcTestConfigNotFound {
	return &CreateOidcTestConfigNotFound{}
}

/* CreateOidcTestConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateOidcTestConfigNotFound struct {
	Payload *models.Error
}

func (o *CreateOidcTestConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /oidc_test_configs][%d] createOidcTestConfigNotFound  %+v", 404, o.Payload)
}
func (o *CreateOidcTestConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateOidcTestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOidcTestConfigUnprocessableEntity creates a CreateOidcTestConfigUnprocessableEntity with default headers values
func NewCreateOidcTestConfigUnprocessableEntity() *CreateOidcTestConfigUnprocessableEntity {
	return &CreateOidcTestConfigUnprocessableEntity{}
}

/* CreateOidcTestConfigUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateOidcTestConfigUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateOidcTestConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /oidc_test_configs][%d] createOidcTestConfigUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateOidcTestConfigUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateOidcTestConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
