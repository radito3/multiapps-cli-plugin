// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	baseclient "github.com/cloudfoundry-incubator/multiapps-cli-plugin/clients/baseclient"
	"github.com/cloudfoundry-incubator/multiapps-cli-plugin/clients/models"
)

// GetMtaOperationLogsReader is a Reader for the GetMtaOperationLogs structure.
type GetMtaOperationLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMtaOperationLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMtaOperationLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, baseclient.BuildErrorResponse(response, consumer, o.formats)
	}
}

// NewGetMtaOperationLogsOK creates a GetMtaOperationLogsOK with default headers values
func NewGetMtaOperationLogsOK() *GetMtaOperationLogsOK {
	return &GetMtaOperationLogsOK{}
}

/*GetMtaOperationLogsOK handles this case with default header values.

OK
*/
type GetMtaOperationLogsOK struct {
	Payload models.GetMtaOperationLogsOKBody
}

func (o *GetMtaOperationLogsOK) Error() string {
	return fmt.Sprintf("[GET /operations/{operationId}/logs][%d] getMtaOperationLogsOK  %+v", 200, o.Payload)
}

func (o *GetMtaOperationLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
