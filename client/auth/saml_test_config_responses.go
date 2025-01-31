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

// SamlTestConfigReader is a Reader for the SamlTestConfig structure.
type SamlTestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SamlTestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSamlTestConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSamlTestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSamlTestConfigOK creates a SamlTestConfigOK with default headers values
func NewSamlTestConfigOK() *SamlTestConfigOK {
	return &SamlTestConfigOK{}
}

/* SamlTestConfigOK describes a response with status code 200, with default header values.

SAML test config.
*/
type SamlTestConfigOK struct {
	Payload *models.SamlConfig
}

func (o *SamlTestConfigOK) Error() string {
	return fmt.Sprintf("[GET /saml_test_configs/{test_slug}][%d] samlTestConfigOK  %+v", 200, o.Payload)
}
func (o *SamlTestConfigOK) GetPayload() *models.SamlConfig {
	return o.Payload
}

func (o *SamlTestConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SamlConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSamlTestConfigNotFound creates a SamlTestConfigNotFound with default headers values
func NewSamlTestConfigNotFound() *SamlTestConfigNotFound {
	return &SamlTestConfigNotFound{}
}

/* SamlTestConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SamlTestConfigNotFound struct {
	Payload *models.Error
}

func (o *SamlTestConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /saml_test_configs/{test_slug}][%d] samlTestConfigNotFound  %+v", 404, o.Payload)
}
func (o *SamlTestConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SamlTestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
