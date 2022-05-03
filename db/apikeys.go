package db

import (
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) InsertApiKey(key string) error {

	s := &models.APIKey{KeyVal: key}

	err := s.Insert(Ctx, DB, boil.Infer())
	return err
}

func (q *SqlQuerier) FindApiKeyByName(key string) (*models.APIKey, error) {
	s, err := models.APIKeys(qm.Where("key_val=?", key)).One(Ctx, DB)
	if err != nil {
		return nil, err
	}
	return s, err
}
