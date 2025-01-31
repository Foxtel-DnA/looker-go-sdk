// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Foxtel-DnA/looker-go-sdk/models"
)

// FolderAncestorsReader is a Reader for the FolderAncestors structure.
type FolderAncestorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FolderAncestorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFolderAncestorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFolderAncestorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFolderAncestorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFolderAncestorsOK creates a FolderAncestorsOK with default headers values
func NewFolderAncestorsOK() *FolderAncestorsOK {
	return &FolderAncestorsOK{}
}

/* FolderAncestorsOK describes a response with status code 200, with default header values.

Folders
*/
type FolderAncestorsOK struct {
	Payload []*models.Folder
}

func (o *FolderAncestorsOK) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_id}/ancestors][%d] folderAncestorsOK  %+v", 200, o.Payload)
}
func (o *FolderAncestorsOK) GetPayload() []*models.Folder {
	return o.Payload
}

func (o *FolderAncestorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFolderAncestorsBadRequest creates a FolderAncestorsBadRequest with default headers values
func NewFolderAncestorsBadRequest() *FolderAncestorsBadRequest {
	return &FolderAncestorsBadRequest{}
}

/* FolderAncestorsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FolderAncestorsBadRequest struct {
	Payload *models.Error
}

func (o *FolderAncestorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_id}/ancestors][%d] folderAncestorsBadRequest  %+v", 400, o.Payload)
}
func (o *FolderAncestorsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *FolderAncestorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFolderAncestorsNotFound creates a FolderAncestorsNotFound with default headers values
func NewFolderAncestorsNotFound() *FolderAncestorsNotFound {
	return &FolderAncestorsNotFound{}
}

/* FolderAncestorsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FolderAncestorsNotFound struct {
	Payload *models.Error
}

func (o *FolderAncestorsNotFound) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_id}/ancestors][%d] folderAncestorsNotFound  %+v", 404, o.Payload)
}
func (o *FolderAncestorsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *FolderAncestorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
