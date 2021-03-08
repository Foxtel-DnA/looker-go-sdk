// Code generated by go-swagger; DO NOT EDIT.

package theme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// SetDefaultThemeReader is a Reader for the SetDefaultTheme structure.
type SetDefaultThemeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDefaultThemeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDefaultThemeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetDefaultThemeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetDefaultThemeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSetDefaultThemeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSetDefaultThemeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetDefaultThemeOK creates a SetDefaultThemeOK with default headers values
func NewSetDefaultThemeOK() *SetDefaultThemeOK {
	return &SetDefaultThemeOK{}
}

/* SetDefaultThemeOK describes a response with status code 200, with default header values.

Theme
*/
type SetDefaultThemeOK struct {
	Payload *models.Theme
}

func (o *SetDefaultThemeOK) Error() string {
	return fmt.Sprintf("[PUT /themes/default][%d] setDefaultThemeOK  %+v", 200, o.Payload)
}
func (o *SetDefaultThemeOK) GetPayload() *models.Theme {
	return o.Payload
}

func (o *SetDefaultThemeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Theme)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultThemeBadRequest creates a SetDefaultThemeBadRequest with default headers values
func NewSetDefaultThemeBadRequest() *SetDefaultThemeBadRequest {
	return &SetDefaultThemeBadRequest{}
}

/* SetDefaultThemeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SetDefaultThemeBadRequest struct {
	Payload *models.Error
}

func (o *SetDefaultThemeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /themes/default][%d] setDefaultThemeBadRequest  %+v", 400, o.Payload)
}
func (o *SetDefaultThemeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDefaultThemeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultThemeNotFound creates a SetDefaultThemeNotFound with default headers values
func NewSetDefaultThemeNotFound() *SetDefaultThemeNotFound {
	return &SetDefaultThemeNotFound{}
}

/* SetDefaultThemeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SetDefaultThemeNotFound struct {
	Payload *models.Error
}

func (o *SetDefaultThemeNotFound) Error() string {
	return fmt.Sprintf("[PUT /themes/default][%d] setDefaultThemeNotFound  %+v", 404, o.Payload)
}
func (o *SetDefaultThemeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDefaultThemeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultThemeUnprocessableEntity creates a SetDefaultThemeUnprocessableEntity with default headers values
func NewSetDefaultThemeUnprocessableEntity() *SetDefaultThemeUnprocessableEntity {
	return &SetDefaultThemeUnprocessableEntity{}
}

/* SetDefaultThemeUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type SetDefaultThemeUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *SetDefaultThemeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /themes/default][%d] setDefaultThemeUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *SetDefaultThemeUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *SetDefaultThemeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultThemeTooManyRequests creates a SetDefaultThemeTooManyRequests with default headers values
func NewSetDefaultThemeTooManyRequests() *SetDefaultThemeTooManyRequests {
	return &SetDefaultThemeTooManyRequests{}
}

/* SetDefaultThemeTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type SetDefaultThemeTooManyRequests struct {
	Payload *models.Error
}

func (o *SetDefaultThemeTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /themes/default][%d] setDefaultThemeTooManyRequests  %+v", 429, o.Payload)
}
func (o *SetDefaultThemeTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDefaultThemeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}