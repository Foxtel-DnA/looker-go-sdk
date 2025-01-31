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

// SearchUsersReader is a Reader for the SearchUsers structure.
type SearchUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchUsersOK creates a SearchUsersOK with default headers values
func NewSearchUsersOK() *SearchUsersOK {
	return &SearchUsersOK{}
}

/* SearchUsersOK describes a response with status code 200, with default header values.

Matching users.
*/
type SearchUsersOK struct {
	Payload []*models.User
}

func (o *SearchUsersOK) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersOK  %+v", 200, o.Payload)
}
func (o *SearchUsersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *SearchUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersBadRequest creates a SearchUsersBadRequest with default headers values
func NewSearchUsersBadRequest() *SearchUsersBadRequest {
	return &SearchUsersBadRequest{}
}

/* SearchUsersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchUsersBadRequest struct {
	Payload *models.Error
}

func (o *SearchUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersBadRequest  %+v", 400, o.Payload)
}
func (o *SearchUsersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersNotFound creates a SearchUsersNotFound with default headers values
func NewSearchUsersNotFound() *SearchUsersNotFound {
	return &SearchUsersNotFound{}
}

/* SearchUsersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchUsersNotFound struct {
	Payload *models.Error
}

func (o *SearchUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersNotFound  %+v", 404, o.Payload)
}
func (o *SearchUsersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
