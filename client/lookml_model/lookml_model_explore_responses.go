// Code generated by go-swagger; DO NOT EDIT.

package lookml_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// LookmlModelExploreReader is a Reader for the LookmlModelExplore structure.
type LookmlModelExploreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookmlModelExploreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLookmlModelExploreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLookmlModelExploreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLookmlModelExploreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLookmlModelExploreOK creates a LookmlModelExploreOK with default headers values
func NewLookmlModelExploreOK() *LookmlModelExploreOK {
	return &LookmlModelExploreOK{}
}

/* LookmlModelExploreOK describes a response with status code 200, with default header values.

LookML Model Explore
*/
type LookmlModelExploreOK struct {
	Payload *models.LookmlModelExplore
}

func (o *LookmlModelExploreOK) Error() string {
	return fmt.Sprintf("[GET /lookml_models/{lookml_model_name}/explores/{explore_name}][%d] lookmlModelExploreOK  %+v", 200, o.Payload)
}
func (o *LookmlModelExploreOK) GetPayload() *models.LookmlModelExplore {
	return o.Payload
}

func (o *LookmlModelExploreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookmlModelExplore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookmlModelExploreBadRequest creates a LookmlModelExploreBadRequest with default headers values
func NewLookmlModelExploreBadRequest() *LookmlModelExploreBadRequest {
	return &LookmlModelExploreBadRequest{}
}

/* LookmlModelExploreBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type LookmlModelExploreBadRequest struct {
	Payload *models.Error
}

func (o *LookmlModelExploreBadRequest) Error() string {
	return fmt.Sprintf("[GET /lookml_models/{lookml_model_name}/explores/{explore_name}][%d] lookmlModelExploreBadRequest  %+v", 400, o.Payload)
}
func (o *LookmlModelExploreBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *LookmlModelExploreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookmlModelExploreNotFound creates a LookmlModelExploreNotFound with default headers values
func NewLookmlModelExploreNotFound() *LookmlModelExploreNotFound {
	return &LookmlModelExploreNotFound{}
}

/* LookmlModelExploreNotFound describes a response with status code 404, with default header values.

Not Found
*/
type LookmlModelExploreNotFound struct {
	Payload *models.Error
}

func (o *LookmlModelExploreNotFound) Error() string {
	return fmt.Sprintf("[GET /lookml_models/{lookml_model_name}/explores/{explore_name}][%d] lookmlModelExploreNotFound  %+v", 404, o.Payload)
}
func (o *LookmlModelExploreNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *LookmlModelExploreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
