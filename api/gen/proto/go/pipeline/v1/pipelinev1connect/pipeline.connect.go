// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pipeline/v1/pipeline.proto

package pipelinev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/grafana/fleet-management-api/api/gen/proto/go/pipeline/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PipelineServiceName is the fully-qualified name of the PipelineService service.
	PipelineServiceName = "pipeline.v1.PipelineService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PipelineServiceGetPipelineProcedure is the fully-qualified name of the PipelineService's
	// GetPipeline RPC.
	PipelineServiceGetPipelineProcedure = "/pipeline.v1.PipelineService/GetPipeline"
	// PipelineServiceGetPipelineIDProcedure is the fully-qualified name of the PipelineService's
	// GetPipelineID RPC.
	PipelineServiceGetPipelineIDProcedure = "/pipeline.v1.PipelineService/GetPipelineID"
	// PipelineServiceListPipelinesProcedure is the fully-qualified name of the PipelineService's
	// ListPipelines RPC.
	PipelineServiceListPipelinesProcedure = "/pipeline.v1.PipelineService/ListPipelines"
	// PipelineServiceCreateAlloyPipelinesProcedure is the fully-qualified name of the PipelineService's
	// CreateAlloyPipelines RPC.
	PipelineServiceCreateAlloyPipelinesProcedure = "/pipeline.v1.PipelineService/CreateAlloyPipelines"
	// PipelineServiceCreatePipelineProcedure is the fully-qualified name of the PipelineService's
	// CreatePipeline RPC.
	PipelineServiceCreatePipelineProcedure = "/pipeline.v1.PipelineService/CreatePipeline"
	// PipelineServiceUpdatePipelineProcedure is the fully-qualified name of the PipelineService's
	// UpdatePipeline RPC.
	PipelineServiceUpdatePipelineProcedure = "/pipeline.v1.PipelineService/UpdatePipeline"
	// PipelineServiceUpsertPipelineProcedure is the fully-qualified name of the PipelineService's
	// UpsertPipeline RPC.
	PipelineServiceUpsertPipelineProcedure = "/pipeline.v1.PipelineService/UpsertPipeline"
	// PipelineServiceDeletePipelineProcedure is the fully-qualified name of the PipelineService's
	// DeletePipeline RPC.
	PipelineServiceDeletePipelineProcedure = "/pipeline.v1.PipelineService/DeletePipeline"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	pipelineServiceServiceDescriptor                    = v1.File_pipeline_v1_pipeline_proto.Services().ByName("PipelineService")
	pipelineServiceGetPipelineMethodDescriptor          = pipelineServiceServiceDescriptor.Methods().ByName("GetPipeline")
	pipelineServiceGetPipelineIDMethodDescriptor        = pipelineServiceServiceDescriptor.Methods().ByName("GetPipelineID")
	pipelineServiceListPipelinesMethodDescriptor        = pipelineServiceServiceDescriptor.Methods().ByName("ListPipelines")
	pipelineServiceCreateAlloyPipelinesMethodDescriptor = pipelineServiceServiceDescriptor.Methods().ByName("CreateAlloyPipelines")
	pipelineServiceCreatePipelineMethodDescriptor       = pipelineServiceServiceDescriptor.Methods().ByName("CreatePipeline")
	pipelineServiceUpdatePipelineMethodDescriptor       = pipelineServiceServiceDescriptor.Methods().ByName("UpdatePipeline")
	pipelineServiceUpsertPipelineMethodDescriptor       = pipelineServiceServiceDescriptor.Methods().ByName("UpsertPipeline")
	pipelineServiceDeletePipelineMethodDescriptor       = pipelineServiceServiceDescriptor.Methods().ByName("DeletePipeline")
)

// PipelineServiceClient is a client for the pipeline.v1.PipelineService service.
type PipelineServiceClient interface {
	// GetPipeline returns a pipeline by name.
	GetPipeline(context.Context, *connect.Request[v1.GetPipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// GetPipelineID returns a pipeline ID by name.
	GetPipelineID(context.Context, *connect.Request[v1.GetPipelineIDRequest]) (*connect.Response[v1.GetPipelineIDResponse], error)
	// ListPipelines returns all pipelines.
	ListPipelines(context.Context, *connect.Request[v1.ListPipelinesRequest]) (*connect.Response[v1.Pipelines], error)
	// CreateAlloyPipelines creates initial alloy pipelines.
	CreateAlloyPipelines(context.Context, *connect.Request[v1.CreateAlloyPipelinesRequest]) (*connect.Response[v1.CreateAlloyPipelinesResponse], error)
	// CreatePipeline creates a new pipeline and returns it.
	CreatePipeline(context.Context, *connect.Request[v1.CreatePipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// UpdatePipeline updates an existing pipeline and returns it.
	UpdatePipeline(context.Context, *connect.Request[v1.UpdatePipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// UpsertPipeline creates a new pipeline or updates an existing one and returns it.
	UpsertPipeline(context.Context, *connect.Request[v1.UpsertPipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// DeletePipeline deletes a pipeline by name.
	DeletePipeline(context.Context, *connect.Request[v1.DeletePipelineRequest]) (*connect.Response[v1.DeletePipelineResponse], error)
}

// NewPipelineServiceClient constructs a client for the pipeline.v1.PipelineService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPipelineServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PipelineServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pipelineServiceClient{
		getPipeline: connect.NewClient[v1.GetPipelineRequest, v1.Pipeline](
			httpClient,
			baseURL+PipelineServiceGetPipelineProcedure,
			connect.WithSchema(pipelineServiceGetPipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPipelineID: connect.NewClient[v1.GetPipelineIDRequest, v1.GetPipelineIDResponse](
			httpClient,
			baseURL+PipelineServiceGetPipelineIDProcedure,
			connect.WithSchema(pipelineServiceGetPipelineIDMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listPipelines: connect.NewClient[v1.ListPipelinesRequest, v1.Pipelines](
			httpClient,
			baseURL+PipelineServiceListPipelinesProcedure,
			connect.WithSchema(pipelineServiceListPipelinesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createAlloyPipelines: connect.NewClient[v1.CreateAlloyPipelinesRequest, v1.CreateAlloyPipelinesResponse](
			httpClient,
			baseURL+PipelineServiceCreateAlloyPipelinesProcedure,
			connect.WithSchema(pipelineServiceCreateAlloyPipelinesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createPipeline: connect.NewClient[v1.CreatePipelineRequest, v1.Pipeline](
			httpClient,
			baseURL+PipelineServiceCreatePipelineProcedure,
			connect.WithSchema(pipelineServiceCreatePipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updatePipeline: connect.NewClient[v1.UpdatePipelineRequest, v1.Pipeline](
			httpClient,
			baseURL+PipelineServiceUpdatePipelineProcedure,
			connect.WithSchema(pipelineServiceUpdatePipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		upsertPipeline: connect.NewClient[v1.UpsertPipelineRequest, v1.Pipeline](
			httpClient,
			baseURL+PipelineServiceUpsertPipelineProcedure,
			connect.WithSchema(pipelineServiceUpsertPipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deletePipeline: connect.NewClient[v1.DeletePipelineRequest, v1.DeletePipelineResponse](
			httpClient,
			baseURL+PipelineServiceDeletePipelineProcedure,
			connect.WithSchema(pipelineServiceDeletePipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// pipelineServiceClient implements PipelineServiceClient.
type pipelineServiceClient struct {
	getPipeline          *connect.Client[v1.GetPipelineRequest, v1.Pipeline]
	getPipelineID        *connect.Client[v1.GetPipelineIDRequest, v1.GetPipelineIDResponse]
	listPipelines        *connect.Client[v1.ListPipelinesRequest, v1.Pipelines]
	createAlloyPipelines *connect.Client[v1.CreateAlloyPipelinesRequest, v1.CreateAlloyPipelinesResponse]
	createPipeline       *connect.Client[v1.CreatePipelineRequest, v1.Pipeline]
	updatePipeline       *connect.Client[v1.UpdatePipelineRequest, v1.Pipeline]
	upsertPipeline       *connect.Client[v1.UpsertPipelineRequest, v1.Pipeline]
	deletePipeline       *connect.Client[v1.DeletePipelineRequest, v1.DeletePipelineResponse]
}

// GetPipeline calls pipeline.v1.PipelineService.GetPipeline.
func (c *pipelineServiceClient) GetPipeline(ctx context.Context, req *connect.Request[v1.GetPipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return c.getPipeline.CallUnary(ctx, req)
}

// GetPipelineID calls pipeline.v1.PipelineService.GetPipelineID.
func (c *pipelineServiceClient) GetPipelineID(ctx context.Context, req *connect.Request[v1.GetPipelineIDRequest]) (*connect.Response[v1.GetPipelineIDResponse], error) {
	return c.getPipelineID.CallUnary(ctx, req)
}

// ListPipelines calls pipeline.v1.PipelineService.ListPipelines.
func (c *pipelineServiceClient) ListPipelines(ctx context.Context, req *connect.Request[v1.ListPipelinesRequest]) (*connect.Response[v1.Pipelines], error) {
	return c.listPipelines.CallUnary(ctx, req)
}

// CreateAlloyPipelines calls pipeline.v1.PipelineService.CreateAlloyPipelines.
func (c *pipelineServiceClient) CreateAlloyPipelines(ctx context.Context, req *connect.Request[v1.CreateAlloyPipelinesRequest]) (*connect.Response[v1.CreateAlloyPipelinesResponse], error) {
	return c.createAlloyPipelines.CallUnary(ctx, req)
}

// CreatePipeline calls pipeline.v1.PipelineService.CreatePipeline.
func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, req *connect.Request[v1.CreatePipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return c.createPipeline.CallUnary(ctx, req)
}

// UpdatePipeline calls pipeline.v1.PipelineService.UpdatePipeline.
func (c *pipelineServiceClient) UpdatePipeline(ctx context.Context, req *connect.Request[v1.UpdatePipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return c.updatePipeline.CallUnary(ctx, req)
}

// UpsertPipeline calls pipeline.v1.PipelineService.UpsertPipeline.
func (c *pipelineServiceClient) UpsertPipeline(ctx context.Context, req *connect.Request[v1.UpsertPipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return c.upsertPipeline.CallUnary(ctx, req)
}

// DeletePipeline calls pipeline.v1.PipelineService.DeletePipeline.
func (c *pipelineServiceClient) DeletePipeline(ctx context.Context, req *connect.Request[v1.DeletePipelineRequest]) (*connect.Response[v1.DeletePipelineResponse], error) {
	return c.deletePipeline.CallUnary(ctx, req)
}

// PipelineServiceHandler is an implementation of the pipeline.v1.PipelineService service.
type PipelineServiceHandler interface {
	// GetPipeline returns a pipeline by name.
	GetPipeline(context.Context, *connect.Request[v1.GetPipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// GetPipelineID returns a pipeline ID by name.
	GetPipelineID(context.Context, *connect.Request[v1.GetPipelineIDRequest]) (*connect.Response[v1.GetPipelineIDResponse], error)
	// ListPipelines returns all pipelines.
	ListPipelines(context.Context, *connect.Request[v1.ListPipelinesRequest]) (*connect.Response[v1.Pipelines], error)
	// CreateAlloyPipelines creates initial alloy pipelines.
	CreateAlloyPipelines(context.Context, *connect.Request[v1.CreateAlloyPipelinesRequest]) (*connect.Response[v1.CreateAlloyPipelinesResponse], error)
	// CreatePipeline creates a new pipeline and returns it.
	CreatePipeline(context.Context, *connect.Request[v1.CreatePipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// UpdatePipeline updates an existing pipeline and returns it.
	UpdatePipeline(context.Context, *connect.Request[v1.UpdatePipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// UpsertPipeline creates a new pipeline or updates an existing one and returns it.
	UpsertPipeline(context.Context, *connect.Request[v1.UpsertPipelineRequest]) (*connect.Response[v1.Pipeline], error)
	// DeletePipeline deletes a pipeline by name.
	DeletePipeline(context.Context, *connect.Request[v1.DeletePipelineRequest]) (*connect.Response[v1.DeletePipelineResponse], error)
}

// NewPipelineServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPipelineServiceHandler(svc PipelineServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	pipelineServiceGetPipelineHandler := connect.NewUnaryHandler(
		PipelineServiceGetPipelineProcedure,
		svc.GetPipeline,
		connect.WithSchema(pipelineServiceGetPipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceGetPipelineIDHandler := connect.NewUnaryHandler(
		PipelineServiceGetPipelineIDProcedure,
		svc.GetPipelineID,
		connect.WithSchema(pipelineServiceGetPipelineIDMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceListPipelinesHandler := connect.NewUnaryHandler(
		PipelineServiceListPipelinesProcedure,
		svc.ListPipelines,
		connect.WithSchema(pipelineServiceListPipelinesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceCreateAlloyPipelinesHandler := connect.NewUnaryHandler(
		PipelineServiceCreateAlloyPipelinesProcedure,
		svc.CreateAlloyPipelines,
		connect.WithSchema(pipelineServiceCreateAlloyPipelinesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceCreatePipelineHandler := connect.NewUnaryHandler(
		PipelineServiceCreatePipelineProcedure,
		svc.CreatePipeline,
		connect.WithSchema(pipelineServiceCreatePipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceUpdatePipelineHandler := connect.NewUnaryHandler(
		PipelineServiceUpdatePipelineProcedure,
		svc.UpdatePipeline,
		connect.WithSchema(pipelineServiceUpdatePipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceUpsertPipelineHandler := connect.NewUnaryHandler(
		PipelineServiceUpsertPipelineProcedure,
		svc.UpsertPipeline,
		connect.WithSchema(pipelineServiceUpsertPipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceDeletePipelineHandler := connect.NewUnaryHandler(
		PipelineServiceDeletePipelineProcedure,
		svc.DeletePipeline,
		connect.WithSchema(pipelineServiceDeletePipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/pipeline.v1.PipelineService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PipelineServiceGetPipelineProcedure:
			pipelineServiceGetPipelineHandler.ServeHTTP(w, r)
		case PipelineServiceGetPipelineIDProcedure:
			pipelineServiceGetPipelineIDHandler.ServeHTTP(w, r)
		case PipelineServiceListPipelinesProcedure:
			pipelineServiceListPipelinesHandler.ServeHTTP(w, r)
		case PipelineServiceCreateAlloyPipelinesProcedure:
			pipelineServiceCreateAlloyPipelinesHandler.ServeHTTP(w, r)
		case PipelineServiceCreatePipelineProcedure:
			pipelineServiceCreatePipelineHandler.ServeHTTP(w, r)
		case PipelineServiceUpdatePipelineProcedure:
			pipelineServiceUpdatePipelineHandler.ServeHTTP(w, r)
		case PipelineServiceUpsertPipelineProcedure:
			pipelineServiceUpsertPipelineHandler.ServeHTTP(w, r)
		case PipelineServiceDeletePipelineProcedure:
			pipelineServiceDeletePipelineHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPipelineServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPipelineServiceHandler struct{}

func (UnimplementedPipelineServiceHandler) GetPipeline(context.Context, *connect.Request[v1.GetPipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.GetPipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) GetPipelineID(context.Context, *connect.Request[v1.GetPipelineIDRequest]) (*connect.Response[v1.GetPipelineIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.GetPipelineID is not implemented"))
}

func (UnimplementedPipelineServiceHandler) ListPipelines(context.Context, *connect.Request[v1.ListPipelinesRequest]) (*connect.Response[v1.Pipelines], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.ListPipelines is not implemented"))
}

func (UnimplementedPipelineServiceHandler) CreateAlloyPipelines(context.Context, *connect.Request[v1.CreateAlloyPipelinesRequest]) (*connect.Response[v1.CreateAlloyPipelinesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.CreateAlloyPipelines is not implemented"))
}

func (UnimplementedPipelineServiceHandler) CreatePipeline(context.Context, *connect.Request[v1.CreatePipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.CreatePipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) UpdatePipeline(context.Context, *connect.Request[v1.UpdatePipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.UpdatePipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) UpsertPipeline(context.Context, *connect.Request[v1.UpsertPipelineRequest]) (*connect.Response[v1.Pipeline], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.UpsertPipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) DeletePipeline(context.Context, *connect.Request[v1.DeletePipelineRequest]) (*connect.Response[v1.DeletePipelineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pipeline.v1.PipelineService.DeletePipeline is not implemented"))
}
