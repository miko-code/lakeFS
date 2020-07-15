package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/treeverse/lakefs/api/gen/models"
	"github.com/treeverse/lakefs/api/gen/restapi/operations"
	authop "github.com/treeverse/lakefs/api/gen/restapi/operations/auth"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/branches"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/commits"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/objects"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/refs"
	"github.com/treeverse/lakefs/api/gen/restapi/operations/repositories"
	retentionop "github.com/treeverse/lakefs/api/gen/restapi/operations/retention"
	"github.com/treeverse/lakefs/auth"
	"github.com/treeverse/lakefs/auth/model"
	"github.com/treeverse/lakefs/block"
	"github.com/treeverse/lakefs/catalog"
	"github.com/treeverse/lakefs/db"
	"github.com/treeverse/lakefs/httputil"
	"github.com/treeverse/lakefs/logging"
	"github.com/treeverse/lakefs/onboard"
	"github.com/treeverse/lakefs/permissions"
	"github.com/treeverse/lakefs/retention"
	"github.com/treeverse/lakefs/stats"
	"github.com/treeverse/lakefs/upload"
)

const (
	// Maximum amount of results returned for paginated queries to the API
	MaxResultsPerPage int64 = 1000
)

type Dependencies struct {
	ctx          context.Context
	Cataloger    catalog.Cataloger
	Auth         auth.Service
	BlockAdapter block.Adapter
	Stats        stats.Collector
	Retention    retention.Service
	Dedup        *DedupHandler
	logger       logging.Logger
}

func (d *Dependencies) WithContext(ctx context.Context) *Dependencies {
	return &Dependencies{
		ctx:          ctx,
		Cataloger:    d.Cataloger,
		Auth:         d.Auth,
		BlockAdapter: d.BlockAdapter.WithContext(ctx),
		Stats:        d.Stats,
		Retention:    d.Retention,
		Dedup:        d.Dedup,
		logger:       d.logger.WithContext(ctx),
	}
}

func (d *Dependencies) LogAction(action string) {
	logging.FromContext(d.ctx).
		WithField("action", action).
		WithField("message_type", "action").
		Debug("performing API action")
	d.Stats.CollectEvent("api_server", action)
}

type Controller struct {
	deps *Dependencies
}

func NewController(cataloger catalog.Cataloger, auth auth.Service, blockAdapter block.Adapter, stats stats.Collector, retention retention.Service, logger logging.Logger) *Controller {
	c := &Controller{
		deps: &Dependencies{
			ctx:          context.Background(),
			Cataloger:    cataloger,
			Auth:         auth,
			BlockAdapter: blockAdapter,
			Stats:        stats,
			Retention:    retention,
			Dedup:        NewDedupHandler(blockAdapter),
			logger:       logger,
		},
	}
	c.deps.Dedup.Start()
	return c
}

func (c *Controller) Close() error {
	if c == nil || c.deps == nil {
		return nil
	}
	return c.deps.Dedup.Close()
}

func (c *Controller) Context() context.Context {
	if c.deps.ctx != nil {
		return c.deps.ctx
	}
	return context.Background()
}

// Configure attaches our API operations to a generated swagger API stub
// Adding new handlers requires also adding them here so that the generated server will use them
func (c *Controller) Configure(api *operations.LakefsAPI) {

	// Register operations here
	api.AuthGetCurrentUserHandler = c.GetCurrentUserHandler()
	api.AuthListUsersHandler = c.ListUsersHandler()
	api.AuthGetUserHandler = c.GetUserHandler()
	api.AuthCreateUserHandler = c.CreateUserHandler()
	api.AuthDeleteUserHandler = c.DeleteUserHandler()
	api.AuthGetGroupHandler = c.GetGroupHandler()
	api.AuthListGroupsHandler = c.ListGroupsHandler()
	api.AuthCreateGroupHandler = c.CreateGroupHandler()
	api.AuthDeleteGroupHandler = c.DeleteGroupHandler()
	api.AuthListPoliciesHandler = c.ListPoliciesHandler()
	api.AuthCreatePolicyHandler = c.CreatePolicyHandler()
	api.AuthGetPolicyHandler = c.GetPolicyHandler()
	api.AuthDeletePolicyHandler = c.DeletePolicyHandler()
	api.AuthUpdatePolicyHandler = c.UpdatePolicyHandler()
	api.AuthListGroupMembersHandler = c.ListGroupMembersHandler()
	api.AuthAddGroupMembershipHandler = c.AddGroupMembershipHandler()
	api.AuthDeleteGroupMembershipHandler = c.DeleteGroupMembershipHandler()
	api.AuthListUserCredentialsHandler = c.ListUserCredentialsHandler()
	api.AuthCreateCredentialsHandler = c.CreateCredentialsHandler()
	api.AuthDeleteCredentialsHandler = c.DeleteCredentialsHandler()
	api.AuthGetCredentialsHandler = c.GetCredentialsHandler()
	api.AuthListUserGroupsHandler = c.ListUserGroupsHandler()
	api.AuthListUserPoliciesHandler = c.ListUserPoliciesHandler()
	api.AuthAttachPolicyToUserHandler = c.AttachPolicyToUserHandler()
	api.AuthDetachPolicyFromUserHandler = c.DetachPolicyFromUserHandler()
	api.AuthListGroupPoliciesHandler = c.ListGroupPoliciesHandler()
	api.AuthAttachPolicyToGroupHandler = c.AttachPolicyToGroupHandler()
	api.AuthDetachPolicyFromGroupHandler = c.DetachPolicyFromGroupHandler()

	api.RepositoriesListRepositoriesHandler = c.ListRepositoriesHandler()
	api.RepositoriesGetRepositoryHandler = c.GetRepoHandler()
	api.RepositoriesCreateRepositoryHandler = c.CreateRepositoryHandler()
	api.RepositoriesDeleteRepositoryHandler = c.DeleteRepositoryHandler()

	api.BranchesListBranchesHandler = c.ListBranchesHandler()
	api.BranchesGetBranchHandler = c.GetBranchHandler()
	api.BranchesCreateBranchHandler = c.CreateBranchHandler()
	api.BranchesDeleteBranchHandler = c.DeleteBranchHandler()
	api.BranchesRevertBranchHandler = c.RevertBranchHandler()

	api.CommitsCommitHandler = c.CommitHandler()
	api.CommitsGetCommitHandler = c.GetCommitHandler()
	api.CommitsGetBranchCommitLogHandler = c.CommitsGetBranchCommitLogHandler()

	api.RefsDiffRefsHandler = c.RefsDiffRefsHandler()
	api.BranchesDiffBranchHandler = c.BranchesDiffBranchHandler()
	api.RefsMergeIntoBranchHandler = c.MergeMergeIntoBranchHandler()

	api.ObjectsStatObjectHandler = c.ObjectsStatObjectHandler()
	api.ObjectsGetUnderlyingPropertiesHandler = c.ObjectsGetUnderlyingPropertiesHandler()
	api.ObjectsListObjectsHandler = c.ObjectsListObjectsHandler()
	api.ObjectsGetObjectHandler = c.ObjectsGetObjectHandler()
	api.ObjectsUploadObjectHandler = c.ObjectsUploadObjectHandler()
	api.ObjectsDeleteObjectHandler = c.ObjectsDeleteObjectHandler()

	api.RetentionGetRetentionPolicyHandler = c.RetentionGetRetentionPolicyHandler()
	api.RetentionUpdateRetentionPolicyHandler = c.RetentionUpdateRetentionPolicyHandler()
}

func (c *Controller) setupRequest(user *models.User, r *http.Request, permissions []permissions.Permission) (*Dependencies, error) {
	// add user to context
	ctx := logging.AddFields(r.Context(), logging.Fields{"user": user.ID})
	ctx = context.WithValue(ctx, "user", user)
	deps := c.deps.WithContext(ctx)
	return deps, authorize(deps.Auth, user, permissions)
}

func createPaginator(nextToken string, amountResults int) *models.Pagination {
	return &models.Pagination{
		HasMore:    swag.Bool(nextToken != ""),
		MaxPerPage: swag.Int64(MaxResultsPerPage),
		NextOffset: nextToken,
		Results:    swag.Int64(int64(amountResults)),
	}
}

func pageAmount(i *int64) int {
	inti := int(swag.Int64Value(i))
	if inti > int(MaxResultsPerPage) {
		return int(MaxResultsPerPage)
	}
	if inti <= 0 {
		return 100
	}
	return inti
}

func (c *Controller) GetCurrentUserHandler() authop.GetCurrentUserHandler {
	return authop.GetCurrentUserHandlerFunc(func(params authop.GetCurrentUserParams, user *models.User) middleware.Responder {
		return authop.NewGetCurrentUserOK().WithPayload(&authop.GetCurrentUserOKBody{
			User: user,
		})
	})
}

func (c *Controller) ListRepositoriesHandler() repositories.ListRepositoriesHandler {
	return repositories.ListRepositoriesHandlerFunc(func(params repositories.ListRepositoriesParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListRepositoriesAction,
				Resource: permissions.All,
			},
		})

		if err != nil {
			return repositories.NewListRepositoriesUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("list_repos")

		after, amount := getPaginationParams(params.After, params.Amount)

		repos, hasMore, err := deps.Cataloger.ListRepositories(c.Context(), amount, after)
		if err != nil {
			return repositories.NewListRepositoriesDefault(http.StatusInternalServerError).
				WithPayload(responseError("error listing repositories: %s", err))
		}

		repoList := make([]*models.Repository, len(repos))
		var lastID string
		for i, repo := range repos {
			repoList[i] = &models.Repository{
				StorageNamespace: repo.StorageNamespace,
				CreationDate:     repo.CreationDate.Unix(),
				DefaultBranch:    repo.DefaultBranch,
				ID:               repo.Name,
			}
			lastID = repo.Name
		}
		returnValue := repositories.NewListRepositoriesOK().WithPayload(&repositories.ListRepositoriesOKBody{
			Pagination: &models.Pagination{
				HasMore:    swag.Bool(hasMore),
				Results:    swag.Int64(int64(len(repoList))),
				MaxPerPage: swag.Int64(MaxResultsPerPage),
			},
			Results: repoList,
		})
		if hasMore {
			returnValue.Payload.Pagination.NextOffset = lastID
		}

		return returnValue
	})
}

func getPaginationParams(swagAfter *string, swagAmount *int64) (string, int) {
	// amount
	after := ""
	amount := MaxResultsPerPage
	if swagAmount != nil {
		amount = swag.Int64Value(swagAmount)
	}

	// paginate after
	if swagAfter != nil {
		after = swag.StringValue(swagAfter)
	}
	return after, int(amount)
}

func (c *Controller) GetRepoHandler() repositories.GetRepositoryHandler {
	return repositories.GetRepositoryHandlerFunc(func(params repositories.GetRepositoryParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadRepositoryAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return repositories.NewGetRepositoryUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_repo")
		repo, err := deps.Cataloger.GetRepository(c.Context(), params.Repository)
		if errors.Is(err, db.ErrNotFound) {
			return repositories.NewGetRepositoryNotFound().
				WithPayload(responseError("repository not found"))
		}
		if err != nil {
			return repositories.NewGetRepositoryDefault(http.StatusInternalServerError).
				WithPayload(responseError("error fetching repository: %s", err))
		}

		return repositories.NewGetRepositoryOK().
			WithPayload(&models.Repository{
				StorageNamespace: repo.StorageNamespace,
				CreationDate:     repo.CreationDate.Unix(),
				DefaultBranch:    repo.DefaultBranch,
				ID:               repo.Name,
			})
	})
}

func (c *Controller) GetCommitHandler() commits.GetCommitHandler {
	return commits.GetCommitHandlerFunc(func(params commits.GetCommitParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadCommitAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return commits.NewGetCommitUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_commit")
		commit, err := deps.Cataloger.GetCommit(c.Context(), params.Repository, params.CommitID)
		if errors.Is(err, db.ErrNotFound) {
			return commits.NewGetCommitNotFound().WithPayload(responseError("commit not found"))
		}
		if err != nil {
			return commits.NewGetCommitDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}
		return commits.NewGetCommitOK().WithPayload(&models.Commit{
			Committer:    commit.Committer,
			CreationDate: commit.CreationDate.Unix(),
			ID:           params.CommitID,
			Message:      commit.Message,
			Metadata:     commit.Metadata,
			Parents:      commit.Parents,
		})
	})
}

func (c *Controller) CommitHandler() commits.CommitHandler {
	return commits.CommitHandlerFunc(func(params commits.CommitParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateCommitAction,
				Resource: permissions.BranchArn(params.Repository, params.Branch),
			},
		})
		if err != nil {
			return commits.NewCommitUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("create_commit")
		userModel, err := c.deps.Auth.GetUser(user.ID)
		if err != nil {
			return commits.NewCommitUnauthorized().WithPayload(responseErrorFrom(err))
		}
		committer := userModel.DisplayName
		commitMessage := swag.StringValue(params.Commit.Message)
		commit, err := deps.Cataloger.Commit(c.Context(), params.Repository,
			params.Branch, commitMessage, committer, params.Commit.Metadata)
		if err != nil {
			return commits.NewCommitDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}
		return commits.NewCommitCreated().WithPayload(&models.Commit{
			Committer:    commit.Committer,
			CreationDate: commit.CreationDate.Unix(),
			ID:           commit.Reference,
			Message:      commit.Message,
			Metadata:     commit.Metadata,
			Parents:      commit.Parents,
		})
	})
}

func (c *Controller) CommitsGetBranchCommitLogHandler() commits.GetBranchCommitLogHandler {
	return commits.GetBranchCommitLogHandlerFunc(func(params commits.GetBranchCommitLogParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadBranchAction,
				Resource: permissions.BranchArn(params.Repository, params.Branch),
			},
		})
		if err != nil {
			return commits.NewGetBranchCommitLogUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_branch")
		cataloger := deps.Cataloger

		after, amount := getPaginationParams(params.After, params.Amount)
		// get commit log
		commitLog, hasMore, err := cataloger.ListCommits(c.Context(), params.Repository, params.Branch, after, amount)
		if err != nil {
			return commits.NewGetBranchCommitLogDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		serializedCommits := make([]*models.Commit, len(commitLog))
		lastId := ""
		for i, commit := range commitLog {
			serializedCommits[i] = &models.Commit{
				Committer:    commit.Committer,
				CreationDate: commit.CreationDate.Unix(),
				ID:           commit.Reference,
				Message:      commit.Message,
				Metadata:     commit.Metadata,
				Parents:      commit.Parents,
			}
			lastId = commit.Reference
		}

		returnValue := commits.NewGetBranchCommitLogOK().WithPayload(&commits.GetBranchCommitLogOKBody{
			Pagination: &models.Pagination{
				HasMore:    swag.Bool(hasMore),
				Results:    swag.Int64(int64(len(serializedCommits))),
				MaxPerPage: swag.Int64(MaxResultsPerPage),
			},
			Results: serializedCommits,
		})
		if hasMore {
			returnValue.Payload.Pagination.NextOffset = lastId
		}
		return returnValue
	})
}

func ensureStorageNamespaceRW(adapter block.Adapter, storageNamespace string) error {
	const (
		dummyKey  = "dummy"
		dummyData = "this is dummy data - created by lakefs in order to check accessibility "
	)

	err := adapter.Put(block.ObjectPointer{StorageNamespace: storageNamespace, Identifier: dummyKey}, int64(len(dummyData)), bytes.NewReader([]byte(dummyData)), block.PutOpts{})
	if err != nil {
		return err
	}

	_, err = adapter.Get(block.ObjectPointer{StorageNamespace: storageNamespace, Identifier: dummyKey}, int64(len(dummyData)))
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) CreateRepositoryHandler() repositories.CreateRepositoryHandler {
	return repositories.CreateRepositoryHandlerFunc(func(params repositories.CreateRepositoryParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateRepositoryAction,
				Resource: permissions.RepoArn(swag.StringValue(params.Repository.ID)),
			},
		})
		if err != nil {
			return repositories.NewCreateRepositoryUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("create_repo")

		err = ensureStorageNamespaceRW(deps.BlockAdapter, swag.StringValue(params.Repository.StorageNamespace))
		if err != nil {
			return repositories.NewCreateRepositoryBadRequest().
				WithPayload(responseError("error creating repository: could not access storage namespace"))
		}
		err = deps.Cataloger.CreateRepository(c.Context(),
			swag.StringValue(params.Repository.ID),
			swag.StringValue(params.Repository.StorageNamespace),
			params.Repository.DefaultBranch)
		if err != nil {
			return repositories.NewGetRepositoryDefault(http.StatusInternalServerError).
				WithPayload(responseError(fmt.Sprintf("error creating repository: %s", err)))
		}

		repo, err := deps.Cataloger.GetRepository(c.Context(), swag.StringValue(params.Repository.ID))
		if err != nil {
			return repositories.NewGetRepositoryDefault(http.StatusInternalServerError).
				WithPayload(responseError(fmt.Sprintf("error creating repository: %s", err)))
		}

		return repositories.NewCreateRepositoryCreated().WithPayload(&models.Repository{
			StorageNamespace: repo.StorageNamespace,
			CreationDate:     repo.CreationDate.Unix(),
			DefaultBranch:    repo.DefaultBranch,
			ID:               repo.Name,
		})
	})
}

func (c *Controller) DeleteRepositoryHandler() repositories.DeleteRepositoryHandler {
	return repositories.DeleteRepositoryHandlerFunc(func(params repositories.DeleteRepositoryParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DeleteRepositoryAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return repositories.NewDeleteRepositoryUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("delete_repo")
		cataloger := deps.Cataloger
		err = cataloger.DeleteRepository(c.Context(), params.Repository)
		if errors.Is(err, db.ErrNotFound) {
			return repositories.NewDeleteRepositoryNotFound().
				WithPayload(responseError("repository not found"))
		}
		if err != nil {
			return repositories.NewDeleteRepositoryDefault(http.StatusInternalServerError).
				WithPayload(responseError("error deleting repository"))
		}

		return repositories.NewDeleteRepositoryNoContent()
	})
}

func (c *Controller) ListBranchesHandler() branches.ListBranchesHandler {
	return branches.ListBranchesHandlerFunc(func(params branches.ListBranchesParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListBranchesAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return branches.NewListBranchesUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("list_branches")
		cataloger := deps.Cataloger

		after, amount := getPaginationParams(params.After, params.Amount)

		res, hasMore, err := cataloger.ListBranches(c.Context(), params.Repository, "", amount, after)
		if err != nil {
			return branches.NewListBranchesDefault(http.StatusInternalServerError).
				WithPayload(responseError("could not list branches: %s", err))
		}

		branchList := make([]string, len(res))
		var lastId string
		for i, branch := range res {
			branchList[i] = branch.Name
			lastId = branch.Name
		}
		returnValue := branches.NewListBranchesOK().WithPayload(&branches.ListBranchesOKBody{
			Pagination: &models.Pagination{
				HasMore:    swag.Bool(hasMore),
				Results:    swag.Int64(int64(len(branchList))),
				MaxPerPage: swag.Int64(MaxResultsPerPage),
			},
			Results: branchList,
		})

		if hasMore {
			returnValue.Payload.Pagination.NextOffset = lastId
		}

		return returnValue
	})
}

func (c *Controller) GetBranchHandler() branches.GetBranchHandler {
	return branches.GetBranchHandlerFunc(func(params branches.GetBranchParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadBranchAction,
				Resource: permissions.BranchArn(params.Repository, params.Branch),
			},
		})
		if err != nil {
			return branches.NewGetBranchUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_branch")
		reference, err := deps.Cataloger.GetBranchReference(c.Context(), params.Repository, params.Branch)
		if errors.Is(err, db.ErrNotFound) {
			return branches.NewGetBranchNotFound().
				WithPayload(responseError("branch not found"))
		}
		if err != nil {
			return branches.NewGetBranchDefault(http.StatusInternalServerError).
				WithPayload(responseError("error fetching branch: %s", err))
		}

		return branches.NewGetBranchOK().WithPayload(reference)
	})
}

func (c *Controller) CreateBranchHandler() branches.CreateBranchHandler {
	return branches.CreateBranchHandlerFunc(func(params branches.CreateBranchParams, user *models.User) middleware.Responder {
		repository := params.Repository
		branch := swag.StringValue(params.Branch.Name)
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateBranchAction,
				Resource: permissions.BranchArn(repository, branch),
			},
		})
		if err != nil {
			return branches.NewCreateBranchUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("create_branch")
		cataloger := deps.Cataloger
		sourceBranch := swag.StringValue(params.Branch.Source)
		err = cataloger.CreateBranch(c.Context(), repository, branch, sourceBranch)
		if err != nil {
			return branches.NewCreateBranchDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}
		// TODO(barak): create branch should return the reference of the new branch's commit
		return branches.NewCreateBranchCreated().WithPayload(branch)
	})
}

func (c *Controller) DeleteBranchHandler() branches.DeleteBranchHandler {
	return branches.DeleteBranchHandlerFunc(func(params branches.DeleteBranchParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DeleteBranchAction,
				Resource: permissions.BranchArn(params.Repository, params.Branch),
			},
		})
		if err != nil {
			return branches.NewDeleteBranchUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("delete_branch")
		cataloger := deps.Cataloger
		err = cataloger.DeleteBranch(c.Context(), params.Repository, params.Branch)
		if errors.Is(err, db.ErrNotFound) {
			return branches.NewDeleteBranchNotFound().
				WithPayload(responseError("branch not found"))
		}
		if err != nil {
			return branches.NewDeleteBranchDefault(http.StatusInternalServerError).
				WithPayload(responseError("error fetching branch: %s", err))
		}

		return branches.NewDeleteBranchNoContent()
	})
}

func (c *Controller) MergeMergeIntoBranchHandler() refs.MergeIntoBranchHandler {
	return refs.MergeIntoBranchHandlerFunc(func(params refs.MergeIntoBranchParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateCommitAction,
				Resource: permissions.BranchArn(params.Repository, params.DestinationRef),
			},
		})
		if err != nil {
			return refs.NewMergeIntoBranchUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("merge_branches")
		userModel, err := deps.Auth.GetUser(user.ID)
		if err != nil {
			return refs.NewMergeIntoBranchUnauthorized().WithPayload(responseErrorFrom(err))
		}
		var message string
		var metadata map[string]string
		if params.Merge != nil {
			message = params.Merge.Message
			metadata = params.Merge.Metadata
		}
		res, err := deps.Cataloger.Merge(c.Context(),
			params.Repository, params.SourceRef, params.DestinationRef,
			userModel.DisplayName,
			message,
			metadata)

		// convert merge differences into merge results
		var mergeResults []*models.MergeResult
		if res != nil {
			mergeResults = make([]*models.MergeResult, len(res.Differences))
			for i, d := range res.Differences {
				mergeResults[i] = transformDifferenceToMergeResult(d)
			}
		}

		switch err {
		case nil:
			pl := new(refs.MergeIntoBranchOKBody)
			pl.Results = mergeResults
			return refs.NewMergeIntoBranchOK().WithPayload(pl)
		case catalog.ErrUnsupportedRelation:
			return refs.NewMergeIntoBranchDefault(http.StatusInternalServerError).WithPayload(responseError("branches have no common base"))
		case catalog.ErrBranchNotFound:
			return refs.NewMergeIntoBranchDefault(http.StatusInternalServerError).WithPayload(responseError("a branch does not exist "))
		case catalog.ErrConflictFound:
			pl := new(refs.MergeIntoBranchConflictBody)
			pl.Results = mergeResults
			return refs.NewMergeIntoBranchConflict().WithPayload(pl)
		case catalog.ErrNoDifferenceWasFound:
			return refs.NewMergeIntoBranchDefault(http.StatusInternalServerError).WithPayload(responseError("no difference was found"))
		default:
			return refs.NewMergeIntoBranchDefault(http.StatusInternalServerError).WithPayload(responseError("internal error"))
		}
	})
}

func (c *Controller) BranchesDiffBranchHandler() branches.DiffBranchHandler {
	return branches.DiffBranchHandlerFunc(func(params branches.DiffBranchParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListObjectsAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return branches.NewDiffBranchUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("diff_workspace")
		cataloger := deps.Cataloger
		diff, err := cataloger.DiffUncommitted(c.Context(), params.Repository, params.Branch)
		if err != nil {
			return branches.NewDiffBranchDefault(http.StatusInternalServerError).
				WithPayload(responseError("could not diff branch: %s", err))
		}

		results := make([]*models.Diff, len(diff))
		for i, d := range diff {
			results[i] = transformDifferenceToDiff(d)
		}

		return branches.NewDiffBranchOK().WithPayload(&branches.DiffBranchOKBody{Results: results})
	})
}

func (c *Controller) RefsDiffRefsHandler() refs.DiffRefsHandler {
	return refs.DiffRefsHandlerFunc(func(params refs.DiffRefsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListObjectsAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return refs.NewDiffRefsUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("diff_refs")
		cataloger := deps.Cataloger
		diff, err := cataloger.Diff(c.Context(), params.Repository, params.LeftRef, params.RightRef)
		if err != nil {
			return refs.NewDiffRefsDefault(http.StatusInternalServerError).
				WithPayload(responseError("could not diff references: %s", err))
		}

		results := make([]*models.Diff, len(diff))
		for i, d := range diff {
			results[i] = transformDifferenceToDiff(d)
		}
		return refs.NewDiffRefsOK().WithPayload(&refs.DiffRefsOKBody{Results: results})
	})
}

func (c *Controller) ObjectsStatObjectHandler() objects.StatObjectHandler {
	return objects.StatObjectHandlerFunc(func(params objects.StatObjectParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadObjectAction,
				Resource: permissions.ObjectArn(params.Repository, params.Path),
			},
		})
		if err != nil {
			return objects.NewStatObjectUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("stat_object")
		cataloger := deps.Cataloger

		entry, err := cataloger.GetEntry(c.Context(), params.Repository, params.Ref, params.Path)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewStatObjectNotFound().WithPayload(responseError("resource not found"))
		}
		if err != nil {
			return objects.NewStatObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		// serialize entry
		return objects.NewStatObjectOK().WithPayload(&models.ObjectStats{
			Checksum:  entry.Checksum,
			Mtime:     entry.CreationDate.Unix(),
			Path:      params.Path,
			PathType:  models.ObjectStatsPathTypeOBJECT,
			SizeBytes: entry.Size,
		})
	})
}

func (c *Controller) ObjectsGetUnderlyingPropertiesHandler() objects.GetUnderlyingPropertiesHandler {
	return objects.GetUnderlyingPropertiesHandlerFunc(func(params objects.GetUnderlyingPropertiesParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadObjectAction,
				Resource: permissions.ObjectArn(params.Repository, params.Path),
			},
		})
		if err != nil {
			return objects.NewGetUnderlyingPropertiesUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("object_underlying_properties")
		cataloger := deps.Cataloger

		// read repo
		repo, err := cataloger.GetRepository(c.Context(), params.Repository)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewGetObjectNotFound().WithPayload(responseError("resource not found"))
		}
		if err != nil {
			return objects.NewGetObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		entry, err := cataloger.GetEntry(c.Context(),
			params.Repository, params.Ref, params.Path)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewGetUnderlyingPropertiesNotFound().WithPayload(responseError("resource not found"))
		}
		if err != nil {
			return objects.NewGetUnderlyingPropertiesDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		// read object properties from underlying storage
		properties, err := c.deps.BlockAdapter.GetProperties(block.ObjectPointer{StorageNamespace: repo.StorageNamespace, Identifier: entry.PhysicalAddress})
		if err != nil {
			return objects.NewGetUnderlyingPropertiesDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		// serialize properties
		return objects.NewGetUnderlyingPropertiesOK().WithPayload(&models.UnderlyingObjectProperties{
			StorageClass: properties.StorageClass,
		})
	})
}

func (c *Controller) ObjectsGetObjectHandler() objects.GetObjectHandler {
	return objects.GetObjectHandlerFunc(func(params objects.GetObjectParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadObjectAction,
				Resource: permissions.ObjectArn(params.Repository, params.Path),
			},
		})
		if err != nil {
			return objects.NewGetObjectUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_object")
		cataloger := deps.Cataloger

		// read repo
		repo, err := cataloger.GetRepository(c.Context(), params.Repository)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewGetObjectNotFound().WithPayload(responseError("resource not found"))
		}
		if err != nil {
			return objects.NewGetObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		// read the FS entry
		entry, err := cataloger.GetEntry(c.Context(), params.Repository, params.Ref, params.Path)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewGetObjectNotFound().WithPayload(responseError("resource not found"))
		}
		if err != nil {
			return objects.NewGetObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}
		// setup response
		res := objects.NewGetObjectOK()
		res.ETag = httputil.ETag(entry.Checksum)
		res.LastModified = httputil.HeaderTimestamp(entry.CreationDate)
		res.ContentDisposition = fmt.Sprintf("filename=\"%s\"", filepath.Base(entry.Path))

		// build a response as a multi-reader
		res.ContentLength = entry.Size
		reader, err := deps.BlockAdapter.Get(block.ObjectPointer{StorageNamespace: repo.StorageNamespace, Identifier: entry.PhysicalAddress}, entry.Size)
		if err != nil {
			return objects.NewGetObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		// done
		res.Payload = reader
		return res
	})
}

func (c *Controller) ObjectsListObjectsHandler() objects.ListObjectsHandler {
	return objects.ListObjectsHandlerFunc(func(params objects.ListObjectsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListObjectsAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return objects.NewListObjectsUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("list_objects")
		cataloger := deps.Cataloger

		after, amount := getPaginationParams(params.After, params.Amount)

		res, hasMore, err := cataloger.ListEntries(
			c.Context(),
			params.Repository,
			params.Ref,
			swag.StringValue(params.Tree),
			after,
			amount)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewListObjectsNotFound().WithPayload(responseError("could not find requested path"))
		}
		if err != nil {
			return objects.NewListObjectsDefault(http.StatusInternalServerError).
				WithPayload(responseError("error while listing objects: %s", err))
		}

		objList := make([]*models.ObjectStats, len(res))
		var lastId string
		for i, entry := range res {
			typ := models.ObjectStatsPathTypeOBJECT
			mtime := entry.CreationDate.Unix()
			if entry.CreationDate.IsZero() {
				mtime = 0
			}
			objList[i] = &models.ObjectStats{
				Checksum:  entry.Checksum,
				Mtime:     mtime,
				Path:      entry.Path,
				PathType:  typ,
				SizeBytes: entry.Size,
			}
			lastId = entry.Path
		}
		returnValue := objects.NewListObjectsOK().WithPayload(&objects.ListObjectsOKBody{
			Pagination: &models.Pagination{
				HasMore:    swag.Bool(hasMore),
				Results:    swag.Int64(int64(len(objList))),
				MaxPerPage: swag.Int64(MaxResultsPerPage),
			},
			Results: objList,
		})

		if hasMore {
			returnValue.Payload.Pagination.NextOffset = lastId
		}
		return returnValue
	})
}

const noopUploadObject = false
const noopCreateEntry = false

func (c *Controller) ObjectsUploadObjectHandler() objects.UploadObjectHandler {
	return objects.UploadObjectHandlerFunc(func(params objects.UploadObjectParams, user *models.User) middleware.Responder {
		if noopUploadObject {
			return objects.NewUploadObjectCreated().WithPayload(&models.ObjectStats{
				Checksum:  "cc",
				Mtime:     time.Now().UTC().Unix(),
				Path:      params.Path,
				PathType:  models.ObjectStatsPathTypeOBJECT,
				SizeBytes: 1,
			})
		}
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.WriteObjectAction,
				Resource: permissions.ObjectArn(params.Repository, params.Path),
			},
		})
		if err != nil {
			return objects.NewUploadObjectUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("put_object")
		cataloger := deps.Cataloger

		repo, err := cataloger.GetRepository(c.Context(), params.Repository)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewUploadObjectNotFound().WithPayload(responseError("resource not found"))
		}
		if err != nil {
			return objects.NewUploadObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}
		// workaround in order to extract file content-length using swagger
		file, ok := params.Content.(*runtime.File)
		if !ok {
			return objects.NewUploadObjectNotFound().WithPayload(responseError("failed extracting size from file"))
		}
		byteSize := file.Header.Size

		// read the content
		blob, err := upload.WriteBlob(deps.BlockAdapter, repo.StorageNamespace, params.Content, byteSize, block.PutOpts{StorageClass: params.StorageClass})
		if err != nil {
			return objects.NewUploadObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		if noopCreateEntry {
			return objects.NewUploadObjectCreated().WithPayload(&models.ObjectStats{
				Checksum:  blob.Checksum,
				Mtime:     time.Now().UTC().Unix(),
				Path:      params.Path,
				PathType:  models.ObjectStatsPathTypeOBJECT,
				SizeBytes: blob.Size,
			})
		}
		// write metadata
		writeTime := time.Now()
		entry := catalog.Entry{
			Path:            params.Path,
			PhysicalAddress: blob.PhysicalAddress,
			CreationDate:    writeTime,
			Size:            blob.Size,
			Checksum:        blob.Checksum,
		}
		err = cataloger.CreateEntryDedup(c.Context(), repo.Name, params.Branch, entry, catalog.DedupParams{
			ID:               blob.DedupID,
			Ch:               deps.Dedup.Channel(),
			StorageNamespace: repo.StorageNamespace,
		})
		if err != nil {
			return objects.NewUploadObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}
		return objects.NewUploadObjectCreated().WithPayload(&models.ObjectStats{
			Checksum:  blob.Checksum,
			Mtime:     writeTime.Unix(),
			Path:      params.Path,
			PathType:  models.ObjectStatsPathTypeOBJECT,
			SizeBytes: blob.Size,
		})
	})
}

func (c *Controller) ObjectsDeleteObjectHandler() objects.DeleteObjectHandler {
	return objects.DeleteObjectHandlerFunc(func(params objects.DeleteObjectParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DeleteObjectAction,
				Resource: permissions.ObjectArn(params.Repository, params.Path),
			},
		})
		if err != nil {
			return objects.NewDeleteObjectUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("delete_object")
		cataloger := deps.Cataloger

		err = cataloger.DeleteEntry(c.Context(), params.Repository, params.Branch, params.Path)
		if errors.Is(err, db.ErrNotFound) {
			return objects.NewDeleteObjectNotFound().WithPayload(responseError("resource not found"))
		}
		if err != nil {
			return objects.NewDeleteObjectDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		return objects.NewDeleteObjectNoContent()
	})
}
func (c *Controller) RevertBranchHandler() branches.RevertBranchHandler {
	return branches.RevertBranchHandlerFunc(func(params branches.RevertBranchParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.RevertBranchAction,
				Resource: permissions.BranchArn(params.Repository, params.Branch),
			},
		})
		if err != nil {
			return branches.NewRevertBranchUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("revert_branch")
		cataloger := deps.Cataloger

		ctx := c.Context()
		switch swag.StringValue(params.Revert.Type) {
		case models.RevertCreationTypeCOMMIT:
			err = cataloger.RollbackCommit(ctx, params.Repository, params.Revert.Commit)
		case models.RevertCreationTypeTREE:
			err = cataloger.ResetEntries(ctx, params.Repository, params.Branch, params.Revert.Path)
		case models.RevertCreationTypeRESET:
			err = cataloger.ResetBranch(ctx, params.Repository, params.Branch)
		case models.RevertCreationTypeOBJECT:
			err = cataloger.ResetEntry(ctx, params.Repository, params.Branch, params.Revert.Path)
		default:
			return branches.NewRevertBranchNotFound().
				WithPayload(responseError("revert type not found"))
		}
		if errors.Is(err, db.ErrNotFound) {
			return branches.NewRevertBranchNotFound().WithPayload(responseError("branch not found"))
		}
		if err != nil {
			return branches.NewRevertBranchDefault(http.StatusInternalServerError).WithPayload(responseErrorFrom(err))
		}

		return branches.NewRevertBranchNoContent()
	})
}

func (c *Controller) CreateUserHandler() authop.CreateUserHandler {
	return authop.CreateUserHandlerFunc(func(params authop.CreateUserParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateUserAction,
				Resource: permissions.UserArn(swag.StringValue(params.User.ID)),
			},
		})
		if err != nil {
			return authop.NewCreateUserUnauthorized().
				WithPayload(responseErrorFrom(err))
		}
		u := &model.User{
			CreatedAt:   time.Now(),
			DisplayName: swag.StringValue(params.User.ID),
		}
		err = deps.Auth.CreateUser(u)
		deps.LogAction("create_user")
		if err != nil {
			return authop.NewCreateUserDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewCreateUserCreated().
			WithPayload(&models.User{
				CreationDate: u.CreatedAt.Unix(),
				ID:           u.DisplayName,
			})
	})
}

func (c *Controller) ListUsersHandler() authop.ListUsersHandler {
	return authop.ListUsersHandlerFunc(func(params authop.ListUsersParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListUsersAction,
				Resource: permissions.All,
			},
		})
		if err != nil {
			return authop.NewListUsersUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_users")
		users, paginator, err := deps.Auth.ListUsers(&model.PaginationParams{
			After:  swag.StringValue(params.After),
			Amount: pageAmount(params.Amount),
		})
		if err != nil {
			return authop.NewListUsersDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.User, len(users))
		for i, u := range users {
			response[i] = &models.User{
				CreationDate: u.CreatedAt.Unix(),
				ID:           u.DisplayName,
			}
		}

		return authop.NewListUsersOK().
			WithPayload(&authop.ListUsersOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) GetUserHandler() authop.GetUserHandler {
	return authop.GetUserHandlerFunc(func(params authop.GetUserParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadUserAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewGetUserUnauthorized().
				WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_user")
		u, err := deps.Auth.GetUser(params.UserID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewGetUserNotFound().
				WithPayload(responseError("user not found"))
		}
		if err != nil {
			return authop.NewGetUserDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewGetUserOK().
			WithPayload(&models.User{
				CreationDate: u.CreatedAt.Unix(),
				ID:           u.DisplayName,
			})
	})
}

func (c *Controller) DeleteUserHandler() authop.DeleteUserHandler {
	return authop.DeleteUserHandlerFunc(func(params authop.DeleteUserParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DeleteUserAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewDeleteUserUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("delete_user")
		err = deps.Auth.DeleteUser(params.UserID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewDeleteUserNotFound().
				WithPayload(responseError("user not found"))
		}
		if err != nil {
			return authop.NewDeleteUserDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewDeleteUserNoContent()
	})
}

func (c *Controller) GetGroupHandler() authop.GetGroupHandler {
	return authop.GetGroupHandlerFunc(func(params authop.GetGroupParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadGroupAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewGetGroupUnauthorized().
				WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_group")
		g, err := deps.Auth.GetGroup(params.GroupID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewGetGroupNotFound().
				WithPayload(responseError("group not found"))
		}
		if err != nil {
			return authop.NewGetGroupDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewGetGroupOK().
			WithPayload(&models.Group{
				CreationDate: g.CreatedAt.Unix(),
				ID:           g.DisplayName,
			})
	})
}

func (c *Controller) ListGroupsHandler() authop.ListGroupsHandler {
	return authop.ListGroupsHandlerFunc(func(params authop.ListGroupsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListGroupsAction,
				Resource: permissions.All,
			},
		})
		if err != nil {
			return authop.NewListGroupsUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_groups")
		groups, paginator, err := deps.Auth.ListGroups(&model.PaginationParams{
			After:  swag.StringValue(params.After),
			Amount: pageAmount(params.Amount),
		})

		if err != nil {
			return authop.NewListGroupsDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.Group, len(groups))
		for i, g := range groups {
			response[i] = &models.Group{
				CreationDate: g.CreatedAt.Unix(),
				ID:           g.DisplayName,
			}
		}

		return authop.NewListGroupsOK().
			WithPayload(&authop.ListGroupsOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) CreateGroupHandler() authop.CreateGroupHandler {
	return authop.CreateGroupHandlerFunc(func(params authop.CreateGroupParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateGroupAction,
				Resource: permissions.GroupArn(swag.StringValue(params.Group.ID)),
			},
		})
		if err != nil {
			return authop.NewCreateGroupUnauthorized().
				WithPayload(responseErrorFrom(err))
		}
		g := &model.Group{
			CreatedAt:   time.Now(),
			DisplayName: swag.StringValue(params.Group.ID),
		}

		deps.LogAction("create_group")
		err = deps.Auth.CreateGroup(g)
		if err != nil {
			return authop.NewCreateGroupDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewCreateGroupCreated().
			WithPayload(&models.Group{
				CreationDate: g.CreatedAt.Unix(),
				ID:           g.DisplayName,
			})
	})
}

func (c *Controller) DeleteGroupHandler() authop.DeleteGroupHandler {
	return authop.DeleteGroupHandlerFunc(func(params authop.DeleteGroupParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DeleteGroupAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewDeleteGroupUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("delete_group")
		err = deps.Auth.DeleteGroup(params.GroupID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewDeleteGroupNotFound().
				WithPayload(responseError("group not found"))
		}
		if err != nil {
			return authop.NewDeleteGroupDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}
		return authop.NewDeleteGroupNoContent()
	})
}

func serializePolicy(p *model.Policy) *models.Policy {
	stmts := make([]*models.Statement, len(p.Statement))
	for i, s := range p.Statement {
		stmts[i] = &models.Statement{
			Action:   s.Action,
			Effect:   swag.String(s.Effect),
			Resource: swag.String(s.Resource),
		}
	}
	return &models.Policy{
		ID:           swag.String(p.DisplayName),
		CreationDate: p.CreatedAt.Unix(),
		Statement:    stmts,
	}
}

func (c *Controller) ListPoliciesHandler() authop.ListPoliciesHandler {
	return authop.ListPoliciesHandlerFunc(func(params authop.ListPoliciesParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListPoliciesAction,
				Resource: permissions.All,
			},
		})
		if err != nil {
			return authop.NewListPoliciesUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_policies")
		policies, paginator, err := deps.Auth.ListPolicies(&model.PaginationParams{
			After:  swag.StringValue(params.After),
			Amount: pageAmount(params.Amount),
		})
		if err != nil {
			return authop.NewListPoliciesDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.Policy, len(policies))
		for i, p := range policies {
			response[i] = serializePolicy(p)
		}

		return authop.NewListPoliciesOK().
			WithPayload(&authop.ListPoliciesOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) CreatePolicyHandler() authop.CreatePolicyHandler {
	return authop.CreatePolicyHandlerFunc(func(params authop.CreatePolicyParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreatePolicyAction,
				Resource: permissions.PolicyArn(swag.StringValue(params.Policy.ID)),
			},
		})
		if err != nil {
			return authop.NewCreatePolicyUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		stmts := make(model.Statements, len(params.Policy.Statement))
		for i, apiStatement := range params.Policy.Statement {
			stmts[i] = model.Statement{
				Effect:   swag.StringValue(apiStatement.Effect),
				Action:   apiStatement.Action,
				Resource: swag.StringValue(apiStatement.Resource),
			}
		}

		p := &model.Policy{
			CreatedAt:   time.Now(),
			DisplayName: swag.StringValue(params.Policy.ID),
			Statement:   stmts,
		}

		deps.LogAction("create_policy")
		err = deps.Auth.WritePolicy(p)
		if err != nil {
			return authop.NewCreatePolicyDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewCreatePolicyCreated().
			WithPayload(serializePolicy(p))
	})
}

func (c *Controller) GetPolicyHandler() authop.GetPolicyHandler {
	return authop.GetPolicyHandlerFunc(func(params authop.GetPolicyParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadPolicyAction,
				Resource: permissions.PolicyArn(params.PolicyID),
			},
		})
		if err != nil {
			return authop.NewGetPolicyUnauthorized().
				WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_policy")
		p, err := deps.Auth.GetPolicy(params.PolicyID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewGetPolicyNotFound().
				WithPayload(responseError("policy not found"))
		}
		if err != nil {
			return authop.NewGetPolicyDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewGetPolicyOK().
			WithPayload(serializePolicy(p))
	})
}

func (c *Controller) UpdatePolicyHandler() authop.UpdatePolicyHandler {
	return authop.UpdatePolicyHandlerFunc(func(params authop.UpdatePolicyParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.UpdatePolicyAction,
				Resource: permissions.PolicyArn(params.PolicyID),
			},
		})
		if err != nil {
			return authop.NewUpdatePolicyUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		stmts := make(model.Statements, len(params.Policy.Statement))
		for i, apiStatement := range params.Policy.Statement {
			stmts[i] = model.Statement{
				Effect:   swag.StringValue(apiStatement.Effect),
				Action:   apiStatement.Action,
				Resource: swag.StringValue(apiStatement.Resource),
			}
		}

		p := &model.Policy{
			CreatedAt:   time.Now(),
			DisplayName: swag.StringValue(params.Policy.ID),
			Statement:   stmts,
		}

		deps.LogAction("update_policy")
		err = deps.Auth.WritePolicy(p)
		if err != nil {
			return authop.NewUpdatePolicyDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewUpdatePolicyOK().
			WithPayload(serializePolicy(p))
	})
}

func (c *Controller) DeletePolicyHandler() authop.DeletePolicyHandler {
	return authop.DeletePolicyHandlerFunc(func(params authop.DeletePolicyParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DeletePolicyAction,
				Resource: permissions.PolicyArn(params.PolicyID),
			},
		})
		if err != nil {
			return authop.NewDeletePolicyUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("delete_policy")
		err = deps.Auth.DeletePolicy(params.PolicyID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewDeletePolicyNotFound().
				WithPayload(responseError("policy not found"))
		}
		if err != nil {
			return authop.NewDeletePolicyDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}
		return authop.NewDeletePolicyNoContent()
	})
}

func (c *Controller) ListGroupMembersHandler() authop.ListGroupMembersHandler {
	return authop.ListGroupMembersHandlerFunc(func(params authop.ListGroupMembersParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadGroupAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewListGroupMembersUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_group_users")
		users, paginator, err := deps.Auth.ListGroupUsers(params.GroupID, &model.PaginationParams{
			After:  swag.StringValue(params.After),
			Amount: pageAmount(params.Amount),
		})
		if err != nil {
			return authop.NewListGroupMembersDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.User, len(users))
		for i, u := range users {
			response[i] = &models.User{
				CreationDate: u.CreatedAt.Unix(),
				ID:           u.DisplayName,
			}
		}

		return authop.NewListGroupMembersOK().
			WithPayload(&authop.ListGroupMembersOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) AddGroupMembershipHandler() authop.AddGroupMembershipHandler {
	return authop.AddGroupMembershipHandlerFunc(func(params authop.AddGroupMembershipParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.AddGroupMemberAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewAddGroupMembershipUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("add_user_to_group")
		err = deps.Auth.AddUserToGroup(params.UserID, params.GroupID)
		if err != nil {
			return authop.NewAddGroupMembershipDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewAddGroupMembershipCreated()
	})
}

func (c *Controller) DeleteGroupMembershipHandler() authop.DeleteGroupMembershipHandler {
	return authop.DeleteGroupMembershipHandlerFunc(func(params authop.DeleteGroupMembershipParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.RemoveGroupMemberAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewDeleteGroupMembershipUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("remove_user_from_group")
		err = deps.Auth.RemoveUserFromGroup(params.UserID, params.GroupID)
		if err != nil {
			return authop.NewDeleteGroupMembershipDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewDeleteGroupMembershipNoContent()
	})
}

func (c *Controller) ListUserCredentialsHandler() authop.ListUserCredentialsHandler {
	return authop.ListUserCredentialsHandlerFunc(func(params authop.ListUserCredentialsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ListCredentialsAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewListUserCredentialsUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_user_credentials")
		credentials, paginator, err := deps.Auth.ListUserCredentials(params.UserID, &model.PaginationParams{
			After:  swag.StringValue(params.After),
			Amount: pageAmount(params.Amount),
		})
		if err != nil {
			return authop.NewListUserCredentialsDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.Credentials, len(credentials))
		for i, c := range credentials {
			response[i] = &models.Credentials{
				AccessKeyID:  c.AccessKeyId,
				CreationDate: c.IssuedDate.Unix(),
			}
		}

		return authop.NewListUserCredentialsOK().
			WithPayload(&authop.ListUserCredentialsOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) CreateCredentialsHandler() authop.CreateCredentialsHandler {
	return authop.CreateCredentialsHandlerFunc(func(params authop.CreateCredentialsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateCredentialsAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewCreateCredentialsUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("create_credentials")
		credentials, err := deps.Auth.CreateCredentials(params.UserID)
		if err != nil {
			return authop.NewCreateCredentialsDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewCreateCredentialsCreated().
			WithPayload(&models.CredentialsWithSecret{
				AccessKeyID:     credentials.AccessKeyId,
				AccessSecretKey: credentials.AccessSecretKey,
				CreationDate:    credentials.IssuedDate.Unix(),
			})
	})
}

func (c *Controller) DeleteCredentialsHandler() authop.DeleteCredentialsHandler {
	return authop.DeleteCredentialsHandlerFunc(func(params authop.DeleteCredentialsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DeleteCredentialsAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewDeleteCredentialsUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("delete_credentials")
		err = deps.Auth.DeleteCredentials(params.UserID, params.AccessKeyID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewDeleteCredentialsNotFound().
				WithPayload(responseError("credentials not found"))
		}
		if err != nil {
			return authop.NewDeleteCredentialsDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewDeleteCredentialsNoContent()
	})
}

func (c *Controller) GetCredentialsHandler() authop.GetCredentialsHandler {
	return authop.GetCredentialsHandlerFunc(func(params authop.GetCredentialsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadCredentialsAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewGetCredentialsUnauthorized().
				WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("get_credentials_for_user")
		credentials, err := deps.Auth.GetCredentialsForUser(params.UserID, params.AccessKeyID)
		if errors.Is(err, db.ErrNotFound) {
			return authop.NewGetCredentialsNotFound().
				WithPayload(responseError("credentials not found"))
		}
		if err != nil {
			return authop.NewGetCredentialsDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewGetCredentialsOK().
			WithPayload(&models.Credentials{
				AccessKeyID:  credentials.AccessKeyId,
				CreationDate: credentials.IssuedDate.Unix(),
			})
	})
}

func (c *Controller) ListUserGroupsHandler() authop.ListUserGroupsHandler {
	return authop.ListUserGroupsHandlerFunc(func(params authop.ListUserGroupsParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadUserAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewListUserGroupsUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_user_groups")
		groups, paginator, err := deps.Auth.ListUserGroups(params.UserID, &model.PaginationParams{
			After:  swag.StringValue(params.After),
			Amount: pageAmount(params.Amount),
		})
		if err != nil {
			return authop.NewListUserGroupsDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.Group, len(groups))
		for i, g := range groups {
			response[i] = &models.Group{
				CreationDate: g.CreatedAt.Unix(),
				ID:           g.DisplayName,
			}
		}

		return authop.NewListUserGroupsOK().
			WithPayload(&authop.ListUserGroupsOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) ListUserPoliciesHandler() authop.ListUserPoliciesHandler {
	return authop.ListUserPoliciesHandlerFunc(func(params authop.ListUserPoliciesParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadUserAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewListUserPoliciesUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_user_policies")
		var policies []*model.Policy
		var paginator *model.Paginator
		if swag.BoolValue(params.Effective) {
			policies, paginator, err = deps.Auth.ListEffectivePolicies(params.UserID, &model.PaginationParams{
				After:  swag.StringValue(params.After),
				Amount: pageAmount(params.Amount),
			})
		} else {
			policies, paginator, err = deps.Auth.ListUserPolicies(params.UserID, &model.PaginationParams{
				After:  swag.StringValue(params.After),
				Amount: pageAmount(params.Amount),
			})
		}

		if err != nil {
			return authop.NewListUserPoliciesDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.Policy, len(policies))
		for i, p := range policies {
			response[i] = serializePolicy(p)
		}

		return authop.NewListUserPoliciesOK().
			WithPayload(&authop.ListUserPoliciesOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) AttachPolicyToUserHandler() authop.AttachPolicyToUserHandler {
	return authop.AttachPolicyToUserHandlerFunc(func(params authop.AttachPolicyToUserParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.AttachPolicyAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewAttachPolicyToUserUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("attach_policy_to_user")
		err = deps.Auth.AttachPolicyToUser(params.PolicyID, params.UserID)
		if err != nil {
			return authop.NewAttachPolicyToUserDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewAttachPolicyToUserCreated()
	})
}

func (c *Controller) DetachPolicyFromUserHandler() authop.DetachPolicyFromUserHandler {
	return authop.DetachPolicyFromUserHandlerFunc(func(params authop.DetachPolicyFromUserParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DetachPolicyAction,
				Resource: permissions.UserArn(params.UserID),
			},
		})
		if err != nil {
			return authop.NewDetachPolicyFromUserUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("detach_policy_from_user")
		err = deps.Auth.DetachPolicyFromUser(params.PolicyID, params.UserID)
		if err != nil {
			return authop.NewDetachPolicyFromUserDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewDetachPolicyFromUserNoContent()
	})
}

func (c *Controller) ListGroupPoliciesHandler() authop.ListGroupPoliciesHandler {
	return authop.ListGroupPoliciesHandlerFunc(func(params authop.ListGroupPoliciesParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.ReadGroupAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewListGroupPoliciesUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("list_user_policies")
		policies, paginator, err := deps.Auth.ListGroupPolicies(params.GroupID, &model.PaginationParams{
			After:  swag.StringValue(params.After),
			Amount: pageAmount(params.Amount),
		})
		if err != nil {
			return authop.NewListGroupPoliciesDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		response := make([]*models.Policy, len(policies))
		for i, p := range policies {
			response[i] = serializePolicy(p)
		}

		return authop.NewListGroupPoliciesOK().
			WithPayload(&authop.ListGroupPoliciesOKBody{
				Pagination: createPaginator(paginator.NextPageToken, len(response)),
				Results:    response,
			})
	})
}

func (c *Controller) AttachPolicyToGroupHandler() authop.AttachPolicyToGroupHandler {
	return authop.AttachPolicyToGroupHandlerFunc(func(params authop.AttachPolicyToGroupParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.AttachPolicyAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewAttachPolicyToGroupUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("attach_policy_to_group")
		err = deps.Auth.AttachPolicyToGroup(params.PolicyID, params.GroupID)
		if err != nil {
			return authop.NewAttachPolicyToGroupDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewAttachPolicyToGroupCreated()
	})
}

func (c *Controller) DetachPolicyFromGroupHandler() authop.DetachPolicyFromGroupHandler {
	return authop.DetachPolicyFromGroupHandlerFunc(func(params authop.DetachPolicyFromGroupParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.DetachPolicyAction,
				Resource: permissions.GroupArn(params.GroupID),
			},
		})
		if err != nil {
			return authop.NewDetachPolicyFromGroupUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("detach_policy_from_group")
		err = deps.Auth.DetachPolicyFromGroup(params.PolicyID, params.GroupID)
		if err != nil {
			return authop.NewDetachPolicyFromGroupDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}

		return authop.NewDetachPolicyFromGroupNoContent()
	})
}

func (c *Controller) RetentionGetRetentionPolicyHandler() retentionop.GetRetentionPolicyHandler {
	return retentionop.GetRetentionPolicyHandlerFunc(func(params retentionop.GetRetentionPolicyParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.RetentionReadPolicyAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})

		if err != nil {
			return retentionop.NewGetRetentionPolicyUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		deps.LogAction("get_retention_policy")

		policy, err := deps.Retention.GetPolicy(params.Repository)
		if err != nil {
			return retentionop.NewGetRetentionPolicyDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}
		return retentionop.NewGetRetentionPolicyOK().WithPayload(policy)
	})
}

func (c *Controller) RetentionUpdateRetentionPolicyHandler() retentionop.UpdateRetentionPolicyHandler {
	return retentionop.UpdateRetentionPolicyHandlerFunc(func(params retentionop.UpdateRetentionPolicyParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.RetentionWritePolicyAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return retentionop.NewUpdateRetentionPolicyUnauthorized().
				WithPayload(responseErrorFrom(err))
		}

		err = deps.Retention.UpdatePolicy(params.Repository, params.Policy)
		if err != nil {
			return retentionop.NewUpdateRetentionPolicyDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}
		return retentionop.NewUpdateRetentionPolicyCreated()
	})
}

func (c *Controller) ImportFromS3InventoryHandler() repositories.ImportFromS3InventoryHandler {
	return repositories.ImportFromS3InventoryHandlerFunc(func(params repositories.ImportFromS3InventoryParams, user *models.User) middleware.Responder {
		deps, err := c.setupRequest(user, params.HTTPRequest, []permissions.Permission{
			{
				Action:   permissions.CreateRepositoryAction,
				Resource: permissions.RepoArn(params.Repository),
			},
		})
		if err != nil {
			return repositories.NewImportFromS3InventoryUnauthorized().WithPayload(responseErrorFrom(err))
		}
		deps.LogAction("import_from_s3_inventory")
		userModel, err := c.deps.Auth.GetUser(user.ID)
		username := "lakeFS"
		if err == nil {
			username = userModel.DisplayName
		}
		importer, err := onboard.CreateImporter(deps.Cataloger, deps.BlockAdapter, username, params.ManifestURL, params.Repository)
		if err != nil {
			return repositories.NewImportFromS3InventoryDefault(http.StatusInternalServerError).
				WithPayload(responseErrorFrom(err))
		}
		var diff *onboard.InventoryDiff
		if *params.DryRun {
			diff, err = importer.Import(deps.ctx, true)
			if err != nil {
				return repositories.NewImportFromS3InventoryDefault(http.StatusInternalServerError).
					WithPayload(responseErrorFrom(err))
			}
		} else {
			repo, err := deps.Cataloger.GetRepository(c.Context(), params.Repository)
			if err != nil {
				return repositories.NewImportFromS3InventoryNotFound().
					WithPayload(responseErrorFrom(err))
			}
			_, err = deps.Cataloger.GetBranchReference(deps.ctx, params.Repository, onboard.DefaultBranchName)
			if errors.Is(err, db.ErrNotFound) {
				err = deps.Cataloger.CreateBranch(deps.ctx, params.Repository, onboard.DefaultBranchName, repo.DefaultBranch)
				if err != nil {
					return repositories.NewImportFromS3InventoryDefault(http.StatusInternalServerError).
						WithPayload(responseErrorFrom(err))
				}
			} else if err != nil {
				return repositories.NewImportFromS3InventoryDefault(http.StatusInternalServerError).
					WithPayload(responseErrorFrom(err))
			}
			diff, err = importer.Import(params.HTTPRequest.Context(), false)
			if err != nil {
				return repositories.NewImportFromS3InventoryDefault(http.StatusInternalServerError).
					WithPayload(responseErrorFrom(err))
			}
		}
		return repositories.NewImportFromS3InventoryCreated().WithPayload(&repositories.ImportFromS3InventoryCreatedBody{
			IsDryRun:           *params.DryRun,
			PreviousImportDate: diff.PreviousImportDate.Unix(),
			PreviousManifest:   diff.PreviousInventoryURL,
			AddedOrChanged:     int64(len(diff.AddedOrChanged)),
			Deleted:            int64(len(diff.Deleted)),
		})
	})
}