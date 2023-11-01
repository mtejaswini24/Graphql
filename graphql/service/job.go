package service

import (
	"errors"
	"graphql/graph/model"
	"graphql/models"
	"strconv"

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
	jid := strconv.FormatUint(uint64(j.ID), 10)
	j1 := model.Job{
		Jid:      jid,
		JobTitle: j.JobTitle,
		Salary:   j.Salary,
		Company:  company,
	}

	// Successfully created the record, return the user.
	return &j1, nil
}
func (s *Conn) FetchAllJobs() ([]*model.Job, error) {
	var jobDetails []*model.Job
	result := s.db.Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("job details are not there")
	}
	return jobDetails, nil
}
func (s *Conn) FetchJobByCompanyID(cid string) ([]*model.Job, error) {
	var jobDetails []*model.Job
	result := s.db.Where("id=?", cid).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("job is not there with that company id")
	}
	return jobDetails, nil
}
func (s *Conn) FetchJobByID(jid string) (*model.Job, error) {
	var jobById model.Job
	result := s.db.Where("jid=?", jid).First(&jobById)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("job is not there with that id")
	}
	return &jobById, nil
}
