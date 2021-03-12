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

// FolderChildrenReader is a Reader for the FolderChildren structure.
type FolderChildrenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FolderChildrenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFolderChildrenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFolderChildrenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFolderChildrenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFolderChildrenOK creates a FolderChildrenOK with default headers values
func NewFolderChildrenOK() *FolderChildrenOK {
	return &FolderChildrenOK{}
}

/* FolderChildrenOK describes a response with status code 200, with default header values.

Folders
*/
type FolderChildrenOK struct {
	Payload []*models.Folder
}

func (o *FolderChildrenOK) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_id}/children][%d] folderChildrenOK  %+v", 200, o.Payload)
}
func (o *FolderChildrenOK) GetPayload() []*models.Folder {
	return o.Payload
}

func (o *FolderChildrenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFolderChildrenBadRequest creates a FolderChildrenBadRequest with default headers values
func NewFolderChildrenBadRequest() *FolderChildrenBadRequest {
	return &FolderChildrenBadRequest{}
}

/* FolderChildrenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FolderChildrenBadRequest struct {
	Payload *models.Error
}

func (o *FolderChildrenBadRequest) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_id}/children][%d] folderChildrenBadRequest  %+v", 400, o.Payload)
}
func (o *FolderChildrenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *FolderChildrenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFolderChildrenNotFound creates a FolderChildrenNotFound with default headers values
func NewFolderChildrenNotFound() *FolderChildrenNotFound {
	return &FolderChildrenNotFound{}
}

/* FolderChildrenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FolderChildrenNotFound struct {
	Payload *models.Error
}

func (o *FolderChildrenNotFound) Error() string {
	return fmt.Sprintf("[GET /folders/{folder_id}/children][%d] folderChildrenNotFound  %+v", 404, o.Payload)
}
func (o *FolderChildrenNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *FolderChildrenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
