// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// ProjectReader is a Reader for the Project structure.
type ProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProjectOK creates a ProjectOK with default headers values
func NewProjectOK() *ProjectOK {
	return &ProjectOK{}
}

/*ProjectOK handles this case with default header values.

Project
*/
type ProjectOK struct {
	Payload *models.Project
}

func (o *ProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}][%d] projectOK  %+v", 200, o.Payload)
}

func (o *ProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectBadRequest creates a ProjectBadRequest with default headers values
func NewProjectBadRequest() *ProjectBadRequest {
	return &ProjectBadRequest{}
}

/*ProjectBadRequest handles this case with default header values.

Bad Request
*/
type ProjectBadRequest struct {
	Payload *models.Error
}

func (o *ProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}][%d] projectBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectNotFound creates a ProjectNotFound with default headers values
func NewProjectNotFound() *ProjectNotFound {
	return &ProjectNotFound{}
}

/*ProjectNotFound handles this case with default header values.

Not Found
*/
type ProjectNotFound struct {
	Payload *models.Error
}

func (o *ProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}][%d] projectNotFound  %+v", 404, o.Payload)
}

func (o *ProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
