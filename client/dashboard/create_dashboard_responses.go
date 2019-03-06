// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// CreateDashboardReader is a Reader for the CreateDashboard structure.
type CreateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateDashboardConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateDashboardUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDashboardOK creates a CreateDashboardOK with default headers values
func NewCreateDashboardOK() *CreateDashboardOK {
	return &CreateDashboardOK{}
}

/*CreateDashboardOK handles this case with default header values.

Dashboard
*/
type CreateDashboardOK struct {
	Payload *models.Dashboard
}

func (o *CreateDashboardOK) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardBadRequest creates a CreateDashboardBadRequest with default headers values
func NewCreateDashboardBadRequest() *CreateDashboardBadRequest {
	return &CreateDashboardBadRequest{}
}

/*CreateDashboardBadRequest handles this case with default header values.

Bad Request
*/
type CreateDashboardBadRequest struct {
	Payload *models.Error
}

func (o *CreateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardNotFound creates a CreateDashboardNotFound with default headers values
func NewCreateDashboardNotFound() *CreateDashboardNotFound {
	return &CreateDashboardNotFound{}
}

/*CreateDashboardNotFound handles this case with default header values.

Not Found
*/
type CreateDashboardNotFound struct {
	Payload *models.Error
}

func (o *CreateDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardConflict creates a CreateDashboardConflict with default headers values
func NewCreateDashboardConflict() *CreateDashboardConflict {
	return &CreateDashboardConflict{}
}

/*CreateDashboardConflict handles this case with default header values.

Resource Already Exists
*/
type CreateDashboardConflict struct {
	Payload *models.Error
}

func (o *CreateDashboardConflict) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardConflict  %+v", 409, o.Payload)
}

func (o *CreateDashboardConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardUnprocessableEntity creates a CreateDashboardUnprocessableEntity with default headers values
func NewCreateDashboardUnprocessableEntity() *CreateDashboardUnprocessableEntity {
	return &CreateDashboardUnprocessableEntity{}
}

/*CreateDashboardUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateDashboardUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateDashboardUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /dashboards][%d] createDashboardUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateDashboardUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
