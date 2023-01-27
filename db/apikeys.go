package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) InsertApiKey(ctx context.Context, key string) error {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.InsertApiKey")
	defer span.Finish()

	s := &models.APIKey{KeyVal: key}

	err := s.Insert(ctx, q.DB, boil.Infer())
	if err != nil {
		return fmt.Errorf("failed to InsertApiKey: %w", err)
	}
	return nil
}

func (q *SqlQuerier) FindApiKeyByName(ctx context.Context, key string) (*models.APIKey, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.FindApiKeyByName")
	defer span.Finish()

	s, err := models.APIKeys(qm.Where("key_val=?", key)).One(ctx, q.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to FindApiKeyByName: %w", err)
	}

	q.logger.Debugf(ctx, "FindApiKeyByName got: ", s)
	return s, nil
}
