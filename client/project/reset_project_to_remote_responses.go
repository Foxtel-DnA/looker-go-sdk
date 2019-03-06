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

// ResetProjectToRemoteReader is a Reader for the ResetProjectToRemote structure.
type ResetProjectToRemoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetProjectToRemoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResetProjectToRemoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewResetProjectToRemoteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewResetProjectToRemoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewResetProjectToRemoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResetProjectToRemoteOK creates a ResetProjectToRemoteOK with default headers values
func NewResetProjectToRemoteOK() *ResetProjectToRemoteOK {
	return &ResetProjectToRemoteOK{}
}

/*ResetProjectToRemoteOK handles this case with default header values.

Project
*/
type ResetProjectToRemoteOK struct {
	Payload string
}

func (o *ResetProjectToRemoteOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_remote][%d] resetProjectToRemoteOK  %+v", 200, o.Payload)
}

func (o *ResetProjectToRemoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetProjectToRemoteNoContent creates a ResetProjectToRemoteNoContent with default headers values
func NewResetProjectToRemoteNoContent() *ResetProjectToRemoteNoContent {
	return &ResetProjectToRemoteNoContent{}
}

/*ResetProjectToRemoteNoContent handles this case with default header values.

Returns 204 if project was successfully reset, otherwise 400 with an error message
*/
type ResetProjectToRemoteNoContent struct {
}

func (o *ResetProjectToRemoteNoContent) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_remote][%d] resetProjectToRemoteNoContent ", 204)
}

func (o *ResetProjectToRemoteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetProjectToRemoteBadRequest creates a ResetProjectToRemoteBadRequest with default headers values
func NewResetProjectToRemoteBadRequest() *ResetProjectToRemoteBadRequest {
	return &ResetProjectToRemoteBadRequest{}
}

/*ResetProjectToRemoteBadRequest handles this case with default header values.

Bad Request
*/
type ResetProjectToRemoteBadRequest struct {
	Payload *models.Error
}

func (o *ResetProjectToRemoteBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_remote][%d] resetProjectToRemoteBadRequest  %+v", 400, o.Payload)
}

func (o *ResetProjectToRemoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetProjectToRemoteNotFound creates a ResetProjectToRemoteNotFound with default headers values
func NewResetProjectToRemoteNotFound() *ResetProjectToRemoteNotFound {
	return &ResetProjectToRemoteNotFound{}
}

/*ResetProjectToRemoteNotFound handles this case with default header values.

Not Found
*/
type ResetProjectToRemoteNotFound struct {
	Payload *models.Error
}

func (o *ResetProjectToRemoteNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/reset_to_remote][%d] resetProjectToRemoteNotFound  %+v", 404, o.Payload)
}

func (o *ResetProjectToRemoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
