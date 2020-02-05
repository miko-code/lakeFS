// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/treeverse/lakefs/api/gen/models"
)

// DeleteBranchReader is a Reader for the DeleteBranch structure.
type DeleteBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBranchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteBranchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteBranchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteBranchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBranchNoContent creates a DeleteBranchNoContent with default headers values
func NewDeleteBranchNoContent() *DeleteBranchNoContent {
	return &DeleteBranchNoContent{}
}

/*DeleteBranchNoContent handles this case with default header values.

branch deleted successfully
*/
type DeleteBranchNoContent struct {
}

func (o *DeleteBranchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repositoryId}/branches/{branchId}][%d] deleteBranchNoContent ", 204)
}

func (o *DeleteBranchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBranchUnauthorized creates a DeleteBranchUnauthorized with default headers values
func NewDeleteBranchUnauthorized() *DeleteBranchUnauthorized {
	return &DeleteBranchUnauthorized{}
}

/*DeleteBranchUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteBranchUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteBranchUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repositoryId}/branches/{branchId}][%d] deleteBranchUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteBranchUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBranchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBranchNotFound creates a DeleteBranchNotFound with default headers values
func NewDeleteBranchNotFound() *DeleteBranchNotFound {
	return &DeleteBranchNotFound{}
}

/*DeleteBranchNotFound handles this case with default header values.

branch not found
*/
type DeleteBranchNotFound struct {
	Payload *models.Error
}

func (o *DeleteBranchNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repositoryId}/branches/{branchId}][%d] deleteBranchNotFound  %+v", 404, o.Payload)
}

func (o *DeleteBranchNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBranchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBranchDefault creates a DeleteBranchDefault with default headers values
func NewDeleteBranchDefault(code int) *DeleteBranchDefault {
	return &DeleteBranchDefault{
		_statusCode: code,
	}
}

/*DeleteBranchDefault handles this case with default header values.

generic error response
*/
type DeleteBranchDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete branch default response
func (o *DeleteBranchDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBranchDefault) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repositoryId}/branches/{branchId}][%d] deleteBranch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteBranchDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBranchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
