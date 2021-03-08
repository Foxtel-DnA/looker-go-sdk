// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdateDashboardLayoutComponentReader is a Reader for the UpdateDashboardLayoutComponent structure.
type UpdateDashboardLayoutComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardLayoutComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDashboardLayoutComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDashboardLayoutComponentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDashboardLayoutComponentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateDashboardLayoutComponentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateDashboardLayoutComponentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDashboardLayoutComponentOK creates a UpdateDashboardLayoutComponentOK with default headers values
func NewUpdateDashboardLayoutComponentOK() *UpdateDashboardLayoutComponentOK {
	return &UpdateDashboardLayoutComponentOK{}
}

/* UpdateDashboardLayoutComponentOK describes a response with status code 200, with default header values.

DashboardLayoutComponent
*/
type UpdateDashboardLayoutComponentOK struct {
	Payload *models.DashboardLayoutComponent
}

func (o *UpdateDashboardLayoutComponentOK) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_layout_components/{dashboard_layout_component_id}][%d] updateDashboardLayoutComponentOK  %+v", 200, o.Payload)
}
func (o *UpdateDashboardLayoutComponentOK) GetPayload() *models.DashboardLayoutComponent {
	return o.Payload
}

func (o *UpdateDashboardLayoutComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardLayoutComponent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardLayoutComponentBadRequest creates a UpdateDashboardLayoutComponentBadRequest with default headers values
func NewUpdateDashboardLayoutComponentBadRequest() *UpdateDashboardLayoutComponentBadRequest {
	return &UpdateDashboardLayoutComponentBadRequest{}
}

/* UpdateDashboardLayoutComponentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateDashboardLayoutComponentBadRequest struct {
	Payload *models.Error
}

func (o *UpdateDashboardLayoutComponentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_layout_components/{dashboard_layout_component_id}][%d] updateDashboardLayoutComponentBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateDashboardLayoutComponentBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDashboardLayoutComponentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardLayoutComponentNotFound creates a UpdateDashboardLayoutComponentNotFound with default headers values
func NewUpdateDashboardLayoutComponentNotFound() *UpdateDashboardLayoutComponentNotFound {
	return &UpdateDashboardLayoutComponentNotFound{}
}

/* UpdateDashboardLayoutComponentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateDashboardLayoutComponentNotFound struct {
	Payload *models.Error
}

func (o *UpdateDashboardLayoutComponentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_layout_components/{dashboard_layout_component_id}][%d] updateDashboardLayoutComponentNotFound  %+v", 404, o.Payload)
}
func (o *UpdateDashboardLayoutComponentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDashboardLayoutComponentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardLayoutComponentUnprocessableEntity creates a UpdateDashboardLayoutComponentUnprocessableEntity with default headers values
func NewUpdateDashboardLayoutComponentUnprocessableEntity() *UpdateDashboardLayoutComponentUnprocessableEntity {
	return &UpdateDashboardLayoutComponentUnprocessableEntity{}
}

/* UpdateDashboardLayoutComponentUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateDashboardLayoutComponentUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateDashboardLayoutComponentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_layout_components/{dashboard_layout_component_id}][%d] updateDashboardLayoutComponentUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateDashboardLayoutComponentUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateDashboardLayoutComponentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardLayoutComponentTooManyRequests creates a UpdateDashboardLayoutComponentTooManyRequests with default headers values
func NewUpdateDashboardLayoutComponentTooManyRequests() *UpdateDashboardLayoutComponentTooManyRequests {
	return &UpdateDashboardLayoutComponentTooManyRequests{}
}

/* UpdateDashboardLayoutComponentTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDashboardLayoutComponentTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateDashboardLayoutComponentTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_layout_components/{dashboard_layout_component_id}][%d] updateDashboardLayoutComponentTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateDashboardLayoutComponentTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDashboardLayoutComponentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}