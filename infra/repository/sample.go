package repository

import (
	"context"
	"database/sql"

	"github.com/anneau/go-template/domain/model"
	"github.com/anneau/go-template/infra/database/table"
)

var _ model.SampleRepository = (*sampleRepository)(nil)

type sampleRepository struct {
	ctx context.Context
	t   sql.Tx
}

func NewSampleRepository(ctx context.Context) model.SampleRepository {
	return &sampleRepository{ctx: ctx}
}

func (r *sampleRepository) Find(id model.SampleID) (*model.Sample, error) {
	data := &table.SampleTable{}
	err := r.t.QueryRowContext(r.ctx, "SELECT * FROM sample WHERE id = :id;", sql.Named("id", id.Value)).Scan(&data)

	if err != nil {
		return nil, err
	}

	id, err = model.NewIDFrom[model.Sample](data.ID)

	if err != nil {
		return nil, err
	}

	return &model.Sample{
		SampleID: id,
		Name:     data.Name,
	}, nil
}
