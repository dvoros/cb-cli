// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/*GetStatusOK handles this case with default header values.

ok
*/
type GetStatusOK struct {
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusOK ", 200)
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStatusInternalServerError creates a GetStatusInternalServerError with default headers values
func NewGetStatusInternalServerError() *GetStatusInternalServerError {
	return &GetStatusInternalServerError{}
}

/*GetStatusInternalServerError handles this case with default header values.

Internal server error
*/
type GetStatusInternalServerError struct {
}

func (o *GetStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusInternalServerError ", 500)
}

func (o *GetStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}