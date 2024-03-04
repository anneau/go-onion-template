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
	tx  *sql.Tx
}

func NewSampleRepository(ctx context.Context, tx *sql.Tx) model.SampleRepository {
	return &sampleRepository{ctx: ctx, tx: tx}
}

func (r *sampleRepository) Find(id model.SampleID) (*model.Sample, error) {
	table := &table.SampleTable{}
	err := r.tx.QueryRowContext(r.ctx, "SELECT s.id, s.name FROM samples as s WHERE s.id = $1;", id.Value).Scan(&table.ID, &table.Name)

	if err != nil {
		return nil, err
	}

	id, err = model.NewIDFrom[model.Sample](table.ID)

	if err != nil {
		return nil, err
	}

	return &model.Sample{
		SampleID: id,
		Name:     table.Name,
	}, nil
}
