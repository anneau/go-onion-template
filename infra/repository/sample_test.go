package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/anneau/go-template/domain/model"
	errctrl "github.com/anneau/go-template/errorctrl"
	"github.com/anneau/go-template/infra/database"
	"github.com/anneau/go-template/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestSampleRepositoryFind(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(s *sql.Tx) error
		arg     model.SampleID
		want    *model.Sample
		wantErr error
	}{
		{
			name: "idを条件にデータを取得できる",
			setup: func(s *sql.Tx) error {
				_, err := s.Exec("INSERT INTO samples (id, name) VALUES ('0188fc55-7873-70bf-9b11-a9be5407b14e', 'sample')")
				if err != nil {
					return err
				}
				return nil
			},
			arg: errctrl.Must(model.NewIDFrom[model.Sample]("0188fc55-7873-70bf-9b11-a9be5407b14e")),
			want: &model.Sample{
				SampleID: errctrl.Must(model.NewIDFrom[model.Sample]("0188fc55-7873-70bf-9b11-a9be5407b14e")),
				Name:     "sample",
			},
			wantErr: nil,
		},
		{
			name: "idが存在しない場合はエラーを返す",
			setup: func(s *sql.Tx) error {
				return nil
			},
			arg:     errctrl.Must(model.NewIDFrom[model.Sample]("0188fc55-7873-70bf-9b11-a9be5407b14e")),
			want:    nil,
			wantErr: sql.ErrNoRows,
		},
	}

	ctx := context.Background()
	conn := database.NewTestConnection()

	_ = errctrl.Must(conn.Exec("TRUNCATE samples"))

	for _, tt := range tests {
		tx := errctrl.Must(conn.Begin())
		repo := repository.NewSampleRepository(ctx, tx)
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				errctrl.MustExec(tx.Rollback())
			})

			if tt.setup != nil {
				err := tt.setup(tx)
				assert.NoError(t, err)
			}

			actual, err := repo.Find(tt.arg)

			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
