// Code generated by go-swagger; DO NOT EDIT.

package datagroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// UpdateDatagroupReader is a Reader for the UpdateDatagroup structure.
type UpdateDatagroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDatagroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDatagroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDatagroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDatagroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDatagroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateDatagroupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateDatagroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDatagroupOK creates a UpdateDatagroupOK with default headers values
func NewUpdateDatagroupOK() *UpdateDatagroupOK {
	return &UpdateDatagroupOK{}
}

/* UpdateDatagroupOK describes a response with status code 200, with default header values.

Datagroup
*/
type UpdateDatagroupOK struct {
	Payload *models.Datagroup
}

func (o *UpdateDatagroupOK) Error() string {
	return fmt.Sprintf("[PATCH /datagroups/{datagroup_id}][%d] updateDatagroupOK  %+v", 200, o.Payload)
}
func (o *UpdateDatagroupOK) GetPayload() *models.Datagroup {
	return o.Payload
}

func (o *UpdateDatagroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datagroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatagroupBadRequest creates a UpdateDatagroupBadRequest with default headers values
func NewUpdateDatagroupBadRequest() *UpdateDatagroupBadRequest {
	return &UpdateDatagroupBadRequest{}
}

/* UpdateDatagroupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateDatagroupBadRequest struct {
	Payload *models.Error
}

func (o *UpdateDatagroupBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /datagroups/{datagroup_id}][%d] updateDatagroupBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateDatagroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDatagroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatagroupNotFound creates a UpdateDatagroupNotFound with default headers values
func NewUpdateDatagroupNotFound() *UpdateDatagroupNotFound {
	return &UpdateDatagroupNotFound{}
}

/* UpdateDatagroupNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateDatagroupNotFound struct {
	Payload *models.Error
}

func (o *UpdateDatagroupNotFound) Error() string {
	return fmt.Sprintf("[PATCH /datagroups/{datagroup_id}][%d] updateDatagroupNotFound  %+v", 404, o.Payload)
}
func (o *UpdateDatagroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDatagroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatagroupConflict creates a UpdateDatagroupConflict with default headers values
func NewUpdateDatagroupConflict() *UpdateDatagroupConflict {
	return &UpdateDatagroupConflict{}
}

/* UpdateDatagroupConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type UpdateDatagroupConflict struct {
	Payload *models.Error
}

func (o *UpdateDatagroupConflict) Error() string {
	return fmt.Sprintf("[PATCH /datagroups/{datagroup_id}][%d] updateDatagroupConflict  %+v", 409, o.Payload)
}
func (o *UpdateDatagroupConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDatagroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatagroupUnprocessableEntity creates a UpdateDatagroupUnprocessableEntity with default headers values
func NewUpdateDatagroupUnprocessableEntity() *UpdateDatagroupUnprocessableEntity {
	return &UpdateDatagroupUnprocessableEntity{}
}

/* UpdateDatagroupUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateDatagroupUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateDatagroupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /datagroups/{datagroup_id}][%d] updateDatagroupUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateDatagroupUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateDatagroupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatagroupTooManyRequests creates a UpdateDatagroupTooManyRequests with default headers values
func NewUpdateDatagroupTooManyRequests() *UpdateDatagroupTooManyRequests {
	return &UpdateDatagroupTooManyRequests{}
}

/* UpdateDatagroupTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDatagroupTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateDatagroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /datagroups/{datagroup_id}][%d] updateDatagroupTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateDatagroupTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDatagroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
