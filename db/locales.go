package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	tracer "github.com/opentracing/opentracing-go"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) InsertLocale(ctx context.Context, name string) error {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.InsertLocale")
	defer span.Finish()

	s := &models.Locale{Name: name}

	err := s.Insert(ctx, q.DB, boil.Infer())
	if err != nil {
		return fmt.Errorf("failed to InsertLocale: %w", err)
	}
	return nil
}

func (q *SqlQuerier) FindLocaleByName(ctx context.Context, name string) (*models.Locale, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.FindLocaleByName")
	defer span.Finish()

	s, err := models.Locales(qm.Where("name=?", name)).One(ctx, q.DB)
	if err != nil {
		return nil, fmt.Errorf("FindApiKeyByName: %w", err)
	}

	q.logger.Debugf(ctx, "FindLocaleByName got: ", s)
	return s, nil
}
