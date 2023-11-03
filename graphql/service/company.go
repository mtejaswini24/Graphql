package service

import (
	"errors"
	"graphql/graph/model"
	"graphql/models"

	"github.com/rs/zerolog/log"
)

func (s *Conn) CreateCompany(nc model.NewCompany) (*model.Company, error) {
	c := models.Company{
		Name:     nc.Name,
		Location: nc.Location,
	}
	//calling default create method
	err := s.db.Create(&c).Error
	if err != nil {
		return &model.Company{}, err
	}
	c1 := model.Company{
		ID:       int(c.ID),
		Name:     c.Name,
		Location: c.Location,
	}
	// Successfully created the record, return the user.
	return &c1, nil
}
func (s *Conn) FetchAllCompanies() ([]*model.Company, error) {
	var companyDetails []*model.Company

	result := s.db.Find(&companyDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("company details are not there")
	}
	return companyDetails, nil
}
func (s *Conn) FetchCompanyByID(cid string) (*model.Company, error) {
	var companyById model.Company
	result := s.db.Where("id=?", cid).First(&companyById)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("company is not there with that id")
	}
	return &companyById, nil
}
