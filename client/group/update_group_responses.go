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

// UpdateGroupReader is a Reader for the UpdateGroup structure.
type UpdateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateGroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGroupOK creates a UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {
	return &UpdateGroupOK{}
}

/* UpdateGroupOK describes a response with status code 200, with default header values.

Group
*/
type UpdateGroupOK struct {
	Payload *models.Group
}

func (o *UpdateGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /groups/{group_id}][%d] updateGroupOK  %+v", 200, o.Payload)
}
func (o *UpdateGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UpdateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupBadRequest creates a UpdateGroupBadRequest with default headers values
func NewUpdateGroupBadRequest() *UpdateGroupBadRequest {
	return &UpdateGroupBadRequest{}
}

/* UpdateGroupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateGroupBadRequest struct {
	Payload *models.Error
}

func (o *UpdateGroupBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /groups/{group_id}][%d] updateGroupBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupNotFound creates a UpdateGroupNotFound with default headers values
func NewUpdateGroupNotFound() *UpdateGroupNotFound {
	return &UpdateGroupNotFound{}
}

/* UpdateGroupNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateGroupNotFound struct {
	Payload *models.Error
}

func (o *UpdateGroupNotFound) Error() string {
	return fmt.Sprintf("[PATCH /groups/{group_id}][%d] updateGroupNotFound  %+v", 404, o.Payload)
}
func (o *UpdateGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupUnprocessableEntity creates a UpdateGroupUnprocessableEntity with default headers values
func NewUpdateGroupUnprocessableEntity() *UpdateGroupUnprocessableEntity {
	return &UpdateGroupUnprocessableEntity{}
}

/* UpdateGroupUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateGroupUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateGroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /groups/{group_id}][%d] updateGroupUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateGroupUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateGroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupTooManyRequests creates a UpdateGroupTooManyRequests with default headers values
func NewUpdateGroupTooManyRequests() *UpdateGroupTooManyRequests {
	return &UpdateGroupTooManyRequests{}
}

/* UpdateGroupTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateGroupTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /groups/{group_id}][%d] updateGroupTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateGroupTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
