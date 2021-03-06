package j_account

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

	"koding/remoteapi/models"
)

// NewJAccountFetchMySessionsParams creates a new JAccountFetchMySessionsParams object
// with the default values initialized.
func NewJAccountFetchMySessionsParams() *JAccountFetchMySessionsParams {
	var ()
	return &JAccountFetchMySessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountFetchMySessionsParamsWithTimeout creates a new JAccountFetchMySessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountFetchMySessionsParamsWithTimeout(timeout time.Duration) *JAccountFetchMySessionsParams {
	var ()
	return &JAccountFetchMySessionsParams{

		timeout: timeout,
	}
}

// NewJAccountFetchMySessionsParamsWithContext creates a new JAccountFetchMySessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountFetchMySessionsParamsWithContext(ctx context.Context) *JAccountFetchMySessionsParams {
	var ()
	return &JAccountFetchMySessionsParams{

		Context: ctx,
	}
}

/*JAccountFetchMySessionsParams contains all the parameters to send to the API endpoint
for the j account fetch my sessions operation typically these are written to a http.Request
*/
type JAccountFetchMySessionsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) WithTimeout(timeout time.Duration) *JAccountFetchMySessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) WithContext(ctx context.Context) *JAccountFetchMySessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) WithBody(body models.DefaultSelector) *JAccountFetchMySessionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) WithID(id string) *JAccountFetchMySessionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j account fetch my sessions params
func (o *JAccountFetchMySessionsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountFetchMySessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
