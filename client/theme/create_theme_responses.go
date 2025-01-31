// Code generated by go-swagger; DO NOT EDIT.

package theme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// CreateThemeReader is a Reader for the CreateTheme structure.
type CreateThemeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateThemeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateThemeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateThemeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateThemeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateThemeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateThemeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateThemeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateThemeOK creates a CreateThemeOK with default headers values
func NewCreateThemeOK() *CreateThemeOK {
	return &CreateThemeOK{}
}

/* CreateThemeOK describes a response with status code 200, with default header values.

Theme
*/
type CreateThemeOK struct {
	Payload *models.Theme
}

func (o *CreateThemeOK) Error() string {
	return fmt.Sprintf("[POST /themes][%d] createThemeOK  %+v", 200, o.Payload)
}
func (o *CreateThemeOK) GetPayload() *models.Theme {
	return o.Payload
}

func (o *CreateThemeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Theme)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateThemeBadRequest creates a CreateThemeBadRequest with default headers values
func NewCreateThemeBadRequest() *CreateThemeBadRequest {
	return &CreateThemeBadRequest{}
}

/* CreateThemeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateThemeBadRequest struct {
	Payload *models.Error
}

func (o *CreateThemeBadRequest) Error() string {
	return fmt.Sprintf("[POST /themes][%d] createThemeBadRequest  %+v", 400, o.Payload)
}
func (o *CreateThemeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateThemeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateThemeNotFound creates a CreateThemeNotFound with default headers values
func NewCreateThemeNotFound() *CreateThemeNotFound {
	return &CreateThemeNotFound{}
}

/* CreateThemeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateThemeNotFound struct {
	Payload *models.Error
}

func (o *CreateThemeNotFound) Error() string {
	return fmt.Sprintf("[POST /themes][%d] createThemeNotFound  %+v", 404, o.Payload)
}
func (o *CreateThemeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateThemeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateThemeConflict creates a CreateThemeConflict with default headers values
func NewCreateThemeConflict() *CreateThemeConflict {
	return &CreateThemeConflict{}
}

/* CreateThemeConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateThemeConflict struct {
	Payload *models.Error
}

func (o *CreateThemeConflict) Error() string {
	return fmt.Sprintf("[POST /themes][%d] createThemeConflict  %+v", 409, o.Payload)
}
func (o *CreateThemeConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateThemeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateThemeUnprocessableEntity creates a CreateThemeUnprocessableEntity with default headers values
func NewCreateThemeUnprocessableEntity() *CreateThemeUnprocessableEntity {
	return &CreateThemeUnprocessableEntity{}
}

/* CreateThemeUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateThemeUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateThemeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /themes][%d] createThemeUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateThemeUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateThemeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateThemeTooManyRequests creates a CreateThemeTooManyRequests with default headers values
func NewCreateThemeTooManyRequests() *CreateThemeTooManyRequests {
	return &CreateThemeTooManyRequests{}
}

/* CreateThemeTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateThemeTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateThemeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /themes][%d] createThemeTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateThemeTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateThemeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
