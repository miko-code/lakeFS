// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/treeverse/lakefs/api/gen/models"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/authentication"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/branches"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/commits"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/objects"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/repositories"
)

// NewLakefsAPI creates a new Lakefs instance
func NewLakefsAPI(spec *loads.Document) *LakefsAPI {
	return &LakefsAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		PreServerShutdown:     func() {},
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		BinProducer:           runtime.ByteStreamProducer(),
		JSONProducer:          runtime.JSONProducer(),
		AuthenticationGetAuthenticationHandler: authentication.GetAuthenticationHandlerFunc(func(params authentication.GetAuthenticationParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation authentication.GetAuthentication has not yet been implemented")
		}),
		CommitsCommitHandler: commits.CommitHandlerFunc(func(params commits.CommitParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation commits.Commit has not yet been implemented")
		}),
		BranchesCreateBranchHandler: branches.CreateBranchHandlerFunc(func(params branches.CreateBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation branches.CreateBranch has not yet been implemented")
		}),
		RepositoriesCreateRepositoryHandler: repositories.CreateRepositoryHandlerFunc(func(params repositories.CreateRepositoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation repositories.CreateRepository has not yet been implemented")
		}),
		BranchesDeleteBranchHandler: branches.DeleteBranchHandlerFunc(func(params branches.DeleteBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation branches.DeleteBranch has not yet been implemented")
		}),
		ObjectsDeleteObjectHandler: objects.DeleteObjectHandlerFunc(func(params objects.DeleteObjectParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation objects.DeleteObject has not yet been implemented")
		}),
		RepositoriesDeleteRepositoryHandler: repositories.DeleteRepositoryHandlerFunc(func(params repositories.DeleteRepositoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation repositories.DeleteRepository has not yet been implemented")
		}),
		BranchesDiffBranchHandler: branches.DiffBranchHandlerFunc(func(params branches.DiffBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation branches.DiffBranch has not yet been implemented")
		}),
		BranchesDiffBranchesHandler: branches.DiffBranchesHandlerFunc(func(params branches.DiffBranchesParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation branches.DiffBranches has not yet been implemented")
		}),
		BranchesGetBranchHandler: branches.GetBranchHandlerFunc(func(params branches.GetBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation branches.GetBranch has not yet been implemented")
		}),
		CommitsGetBranchCommitLogHandler: commits.GetBranchCommitLogHandlerFunc(func(params commits.GetBranchCommitLogParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation commits.GetBranchCommitLog has not yet been implemented")
		}),
		CommitsGetCommitHandler: commits.GetCommitHandlerFunc(func(params commits.GetCommitParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation commits.GetCommit has not yet been implemented")
		}),
		ObjectsGetObjectHandler: objects.GetObjectHandlerFunc(func(params objects.GetObjectParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation objects.GetObject has not yet been implemented")
		}),
		RepositoriesGetRepositoryHandler: repositories.GetRepositoryHandlerFunc(func(params repositories.GetRepositoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation repositories.GetRepository has not yet been implemented")
		}),
		BranchesListBranchesHandler: branches.ListBranchesHandlerFunc(func(params branches.ListBranchesParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation branches.ListBranches has not yet been implemented")
		}),
		ObjectsListObjectsHandler: objects.ListObjectsHandlerFunc(func(params objects.ListObjectsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation objects.ListObjects has not yet been implemented")
		}),
		RepositoriesListRepositoriesHandler: repositories.ListRepositoriesHandlerFunc(func(params repositories.ListRepositoriesParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation repositories.ListRepositories has not yet been implemented")
		}),
		BranchesRevertBranchHandler: branches.RevertBranchHandlerFunc(func(params branches.RevertBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation branches.RevertBranch has not yet been implemented")
		}),
		ObjectsStatObjectHandler: objects.StatObjectHandlerFunc(func(params objects.StatObjectParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation objects.StatObject has not yet been implemented")
		}),
		ObjectsUploadObjectHandler: objects.UploadObjectHandlerFunc(func(params objects.UploadObjectParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation objects.UploadObject has not yet been implemented")
		}), // Applies when the Authorization header is set with the Basic scheme
		BasicAuthAuth: func(user string, pass string) (*models.User, error) {
			return nil, errors.NotImplemented("basic auth  (basic_auth) has not yet been implemented")
		},

		// Applies when the "token" query is set
		DownloadTokenAuth: func(token string) (*models.User, error) {
			return nil, errors.NotImplemented("api key auth (download_token) token from query param [token] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*LakefsAPI lakeFS HTTP API */
type LakefsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator
	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer
	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BasicAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	BasicAuthAuth func(string, string) (*models.User, error)

	// DownloadTokenAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key token provided in the query
	DownloadTokenAuth func(string) (*models.User, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AuthenticationGetAuthenticationHandler sets the operation handler for the get authentication operation
	AuthenticationGetAuthenticationHandler authentication.GetAuthenticationHandler
	// CommitsCommitHandler sets the operation handler for the commit operation
	CommitsCommitHandler commits.CommitHandler
	// BranchesCreateBranchHandler sets the operation handler for the create branch operation
	BranchesCreateBranchHandler branches.CreateBranchHandler
	// RepositoriesCreateRepositoryHandler sets the operation handler for the create repository operation
	RepositoriesCreateRepositoryHandler repositories.CreateRepositoryHandler
	// BranchesDeleteBranchHandler sets the operation handler for the delete branch operation
	BranchesDeleteBranchHandler branches.DeleteBranchHandler
	// ObjectsDeleteObjectHandler sets the operation handler for the delete object operation
	ObjectsDeleteObjectHandler objects.DeleteObjectHandler
	// RepositoriesDeleteRepositoryHandler sets the operation handler for the delete repository operation
	RepositoriesDeleteRepositoryHandler repositories.DeleteRepositoryHandler
	// BranchesDiffBranchHandler sets the operation handler for the diff branch operation
	BranchesDiffBranchHandler branches.DiffBranchHandler
	// BranchesDiffBranchesHandler sets the operation handler for the diff branches operation
	BranchesDiffBranchesHandler branches.DiffBranchesHandler
	// BranchesGetBranchHandler sets the operation handler for the get branch operation
	BranchesGetBranchHandler branches.GetBranchHandler
	// CommitsGetBranchCommitLogHandler sets the operation handler for the get branch commit log operation
	CommitsGetBranchCommitLogHandler commits.GetBranchCommitLogHandler
	// CommitsGetCommitHandler sets the operation handler for the get commit operation
	CommitsGetCommitHandler commits.GetCommitHandler
	// ObjectsGetObjectHandler sets the operation handler for the get object operation
	ObjectsGetObjectHandler objects.GetObjectHandler
	// RepositoriesGetRepositoryHandler sets the operation handler for the get repository operation
	RepositoriesGetRepositoryHandler repositories.GetRepositoryHandler
	// BranchesListBranchesHandler sets the operation handler for the list branches operation
	BranchesListBranchesHandler branches.ListBranchesHandler
	// ObjectsListObjectsHandler sets the operation handler for the list objects operation
	ObjectsListObjectsHandler objects.ListObjectsHandler
	// RepositoriesListRepositoriesHandler sets the operation handler for the list repositories operation
	RepositoriesListRepositoriesHandler repositories.ListRepositoriesHandler
	// BranchesRevertBranchHandler sets the operation handler for the revert branch operation
	BranchesRevertBranchHandler branches.RevertBranchHandler
	// ObjectsStatObjectHandler sets the operation handler for the stat object operation
	ObjectsStatObjectHandler objects.StatObjectHandler
	// ObjectsUploadObjectHandler sets the operation handler for the upload object operation
	ObjectsUploadObjectHandler objects.UploadObjectHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *LakefsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *LakefsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *LakefsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *LakefsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *LakefsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *LakefsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *LakefsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the LakefsAPI
func (o *LakefsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BasicAuthAuth == nil {
		unregistered = append(unregistered, "BasicAuthAuth")
	}

	if o.DownloadTokenAuth == nil {
		unregistered = append(unregistered, "TokenAuth")
	}

	if o.AuthenticationGetAuthenticationHandler == nil {
		unregistered = append(unregistered, "Authentication.GetAuthenticationHandler")
	}

	if o.CommitsCommitHandler == nil {
		unregistered = append(unregistered, "Commits.CommitHandler")
	}

	if o.BranchesCreateBranchHandler == nil {
		unregistered = append(unregistered, "Branches.CreateBranchHandler")
	}

	if o.RepositoriesCreateRepositoryHandler == nil {
		unregistered = append(unregistered, "Repositories.CreateRepositoryHandler")
	}

	if o.BranchesDeleteBranchHandler == nil {
		unregistered = append(unregistered, "Branches.DeleteBranchHandler")
	}

	if o.ObjectsDeleteObjectHandler == nil {
		unregistered = append(unregistered, "Objects.DeleteObjectHandler")
	}

	if o.RepositoriesDeleteRepositoryHandler == nil {
		unregistered = append(unregistered, "Repositories.DeleteRepositoryHandler")
	}

	if o.BranchesDiffBranchHandler == nil {
		unregistered = append(unregistered, "Branches.DiffBranchHandler")
	}

	if o.BranchesDiffBranchesHandler == nil {
		unregistered = append(unregistered, "Branches.DiffBranchesHandler")
	}

	if o.BranchesGetBranchHandler == nil {
		unregistered = append(unregistered, "Branches.GetBranchHandler")
	}

	if o.CommitsGetBranchCommitLogHandler == nil {
		unregistered = append(unregistered, "Commits.GetBranchCommitLogHandler")
	}

	if o.CommitsGetCommitHandler == nil {
		unregistered = append(unregistered, "Commits.GetCommitHandler")
	}

	if o.ObjectsGetObjectHandler == nil {
		unregistered = append(unregistered, "Objects.GetObjectHandler")
	}

	if o.RepositoriesGetRepositoryHandler == nil {
		unregistered = append(unregistered, "Repositories.GetRepositoryHandler")
	}

	if o.BranchesListBranchesHandler == nil {
		unregistered = append(unregistered, "Branches.ListBranchesHandler")
	}

	if o.ObjectsListObjectsHandler == nil {
		unregistered = append(unregistered, "Objects.ListObjectsHandler")
	}

	if o.RepositoriesListRepositoriesHandler == nil {
		unregistered = append(unregistered, "Repositories.ListRepositoriesHandler")
	}

	if o.BranchesRevertBranchHandler == nil {
		unregistered = append(unregistered, "Branches.RevertBranchHandler")
	}

	if o.ObjectsStatObjectHandler == nil {
		unregistered = append(unregistered, "Objects.StatObjectHandler")
	}

	if o.ObjectsUploadObjectHandler == nil {
		unregistered = append(unregistered, "Objects.UploadObjectHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *LakefsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *LakefsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {

		case "basic_auth":
			result[name] = o.BasicAuthenticator(func(username, password string) (interface{}, error) {
				return o.BasicAuthAuth(username, password)
			})

		case "download_token":

			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.DownloadTokenAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *LakefsAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *LakefsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *LakefsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *LakefsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the lakefs API
func (o *LakefsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *LakefsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/authentication"] = authentication.NewGetAuthentication(o.context, o.AuthenticationGetAuthenticationHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories/{repositoryId}/branches/{branchId}/commits"] = commits.NewCommit(o.context, o.CommitsCommitHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories/{repositoryId}/branches"] = branches.NewCreateBranch(o.context, o.BranchesCreateBranchHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories"] = repositories.NewCreateRepository(o.context, o.RepositoriesCreateRepositoryHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/repositories/{repositoryId}/branches/{branchId}"] = branches.NewDeleteBranch(o.context, o.BranchesDeleteBranchHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/repositories/{repositoryId}/branches/{branchId}/objects"] = objects.NewDeleteObject(o.context, o.ObjectsDeleteObjectHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/repositories/{repositoryId}"] = repositories.NewDeleteRepository(o.context, o.RepositoriesDeleteRepositoryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}/diff"] = branches.NewDiffBranch(o.context, o.BranchesDiffBranchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}/diff/{otherBranchId}"] = branches.NewDiffBranches(o.context, o.BranchesDiffBranchesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}"] = branches.NewGetBranch(o.context, o.BranchesGetBranchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}/commits"] = commits.NewGetBranchCommitLog(o.context, o.CommitsGetBranchCommitLogHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/commits/{commitId}"] = commits.NewGetCommit(o.context, o.CommitsGetCommitHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}/objects"] = objects.NewGetObject(o.context, o.ObjectsGetObjectHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}"] = repositories.NewGetRepository(o.context, o.RepositoriesGetRepositoryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches"] = branches.NewListBranches(o.context, o.BranchesListBranchesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}/objects/ls"] = objects.NewListObjects(o.context, o.ObjectsListObjectsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories"] = repositories.NewListRepositories(o.context, o.RepositoriesListRepositoriesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/repositories/{repositoryId}/branches/{branchId}"] = branches.NewRevertBranch(o.context, o.BranchesRevertBranchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}/objects/stat"] = objects.NewStatObject(o.context, o.ObjectsStatObjectHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories/{repositoryId}/branches/{branchId}/objects"] = objects.NewUploadObject(o.context, o.ObjectsUploadObjectHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *LakefsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *LakefsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *LakefsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *LakefsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
