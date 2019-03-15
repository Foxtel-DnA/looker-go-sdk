// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// AllContentMetadatasReader is a Reader for the AllContentMetadatas structure.
type AllContentMetadatasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllContentMetadatasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllContentMetadatasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllContentMetadatasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllContentMetadatasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllContentMetadatasOK creates a AllContentMetadatasOK with default headers values
func NewAllContentMetadatasOK() *AllContentMetadatasOK {
	return &AllContentMetadatasOK{}
}

/*AllContentMetadatasOK handles this case with default header values.

Content Metadata
*/
type AllContentMetadatasOK struct {
	Payload []*models.ContentMeta
}

func (o *AllContentMetadatasOK) Error() string {
	return fmt.Sprintf("[GET /content_metadata][%d] allContentMetadatasOK  %+v", 200, o.Payload)
}

func (o *AllContentMetadatasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllContentMetadatasBadRequest creates a AllContentMetadatasBadRequest with default headers values
func NewAllContentMetadatasBadRequest() *AllContentMetadatasBadRequest {
	return &AllContentMetadatasBadRequest{}
}

/*AllContentMetadatasBadRequest handles this case with default header values.

Bad Request
*/
type AllContentMetadatasBadRequest struct {
	Payload *models.Error
}

func (o *AllContentMetadatasBadRequest) Error() string {
	return fmt.Sprintf("[GET /content_metadata][%d] allContentMetadatasBadRequest  %+v", 400, o.Payload)
}

func (o *AllContentMetadatasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllContentMetadatasNotFound creates a AllContentMetadatasNotFound with default headers values
func NewAllContentMetadatasNotFound() *AllContentMetadatasNotFound {
	return &AllContentMetadatasNotFound{}
}

/*AllContentMetadatasNotFound handles this case with default header values.

Not Found
*/
type AllContentMetadatasNotFound struct {
	Payload *models.Error
}

func (o *AllContentMetadatasNotFound) Error() string {
	return fmt.Sprintf("[GET /content_metadata][%d] allContentMetadatasNotFound  %+v", 404, o.Payload)
}

func (o *AllContentMetadatasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}