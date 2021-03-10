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

// UpdateLdapConfigReader is a Reader for the UpdateLdapConfig structure.
type UpdateLdapConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLdapConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLdapConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLdapConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLdapConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateLdapConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLdapConfigOK creates a UpdateLdapConfigOK with default headers values
func NewUpdateLdapConfigOK() *UpdateLdapConfigOK {
	return &UpdateLdapConfigOK{}
}

/* UpdateLdapConfigOK describes a response with status code 200, with default header values.

New state for LDAP Configuration.
*/
type UpdateLdapConfigOK struct {
	Payload *models.LDAPConfig
}

func (o *UpdateLdapConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /ldap_config][%d] updateLdapConfigOK  %+v", 200, o.Payload)
}
func (o *UpdateLdapConfigOK) GetPayload() *models.LDAPConfig {
	return o.Payload
}

func (o *UpdateLdapConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LDAPConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigBadRequest creates a UpdateLdapConfigBadRequest with default headers values
func NewUpdateLdapConfigBadRequest() *UpdateLdapConfigBadRequest {
	return &UpdateLdapConfigBadRequest{}
}

/* UpdateLdapConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateLdapConfigBadRequest struct {
	Payload *models.Error
}

func (o *UpdateLdapConfigBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /ldap_config][%d] updateLdapConfigBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateLdapConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLdapConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigNotFound creates a UpdateLdapConfigNotFound with default headers values
func NewUpdateLdapConfigNotFound() *UpdateLdapConfigNotFound {
	return &UpdateLdapConfigNotFound{}
}

/* UpdateLdapConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateLdapConfigNotFound struct {
	Payload *models.Error
}

func (o *UpdateLdapConfigNotFound) Error() string {
	return fmt.Sprintf("[PATCH /ldap_config][%d] updateLdapConfigNotFound  %+v", 404, o.Payload)
}
func (o *UpdateLdapConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateLdapConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigUnprocessableEntity creates a UpdateLdapConfigUnprocessableEntity with default headers values
func NewUpdateLdapConfigUnprocessableEntity() *UpdateLdapConfigUnprocessableEntity {
	return &UpdateLdapConfigUnprocessableEntity{}
}

/* UpdateLdapConfigUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateLdapConfigUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateLdapConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /ldap_config][%d] updateLdapConfigUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateLdapConfigUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateLdapConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
