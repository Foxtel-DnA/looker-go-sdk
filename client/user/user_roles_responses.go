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

// UserRolesReader is a Reader for the UserRoles structure.
type UserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserRolesOK creates a UserRolesOK with default headers values
func NewUserRolesOK() *UserRolesOK {
	return &UserRolesOK{}
}

/* UserRolesOK describes a response with status code 200, with default header values.

Roles of user.
*/
type UserRolesOK struct {
	Payload []*models.Role
}

func (o *UserRolesOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] userRolesOK  %+v", 200, o.Payload)
}
func (o *UserRolesOK) GetPayload() []*models.Role {
	return o.Payload
}

func (o *UserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserRolesBadRequest creates a UserRolesBadRequest with default headers values
func NewUserRolesBadRequest() *UserRolesBadRequest {
	return &UserRolesBadRequest{}
}

/* UserRolesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserRolesBadRequest struct {
	Payload *models.Error
}

func (o *UserRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] userRolesBadRequest  %+v", 400, o.Payload)
}
func (o *UserRolesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserRolesNotFound creates a UserRolesNotFound with default headers values
func NewUserRolesNotFound() *UserRolesNotFound {
	return &UserRolesNotFound{}
}

/* UserRolesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserRolesNotFound struct {
	Payload *models.Error
}

func (o *UserRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/roles][%d] userRolesNotFound  %+v", 404, o.Payload)
}
func (o *UserRolesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
