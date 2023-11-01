package service

import (
	"graphql/graph/model"
)

type Service interface {
	CreateUser(nu model.NewUser) (*model.User, error)
	CreateCompany(nc model.NewCompany) (*model.Company, error)
	FetchAllCompanies() ([]*model.Company, error)
	FetchCompanyByID(cid string) (*model.Company, error)
	CreateJob(nj model.NewJob) (*model.Job, error)
	FetchAllJobs() ([]*model.Job, error)
	FetchJobByCompanyID(cid string) ([]*model.Job, error)
	FetchJobByID(jid string) (*model.Job, error)
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
