package repository

import "github.com/anneau/go-template/domain/model"

var _ model.SampleRepository = (*sampleRepository)(nil)

type sampleRepository struct{}

func NewSampleRepository() model.SampleRepository {
	return &sampleRepository{}
}

func (r *sampleRepository) Find(id model.SampleID) (*model.Sample, error) {
	return nil, nil
}
