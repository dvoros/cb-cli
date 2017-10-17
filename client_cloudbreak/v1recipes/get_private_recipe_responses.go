// Code generated by go-swagger; DO NOT EDIT.

package v1recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivateRecipeReader is a Reader for the GetPrivateRecipe structure.
type GetPrivateRecipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateRecipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateRecipeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateRecipeOK creates a GetPrivateRecipeOK with default headers values
func NewGetPrivateRecipeOK() *GetPrivateRecipeOK {
	return &GetPrivateRecipeOK{}
}

/*GetPrivateRecipeOK handles this case with default header values.

successful operation
*/
type GetPrivateRecipeOK struct {
	Payload *models_cloudbreak.RecipeResponse
}

func (o *GetPrivateRecipeOK) Error() string {
	return fmt.Sprintf("[GET /v1/recipes/user/{name}][%d] getPrivateRecipeOK  %+v", 200, o.Payload)
}

func (o *GetPrivateRecipeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RecipeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
