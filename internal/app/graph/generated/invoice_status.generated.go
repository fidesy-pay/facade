// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNInvoiceStatus2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceStatus(ctx context.Context, v interface{}) (model.InvoiceStatus, error) {
	var res model.InvoiceStatus
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNInvoiceStatus2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceStatus(ctx context.Context, sel ast.SelectionSet, v model.InvoiceStatus) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalOInvoiceStatus2ᚕgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceStatusᚄ(ctx context.Context, v interface{}) ([]model.InvoiceStatus, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]model.InvoiceStatus, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNInvoiceStatus2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceStatus(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOInvoiceStatus2ᚕgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceStatusᚄ(ctx context.Context, sel ast.SelectionSet, v []model.InvoiceStatus) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNInvoiceStatus2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐInvoiceStatus(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

// endregion ***************************** type.gotpl *****************************
