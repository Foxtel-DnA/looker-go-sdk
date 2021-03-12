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

// UpdateCloudStorageConfigurationReader is a Reader for the UpdateCloudStorageConfiguration structure.
type UpdateCloudStorageConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCloudStorageConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCloudStorageConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCloudStorageConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCloudStorageConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateCloudStorageConfigurationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCloudStorageConfigurationOK creates a UpdateCloudStorageConfigurationOK with default headers values
func NewUpdateCloudStorageConfigurationOK() *UpdateCloudStorageConfigurationOK {
	return &UpdateCloudStorageConfigurationOK{}
}

/* UpdateCloudStorageConfigurationOK describes a response with status code 200, with default header values.

New state for specified model set.
*/
type UpdateCloudStorageConfigurationOK struct {
	Payload *models.BackupConfiguration
}

func (o *UpdateCloudStorageConfigurationOK) Error() string {
	return fmt.Sprintf("[PATCH /cloud_storage][%d] updateCloudStorageConfigurationOK  %+v", 200, o.Payload)
}
func (o *UpdateCloudStorageConfigurationOK) GetPayload() *models.BackupConfiguration {
	return o.Payload
}

func (o *UpdateCloudStorageConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudStorageConfigurationBadRequest creates a UpdateCloudStorageConfigurationBadRequest with default headers values
func NewUpdateCloudStorageConfigurationBadRequest() *UpdateCloudStorageConfigurationBadRequest {
	return &UpdateCloudStorageConfigurationBadRequest{}
}

/* UpdateCloudStorageConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCloudStorageConfigurationBadRequest struct {
	Payload *models.Error
}

func (o *UpdateCloudStorageConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /cloud_storage][%d] updateCloudStorageConfigurationBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCloudStorageConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCloudStorageConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudStorageConfigurationNotFound creates a UpdateCloudStorageConfigurationNotFound with default headers values
func NewUpdateCloudStorageConfigurationNotFound() *UpdateCloudStorageConfigurationNotFound {
	return &UpdateCloudStorageConfigurationNotFound{}
}

/* UpdateCloudStorageConfigurationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateCloudStorageConfigurationNotFound struct {
	Payload *models.Error
}

func (o *UpdateCloudStorageConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /cloud_storage][%d] updateCloudStorageConfigurationNotFound  %+v", 404, o.Payload)
}
func (o *UpdateCloudStorageConfigurationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCloudStorageConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudStorageConfigurationUnprocessableEntity creates a UpdateCloudStorageConfigurationUnprocessableEntity with default headers values
func NewUpdateCloudStorageConfigurationUnprocessableEntity() *UpdateCloudStorageConfigurationUnprocessableEntity {
	return &UpdateCloudStorageConfigurationUnprocessableEntity{}
}

/* UpdateCloudStorageConfigurationUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateCloudStorageConfigurationUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateCloudStorageConfigurationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /cloud_storage][%d] updateCloudStorageConfigurationUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateCloudStorageConfigurationUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateCloudStorageConfigurationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
