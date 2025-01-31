// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// LdapConfigReader is a Reader for the LdapConfig structure.
type LdapConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LdapConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLdapConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLdapConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLdapConfigOK creates a LdapConfigOK with default headers values
func NewLdapConfigOK() *LdapConfigOK {
	return &LdapConfigOK{}
}

/* LdapConfigOK describes a response with status code 200, with default header values.

LDAP Configuration.
*/
type LdapConfigOK struct {
	Payload *models.LDAPConfig
}

func (o *LdapConfigOK) Error() string {
	return fmt.Sprintf("[GET /ldap_config][%d] ldapConfigOK  %+v", 200, o.Payload)
}
func (o *LdapConfigOK) GetPayload() *models.LDAPConfig {
	return o.Payload
}

func (o *LdapConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LDAPConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLdapConfigNotFound creates a LdapConfigNotFound with default headers values
func NewLdapConfigNotFound() *LdapConfigNotFound {
	return &LdapConfigNotFound{}
}

/* LdapConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type LdapConfigNotFound struct {
	Payload *models.Error
}

func (o *LdapConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /ldap_config][%d] ldapConfigNotFound  %+v", 404, o.Payload)
}
func (o *LdapConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *LdapConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
