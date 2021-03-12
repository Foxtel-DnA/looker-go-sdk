// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// AllRolesReader is a Reader for the AllRoles structure.
type AllRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllRolesOK creates a AllRolesOK with default headers values
func NewAllRolesOK() *AllRolesOK {
	return &AllRolesOK{}
}

/* AllRolesOK describes a response with status code 200, with default header values.

Role
*/
type AllRolesOK struct {
	Payload []*models.Role
}

func (o *AllRolesOK) Error() string {
	return fmt.Sprintf("[GET /roles][%d] allRolesOK  %+v", 200, o.Payload)
}
func (o *AllRolesOK) GetPayload() []*models.Role {
	return o.Payload
}

func (o *AllRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllRolesBadRequest creates a AllRolesBadRequest with default headers values
func NewAllRolesBadRequest() *AllRolesBadRequest {
	return &AllRolesBadRequest{}
}

/* AllRolesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllRolesBadRequest struct {
	Payload *models.Error
}

func (o *AllRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /roles][%d] allRolesBadRequest  %+v", 400, o.Payload)
}
func (o *AllRolesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllRolesNotFound creates a AllRolesNotFound with default headers values
func NewAllRolesNotFound() *AllRolesNotFound {
	return &AllRolesNotFound{}
}

/* AllRolesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllRolesNotFound struct {
	Payload *models.Error
}

func (o *AllRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /roles][%d] allRolesNotFound  %+v", 404, o.Payload)
}
func (o *AllRolesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
