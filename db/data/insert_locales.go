package data

import (
	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InsertLocales() {
	a := &models.Locale{
		Name: "en",
	}
	a.Insert(Ctx, DB, boil.Infer())

	b := &models.Locale{
		Name: "ja",
	}
	b.Insert(Ctx, DB, boil.Infer())
}
