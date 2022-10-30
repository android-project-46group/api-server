package data

import (
	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InsertApiKey() {
	key := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	a := &models.APIKey{
		KeyVal: key,
	}
	a.Insert(Ctx, DB, boil.Infer())
}
