// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/billtrust/looker-go-sdk/models"
)

// UpdateSessionConfigReader is a Reader for the UpdateSessionConfig structure.
type UpdateSessionConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSessionConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSessionConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSessionConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSessionConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateSessionConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateSessionConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSessionConfigOK creates a UpdateSessionConfigOK with default headers values
func NewUpdateSessionConfigOK() *UpdateSessionConfigOK {
	return &UpdateSessionConfigOK{}
}

/* UpdateSessionConfigOK describes a response with status code 200, with default header values.

Session Config
*/
type UpdateSessionConfigOK struct {
	Payload *models.SessionConfig
}

func (o *UpdateSessionConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /session_config][%d] updateSessionConfigOK  %+v", 200, o.Payload)
}
func (o *UpdateSessionConfigOK) GetPayload() *models.SessionConfig {
	return o.Payload
}

func (o *UpdateSessionConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSessionConfigBadRequest creates a UpdateSessionConfigBadRequest with default headers values
func NewUpdateSessionConfigBadRequest() *UpdateSessionConfigBadRequest {
	return &UpdateSessionConfigBadRequest{}
}

/* UpdateSessionConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateSessionConfigBadRequest struct {
	Payload *models.Error
}

func (o *UpdateSessionConfigBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /session_config][%d] updateSessionConfigBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateSessionConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSessionConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSessionConfigNotFound creates a UpdateSessionConfigNotFound with default headers values
func NewUpdateSessionConfigNotFound() *UpdateSessionConfigNotFound {
	return &UpdateSessionConfigNotFound{}
}

/* UpdateSessionConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateSessionConfigNotFound struct {
	Payload *models.Error
}

func (o *UpdateSessionConfigNotFound) Error() string {
	return fmt.Sprintf("[PATCH /session_config][%d] updateSessionConfigNotFound  %+v", 404, o.Payload)
}
func (o *UpdateSessionConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSessionConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSessionConfigUnprocessableEntity creates a UpdateSessionConfigUnprocessableEntity with default headers values
func NewUpdateSessionConfigUnprocessableEntity() *UpdateSessionConfigUnprocessableEntity {
	return &UpdateSessionConfigUnprocessableEntity{}
}

/* UpdateSessionConfigUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateSessionConfigUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateSessionConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /session_config][%d] updateSessionConfigUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateSessionConfigUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateSessionConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSessionConfigTooManyRequests creates a UpdateSessionConfigTooManyRequests with default headers values
func NewUpdateSessionConfigTooManyRequests() *UpdateSessionConfigTooManyRequests {
	return &UpdateSessionConfigTooManyRequests{}
}

/* UpdateSessionConfigTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateSessionConfigTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateSessionConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /session_config][%d] updateSessionConfigTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateSessionConfigTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSessionConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
