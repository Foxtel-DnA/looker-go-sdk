// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// DeleteUserCredentialsLdapReader is a Reader for the DeleteUserCredentialsLdap structure.
type DeleteUserCredentialsLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserCredentialsLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserCredentialsLdapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserCredentialsLdapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserCredentialsLdapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserCredentialsLdapTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserCredentialsLdapNoContent creates a DeleteUserCredentialsLdapNoContent with default headers values
func NewDeleteUserCredentialsLdapNoContent() *DeleteUserCredentialsLdapNoContent {
	return &DeleteUserCredentialsLdapNoContent{}
}

/* DeleteUserCredentialsLdapNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsLdapNoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsLdapNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_ldap][%d] deleteUserCredentialsLdapNoContent  %+v", 204, o.Payload)
}
func (o *DeleteUserCredentialsLdapNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteUserCredentialsLdapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsLdapBadRequest creates a DeleteUserCredentialsLdapBadRequest with default headers values
func NewDeleteUserCredentialsLdapBadRequest() *DeleteUserCredentialsLdapBadRequest {
	return &DeleteUserCredentialsLdapBadRequest{}
}

/* DeleteUserCredentialsLdapBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserCredentialsLdapBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsLdapBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_ldap][%d] deleteUserCredentialsLdapBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserCredentialsLdapBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsLdapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsLdapNotFound creates a DeleteUserCredentialsLdapNotFound with default headers values
func NewDeleteUserCredentialsLdapNotFound() *DeleteUserCredentialsLdapNotFound {
	return &DeleteUserCredentialsLdapNotFound{}
}

/* DeleteUserCredentialsLdapNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserCredentialsLdapNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsLdapNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_ldap][%d] deleteUserCredentialsLdapNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserCredentialsLdapNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsLdapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsLdapTooManyRequests creates a DeleteUserCredentialsLdapTooManyRequests with default headers values
func NewDeleteUserCredentialsLdapTooManyRequests() *DeleteUserCredentialsLdapTooManyRequests {
	return &DeleteUserCredentialsLdapTooManyRequests{}
}

/* DeleteUserCredentialsLdapTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteUserCredentialsLdapTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsLdapTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_ldap][%d] deleteUserCredentialsLdapTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteUserCredentialsLdapTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsLdapTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
