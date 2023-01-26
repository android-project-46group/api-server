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
	return s, nil
}
