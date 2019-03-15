// Code generated by go-swagger; DO NOT EDIT.

package render_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// CreateDashboardRenderTaskReader is a Reader for the CreateDashboardRenderTask structure.
type CreateDashboardRenderTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardRenderTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDashboardRenderTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDashboardRenderTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateDashboardRenderTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateDashboardRenderTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateDashboardRenderTaskUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDashboardRenderTaskOK creates a CreateDashboardRenderTaskOK with default headers values
func NewCreateDashboardRenderTaskOK() *CreateDashboardRenderTaskOK {
	return &CreateDashboardRenderTaskOK{}
}

/*CreateDashboardRenderTaskOK handles this case with default header values.

Render Task
*/
type CreateDashboardRenderTaskOK struct {
	Payload *models.RenderTask
}

func (o *CreateDashboardRenderTaskOK) Error() string {
	return fmt.Sprintf("[POST /render_tasks/dashboards/{dashboard_id}/{result_format}][%d] createDashboardRenderTaskOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardRenderTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RenderTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardRenderTaskBadRequest creates a CreateDashboardRenderTaskBadRequest with default headers values
func NewCreateDashboardRenderTaskBadRequest() *CreateDashboardRenderTaskBadRequest {
	return &CreateDashboardRenderTaskBadRequest{}
}

/*CreateDashboardRenderTaskBadRequest handles this case with default header values.

Bad Request
*/
type CreateDashboardRenderTaskBadRequest struct {
	Payload *models.Error
}

func (o *CreateDashboardRenderTaskBadRequest) Error() string {
	return fmt.Sprintf("[POST /render_tasks/dashboards/{dashboard_id}/{result_format}][%d] createDashboardRenderTaskBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDashboardRenderTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardRenderTaskNotFound creates a CreateDashboardRenderTaskNotFound with default headers values
func NewCreateDashboardRenderTaskNotFound() *CreateDashboardRenderTaskNotFound {
	return &CreateDashboardRenderTaskNotFound{}
}

/*CreateDashboardRenderTaskNotFound handles this case with default header values.

Not Found
*/
type CreateDashboardRenderTaskNotFound struct {
	Payload *models.Error
}

func (o *CreateDashboardRenderTaskNotFound) Error() string {
	return fmt.Sprintf("[POST /render_tasks/dashboards/{dashboard_id}/{result_format}][%d] createDashboardRenderTaskNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardRenderTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardRenderTaskConflict creates a CreateDashboardRenderTaskConflict with default headers values
func NewCreateDashboardRenderTaskConflict() *CreateDashboardRenderTaskConflict {
	return &CreateDashboardRenderTaskConflict{}
}

/*CreateDashboardRenderTaskConflict handles this case with default header values.

Resource Already Exists
*/
type CreateDashboardRenderTaskConflict struct {
	Payload *models.Error
}

func (o *CreateDashboardRenderTaskConflict) Error() string {
	return fmt.Sprintf("[POST /render_tasks/dashboards/{dashboard_id}/{result_format}][%d] createDashboardRenderTaskConflict  %+v", 409, o.Payload)
}

func (o *CreateDashboardRenderTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardRenderTaskUnprocessableEntity creates a CreateDashboardRenderTaskUnprocessableEntity with default headers values
func NewCreateDashboardRenderTaskUnprocessableEntity() *CreateDashboardRenderTaskUnprocessableEntity {
	return &CreateDashboardRenderTaskUnprocessableEntity{}
}

/*CreateDashboardRenderTaskUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateDashboardRenderTaskUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateDashboardRenderTaskUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /render_tasks/dashboards/{dashboard_id}/{result_format}][%d] createDashboardRenderTaskUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateDashboardRenderTaskUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}