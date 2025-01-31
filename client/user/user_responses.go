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

// UserReader is a Reader for the User structure.
type UserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserOK creates a UserOK with default headers values
func NewUserOK() *UserOK {
	return &UserOK{}
}

/* UserOK describes a response with status code 200, with default header values.

Specified user.
*/
type UserOK struct {
	Payload *models.User
}

func (o *UserOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] userOK  %+v", 200, o.Payload)
}
func (o *UserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserBadRequest creates a UserBadRequest with default headers values
func NewUserBadRequest() *UserBadRequest {
	return &UserBadRequest{}
}

/* UserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserBadRequest struct {
	Payload *models.Error
}

func (o *UserBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] userBadRequest  %+v", 400, o.Payload)
}
func (o *UserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserNotFound creates a UserNotFound with default headers values
func NewUserNotFound() *UserNotFound {
	return &UserNotFound{}
}

/* UserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserNotFound struct {
	Payload *models.Error
}

func (o *UserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] userNotFound  %+v", 404, o.Payload)
}
func (o *UserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
