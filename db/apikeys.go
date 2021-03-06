package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) InsertApiKey(key string) error {

	s := &models.APIKey{KeyVal: key}

	err := s.Insert(q.ctx, q.DB, boil.Infer())
	return fmt.Errorf("InsertApiKey: %w", err)
}

func (q *SqlQuerier) FindApiKeyByName(key string) (*models.APIKey, error) {
	s, err := models.APIKeys(qm.Where("key_val=?", key)).One(q.ctx, q.DB)
	if err != nil {
		return nil, fmt.Errorf("FindApiKeyByName: %w", err)
	}
	return s, nil
}
