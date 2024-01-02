// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	invoices_service "github.com/fidesy-pay/facade/pkg/invoices-service"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type InvoiceMutationsResolver interface {
	CreateInvoice(ctx context.Context, obj *model.InvoiceMutations, input model.CreateInvoiceInput) (*model.CreateInvoicePayload, error)
	UpdateInvoice(ctx context.Context, obj *model.InvoiceMutations, input model.UpdateInvoiceInput) (*model.UpdateInvoicePayload, error)
	CheckInvoice(ctx context.Context, obj *model.InvoiceMutations, input model.CheckInvoiceInput) (*model.CheckInvoicePayload, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_InvoiceMutations_checkInvoice_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.CheckInvoiceInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNCheckInvoiceInput2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCheckInvoiceInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_InvoiceMutations_createInvoice_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.CreateInvoiceInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNCreateInvoiceInput2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCreateInvoiceInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_InvoiceMutations_updateInvoice_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.UpdateInvoiceInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNUpdateInvoiceInput2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐUpdateInvoiceInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _CheckInvoicePayload_invoice(ctx context.Context, field graphql.CollectedField, obj *model.CheckInvoicePayload) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CheckInvoicePayload_invoice(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Invoice, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*invoices_service.Invoice)
	fc.Result = res
	return ec.marshalNInvoice2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋpkgᚋinvoicesᚑserviceᚐInvoice(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CheckInvoicePayload_invoice(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CheckInvoicePayload",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Invoice_id(ctx, field)
			case "usd_amount":
				return ec.fieldContext_Invoice_usd_amount(ctx, field)
			case "token_amount":
				return ec.fieldContext_Invoice_token_amount(ctx, field)
			case "chain":
				return ec.fieldContext_Invoice_chain(ctx, field)
			case "token":
				return ec.fieldContext_Invoice_token(ctx, field)
			case "status":
				return ec.fieldContext_Invoice_status(ctx, field)
			case "address":
				return ec.fieldContext_Invoice_address(ctx, field)
			case "created_at":
				return ec.fieldContext_Invoice_created_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Invoice", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _CreateInvoicePayload_id(ctx context.Context, field graphql.CollectedField, obj *model.CreateInvoicePayload) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CreateInvoicePayload_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CreateInvoicePayload_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CreateInvoicePayload",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _InvoiceMutations_createInvoice(ctx context.Context, field graphql.CollectedField, obj *model.InvoiceMutations) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_InvoiceMutations_createInvoice(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.InvoiceMutations().CreateInvoice(rctx, obj, fc.Args["input"].(model.CreateInvoiceInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.CreateInvoicePayload)
	fc.Result = res
	return ec.marshalNCreateInvoicePayload2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCreateInvoicePayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_InvoiceMutations_createInvoice(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "InvoiceMutations",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_CreateInvoicePayload_id(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type CreateInvoicePayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_InvoiceMutations_createInvoice_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _InvoiceMutations_updateInvoice(ctx context.Context, field graphql.CollectedField, obj *model.InvoiceMutations) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_InvoiceMutations_updateInvoice(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.InvoiceMutations().UpdateInvoice(rctx, obj, fc.Args["input"].(model.UpdateInvoiceInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.UpdateInvoicePayload)
	fc.Result = res
	return ec.marshalNUpdateInvoicePayload2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐUpdateInvoicePayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_InvoiceMutations_updateInvoice(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "InvoiceMutations",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "invoice":
				return ec.fieldContext_UpdateInvoicePayload_invoice(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type UpdateInvoicePayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_InvoiceMutations_updateInvoice_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _InvoiceMutations_checkInvoice(ctx context.Context, field graphql.CollectedField, obj *model.InvoiceMutations) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_InvoiceMutations_checkInvoice(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.InvoiceMutations().CheckInvoice(rctx, obj, fc.Args["input"].(model.CheckInvoiceInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.CheckInvoicePayload)
	fc.Result = res
	return ec.marshalNCheckInvoicePayload2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCheckInvoicePayload(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_InvoiceMutations_checkInvoice(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "InvoiceMutations",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "invoice":
				return ec.fieldContext_CheckInvoicePayload_invoice(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type CheckInvoicePayload", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_InvoiceMutations_checkInvoice_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _UpdateInvoicePayload_invoice(ctx context.Context, field graphql.CollectedField, obj *model.UpdateInvoicePayload) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UpdateInvoicePayload_invoice(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Invoice, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*invoices_service.Invoice)
	fc.Result = res
	return ec.marshalNInvoice2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋpkgᚋinvoicesᚑserviceᚐInvoice(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UpdateInvoicePayload_invoice(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UpdateInvoicePayload",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Invoice_id(ctx, field)
			case "usd_amount":
				return ec.fieldContext_Invoice_usd_amount(ctx, field)
			case "token_amount":
				return ec.fieldContext_Invoice_token_amount(ctx, field)
			case "chain":
				return ec.fieldContext_Invoice_chain(ctx, field)
			case "token":
				return ec.fieldContext_Invoice_token(ctx, field)
			case "status":
				return ec.fieldContext_Invoice_status(ctx, field)
			case "address":
				return ec.fieldContext_Invoice_address(ctx, field)
			case "created_at":
				return ec.fieldContext_Invoice_created_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Invoice", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCheckInvoiceInput(ctx context.Context, obj interface{}) (model.CheckInvoiceInput, error) {
	var it model.CheckInvoiceInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputCreateInvoiceInput(ctx context.Context, obj interface{}) (model.CreateInvoiceInput, error) {
	var it model.CreateInvoiceInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"client_id", "usd_amount"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "client_id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("client_id"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ClientID = data
		case "usd_amount":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("usd_amount"))
			data, err := ec.unmarshalNFloat2float64(ctx, v)
			if err != nil {
				return it, err
			}
			it.UsdAmount = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpdateInvoiceInput(ctx context.Context, obj interface{}) (model.UpdateInvoiceInput, error) {
	var it model.UpdateInvoiceInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "chain", "token"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "chain":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("chain"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Chain = data
		case "token":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("token"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Token = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var checkInvoicePayloadImplementors = []string{"CheckInvoicePayload"}

func (ec *executionContext) _CheckInvoicePayload(ctx context.Context, sel ast.SelectionSet, obj *model.CheckInvoicePayload) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, checkInvoicePayloadImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CheckInvoicePayload")
		case "invoice":
			out.Values[i] = ec._CheckInvoicePayload_invoice(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var createInvoicePayloadImplementors = []string{"CreateInvoicePayload"}

func (ec *executionContext) _CreateInvoicePayload(ctx context.Context, sel ast.SelectionSet, obj *model.CreateInvoicePayload) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, createInvoicePayloadImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CreateInvoicePayload")
		case "id":
			out.Values[i] = ec._CreateInvoicePayload_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var invoiceMutationsImplementors = []string{"InvoiceMutations"}

func (ec *executionContext) _InvoiceMutations(ctx context.Context, sel ast.SelectionSet, obj *model.InvoiceMutations) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, invoiceMutationsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("InvoiceMutations")
		case "createInvoice":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._InvoiceMutations_createInvoice(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "updateInvoice":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._InvoiceMutations_updateInvoice(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "checkInvoice":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._InvoiceMutations_checkInvoice(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var updateInvoicePayloadImplementors = []string{"UpdateInvoicePayload"}

func (ec *executionContext) _UpdateInvoicePayload(ctx context.Context, sel ast.SelectionSet, obj *model.UpdateInvoicePayload) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, updateInvoicePayloadImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UpdateInvoicePayload")
		case "invoice":
			out.Values[i] = ec._UpdateInvoicePayload_invoice(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNCheckInvoiceInput2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCheckInvoiceInput(ctx context.Context, v interface{}) (model.CheckInvoiceInput, error) {
	res, err := ec.unmarshalInputCheckInvoiceInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNCheckInvoicePayload2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCheckInvoicePayload(ctx context.Context, sel ast.SelectionSet, v model.CheckInvoicePayload) graphql.Marshaler {
	return ec._CheckInvoicePayload(ctx, sel, &v)
}

func (ec *executionContext) marshalNCheckInvoicePayload2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCheckInvoicePayload(ctx context.Context, sel ast.SelectionSet, v *model.CheckInvoicePayload) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CheckInvoicePayload(ctx, sel, v)
}

func (ec *executionContext) unmarshalNCreateInvoiceInput2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCreateInvoiceInput(ctx context.Context, v interface{}) (model.CreateInvoiceInput, error) {
	res, err := ec.unmarshalInputCreateInvoiceInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNCreateInvoicePayload2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCreateInvoicePayload(ctx context.Context, sel ast.SelectionSet, v model.CreateInvoicePayload) graphql.Marshaler {
	return ec._CreateInvoicePayload(ctx, sel, &v)
}

func (ec *executionContext) marshalNCreateInvoicePayload2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐCreateInvoicePayload(ctx context.Context, sel ast.SelectionSet, v *model.CreateInvoicePayload) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CreateInvoicePayload(ctx, sel, v)
}

func (ec *executionContext) marshalNInvoiceMutations2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceMutations(ctx context.Context, sel ast.SelectionSet, v model.InvoiceMutations) graphql.Marshaler {
	return ec._InvoiceMutations(ctx, sel, &v)
}

func (ec *executionContext) marshalNInvoiceMutations2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceMutations(ctx context.Context, sel ast.SelectionSet, v *model.InvoiceMutations) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._InvoiceMutations(ctx, sel, v)
}

func (ec *executionContext) unmarshalNUpdateInvoiceInput2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐUpdateInvoiceInput(ctx context.Context, v interface{}) (model.UpdateInvoiceInput, error) {
	res, err := ec.unmarshalInputUpdateInvoiceInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNUpdateInvoicePayload2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐUpdateInvoicePayload(ctx context.Context, sel ast.SelectionSet, v model.UpdateInvoicePayload) graphql.Marshaler {
	return ec._UpdateInvoicePayload(ctx, sel, &v)
}

func (ec *executionContext) marshalNUpdateInvoicePayload2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐUpdateInvoicePayload(ctx context.Context, sel ast.SelectionSet, v *model.UpdateInvoicePayload) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._UpdateInvoicePayload(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
