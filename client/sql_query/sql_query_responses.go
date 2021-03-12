// Code generated by go-swagger; DO NOT EDIT.

package sql_query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Foxtel-DnA/looker-go-sdk/models"
)

// SQLQueryReader is a Reader for the SQLQuery structure.
type SQLQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SQLQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSQLQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSQLQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSQLQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSQLQueryOK creates a SQLQueryOK with default headers values
func NewSQLQueryOK() *SQLQueryOK {
	return &SQLQueryOK{}
}

/*SQLQueryOK handles this case with default header values.

SQL Runner Query
*/
type SQLQueryOK struct {
	Payload *models.SQLQuery
}

func (o *SQLQueryOK) Error() string {
	return fmt.Sprintf("[GET /sql_queries/{slug}][%d] sqlQueryOK  %+v", 200, o.Payload)
}

func (o *SQLQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SQLQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSQLQueryBadRequest creates a SQLQueryBadRequest with default headers values
func NewSQLQueryBadRequest() *SQLQueryBadRequest {
	return &SQLQueryBadRequest{}
}

/*SQLQueryBadRequest handles this case with default header values.

Bad Request
*/
type SQLQueryBadRequest struct {
	Payload *models.Error
}

func (o *SQLQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /sql_queries/{slug}][%d] sqlQueryBadRequest  %+v", 400, o.Payload)
}

func (o *SQLQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSQLQueryNotFound creates a SQLQueryNotFound with default headers values
func NewSQLQueryNotFound() *SQLQueryNotFound {
	return &SQLQueryNotFound{}
}

/*SQLQueryNotFound handles this case with default header values.

Not Found
*/
type SQLQueryNotFound struct {
	Payload *models.Error
}

func (o *SQLQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /sql_queries/{slug}][%d] sqlQueryNotFound  %+v", 404, o.Payload)
}

func (o *SQLQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
