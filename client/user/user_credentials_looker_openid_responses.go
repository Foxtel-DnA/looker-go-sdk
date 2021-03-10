// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// UserCredentialsLookerOpenidReader is a Reader for the UserCredentialsLookerOpenid structure.
type UserCredentialsLookerOpenidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCredentialsLookerOpenidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserCredentialsLookerOpenidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserCredentialsLookerOpenidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserCredentialsLookerOpenidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserCredentialsLookerOpenidOK creates a UserCredentialsLookerOpenidOK with default headers values
func NewUserCredentialsLookerOpenidOK() *UserCredentialsLookerOpenidOK {
	return &UserCredentialsLookerOpenidOK{}
}

/* UserCredentialsLookerOpenidOK describes a response with status code 200, with default header values.

Looker OpenId Credential
*/
type UserCredentialsLookerOpenidOK struct {
	Payload *models.CredentialsLookerOpenid
}

func (o *UserCredentialsLookerOpenidOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_looker_openid][%d] userCredentialsLookerOpenidOK  %+v", 200, o.Payload)
}
func (o *UserCredentialsLookerOpenidOK) GetPayload() *models.CredentialsLookerOpenid {
	return o.Payload
}

func (o *UserCredentialsLookerOpenidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsLookerOpenid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsLookerOpenidBadRequest creates a UserCredentialsLookerOpenidBadRequest with default headers values
func NewUserCredentialsLookerOpenidBadRequest() *UserCredentialsLookerOpenidBadRequest {
	return &UserCredentialsLookerOpenidBadRequest{}
}

/* UserCredentialsLookerOpenidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserCredentialsLookerOpenidBadRequest struct {
	Payload *models.Error
}

func (o *UserCredentialsLookerOpenidBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_looker_openid][%d] userCredentialsLookerOpenidBadRequest  %+v", 400, o.Payload)
}
func (o *UserCredentialsLookerOpenidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserCredentialsLookerOpenidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsLookerOpenidNotFound creates a UserCredentialsLookerOpenidNotFound with default headers values
func NewUserCredentialsLookerOpenidNotFound() *UserCredentialsLookerOpenidNotFound {
	return &UserCredentialsLookerOpenidNotFound{}
}

/* UserCredentialsLookerOpenidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserCredentialsLookerOpenidNotFound struct {
	Payload *models.Error
}

func (o *UserCredentialsLookerOpenidNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_looker_openid][%d] userCredentialsLookerOpenidNotFound  %+v", 404, o.Payload)
}
func (o *UserCredentialsLookerOpenidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserCredentialsLookerOpenidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
