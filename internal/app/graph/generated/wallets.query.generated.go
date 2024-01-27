// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	crypto_service "github.com/fidesy-pay/facade/pkg/crypto-service"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _WalletsPagination_items(ctx context.Context, field graphql.CollectedField, obj *model.WalletsPagination) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_WalletsPagination_items(ctx, field)
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
		return obj.Items, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*crypto_service.Wallet)
	fc.Result = res
	return ec.marshalOWallet2ᚕᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋpkgᚋcryptoᚑserviceᚐWalletᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_WalletsPagination_items(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "WalletsPagination",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "address":
				return ec.fieldContext_Wallet_address(ctx, field)
			case "balance":
				return ec.fieldContext_Wallet_balance(ctx, field)
			case "chain":
				return ec.fieldContext_Wallet_chain(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Wallet", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputWalletsFilter(ctx context.Context, obj interface{}) (model.WalletsFilter, error) {
	var it model.WalletsFilter
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"clientIdIn"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "clientIdIn":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("clientIdIn"))
			data, err := ec.unmarshalOString2ᚕstringᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.ClientIDIn = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var walletsPaginationImplementors = []string{"WalletsPagination"}

func (ec *executionContext) _WalletsPagination(ctx context.Context, sel ast.SelectionSet, obj *model.WalletsPagination) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, walletsPaginationImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("WalletsPagination")
		case "items":
			out.Values[i] = ec._WalletsPagination_items(ctx, field, obj)
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

func (ec *executionContext) unmarshalNWalletsFilter2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐWalletsFilter(ctx context.Context, v interface{}) (model.WalletsFilter, error) {
	res, err := ec.unmarshalInputWalletsFilter(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNWalletsPagination2githubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐWalletsPagination(ctx context.Context, sel ast.SelectionSet, v model.WalletsPagination) graphql.Marshaler {
	return ec._WalletsPagination(ctx, sel, &v)
}

func (ec *executionContext) marshalNWalletsPagination2ᚖgithubᚗcomᚋfidesyᚑpayᚋfacadeᚋinternalᚋpkgᚋmodelᚐWalletsPagination(ctx context.Context, sel ast.SelectionSet, v *model.WalletsPagination) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._WalletsPagination(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
