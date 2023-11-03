package service

import (
	"errors"
	"graphql/graph/model"
	"graphql/models"

	"github.com/rs/zerolog/log"
)

func (s *Conn) CreateJob(nj model.NewJob) (*model.Job, error) {
	// We prepare the Company record.
	j := models.Job{
		Cid:      nj.Cid,
		JobTitle: nj.JobTitle,
		Salary:   nj.Salary,
	}
	err := s.db.Create(&j).Error
	if err != nil {
		return &model.Job{}, err
	}
	company, err := s.FetchCompanyByID(j.Cid)
	if err != nil {
		log.Error().Err(err).Msg("not getting company details")
		return nil, errors.New("job creation failed")
	}
	j1 := model.Job{
		ID:       int(j.ID),
		JobTitle: j.JobTitle,
		Salary:   j.Salary,
		Company:  company,
	}

	// Successfully created the record, return the user.
	return &j1, nil
}
func (s *Conn) FetchAllJobs() ([]*models.Job, error) {
	var jobDetails []*models.Job
	result := s.db.Preload("Company").Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("job details are not there")
	}
	return jobDetails, nil
}
func (s *Conn) FetchJobByCompanyID(cid string) ([]*models.Job, error) {
	var jobDetails []*models.Job

	result := s.db.Preload("Company").Where("cid=?", cid).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("job is not there with that company id")
	}
	return jobDetails, nil
}
func (s *Conn) FetchJobByID(jid int) (*models.Job, error) {
	var jobById models.Job
	result := s.db.Preload("Company").Where("id=?", jid).First(&jobById)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("job is not there with that id")
	}
	return &jobById, nil
}
