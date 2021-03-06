// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ResendMailReader is a Reader for the ResendMail structure.
type ResendMailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResendMailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResendMailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewResendMailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewResendMailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewResendMailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewResendMailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResendMailOK creates a ResendMailOK with default headers values
func NewResendMailOK() *ResendMailOK {
	return &ResendMailOK{}
}

/*ResendMailOK handles this case with default header values.

successful operation
*/
type ResendMailOK struct {
	Payload ResendMailOKBody
}

func (o *ResendMailOK) Error() string {
	return fmt.Sprintf("[POST /caas/api/tenants/{name}][%d] resendMailOK  %+v", 200, o.Payload)
}

func (o *ResendMailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResendMailBadRequest creates a ResendMailBadRequest with default headers values
func NewResendMailBadRequest() *ResendMailBadRequest {
	return &ResendMailBadRequest{}
}

/*ResendMailBadRequest handles this case with default header values.

Bad Request
*/
type ResendMailBadRequest struct {
}

func (o *ResendMailBadRequest) Error() string {
	return fmt.Sprintf("[POST /caas/api/tenants/{name}][%d] resendMailBadRequest ", 400)
}

func (o *ResendMailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendMailUnauthorized creates a ResendMailUnauthorized with default headers values
func NewResendMailUnauthorized() *ResendMailUnauthorized {
	return &ResendMailUnauthorized{}
}

/*ResendMailUnauthorized handles this case with default header values.

Unauthorized
*/
type ResendMailUnauthorized struct {
}

func (o *ResendMailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /caas/api/tenants/{name}][%d] resendMailUnauthorized ", 401)
}

func (o *ResendMailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendMailForbidden creates a ResendMailForbidden with default headers values
func NewResendMailForbidden() *ResendMailForbidden {
	return &ResendMailForbidden{}
}

/*ResendMailForbidden handles this case with default header values.

Forbidden
*/
type ResendMailForbidden struct {
}

func (o *ResendMailForbidden) Error() string {
	return fmt.Sprintf("[POST /caas/api/tenants/{name}][%d] resendMailForbidden ", 403)
}

func (o *ResendMailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResendMailInternalServerError creates a ResendMailInternalServerError with default headers values
func NewResendMailInternalServerError() *ResendMailInternalServerError {
	return &ResendMailInternalServerError{}
}

/*ResendMailInternalServerError handles this case with default header values.

Internal server error
*/
type ResendMailInternalServerError struct {
}

func (o *ResendMailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /caas/api/tenants/{name}][%d] resendMailInternalServerError ", 500)
}

func (o *ResendMailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ResendMailOKBody resend mail o k body
swagger:model ResendMailOKBody
*/

type ResendMailOKBody interface{}
