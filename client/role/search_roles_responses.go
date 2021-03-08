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

// SearchRolesReader is a Reader for the SearchRoles structure.
type SearchRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchRolesOK creates a SearchRolesOK with default headers values
func NewSearchRolesOK() *SearchRolesOK {
	return &SearchRolesOK{}
}

/* SearchRolesOK describes a response with status code 200, with default header values.

Role
*/
type SearchRolesOK struct {
	Payload []*models.Role
}

func (o *SearchRolesOK) Error() string {
	return fmt.Sprintf("[GET /roles/search][%d] searchRolesOK  %+v", 200, o.Payload)
}
func (o *SearchRolesOK) GetPayload() []*models.Role {
	return o.Payload
}

func (o *SearchRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRolesBadRequest creates a SearchRolesBadRequest with default headers values
func NewSearchRolesBadRequest() *SearchRolesBadRequest {
	return &SearchRolesBadRequest{}
}

/* SearchRolesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchRolesBadRequest struct {
	Payload *models.Error
}

func (o *SearchRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /roles/search][%d] searchRolesBadRequest  %+v", 400, o.Payload)
}
func (o *SearchRolesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRolesNotFound creates a SearchRolesNotFound with default headers values
func NewSearchRolesNotFound() *SearchRolesNotFound {
	return &SearchRolesNotFound{}
}

/* SearchRolesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchRolesNotFound struct {
	Payload *models.Error
}

func (o *SearchRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /roles/search][%d] searchRolesNotFound  %+v", 404, o.Payload)
}
func (o *SearchRolesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
