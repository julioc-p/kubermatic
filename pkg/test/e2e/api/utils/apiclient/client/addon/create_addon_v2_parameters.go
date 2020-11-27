// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/api/utils/apiclient/models"
)

// NewCreateAddonV2Params creates a new CreateAddonV2Params object
// with the default values initialized.
func NewCreateAddonV2Params() *CreateAddonV2Params {
	var ()
	return &CreateAddonV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAddonV2ParamsWithTimeout creates a new CreateAddonV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAddonV2ParamsWithTimeout(timeout time.Duration) *CreateAddonV2Params {
	var ()
	return &CreateAddonV2Params{

		timeout: timeout,
	}
}

// NewCreateAddonV2ParamsWithContext creates a new CreateAddonV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAddonV2ParamsWithContext(ctx context.Context) *CreateAddonV2Params {
	var ()
	return &CreateAddonV2Params{

		Context: ctx,
	}
}

// NewCreateAddonV2ParamsWithHTTPClient creates a new CreateAddonV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAddonV2ParamsWithHTTPClient(client *http.Client) *CreateAddonV2Params {
	var ()
	return &CreateAddonV2Params{
		HTTPClient: client,
	}
}

/*CreateAddonV2Params contains all the parameters to send to the API endpoint
for the create addon v2 operation typically these are written to a http.Request
*/
type CreateAddonV2Params struct {

	/*Body*/
	Body *models.Addon
	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create addon v2 params
func (o *CreateAddonV2Params) WithTimeout(timeout time.Duration) *CreateAddonV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create addon v2 params
func (o *CreateAddonV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create addon v2 params
func (o *CreateAddonV2Params) WithContext(ctx context.Context) *CreateAddonV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create addon v2 params
func (o *CreateAddonV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create addon v2 params
func (o *CreateAddonV2Params) WithHTTPClient(client *http.Client) *CreateAddonV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create addon v2 params
func (o *CreateAddonV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create addon v2 params
func (o *CreateAddonV2Params) WithBody(body *models.Addon) *CreateAddonV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create addon v2 params
func (o *CreateAddonV2Params) SetBody(body *models.Addon) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create addon v2 params
func (o *CreateAddonV2Params) WithClusterID(clusterID string) *CreateAddonV2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create addon v2 params
func (o *CreateAddonV2Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create addon v2 params
func (o *CreateAddonV2Params) WithProjectID(projectID string) *CreateAddonV2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create addon v2 params
func (o *CreateAddonV2Params) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAddonV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
