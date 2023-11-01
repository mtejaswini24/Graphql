package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"graphql/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.S.CreateUser(input)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {
	return r.S.CreateCompany(input)
}

// CreateJob is the resolver for the createJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.Job, error) {
	return r.S.CreateJob(input)
}

// FetchAllCompanies is the resolver for the fetchAllCompanies field.
func (r *queryResolver) FetchAllCompanies(ctx context.Context) ([]*model.Company, error) {
	return r.S.FetchAllCompanies()
}

// FetchCompanyByID is the resolver for the fetchCompanyById field.
func (r *queryResolver) FetchCompanyByID(ctx context.Context, cid string) (*model.Company, error) {
	return r.S.FetchCompanyByID(cid)
}

// FetchAllJobs is the resolver for the fetchAllJobs field.
func (r *queryResolver) FetchAllJobs(ctx context.Context) ([]*model.Job, error) {
	return r.S.FetchAllJobs()
}

// FetchJobByID is the resolver for the fetchJobById field.
func (r *queryResolver) FetchJobByID(ctx context.Context, jid string) (*model.Job, error) {
	return r.S.FetchJobByID(jid)
}

// FetchJobByCompanyID is the resolver for the fetchJobByCompanyId field.
func (r *queryResolver) FetchJobByCompanyID(ctx context.Context, cid string) ([]*model.Job, error) {
	return r.S.FetchJobByCompanyID(cid)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
