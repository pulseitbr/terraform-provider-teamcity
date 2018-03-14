// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// NewSetProjectsOrderParams creates a new SetProjectsOrderParams object
// with the default values initialized.
func NewSetProjectsOrderParams() *SetProjectsOrderParams {
	var ()
	return &SetProjectsOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetProjectsOrderParamsWithTimeout creates a new SetProjectsOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetProjectsOrderParamsWithTimeout(timeout time.Duration) *SetProjectsOrderParams {
	var ()
	return &SetProjectsOrderParams{

		timeout: timeout,
	}
}

// NewSetProjectsOrderParamsWithContext creates a new SetProjectsOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetProjectsOrderParamsWithContext(ctx context.Context) *SetProjectsOrderParams {
	var ()
	return &SetProjectsOrderParams{

		Context: ctx,
	}
}

// NewSetProjectsOrderParamsWithHTTPClient creates a new SetProjectsOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetProjectsOrderParamsWithHTTPClient(client *http.Client) *SetProjectsOrderParams {
	var ()
	return &SetProjectsOrderParams{
		HTTPClient: client,
	}
}

/*SetProjectsOrderParams contains all the parameters to send to the API endpoint
for the set projects order operation typically these are written to a http.Request
*/
type SetProjectsOrderParams struct {

	/*Body*/
	Body *models.Projects
	/*Field*/
	Field string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set projects order params
func (o *SetProjectsOrderParams) WithTimeout(timeout time.Duration) *SetProjectsOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set projects order params
func (o *SetProjectsOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set projects order params
func (o *SetProjectsOrderParams) WithContext(ctx context.Context) *SetProjectsOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set projects order params
func (o *SetProjectsOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set projects order params
func (o *SetProjectsOrderParams) WithHTTPClient(client *http.Client) *SetProjectsOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set projects order params
func (o *SetProjectsOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set projects order params
func (o *SetProjectsOrderParams) WithBody(body *models.Projects) *SetProjectsOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set projects order params
func (o *SetProjectsOrderParams) SetBody(body *models.Projects) {
	o.Body = body
}

// WithField adds the field to the set projects order params
func (o *SetProjectsOrderParams) WithField(field string) *SetProjectsOrderParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the set projects order params
func (o *SetProjectsOrderParams) SetField(field string) {
	o.Field = field
}

// WithProjectLocator adds the projectLocator to the set projects order params
func (o *SetProjectsOrderParams) WithProjectLocator(projectLocator string) *SetProjectsOrderParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the set projects order params
func (o *SetProjectsOrderParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetProjectsOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}