// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// AllGroupUsersReader is a Reader for the AllGroupUsers structure.
type AllGroupUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllGroupUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllGroupUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllGroupUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllGroupUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllGroupUsersOK creates a AllGroupUsersOK with default headers values
func NewAllGroupUsersOK() *AllGroupUsersOK {
	return &AllGroupUsersOK{}
}

/* AllGroupUsersOK describes a response with status code 200, with default header values.

All users in group.
*/
type AllGroupUsersOK struct {
	Payload []*models.User
}

func (o *AllGroupUsersOK) Error() string {
	return fmt.Sprintf("[GET /groups/{group_id}/users][%d] allGroupUsersOK  %+v", 200, o.Payload)
}
func (o *AllGroupUsersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *AllGroupUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGroupUsersBadRequest creates a AllGroupUsersBadRequest with default headers values
func NewAllGroupUsersBadRequest() *AllGroupUsersBadRequest {
	return &AllGroupUsersBadRequest{}
}

/* AllGroupUsersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllGroupUsersBadRequest struct {
	Payload *models.Error
}

func (o *AllGroupUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /groups/{group_id}/users][%d] allGroupUsersBadRequest  %+v", 400, o.Payload)
}
func (o *AllGroupUsersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllGroupUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGroupUsersNotFound creates a AllGroupUsersNotFound with default headers values
func NewAllGroupUsersNotFound() *AllGroupUsersNotFound {
	return &AllGroupUsersNotFound{}
}

/* AllGroupUsersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllGroupUsersNotFound struct {
	Payload *models.Error
}

func (o *AllGroupUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /groups/{group_id}/users][%d] allGroupUsersNotFound  %+v", 404, o.Payload)
}
func (o *AllGroupUsersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllGroupUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
