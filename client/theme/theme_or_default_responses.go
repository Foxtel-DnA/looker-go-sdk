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

// ThemeOrDefaultReader is a Reader for the ThemeOrDefault structure.
type ThemeOrDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThemeOrDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewThemeOrDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewThemeOrDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewThemeOrDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewThemeOrDefaultOK creates a ThemeOrDefaultOK with default headers values
func NewThemeOrDefaultOK() *ThemeOrDefaultOK {
	return &ThemeOrDefaultOK{}
}

/* ThemeOrDefaultOK describes a response with status code 200, with default header values.

Theme
*/
type ThemeOrDefaultOK struct {
	Payload *models.Theme
}

func (o *ThemeOrDefaultOK) Error() string {
	return fmt.Sprintf("[GET /themes/theme_or_default][%d] themeOrDefaultOK  %+v", 200, o.Payload)
}
func (o *ThemeOrDefaultOK) GetPayload() *models.Theme {
	return o.Payload
}

func (o *ThemeOrDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Theme)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThemeOrDefaultBadRequest creates a ThemeOrDefaultBadRequest with default headers values
func NewThemeOrDefaultBadRequest() *ThemeOrDefaultBadRequest {
	return &ThemeOrDefaultBadRequest{}
}

/* ThemeOrDefaultBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ThemeOrDefaultBadRequest struct {
	Payload *models.Error
}

func (o *ThemeOrDefaultBadRequest) Error() string {
	return fmt.Sprintf("[GET /themes/theme_or_default][%d] themeOrDefaultBadRequest  %+v", 400, o.Payload)
}
func (o *ThemeOrDefaultBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ThemeOrDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThemeOrDefaultNotFound creates a ThemeOrDefaultNotFound with default headers values
func NewThemeOrDefaultNotFound() *ThemeOrDefaultNotFound {
	return &ThemeOrDefaultNotFound{}
}

/* ThemeOrDefaultNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ThemeOrDefaultNotFound struct {
	Payload *models.Error
}

func (o *ThemeOrDefaultNotFound) Error() string {
	return fmt.Sprintf("[GET /themes/theme_or_default][%d] themeOrDefaultNotFound  %+v", 404, o.Payload)
}
func (o *ThemeOrDefaultNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ThemeOrDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
