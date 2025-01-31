// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// UpdateBackupConfigurationReader is a Reader for the UpdateBackupConfiguration structure.
type UpdateBackupConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBackupConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBackupConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateBackupConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateBackupConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateBackupConfigurationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateBackupConfigurationOK creates a UpdateBackupConfigurationOK with default headers values
func NewUpdateBackupConfigurationOK() *UpdateBackupConfigurationOK {
	return &UpdateBackupConfigurationOK{}
}

/* UpdateBackupConfigurationOK describes a response with status code 200, with default header values.

New state for specified model set.
*/
type UpdateBackupConfigurationOK struct {
	Payload *models.BackupConfiguration
}

func (o *UpdateBackupConfigurationOK) Error() string {
	return fmt.Sprintf("[PATCH /backup_configuration][%d] updateBackupConfigurationOK  %+v", 200, o.Payload)
}
func (o *UpdateBackupConfigurationOK) GetPayload() *models.BackupConfiguration {
	return o.Payload
}

func (o *UpdateBackupConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBackupConfigurationBadRequest creates a UpdateBackupConfigurationBadRequest with default headers values
func NewUpdateBackupConfigurationBadRequest() *UpdateBackupConfigurationBadRequest {
	return &UpdateBackupConfigurationBadRequest{}
}

/* UpdateBackupConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateBackupConfigurationBadRequest struct {
	Payload *models.Error
}

func (o *UpdateBackupConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /backup_configuration][%d] updateBackupConfigurationBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateBackupConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateBackupConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBackupConfigurationNotFound creates a UpdateBackupConfigurationNotFound with default headers values
func NewUpdateBackupConfigurationNotFound() *UpdateBackupConfigurationNotFound {
	return &UpdateBackupConfigurationNotFound{}
}

/* UpdateBackupConfigurationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateBackupConfigurationNotFound struct {
	Payload *models.Error
}

func (o *UpdateBackupConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /backup_configuration][%d] updateBackupConfigurationNotFound  %+v", 404, o.Payload)
}
func (o *UpdateBackupConfigurationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateBackupConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBackupConfigurationUnprocessableEntity creates a UpdateBackupConfigurationUnprocessableEntity with default headers values
func NewUpdateBackupConfigurationUnprocessableEntity() *UpdateBackupConfigurationUnprocessableEntity {
	return &UpdateBackupConfigurationUnprocessableEntity{}
}

/* UpdateBackupConfigurationUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateBackupConfigurationUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateBackupConfigurationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /backup_configuration][%d] updateBackupConfigurationUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateBackupConfigurationUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateBackupConfigurationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
