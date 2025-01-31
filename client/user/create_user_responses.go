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

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserOK creates a CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {
	return &CreateUserOK{}
}

/* CreateUserOK describes a response with status code 200, with default header values.

Created User
*/
type CreateUserOK struct {
	Payload *models.User
}

func (o *CreateUserOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserOK  %+v", 200, o.Payload)
}
func (o *CreateUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *CreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/* CreateUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateUserBadRequest struct {
	Payload *models.Error
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserBadRequest  %+v", 400, o.Payload)
}
func (o *CreateUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserNotFound creates a CreateUserNotFound with default headers values
func NewCreateUserNotFound() *CreateUserNotFound {
	return &CreateUserNotFound{}
}

/* CreateUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateUserNotFound struct {
	Payload *models.Error
}

func (o *CreateUserNotFound) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserNotFound  %+v", 404, o.Payload)
}
func (o *CreateUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/* CreateUserConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateUserConflict struct {
	Payload *models.Error
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserConflict  %+v", 409, o.Payload)
}
func (o *CreateUserConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnprocessableEntity creates a CreateUserUnprocessableEntity with default headers values
func NewCreateUserUnprocessableEntity() *CreateUserUnprocessableEntity {
	return &CreateUserUnprocessableEntity{}
}

/* CreateUserUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateUserUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users][%d] createUserUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateUserUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
