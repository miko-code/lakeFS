// Code generated by go-swagger; DO NOT EDIT.

package branches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	"github.com/treeverse/lakefs/api/gen/models"
)

// DiffBranchesHandlerFunc turns a function with the right signature into a diff branches handler
type DiffBranchesHandlerFunc func(DiffBranchesParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DiffBranchesHandlerFunc) Handle(params DiffBranchesParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DiffBranchesHandler interface for that can handle valid diff branches params
type DiffBranchesHandler interface {
	Handle(DiffBranchesParams, *models.User) middleware.Responder
}

// NewDiffBranches creates a new http.Handler for the diff branches operation
func NewDiffBranches(ctx *middleware.Context, handler DiffBranchesHandler) *DiffBranches {
	return &DiffBranches{Context: ctx, Handler: handler}
}

/*DiffBranches swagger:route GET /repositories/{repositoryId}/branches/{branchId}/diff/{otherBranchId} branches diffBranches

diff branches

*/
type DiffBranches struct {
	Context *middleware.Context
	Handler DiffBranchesHandler
}

func (o *DiffBranches) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDiffBranchesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DiffBranchesOKBody diff branches o k body
// swagger:model DiffBranchesOKBody
type DiffBranchesOKBody struct {

	// results
	Results []*models.Diff `json:"results"`
}

// Validate validates this diff branches o k body
func (o *DiffBranchesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DiffBranchesOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("diffBranchesOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DiffBranchesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DiffBranchesOKBody) UnmarshalBinary(b []byte) error {
	var res DiffBranchesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
