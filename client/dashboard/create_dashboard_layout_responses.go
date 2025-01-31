// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// CreateDashboardLayoutReader is a Reader for the CreateDashboardLayout structure.
type CreateDashboardLayoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardLayoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDashboardLayoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDashboardLayoutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDashboardLayoutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDashboardLayoutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateDashboardLayoutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDashboardLayoutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDashboardLayoutOK creates a CreateDashboardLayoutOK with default headers values
func NewCreateDashboardLayoutOK() *CreateDashboardLayoutOK {
	return &CreateDashboardLayoutOK{}
}

/* CreateDashboardLayoutOK describes a response with status code 200, with default header values.

DashboardLayout
*/
type CreateDashboardLayoutOK struct {
	Payload *models.DashboardLayout
}

func (o *CreateDashboardLayoutOK) Error() string {
	return fmt.Sprintf("[POST /dashboard_layouts][%d] createDashboardLayoutOK  %+v", 200, o.Payload)
}
func (o *CreateDashboardLayoutOK) GetPayload() *models.DashboardLayout {
	return o.Payload
}

func (o *CreateDashboardLayoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardLayout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardLayoutBadRequest creates a CreateDashboardLayoutBadRequest with default headers values
func NewCreateDashboardLayoutBadRequest() *CreateDashboardLayoutBadRequest {
	return &CreateDashboardLayoutBadRequest{}
}

/* CreateDashboardLayoutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateDashboardLayoutBadRequest struct {
	Payload *models.Error
}

func (o *CreateDashboardLayoutBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboard_layouts][%d] createDashboardLayoutBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDashboardLayoutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDashboardLayoutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardLayoutNotFound creates a CreateDashboardLayoutNotFound with default headers values
func NewCreateDashboardLayoutNotFound() *CreateDashboardLayoutNotFound {
	return &CreateDashboardLayoutNotFound{}
}

/* CreateDashboardLayoutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateDashboardLayoutNotFound struct {
	Payload *models.Error
}

func (o *CreateDashboardLayoutNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboard_layouts][%d] createDashboardLayoutNotFound  %+v", 404, o.Payload)
}
func (o *CreateDashboardLayoutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDashboardLayoutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardLayoutConflict creates a CreateDashboardLayoutConflict with default headers values
func NewCreateDashboardLayoutConflict() *CreateDashboardLayoutConflict {
	return &CreateDashboardLayoutConflict{}
}

/* CreateDashboardLayoutConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateDashboardLayoutConflict struct {
	Payload *models.Error
}

func (o *CreateDashboardLayoutConflict) Error() string {
	return fmt.Sprintf("[POST /dashboard_layouts][%d] createDashboardLayoutConflict  %+v", 409, o.Payload)
}
func (o *CreateDashboardLayoutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDashboardLayoutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardLayoutUnprocessableEntity creates a CreateDashboardLayoutUnprocessableEntity with default headers values
func NewCreateDashboardLayoutUnprocessableEntity() *CreateDashboardLayoutUnprocessableEntity {
	return &CreateDashboardLayoutUnprocessableEntity{}
}

/* CreateDashboardLayoutUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateDashboardLayoutUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateDashboardLayoutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /dashboard_layouts][%d] createDashboardLayoutUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateDashboardLayoutUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateDashboardLayoutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardLayoutTooManyRequests creates a CreateDashboardLayoutTooManyRequests with default headers values
func NewCreateDashboardLayoutTooManyRequests() *CreateDashboardLayoutTooManyRequests {
	return &CreateDashboardLayoutTooManyRequests{}
}

/* CreateDashboardLayoutTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateDashboardLayoutTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateDashboardLayoutTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dashboard_layouts][%d] createDashboardLayoutTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateDashboardLayoutTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDashboardLayoutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
