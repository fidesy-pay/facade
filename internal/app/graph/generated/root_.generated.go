// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/fidesy-pay/facade/internal/pkg/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Client() ClientResolver
	Invoice() InvoiceResolver
	InvoiceMutations() InvoiceMutationsResolver
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Balance struct {
		Balance    func(childComplexity int) int
		Token      func(childComplexity int) int
		UsdBalance func(childComplexity int) int
	}

	CheckInvoicePayload struct {
		Invoice func(childComplexity int) int
	}

	Client struct {
		ApiKey    func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		Id        func(childComplexity int) int
		Invoices  func(childComplexity int) int
		PhotoUrl  func(childComplexity int) int
		Username  func(childComplexity int) int
		Wallets   func(childComplexity int) int
	}

	CreateInvoicePayload struct {
		ID func(childComplexity int) int
	}

	Invoice struct {
		Address     func(childComplexity int) int
		Chain       func(childComplexity int) int
		CreatedAt   func(childComplexity int) int
		Id          func(childComplexity int) int
		Payer       func(childComplexity int) int
		Status      func(childComplexity int) int
		Token       func(childComplexity int) int
		TokenAmount func(childComplexity int) int
		UsdAmount   func(childComplexity int) int
	}

	InvoiceMutations struct {
		CheckInvoice  func(childComplexity int, input model.CheckInvoiceInput) int
		CreateInvoice func(childComplexity int, input model.CreateInvoiceInput) int
		UpdateInvoice func(childComplexity int, input model.UpdateInvoiceInput) int
	}

	InvoicesPagination struct {
		Items func(childComplexity int) int
	}

	LoginPayload struct {
		Token func(childComplexity int) int
	}

	MainBalance struct {
		UsdBalance func(childComplexity int) int
	}

	Mutation struct {
		InvoiceMutations func(childComplexity int) int
		Login            func(childComplexity int, input model.LoginInput) int
		SignUp           func(childComplexity int, input model.SignUpInput) int
	}

	Query struct {
		Balance     func(childComplexity int, filter model.BalanceFilter) int
		Invoices    func(childComplexity int, filter *model.InvoicesFilter) int
		MainBalance func(childComplexity int) int
		Me          func(childComplexity int) int
		Wallets     func(childComplexity int, filter *model.WalletsFilter) int
	}

	SignUpPayload struct {
		Token func(childComplexity int) int
	}

	UpdateInvoicePayload struct {
		Invoice func(childComplexity int) int
	}

	Wallet struct {
		Address func(childComplexity int) int
		Chain   func(childComplexity int) int
	}

	WalletsPagination struct {
		Items func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Balance.balance":
		if e.complexity.Balance.Balance == nil {
			break
		}

		return e.complexity.Balance.Balance(childComplexity), true

	case "Balance.token":
		if e.complexity.Balance.Token == nil {
			break
		}

		return e.complexity.Balance.Token(childComplexity), true

	case "Balance.usdBalance":
		if e.complexity.Balance.UsdBalance == nil {
			break
		}

		return e.complexity.Balance.UsdBalance(childComplexity), true

	case "CheckInvoicePayload.invoice":
		if e.complexity.CheckInvoicePayload.Invoice == nil {
			break
		}

		return e.complexity.CheckInvoicePayload.Invoice(childComplexity), true

	case "Client.api_key":
		if e.complexity.Client.ApiKey == nil {
			break
		}

		return e.complexity.Client.ApiKey(childComplexity), true

	case "Client.created_at":
		if e.complexity.Client.CreatedAt == nil {
			break
		}

		return e.complexity.Client.CreatedAt(childComplexity), true

	case "Client.id":
		if e.complexity.Client.Id == nil {
			break
		}

		return e.complexity.Client.Id(childComplexity), true

	case "Client.invoices":
		if e.complexity.Client.Invoices == nil {
			break
		}

		return e.complexity.Client.Invoices(childComplexity), true

	case "Client.photo_url":
		if e.complexity.Client.PhotoUrl == nil {
			break
		}

		return e.complexity.Client.PhotoUrl(childComplexity), true

	case "Client.username":
		if e.complexity.Client.Username == nil {
			break
		}

		return e.complexity.Client.Username(childComplexity), true

	case "Client.wallets":
		if e.complexity.Client.Wallets == nil {
			break
		}

		return e.complexity.Client.Wallets(childComplexity), true

	case "CreateInvoicePayload.id":
		if e.complexity.CreateInvoicePayload.ID == nil {
			break
		}

		return e.complexity.CreateInvoicePayload.ID(childComplexity), true

	case "Invoice.address":
		if e.complexity.Invoice.Address == nil {
			break
		}

		return e.complexity.Invoice.Address(childComplexity), true

	case "Invoice.chain":
		if e.complexity.Invoice.Chain == nil {
			break
		}

		return e.complexity.Invoice.Chain(childComplexity), true

	case "Invoice.created_at":
		if e.complexity.Invoice.CreatedAt == nil {
			break
		}

		return e.complexity.Invoice.CreatedAt(childComplexity), true

	case "Invoice.id":
		if e.complexity.Invoice.Id == nil {
			break
		}

		return e.complexity.Invoice.Id(childComplexity), true

	case "Invoice.payer":
		if e.complexity.Invoice.Payer == nil {
			break
		}

		return e.complexity.Invoice.Payer(childComplexity), true

	case "Invoice.status":
		if e.complexity.Invoice.Status == nil {
			break
		}

		return e.complexity.Invoice.Status(childComplexity), true

	case "Invoice.token":
		if e.complexity.Invoice.Token == nil {
			break
		}

		return e.complexity.Invoice.Token(childComplexity), true

	case "Invoice.token_amount":
		if e.complexity.Invoice.TokenAmount == nil {
			break
		}

		return e.complexity.Invoice.TokenAmount(childComplexity), true

	case "Invoice.usd_amount":
		if e.complexity.Invoice.UsdAmount == nil {
			break
		}

		return e.complexity.Invoice.UsdAmount(childComplexity), true

	case "InvoiceMutations.checkInvoice":
		if e.complexity.InvoiceMutations.CheckInvoice == nil {
			break
		}

		args, err := ec.field_InvoiceMutations_checkInvoice_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.InvoiceMutations.CheckInvoice(childComplexity, args["input"].(model.CheckInvoiceInput)), true

	case "InvoiceMutations.createInvoice":
		if e.complexity.InvoiceMutations.CreateInvoice == nil {
			break
		}

		args, err := ec.field_InvoiceMutations_createInvoice_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.InvoiceMutations.CreateInvoice(childComplexity, args["input"].(model.CreateInvoiceInput)), true

	case "InvoiceMutations.updateInvoice":
		if e.complexity.InvoiceMutations.UpdateInvoice == nil {
			break
		}

		args, err := ec.field_InvoiceMutations_updateInvoice_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.InvoiceMutations.UpdateInvoice(childComplexity, args["input"].(model.UpdateInvoiceInput)), true

	case "InvoicesPagination.items":
		if e.complexity.InvoicesPagination.Items == nil {
			break
		}

		return e.complexity.InvoicesPagination.Items(childComplexity), true

	case "LoginPayload.token":
		if e.complexity.LoginPayload.Token == nil {
			break
		}

		return e.complexity.LoginPayload.Token(childComplexity), true

	case "MainBalance.usdBalance":
		if e.complexity.MainBalance.UsdBalance == nil {
			break
		}

		return e.complexity.MainBalance.UsdBalance(childComplexity), true

	case "Mutation.invoiceMutations":
		if e.complexity.Mutation.InvoiceMutations == nil {
			break
		}

		return e.complexity.Mutation.InvoiceMutations(childComplexity), true

	case "Mutation.login":
		if e.complexity.Mutation.Login == nil {
			break
		}

		args, err := ec.field_Mutation_login_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.Login(childComplexity, args["input"].(model.LoginInput)), true

	case "Mutation.signUp":
		if e.complexity.Mutation.SignUp == nil {
			break
		}

		args, err := ec.field_Mutation_signUp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SignUp(childComplexity, args["input"].(model.SignUpInput)), true

	case "Query.balance":
		if e.complexity.Query.Balance == nil {
			break
		}

		args, err := ec.field_Query_balance_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Balance(childComplexity, args["filter"].(model.BalanceFilter)), true

	case "Query.invoices":
		if e.complexity.Query.Invoices == nil {
			break
		}

		args, err := ec.field_Query_invoices_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Invoices(childComplexity, args["filter"].(*model.InvoicesFilter)), true

	case "Query.mainBalance":
		if e.complexity.Query.MainBalance == nil {
			break
		}

		return e.complexity.Query.MainBalance(childComplexity), true

	case "Query.me":
		if e.complexity.Query.Me == nil {
			break
		}

		return e.complexity.Query.Me(childComplexity), true

	case "Query.wallets":
		if e.complexity.Query.Wallets == nil {
			break
		}

		args, err := ec.field_Query_wallets_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Wallets(childComplexity, args["filter"].(*model.WalletsFilter)), true

	case "SignUpPayload.token":
		if e.complexity.SignUpPayload.Token == nil {
			break
		}

		return e.complexity.SignUpPayload.Token(childComplexity), true

	case "UpdateInvoicePayload.invoice":
		if e.complexity.UpdateInvoicePayload.Invoice == nil {
			break
		}

		return e.complexity.UpdateInvoicePayload.Invoice(childComplexity), true

	case "Wallet.address":
		if e.complexity.Wallet.Address == nil {
			break
		}

		return e.complexity.Wallet.Address(childComplexity), true

	case "Wallet.chain":
		if e.complexity.Wallet.Chain == nil {
			break
		}

		return e.complexity.Wallet.Chain(childComplexity), true

	case "WalletsPagination.items":
		if e.complexity.WalletsPagination.Items == nil {
			break
		}

		return e.complexity.WalletsPagination.Items(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputBalanceFilter,
		ec.unmarshalInputCheckInvoiceInput,
		ec.unmarshalInputCreateInvoiceInput,
		ec.unmarshalInputInvoicesFilter,
		ec.unmarshalInputLoginInput,
		ec.unmarshalInputSignUpInput,
		ec.unmarshalInputUpdateInvoiceInput,
		ec.unmarshalInputWalletsFilter,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../../../api/graphql/directives/goModel.graphql", Input: `directive @goModel(model: String, models: [String!]) on OBJECT
    | INPUT_OBJECT
    | SCALAR
    | ENUM
    | INTERFACE
    | UNION

directive @goField(forceResolver: Boolean, name: String, omittable: Boolean) on INPUT_FIELD_DEFINITION
    | FIELD_DEFINITION`, BuiltIn: false},
	{Name: "../../../../api/graphql/enum/invoice_status.graphql", Input: `enum InvoiceStatus {
  UNKNOWN_STATUS
  NEW
  PENDING
  FAILED
  SUCCESS
  EXPIRED
  SENDING_TO_CLIENT
  MANUAL_CONTROL
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/mutation/invoice.mutation.graphql", Input: `extend type Mutation {
    invoiceMutations: InvoiceMutations!
}

type InvoiceMutations {
    createInvoice(input: CreateInvoiceInput!): CreateInvoicePayload!
    updateInvoice(input: UpdateInvoiceInput!): UpdateInvoicePayload!
    checkInvoice(input: CheckInvoiceInput!): CheckInvoicePayload!
}

input CreateInvoiceInput {
    usd_amount: Float!
}
type CreateInvoicePayload {
    id: String!
}

input UpdateInvoiceInput {
    id: String!
    chain: String!
    token: String!
}

type UpdateInvoicePayload {
    invoice: Invoice!
}

input CheckInvoiceInput {
    id: String!
}

type CheckInvoicePayload {
    invoice: Invoice!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/mutation/login.graphql", Input: `extend type Mutation {
    login(input: LoginInput!): LoginPayload!
}

input LoginInput {
    username: String!
    password: String!
}

type LoginPayload {
    token: String!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/mutation/signup.graphql", Input: `extend type Mutation {
    signUp(input: SignUpInput!): SignUpPayload!
}

input SignUpInput {
    username: String!
    password: String!
}

type SignUpPayload {
    token: String!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/query/balance.query.graphql", Input: `extend type Query {
    balance(filter: BalanceFilter!): Balance!
}

input BalanceFilter {
    addressEq: String!
    chainEq: String!
    tokenEq: String!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/query/invoices.query.graphql", Input: `extend type Query {
    invoices(filter: InvoicesFilter): InvoicesPagination!
}

input InvoicesFilter {
    idIn: [String!]
    clientIdIn: [String!]
    statusIn: [InvoiceStatus!]
}

type InvoicesPagination {
    items: [Invoice!]
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/query/main_balance.query.graphql", Input: `extend type Query {
    mainBalance: MainBalance!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/query/me.query.graphql", Input: `extend type Query {
    me: Client!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/query/wallets.query.graphql", Input: `extend type Query {
    wallets(filter: WalletsFilter): WalletsPagination!
}

input WalletsFilter {
    clientIdIn: [String!]
}

type WalletsPagination {
    items: [Wallet!]
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/scalars.graphql", Input: `scalar Time`, BuiltIn: false},
	{Name: "../../../../api/graphql/schema/schema.graphql", Input: `schema {
    query: Query
    mutation: Mutation
}

type Query

type Mutation`, BuiltIn: false},
	{Name: "../../../../api/graphql/types/balance.graphql", Input: `type Balance {
    token: String!
    balance: Float!
    usdBalance: Float!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/types/client.graphql", Input: `type Client @goModel(model: "github.com/fidesy-pay/facade/pkg/clients-service.Client") {
    id: String!
    username: String!
    api_key: String!
    created_at: Time!
    photo_url: String

    wallets: [Wallet!]
    invoices: [Invoice]!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/types/invoice.graphql", Input: `type Invoice @goModel(model: "github.com/fidesy-pay/facade/pkg/invoices-service.Invoice") {
    id: String!
    usd_amount: Float!
    token_amount: Float!
    chain: String!
    token: String!
    status: InvoiceStatus!
    address: String!
    created_at: Time!
    payer: Client
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/types/main_balance.graphql", Input: `type MainBalance {
    usdBalance: Float!
}`, BuiltIn: false},
	{Name: "../../../../api/graphql/types/wallet.graphql", Input: `type Wallet @goModel(model: "github.com/fidesy-pay/facade/pkg/crypto-service.Wallet") {
    address: String!
    chain: String!
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
