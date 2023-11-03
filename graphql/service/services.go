package service

import (
	"graphql/graph/model"
	"graphql/models"
)

type Service interface {
	CreateUser(nu model.NewUser) (*model.User, error)
	CreateCompany(nc model.NewCompany) (*model.Company, error)
	FetchAllCompanies() ([]*model.Company, error)
	FetchCompanyByID(cid string) (*model.Company, error)
	CreateJob(nj model.NewJob) (*model.Job, error)
	FetchAllJobs() ([]*models.Job, error)
	FetchJobByCompanyID(cid string) ([]*models.Job, error)
	FetchJobByID(jid int) (*models.Job, error)
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
