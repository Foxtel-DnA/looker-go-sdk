// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// UpdateScheduledPlanReader is a Reader for the UpdateScheduledPlan structure.
type UpdateScheduledPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScheduledPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateScheduledPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateScheduledPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateScheduledPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateScheduledPlanUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateScheduledPlanTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateScheduledPlanOK creates a UpdateScheduledPlanOK with default headers values
func NewUpdateScheduledPlanOK() *UpdateScheduledPlanOK {
	return &UpdateScheduledPlanOK{}
}

/* UpdateScheduledPlanOK describes a response with status code 200, with default header values.

Scheduled Plan
*/
type UpdateScheduledPlanOK struct {
	Payload *models.ScheduledPlan
}

func (o *UpdateScheduledPlanOK) Error() string {
	return fmt.Sprintf("[PATCH /scheduled_plans/{scheduled_plan_id}][%d] updateScheduledPlanOK  %+v", 200, o.Payload)
}
func (o *UpdateScheduledPlanOK) GetPayload() *models.ScheduledPlan {
	return o.Payload
}

func (o *UpdateScheduledPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduledPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduledPlanBadRequest creates a UpdateScheduledPlanBadRequest with default headers values
func NewUpdateScheduledPlanBadRequest() *UpdateScheduledPlanBadRequest {
	return &UpdateScheduledPlanBadRequest{}
}

/* UpdateScheduledPlanBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateScheduledPlanBadRequest struct {
	Payload *models.Error
}

func (o *UpdateScheduledPlanBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /scheduled_plans/{scheduled_plan_id}][%d] updateScheduledPlanBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateScheduledPlanBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateScheduledPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduledPlanNotFound creates a UpdateScheduledPlanNotFound with default headers values
func NewUpdateScheduledPlanNotFound() *UpdateScheduledPlanNotFound {
	return &UpdateScheduledPlanNotFound{}
}

/* UpdateScheduledPlanNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateScheduledPlanNotFound struct {
	Payload *models.Error
}

func (o *UpdateScheduledPlanNotFound) Error() string {
	return fmt.Sprintf("[PATCH /scheduled_plans/{scheduled_plan_id}][%d] updateScheduledPlanNotFound  %+v", 404, o.Payload)
}
func (o *UpdateScheduledPlanNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateScheduledPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduledPlanUnprocessableEntity creates a UpdateScheduledPlanUnprocessableEntity with default headers values
func NewUpdateScheduledPlanUnprocessableEntity() *UpdateScheduledPlanUnprocessableEntity {
	return &UpdateScheduledPlanUnprocessableEntity{}
}

/* UpdateScheduledPlanUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateScheduledPlanUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateScheduledPlanUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /scheduled_plans/{scheduled_plan_id}][%d] updateScheduledPlanUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateScheduledPlanUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateScheduledPlanUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduledPlanTooManyRequests creates a UpdateScheduledPlanTooManyRequests with default headers values
func NewUpdateScheduledPlanTooManyRequests() *UpdateScheduledPlanTooManyRequests {
	return &UpdateScheduledPlanTooManyRequests{}
}

/* UpdateScheduledPlanTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateScheduledPlanTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateScheduledPlanTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /scheduled_plans/{scheduled_plan_id}][%d] updateScheduledPlanTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateScheduledPlanTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateScheduledPlanTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
