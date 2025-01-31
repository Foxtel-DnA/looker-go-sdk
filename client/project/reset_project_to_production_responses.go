// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// ResetProjectToProductionReader is a Reader for the ResetProjectToProduction structure.
type ResetProjectToProductionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetProjectToProductionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetProjectToProductionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewResetProjectToProductionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResetProjectToProductionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResetProjectToProductionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewResetProjectToProductionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewResetProjectToProductionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResetProjectToProductionOK creates a ResetProjectToProductionOK with default headers values
func NewResetProjectToProductionOK() *ResetProjectToProductionOK {
	return &ResetProjectToProductionOK{}
}

/* ResetProjectToProductionOK describes a response with status code 200, with default header values.

Project
*/
type ResetProjectToProductionOK struct {
	Payload string
}

func (o *ResetProjectToProductionOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_production][%d] resetProjectToProductionOK  %+v", 200, o.Payload)
}
func (o *ResetProjectToProductionOK) GetPayload() string {
	return o.Payload
}

func (o *ResetProjectToProductionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetProjectToProductionNoContent creates a ResetProjectToProductionNoContent with default headers values
func NewResetProjectToProductionNoContent() *ResetProjectToProductionNoContent {
	return &ResetProjectToProductionNoContent{}
}

/* ResetProjectToProductionNoContent describes a response with status code 204, with default header values.

Returns 204 if project was successfully reset, otherwise 400 with an error message
*/
type ResetProjectToProductionNoContent struct {
}

func (o *ResetProjectToProductionNoContent) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_production][%d] resetProjectToProductionNoContent ", 204)
}

func (o *ResetProjectToProductionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetProjectToProductionBadRequest creates a ResetProjectToProductionBadRequest with default headers values
func NewResetProjectToProductionBadRequest() *ResetProjectToProductionBadRequest {
	return &ResetProjectToProductionBadRequest{}
}

/* ResetProjectToProductionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ResetProjectToProductionBadRequest struct {
	Payload *models.Error
}

func (o *ResetProjectToProductionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_production][%d] resetProjectToProductionBadRequest  %+v", 400, o.Payload)
}
func (o *ResetProjectToProductionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResetProjectToProductionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetProjectToProductionNotFound creates a ResetProjectToProductionNotFound with default headers values
func NewResetProjectToProductionNotFound() *ResetProjectToProductionNotFound {
	return &ResetProjectToProductionNotFound{}
}

/* ResetProjectToProductionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ResetProjectToProductionNotFound struct {
	Payload *models.Error
}

func (o *ResetProjectToProductionNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_production][%d] resetProjectToProductionNotFound  %+v", 404, o.Payload)
}
func (o *ResetProjectToProductionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResetProjectToProductionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetProjectToProductionUnprocessableEntity creates a ResetProjectToProductionUnprocessableEntity with default headers values
func NewResetProjectToProductionUnprocessableEntity() *ResetProjectToProductionUnprocessableEntity {
	return &ResetProjectToProductionUnprocessableEntity{}
}

/* ResetProjectToProductionUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ResetProjectToProductionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *ResetProjectToProductionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_production][%d] resetProjectToProductionUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ResetProjectToProductionUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ResetProjectToProductionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetProjectToProductionTooManyRequests creates a ResetProjectToProductionTooManyRequests with default headers values
func NewResetProjectToProductionTooManyRequests() *ResetProjectToProductionTooManyRequests {
	return &ResetProjectToProductionTooManyRequests{}
}

/* ResetProjectToProductionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ResetProjectToProductionTooManyRequests struct {
	Payload *models.Error
}

func (o *ResetProjectToProductionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_production][%d] resetProjectToProductionTooManyRequests  %+v", 429, o.Payload)
}
func (o *ResetProjectToProductionTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResetProjectToProductionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
