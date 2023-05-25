package data

import (
	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertTags() {
	tags := []string{"可愛い", "あざとい"}
	for _, tag := range tags {
		m := &models.Tag{TagName: tag}
		// err := m.Insert(Ctx, DB, boil.Infer())
		m.Insert(Ctx, DB, boil.Infer())
	}
}

func GetTagIdByName(name string) (int, error) {
	ss, err := models.Tags(qm.Where("tag_name = ?", name)).One(Ctx, DB)
	if err != nil {
		return 0, err
	}

	return ss.TagID, nil
}
