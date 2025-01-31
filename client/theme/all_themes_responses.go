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

// AllThemesReader is a Reader for the AllThemes structure.
type AllThemesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllThemesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllThemesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllThemesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllThemesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllThemesOK creates a AllThemesOK with default headers values
func NewAllThemesOK() *AllThemesOK {
	return &AllThemesOK{}
}

/* AllThemesOK describes a response with status code 200, with default header values.

Themes
*/
type AllThemesOK struct {
	Payload []*models.Theme
}

func (o *AllThemesOK) Error() string {
	return fmt.Sprintf("[GET /themes][%d] allThemesOK  %+v", 200, o.Payload)
}
func (o *AllThemesOK) GetPayload() []*models.Theme {
	return o.Payload
}

func (o *AllThemesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllThemesBadRequest creates a AllThemesBadRequest with default headers values
func NewAllThemesBadRequest() *AllThemesBadRequest {
	return &AllThemesBadRequest{}
}

/* AllThemesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllThemesBadRequest struct {
	Payload *models.Error
}

func (o *AllThemesBadRequest) Error() string {
	return fmt.Sprintf("[GET /themes][%d] allThemesBadRequest  %+v", 400, o.Payload)
}
func (o *AllThemesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllThemesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllThemesNotFound creates a AllThemesNotFound with default headers values
func NewAllThemesNotFound() *AllThemesNotFound {
	return &AllThemesNotFound{}
}

/* AllThemesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllThemesNotFound struct {
	Payload *models.Error
}

func (o *AllThemesNotFound) Error() string {
	return fmt.Sprintf("[GET /themes][%d] allThemesNotFound  %+v", 404, o.Payload)
}
func (o *AllThemesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllThemesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
