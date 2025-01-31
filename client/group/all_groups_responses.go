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

// AllGroupsReader is a Reader for the AllGroups structure.
type AllGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllGroupsOK creates a AllGroupsOK with default headers values
func NewAllGroupsOK() *AllGroupsOK {
	return &AllGroupsOK{}
}

/* AllGroupsOK describes a response with status code 200, with default header values.

Group
*/
type AllGroupsOK struct {
	Payload []*models.Group
}

func (o *AllGroupsOK) Error() string {
	return fmt.Sprintf("[GET /groups][%d] allGroupsOK  %+v", 200, o.Payload)
}
func (o *AllGroupsOK) GetPayload() []*models.Group {
	return o.Payload
}

func (o *AllGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGroupsBadRequest creates a AllGroupsBadRequest with default headers values
func NewAllGroupsBadRequest() *AllGroupsBadRequest {
	return &AllGroupsBadRequest{}
}

/* AllGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllGroupsBadRequest struct {
	Payload *models.Error
}

func (o *AllGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /groups][%d] allGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *AllGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGroupsNotFound creates a AllGroupsNotFound with default headers values
func NewAllGroupsNotFound() *AllGroupsNotFound {
	return &AllGroupsNotFound{}
}

/* AllGroupsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllGroupsNotFound struct {
	Payload *models.Error
}

func (o *AllGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /groups][%d] allGroupsNotFound  %+v", 404, o.Payload)
}
func (o *AllGroupsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
