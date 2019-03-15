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

// TestLdapConfigUserAuthReader is a Reader for the TestLdapConfigUserAuth structure.
type TestLdapConfigUserAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestLdapConfigUserAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTestLdapConfigUserAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTestLdapConfigUserAuthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTestLdapConfigUserAuthNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewTestLdapConfigUserAuthUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestLdapConfigUserAuthOK creates a TestLdapConfigUserAuthOK with default headers values
func NewTestLdapConfigUserAuthOK() *TestLdapConfigUserAuthOK {
	return &TestLdapConfigUserAuthOK{}
}

/*TestLdapConfigUserAuthOK handles this case with default header values.

Result info.
*/
type TestLdapConfigUserAuthOK struct {
	Payload *models.LDAPConfigTestResult
}

func (o *TestLdapConfigUserAuthOK) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_auth][%d] testLdapConfigUserAuthOK  %+v", 200, o.Payload)
}

func (o *TestLdapConfigUserAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LDAPConfigTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigUserAuthBadRequest creates a TestLdapConfigUserAuthBadRequest with default headers values
func NewTestLdapConfigUserAuthBadRequest() *TestLdapConfigUserAuthBadRequest {
	return &TestLdapConfigUserAuthBadRequest{}
}

/*TestLdapConfigUserAuthBadRequest handles this case with default header values.

Bad Request
*/
type TestLdapConfigUserAuthBadRequest struct {
	Payload *models.Error
}

func (o *TestLdapConfigUserAuthBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_auth][%d] testLdapConfigUserAuthBadRequest  %+v", 400, o.Payload)
}

func (o *TestLdapConfigUserAuthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigUserAuthNotFound creates a TestLdapConfigUserAuthNotFound with default headers values
func NewTestLdapConfigUserAuthNotFound() *TestLdapConfigUserAuthNotFound {
	return &TestLdapConfigUserAuthNotFound{}
}

/*TestLdapConfigUserAuthNotFound handles this case with default header values.

Not Found
*/
type TestLdapConfigUserAuthNotFound struct {
	Payload *models.Error
}

func (o *TestLdapConfigUserAuthNotFound) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_auth][%d] testLdapConfigUserAuthNotFound  %+v", 404, o.Payload)
}

func (o *TestLdapConfigUserAuthNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigUserAuthUnprocessableEntity creates a TestLdapConfigUserAuthUnprocessableEntity with default headers values
func NewTestLdapConfigUserAuthUnprocessableEntity() *TestLdapConfigUserAuthUnprocessableEntity {
	return &TestLdapConfigUserAuthUnprocessableEntity{}
}

/*TestLdapConfigUserAuthUnprocessableEntity handles this case with default header values.

Validation Error
*/
type TestLdapConfigUserAuthUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *TestLdapConfigUserAuthUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_auth][%d] testLdapConfigUserAuthUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TestLdapConfigUserAuthUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}