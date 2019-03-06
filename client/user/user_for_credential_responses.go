// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// UserForCredentialReader is a Reader for the UserForCredential structure.
type UserForCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserForCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserForCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserForCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserForCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserForCredentialOK creates a UserForCredentialOK with default headers values
func NewUserForCredentialOK() *UserForCredentialOK {
	return &UserForCredentialOK{}
}

/*UserForCredentialOK handles this case with default header values.

Specified user.
*/
type UserForCredentialOK struct {
	Payload *models.User
}

func (o *UserForCredentialOK) Error() string {
	return fmt.Sprintf("[GET /users/credential/{credential_type}/{credential_id}][%d] userForCredentialOK  %+v", 200, o.Payload)
}

func (o *UserForCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserForCredentialBadRequest creates a UserForCredentialBadRequest with default headers values
func NewUserForCredentialBadRequest() *UserForCredentialBadRequest {
	return &UserForCredentialBadRequest{}
}

/*UserForCredentialBadRequest handles this case with default header values.

Bad Request
*/
type UserForCredentialBadRequest struct {
	Payload *models.Error
}

func (o *UserForCredentialBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/credential/{credential_type}/{credential_id}][%d] userForCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *UserForCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserForCredentialNotFound creates a UserForCredentialNotFound with default headers values
func NewUserForCredentialNotFound() *UserForCredentialNotFound {
	return &UserForCredentialNotFound{}
}

/*UserForCredentialNotFound handles this case with default header values.

Not Found
*/
type UserForCredentialNotFound struct {
	Payload *models.Error
}

func (o *UserForCredentialNotFound) Error() string {
	return fmt.Sprintf("[GET /users/credential/{credential_type}/{credential_id}][%d] userForCredentialNotFound  %+v", 404, o.Payload)
}

func (o *UserForCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
