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

// ImportLookmlDashboardReader is a Reader for the ImportLookmlDashboard structure.
type ImportLookmlDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportLookmlDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportLookmlDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewImportLookmlDashboardCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportLookmlDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportLookmlDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImportLookmlDashboardConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewImportLookmlDashboardUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewImportLookmlDashboardTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportLookmlDashboardOK creates a ImportLookmlDashboardOK with default headers values
func NewImportLookmlDashboardOK() *ImportLookmlDashboardOK {
	return &ImportLookmlDashboardOK{}
}

/* ImportLookmlDashboardOK describes a response with status code 200, with default header values.

Dashboard
*/
type ImportLookmlDashboardOK struct {
	Payload *models.Dashboard
}

func (o *ImportLookmlDashboardOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/{lookml_dashboard_id}/import/{space_id}][%d] importLookmlDashboardOK  %+v", 200, o.Payload)
}
func (o *ImportLookmlDashboardOK) GetPayload() *models.Dashboard {
	return o.Payload
}

func (o *ImportLookmlDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLookmlDashboardCreated creates a ImportLookmlDashboardCreated with default headers values
func NewImportLookmlDashboardCreated() *ImportLookmlDashboardCreated {
	return &ImportLookmlDashboardCreated{}
}

/* ImportLookmlDashboardCreated describes a response with status code 201, with default header values.

dashboard
*/
type ImportLookmlDashboardCreated struct {
	Payload *models.Dashboard
}

func (o *ImportLookmlDashboardCreated) Error() string {
	return fmt.Sprintf("[POST /dashboards/{lookml_dashboard_id}/import/{space_id}][%d] importLookmlDashboardCreated  %+v", 201, o.Payload)
}
func (o *ImportLookmlDashboardCreated) GetPayload() *models.Dashboard {
	return o.Payload
}

func (o *ImportLookmlDashboardCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLookmlDashboardBadRequest creates a ImportLookmlDashboardBadRequest with default headers values
func NewImportLookmlDashboardBadRequest() *ImportLookmlDashboardBadRequest {
	return &ImportLookmlDashboardBadRequest{}
}

/* ImportLookmlDashboardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ImportLookmlDashboardBadRequest struct {
	Payload *models.Error
}

func (o *ImportLookmlDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards/{lookml_dashboard_id}/import/{space_id}][%d] importLookmlDashboardBadRequest  %+v", 400, o.Payload)
}
func (o *ImportLookmlDashboardBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportLookmlDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLookmlDashboardNotFound creates a ImportLookmlDashboardNotFound with default headers values
func NewImportLookmlDashboardNotFound() *ImportLookmlDashboardNotFound {
	return &ImportLookmlDashboardNotFound{}
}

/* ImportLookmlDashboardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ImportLookmlDashboardNotFound struct {
	Payload *models.Error
}

func (o *ImportLookmlDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /dashboards/{lookml_dashboard_id}/import/{space_id}][%d] importLookmlDashboardNotFound  %+v", 404, o.Payload)
}
func (o *ImportLookmlDashboardNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportLookmlDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLookmlDashboardConflict creates a ImportLookmlDashboardConflict with default headers values
func NewImportLookmlDashboardConflict() *ImportLookmlDashboardConflict {
	return &ImportLookmlDashboardConflict{}
}

/* ImportLookmlDashboardConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type ImportLookmlDashboardConflict struct {
	Payload *models.Error
}

func (o *ImportLookmlDashboardConflict) Error() string {
	return fmt.Sprintf("[POST /dashboards/{lookml_dashboard_id}/import/{space_id}][%d] importLookmlDashboardConflict  %+v", 409, o.Payload)
}
func (o *ImportLookmlDashboardConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportLookmlDashboardConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLookmlDashboardUnprocessableEntity creates a ImportLookmlDashboardUnprocessableEntity with default headers values
func NewImportLookmlDashboardUnprocessableEntity() *ImportLookmlDashboardUnprocessableEntity {
	return &ImportLookmlDashboardUnprocessableEntity{}
}

/* ImportLookmlDashboardUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ImportLookmlDashboardUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *ImportLookmlDashboardUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /dashboards/{lookml_dashboard_id}/import/{space_id}][%d] importLookmlDashboardUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ImportLookmlDashboardUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ImportLookmlDashboardUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportLookmlDashboardTooManyRequests creates a ImportLookmlDashboardTooManyRequests with default headers values
func NewImportLookmlDashboardTooManyRequests() *ImportLookmlDashboardTooManyRequests {
	return &ImportLookmlDashboardTooManyRequests{}
}

/* ImportLookmlDashboardTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ImportLookmlDashboardTooManyRequests struct {
	Payload *models.Error
}

func (o *ImportLookmlDashboardTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dashboards/{lookml_dashboard_id}/import/{space_id}][%d] importLookmlDashboardTooManyRequests  %+v", 429, o.Payload)
}
func (o *ImportLookmlDashboardTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportLookmlDashboardTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
