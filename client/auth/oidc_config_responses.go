// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// OidcConfigReader is a Reader for the OidcConfig structure.
type OidcConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OidcConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOidcConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewOidcConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOidcConfigOK creates a OidcConfigOK with default headers values
func NewOidcConfigOK() *OidcConfigOK {
	return &OidcConfigOK{}
}

/*OidcConfigOK handles this case with default header values.

OIDC Configuration.
*/
type OidcConfigOK struct {
	Payload *models.OIDCConfig
}

func (o *OidcConfigOK) Error() string {
	return fmt.Sprintf("[GET /oidc_config][%d] oidcConfigOK  %+v", 200, o.Payload)
}

func (o *OidcConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OIDCConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOidcConfigNotFound creates a OidcConfigNotFound with default headers values
func NewOidcConfigNotFound() *OidcConfigNotFound {
	return &OidcConfigNotFound{}
}

/*OidcConfigNotFound handles this case with default header values.

Not Found
*/
type OidcConfigNotFound struct {
	Payload *models.Error
}

func (o *OidcConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /oidc_config][%d] oidcConfigNotFound  %+v", 404, o.Payload)
}

func (o *OidcConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
