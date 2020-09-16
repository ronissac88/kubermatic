// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/api/utils/apiclient/models"
)

// GetClusterHealthV2Reader is a Reader for the GetClusterHealthV2 structure.
type GetClusterHealthV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterHealthV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterHealthV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterHealthV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterHealthV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterHealthV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterHealthV2OK creates a GetClusterHealthV2OK with default headers values
func NewGetClusterHealthV2OK() *GetClusterHealthV2OK {
	return &GetClusterHealthV2OK{}
}

/*GetClusterHealthV2OK handles this case with default header values.

ClusterHealth
*/
type GetClusterHealthV2OK struct {
	Payload *models.ClusterHealth
}

func (o *GetClusterHealthV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/health][%d] getClusterHealthV2OK  %+v", 200, o.Payload)
}

func (o *GetClusterHealthV2OK) GetPayload() *models.ClusterHealth {
	return o.Payload
}

func (o *GetClusterHealthV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterHealthV2Unauthorized creates a GetClusterHealthV2Unauthorized with default headers values
func NewGetClusterHealthV2Unauthorized() *GetClusterHealthV2Unauthorized {
	return &GetClusterHealthV2Unauthorized{}
}

/*GetClusterHealthV2Unauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterHealthV2Unauthorized struct {
}

func (o *GetClusterHealthV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/health][%d] getClusterHealthV2Unauthorized ", 401)
}

func (o *GetClusterHealthV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterHealthV2Forbidden creates a GetClusterHealthV2Forbidden with default headers values
func NewGetClusterHealthV2Forbidden() *GetClusterHealthV2Forbidden {
	return &GetClusterHealthV2Forbidden{}
}

/*GetClusterHealthV2Forbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterHealthV2Forbidden struct {
}

func (o *GetClusterHealthV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/health][%d] getClusterHealthV2Forbidden ", 403)
}

func (o *GetClusterHealthV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterHealthV2Default creates a GetClusterHealthV2Default with default headers values
func NewGetClusterHealthV2Default(code int) *GetClusterHealthV2Default {
	return &GetClusterHealthV2Default{
		_statusCode: code,
	}
}

/*GetClusterHealthV2Default handles this case with default header values.

errorResponse
*/
type GetClusterHealthV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster health v2 default response
func (o *GetClusterHealthV2Default) Code() int {
	return o._statusCode
}

func (o *GetClusterHealthV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/health][%d] getClusterHealthV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterHealthV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterHealthV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
