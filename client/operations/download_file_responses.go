// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DownloadFileReader is a Reader for the DownloadFile structure.
type DownloadFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadFileOK creates a DownloadFileOK with default headers values
func NewDownloadFileOK(writer io.Writer) *DownloadFileOK {
	return &DownloadFileOK{

		Payload: writer,
	}
}

/* DownloadFileOK describes a response with status code 200, with default header values.

Requests succeeded. The file has been downloaded.
*/
type DownloadFileOK struct {
	ContentDisposition string
	ContentType        string

	Payload io.Writer
}

func (o *DownloadFileOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileOK  %+v", 200, o.Payload)
}
func (o *DownloadFileOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Disposition
	hdrContentDisposition := response.GetHeader("Content-Disposition")

	if hdrContentDisposition != "" {
		o.ContentDisposition = hdrContentDisposition
	}

	// hydrates response header Content-Type
	hdrContentType := response.GetHeader("Content-Type")

	if hdrContentType != "" {
		o.ContentType = hdrContentType
	}

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileBadRequest creates a DownloadFileBadRequest with default headers values
func NewDownloadFileBadRequest() *DownloadFileBadRequest {
	return &DownloadFileBadRequest{}
}

/* DownloadFileBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type DownloadFileBadRequest struct {
}

func (o *DownloadFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileBadRequest ", 400)
}

func (o *DownloadFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadFileForbidden creates a DownloadFileForbidden with default headers values
func NewDownloadFileForbidden() *DownloadFileForbidden {
	return &DownloadFileForbidden{}
}

/* DownloadFileForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type DownloadFileForbidden struct {
	Payload *DownloadFileForbiddenBody
}

func (o *DownloadFileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileForbidden  %+v", 403, o.Payload)
}
func (o *DownloadFileForbidden) GetPayload() *DownloadFileForbiddenBody {
	return o.Payload
}

func (o *DownloadFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DownloadFileForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileNotFound creates a DownloadFileNotFound with default headers values
func NewDownloadFileNotFound() *DownloadFileNotFound {
	return &DownloadFileNotFound{}
}

/* DownloadFileNotFound describes a response with status code 404, with default header values.

Request failed. `file_name` does not exist .
*/
type DownloadFileNotFound struct {
	Payload *DownloadFileNotFoundBody
}

func (o *DownloadFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileNotFound  %+v", 404, o.Payload)
}
func (o *DownloadFileNotFound) GetPayload() *DownloadFileNotFoundBody {
	return o.Payload
}

func (o *DownloadFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DownloadFileNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileInternalServerError creates a DownloadFileInternalServerError with default headers values
func NewDownloadFileInternalServerError() *DownloadFileInternalServerError {
	return &DownloadFileInternalServerError{}
}

/* DownloadFileInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type DownloadFileInternalServerError struct {
	Payload *DownloadFileInternalServerErrorBody
}

func (o *DownloadFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileInternalServerError  %+v", 500, o.Payload)
}
func (o *DownloadFileInternalServerError) GetPayload() *DownloadFileInternalServerErrorBody {
	return o.Payload
}

func (o *DownloadFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DownloadFileInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DownloadFileForbiddenBody download file forbidden body
swagger:model DownloadFileForbiddenBody
*/
type DownloadFileForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this download file forbidden body
func (o *DownloadFileForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this download file forbidden body based on context it is used
func (o *DownloadFileForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DownloadFileForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DownloadFileForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DownloadFileForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DownloadFileInternalServerErrorBody download file internal server error body
swagger:model DownloadFileInternalServerErrorBody
*/
type DownloadFileInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this download file internal server error body
func (o *DownloadFileInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this download file internal server error body based on context it is used
func (o *DownloadFileInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DownloadFileInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DownloadFileInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DownloadFileInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DownloadFileNotFoundBody download file not found body
swagger:model DownloadFileNotFoundBody
*/
type DownloadFileNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this download file not found body
func (o *DownloadFileNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this download file not found body based on context it is used
func (o *DownloadFileNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DownloadFileNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DownloadFileNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DownloadFileNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
