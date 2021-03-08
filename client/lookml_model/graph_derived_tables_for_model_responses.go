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

// GraphDerivedTablesForModelReader is a Reader for the GraphDerivedTablesForModel structure.
type GraphDerivedTablesForModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GraphDerivedTablesForModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGraphDerivedTablesForModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGraphDerivedTablesForModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGraphDerivedTablesForModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGraphDerivedTablesForModelOK creates a GraphDerivedTablesForModelOK with default headers values
func NewGraphDerivedTablesForModelOK() *GraphDerivedTablesForModelOK {
	return &GraphDerivedTablesForModelOK{}
}

/* GraphDerivedTablesForModelOK describes a response with status code 200, with default header values.

Derived Table
*/
type GraphDerivedTablesForModelOK struct {
	Payload *models.DependencyGraph
}

func (o *GraphDerivedTablesForModelOK) Error() string {
	return fmt.Sprintf("[GET /derived_table/graph/model/{model}][%d] graphDerivedTablesForModelOK  %+v", 200, o.Payload)
}
func (o *GraphDerivedTablesForModelOK) GetPayload() *models.DependencyGraph {
	return o.Payload
}

func (o *GraphDerivedTablesForModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DependencyGraph)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphDerivedTablesForModelBadRequest creates a GraphDerivedTablesForModelBadRequest with default headers values
func NewGraphDerivedTablesForModelBadRequest() *GraphDerivedTablesForModelBadRequest {
	return &GraphDerivedTablesForModelBadRequest{}
}

/* GraphDerivedTablesForModelBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GraphDerivedTablesForModelBadRequest struct {
	Payload *models.Error
}

func (o *GraphDerivedTablesForModelBadRequest) Error() string {
	return fmt.Sprintf("[GET /derived_table/graph/model/{model}][%d] graphDerivedTablesForModelBadRequest  %+v", 400, o.Payload)
}
func (o *GraphDerivedTablesForModelBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GraphDerivedTablesForModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphDerivedTablesForModelNotFound creates a GraphDerivedTablesForModelNotFound with default headers values
func NewGraphDerivedTablesForModelNotFound() *GraphDerivedTablesForModelNotFound {
	return &GraphDerivedTablesForModelNotFound{}
}

/* GraphDerivedTablesForModelNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GraphDerivedTablesForModelNotFound struct {
	Payload *models.Error
}

func (o *GraphDerivedTablesForModelNotFound) Error() string {
	return fmt.Sprintf("[GET /derived_table/graph/model/{model}][%d] graphDerivedTablesForModelNotFound  %+v", 404, o.Payload)
}
func (o *GraphDerivedTablesForModelNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GraphDerivedTablesForModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
