// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adamluzsi/toggler/lib/go/models"
)

// NewIsFeatureEnabledParams creates a new IsFeatureEnabledParams object
// with the default values initialized.
func NewIsFeatureEnabledParams() *IsFeatureEnabledParams {
	var ()
	return &IsFeatureEnabledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsFeatureEnabledParamsWithTimeout creates a new IsFeatureEnabledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsFeatureEnabledParamsWithTimeout(timeout time.Duration) *IsFeatureEnabledParams {
	var ()
	return &IsFeatureEnabledParams{

		timeout: timeout,
	}
}

// NewIsFeatureEnabledParamsWithContext creates a new IsFeatureEnabledParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsFeatureEnabledParamsWithContext(ctx context.Context) *IsFeatureEnabledParams {
	var ()
	return &IsFeatureEnabledParams{

		Context: ctx,
	}
}

// NewIsFeatureEnabledParamsWithHTTPClient creates a new IsFeatureEnabledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsFeatureEnabledParamsWithHTTPClient(client *http.Client) *IsFeatureEnabledParams {
	var ()
	return &IsFeatureEnabledParams{
		HTTPClient: client,
	}
}

/*IsFeatureEnabledParams contains all the parameters to send to the API endpoint
for the is feature enabled operation typically these are written to a http.Request
*/
type IsFeatureEnabledParams struct {

	/*Body*/
	Body *models.IsFeatureEnabledRequestPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is feature enabled params
func (o *IsFeatureEnabledParams) WithTimeout(timeout time.Duration) *IsFeatureEnabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is feature enabled params
func (o *IsFeatureEnabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is feature enabled params
func (o *IsFeatureEnabledParams) WithContext(ctx context.Context) *IsFeatureEnabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is feature enabled params
func (o *IsFeatureEnabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is feature enabled params
func (o *IsFeatureEnabledParams) WithHTTPClient(client *http.Client) *IsFeatureEnabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is feature enabled params
func (o *IsFeatureEnabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the is feature enabled params
func (o *IsFeatureEnabledParams) WithBody(body *models.IsFeatureEnabledRequestPayload) *IsFeatureEnabledParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the is feature enabled params
func (o *IsFeatureEnabledParams) SetBody(body *models.IsFeatureEnabledRequestPayload) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IsFeatureEnabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}