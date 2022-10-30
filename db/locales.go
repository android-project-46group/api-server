package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) InsertLocale(name string) error {

	s := &models.Locale{Name: name}

	err := s.Insert(q.ctx, q.DB, boil.Infer())
	return fmt.Errorf("InsertLocale: %w", err)
}

func (q *SqlQuerier) FindLocaleByName(name string) (*models.Locale, error) {
	s, err := models.Locales(qm.Where("name=?", name)).One(q.ctx, q.DB)
	if err != nil {
		return nil, fmt.Errorf("FindApiKeyByName: %w", err)
	}
	return s, nil
}
